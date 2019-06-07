package snatsubnet

import (
	"bytes"
	"context"
	"net"
	"reflect"

	mapset "github.com/deckarep/golang-set"
	"github.com/gaurav-dalvi/snat-operator/cmd/manager/utils"
	aciv1 "github.com/gaurav-dalvi/snat-operator/pkg/apis/aci/v1"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const snatsubnetFinalizer = "finalizer.snatsubnet.aci.snat"

var log = logf.Log.WithName("controller_snatsubnet")

// Add creates a new SnatSubnet Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileSnatSubnet{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("snatsubnet-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource SnatSubnet
	err = c.Watch(&source.Kind{Type: &aciv1.SnatSubnet{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileSnatSubnet{}

// ReconcileSnatSubnet reconciles a SnatSubnet object
type ReconcileSnatSubnet struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a SnatSubnet object and makes changes based on the state read
// and what is in the SnatSubnet.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileSnatSubnet) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request:", request.Namespace+"/"+request.Name)
	reqLogger.Info("Reconciling SnatSubnet")

	// Fetch the SnatSubnet instance
	instance := &aciv1.SnatSubnet{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Check if the snatsubnet cr was marked to be deleted
	isSnatSubnetToBeDeleted := instance.GetDeletionTimestamp() != nil
	if isSnatSubnetToBeDeleted {
		if utils.Contains(instance.GetFinalizers(), snatsubnetFinalizer) {
			// Run finalization logic for snatsubnetFinalizer. If the
			// finalization logic fails, don't remove the finalizer so
			// that we can retry during the next reconciliation.
			if err := r.finalizeSnatSubnet(reqLogger, instance); err != nil {
				return reconcile.Result{}, err
			}

			// Remove snatsubnetFinalizer. Once all finalizers have been
			// removed, the object will be deleted.
			instance.SetFinalizers(utils.Remove(instance.GetFinalizers(), snatsubnetFinalizer))
			err := r.client.Update(context.TODO(), instance)
			if err != nil {
				return reconcile.Result{}, err
			}
		}
		return reconcile.Result{}, nil
	}

	// Validation of SnatSpec struct
	validator := utils.Validator{}
	validator.ValidateSnatSubnet(instance)
	if !validator.Validated {
		reqLogger.Error(err, "SnatSpec is not valid - "+validator.ErrorMessage)
		return reconcile.Result{}, err
	}

	// Add finalizer for this CR
	if !utils.Contains(instance.GetFinalizers(), snatsubnetFinalizer) {
		if err := r.addFinalizer(instance); err != nil {
			return reconcile.Result{}, err
		}
	}

	// Only one instance of SnatSubnet object has to be present in the system.
	// Any new creation of snatsubnet CR will be deleted here.
	snatSubnetList := &aciv1.SnatSubnetList{}
	err = r.client.List(context.TODO(), &client.ListOptions{Namespace: ""}, snatSubnetList)
	if err != nil {
		reqLogger.Error(err, "Failed to list existing snatsubnets\n")
		return reconcile.Result{}, err
	}
	if len(snatSubnetList.Items) != 1 {
		reqLogger.Error(err, "Only one instance of snatsubnet should be present in the system, deleting this one: "+request.Name)
		// select the item matching with request name
		for _, item := range snatSubnetList.Items {
			if item.ObjectMeta.Name == request.Name {
				err = r.client.Delete(context.TODO(), &item)
				if err != nil {
					reqLogger.Error(err, "failed to delete a snatsubnet item : "+item.ObjectMeta.Name)
					return reconcile.Result{}, err
				}
				break
			}
		}
		return reconcile.Result{}, err
	}

	// Update the status if necessary
	expandedsnatports := utils.ExpandPortRanges(instance.Spec.Snatports, instance.Spec.Pernodeports)
	if !reflect.DeepEqual(instance.Status.Expandedsnatports, expandedsnatports) {
		instance.Status.Expandedsnatports = expandedsnatports
		err := r.client.Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "failed to update the SnatSubnet")
			return reconcile.Result{}, err
		}
		reqLogger.Info("Updated snatsubnet status", "Status:", instance.Status)
	}

	// In case of update (deletion of subnets which are currently used by snatip)
	if err := r.finalizeSnatSubnet(reqLogger, instance); err != nil {
		return reconcile.Result{}, err
	}

	// If snatIP resource is using any of the IP in snatSubnet, then check that and send appropriate error
	// return r.handleSnatSubnetUpdate(*instance)
	return reconcile.Result{}, nil
}

// Cleanup steps to be done when snatsubnet resource is getting deleted.
// This will delete all the matching snatip CRs which had
// IPs from this snatip resource
func (r *ReconcileSnatSubnet) finalizeSnatSubnet(reqLogger logr.Logger, m *aciv1.SnatSubnet) error {
	// TODO(user): Add the cleanup steps that the operator
	// needs to do before the CR can be deleted

	// Get all snatip CRs
	snatIPList, err := utils.GetAllSnatIps(r.client)
	if err != nil {
		reqLogger.Error(err, "snatsubnet CR could not found")
		return err
	}

	// Check if snatsubnet's CIDR and snatIP's CIDR are subset of each other or not
	for _, item := range m.Spec.Snatipsubnets {
		bIp, _, bErr := net.ParseCIDR(item)
		if bErr != nil {
			reqLogger.Error(err, "Invalid bigger CIDR", "CIDR", item)
			return nil
		}
		for _, snatip := range snatIPList.Items {
			for _, ip := range snatip.Spec.Snatipsubnets {
				sIp, _, sErr := net.ParseCIDR(ip)
				if sErr != nil {
					reqLogger.Error(err, "Invalid smaller CIDR", "CIDR", item)
					return nil
				}
				if bytes.Equal(bIp, sIp) {
					// Add it to set so that we can delete all the snatip objects later
					err := utils.DeleteSnatIPCR(snatip.ObjectMeta.Name, r.client)
					if err != nil {
						reqLogger.Error(err, "Could not find snatip IP", "Name", item)
					}
				}
			}
		}

	}
	reqLogger.Info("Successfully finalized snatsubnet")
	return nil
}

// Add finalizer string to snatsubnet resource to run cleanup logic on delete
func (r *ReconcileSnatSubnet) addFinalizer(m *aciv1.SnatSubnet) error {
	log.Info("Adding Finalizer for the SnatSubnet")
	m.SetFinalizers(append(m.GetFinalizers(), snatsubnetFinalizer))

	// Update CR
	err := r.client.Update(context.TODO(), m)
	if err != nil {
		log.Error(err, "Failed to update SnatSubnet with finalizer")
		return err
	}
	return nil
}

// To handle update of snatsubnet resource. If any of the IP  from the status is in use in Snatip resource,
// then return the error
func (r *ReconcileSnatSubnet) handleSnatSubnetUpdate(instance aciv1.SnatSubnet) (reconcile.Result, error) {
	snatIpList := &aciv1.SnatIPList{}
	err := r.client.List(context.TODO(), &client.ListOptions{Namespace: ""}, snatIpList)
	if err != nil {
		log.Error(err, "Failed to list existing snatip items\n")
		return reconcile.Result{}, err
	}

	// All Ip addresses which are using bu snatsubnet resource
	origAllIps := utils.ExpandCIDRs(instance.Spec.Snatipsubnets)
	origIPSet := mapset.NewSet()
	for _, item := range origAllIps {
		origIPSet.Add(item)
	}
	log.Info("Expanded IPs", "OriginalIPs", origIPSet)

	var temp []string
	// All Ip addresses which are using bu snatIP resources
	for _, item := range snatIpList.Items {
		temp = append(temp, item.Status.Allips...)
	}

	currIPSet := mapset.NewSet()
	for _, item := range temp {
		currIPSet.Add(item)
	}
	log.Info("Expanded IPs", "CurrentIPs", currIPSet)

	diffSet := origIPSet.Difference(currIPSet)
	if diffSet.Difference(currIPSet).Cardinality() != 0 && currIPSet.Cardinality() != 0 {
		log.Error(err, "Can not delete / update snatsubnet resource as IPs are in use")
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

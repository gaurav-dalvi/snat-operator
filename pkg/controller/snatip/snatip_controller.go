package snatip

import (
	"context"
	"reflect"

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

const snatipFinalizer = "finalizer.snatip.aci.snat"

var log = logf.Log.WithName("controller_snatip")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new SnatIP Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileSnatIP{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("snatip-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource SnatIP
	err = c.Watch(&source.Kind{Type: &aciv1.SnatIP{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileSnatIP{}

// ReconcileSnatIP reconciles a SnatIP object
type ReconcileSnatIP struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a SnatIP object and makes changes based on the state read
// and what is in the SnatIP.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileSnatIP) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request:", request.Namespace+"/"+request.Name)
	reqLogger.Info("Reconciling SnatIP")

	// Fetch the SnatIP instance
	instance := &aciv1.SnatIP{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}
	reqLogger.Info("snatip-controller", "resourcetype", instance.Spec.Resourcetype, "name", instance.Spec.Name)

	// Check if the snatip cr was marked to be deleted
	isSnatIPToBeDeleted := instance.GetDeletionTimestamp() != nil
	if isSnatIPToBeDeleted {
		if utils.Contains(instance.GetFinalizers(), snatipFinalizer) {
			// Run finalization logic for snatipFinalizer. If the
			// finalization logic fails, don't remove the finalizer so
			// that we can retry during the next reconciliation.
			if err := r.finalizeSnatIP(reqLogger, instance); err != nil {
				return reconcile.Result{}, err
			}

			// Remove snatipFinalizer. Once all finalizers have been
			// removed, the object will be deleted.
			instance.SetFinalizers(utils.Remove(instance.GetFinalizers(), snatipFinalizer))
			err := r.client.Update(context.TODO(), instance)
			if err != nil {
				return reconcile.Result{}, err
			}
		}
		return reconcile.Result{}, nil
	}

	// Validation of SnatIPSpec struct
	validator := utils.Validator{}
	validator.ValidateSnatIP(instance)
	if !validator.Validated {
		reqLogger.Error(err, "SnatIPSpec is not valid, hence deleting it : "+validator.ErrorMessage)
		// Deleting snatip instance
		err = r.client.Delete(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "failed to delete a santip item : "+instance.ObjectMeta.Name)
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, err
	}

	// Add finalizer for this CR
	if !utils.Contains(instance.GetFinalizers(), snatipFinalizer) {
		if err := r.addFinalizer(instance); err != nil {
			return reconcile.Result{}, err
		}
	}

	// Update the status if necessary
	expandedIPs := utils.ExpandCIDRs(instance.Spec.Snatipsubnets)
	if !reflect.DeepEqual(instance.Status.Allips, expandedIPs) {
		instance.Status.Allips = expandedIPs
		err := r.client.Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "failed to update the SnatIps")
			return reconcile.Result{}, err
		}
		reqLogger.Info("Updated snatip status", "Status:", instance.Status)
	}

	return reconcile.Result{}, nil
}

// Cleanup steps to be done when snatip resource is getting deleted.
// This will delete all the matching snatallocations CRs which had
// IPs from this snatip resource
func (r *ReconcileSnatIP) finalizeSnatIP(reqLogger logr.Logger, m *aciv1.SnatIP) error {
	// TODO(user): Add the cleanup steps that the operator
	// needs to do before the CR can be deleted

	// Get all snatallocation CRs
	allocList, err := utils.GetAllSnatAllocations(r.client)
	if err != nil {
		reqLogger.Error(err, "snat allocations could not found")
		return err
	}
	if len(allocList.Items) == 0 {
		reqLogger.Info("No snatallocation CR present in the system")
		return nil
	}

	// Delete snatallocation CR if there is matching IP in snatIP CR
	for _, item := range allocList.Items {
		if utils.Contains(m.Status.Allips, item.Spec.Snatip) {
			// Delete snatallocation CR
			err := utils.DeleteSnatAllocationCR(item.Spec.Podname, item.Spec.Namespace, r.client)
			if err != nil {
				reqLogger.Error(err, "snat allocation can not be deleted", "name", item.Spec.Podname+"/"+item.Spec.Namespace)
				return err
			}
		}
	}

	reqLogger.Info("Successfully finalized snatip")
	return nil
}

// Add finalizer string to snatip resource to run cleanup logic on delete
func (r *ReconcileSnatIP) addFinalizer(m *aciv1.SnatIP) error {
	log.Info("Adding Finalizer for the SnatIP")
	m.SetFinalizers(append(m.GetFinalizers(), snatipFinalizer))

	// Update CR
	err := r.client.Update(context.TODO(), m)
	if err != nil {
		log.Error(err, "Failed to update SnatIP with finalizer")
		return err
	}
	return nil
}

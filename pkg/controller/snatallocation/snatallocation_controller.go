package snatallocation

import (
	"context"
	"strings"

	"github.com/gaurav-dalvi/snat-operator/cmd/manager/utils"
	noironetworksv1 "github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_snatallocation")

// Add creates a new SnatAllocation Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileSnatAllocation{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("snatallocation-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource SnatAllocation
	err = c.Watch(&source.Kind{Type: &noironetworksv1.SnatAllocation{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestsFromMapFunc{ToRequests: HandlePodsForPodsMapper(mgr.GetClient(), []predicate.Predicate{})})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileSnatAllocation{}

// ReconcileSnatAllocation reconciles a SnatAllocation object
type ReconcileSnatAllocation struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a SnatAllocation object and makes changes based on the state read
// and what is in the SnatAllocation.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileSnatAllocation) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request:", request.Namespace+"/"+request.Name)
	reqLogger.Info("Reconciling SnatAllocation")

	// If pod belongs to namespace in snapt ip cr
	if strings.HasPrefix(request.Name, "snat-namespace-") {
		found_pod := &corev1.Pod{}
		err := r.client.Get(context.TODO(), types.NamespacedName{Name: utils.GetPodNameFromReoncileRequest(request.Name), Namespace: request.Namespace}, found_pod)
		if err != nil && errors.IsNotFound(err) {
			return reconcile.Result{}, err
		} else if err != nil {
			return reconcile.Result{}, err
		}
		reqLogger.Info("********POD found********", "Pod name", found_pod.Name)
		return reconcile.Result{Requeue: true}, nil
	}

	// Fetch the SnatAllocation instance
	instance := &noironetworksv1.SnatAllocation{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			reqLogger.Info("Creating SNATAllocation CR")
			cr := newSnatAllocationCR()
			err = r.client.Create(context.TODO(), cr)
			if err != nil {
				reqLogger.Error(err, "failed to create a snat allocation cr")
				return reconcile.Result{}, err
			}
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	reqLogger.Info("Instance found", "Instance name", instance.Name)

	// Check if this SnatSubnet already exists
	found_snatsubnet := &noironetworksv1.SnatSubnet{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: "foo", Namespace: "default"}, found_snatsubnet)
	if err != nil && errors.IsNotFound(err) {
		// Pod created successfully - don't requeue
		return reconcile.Result{}, err
	} else if err != nil {
		return reconcile.Result{}, err
	}

	reqLogger.Info("snat_allocation", "SnatSubnet.PerNodePorts", found_snatsubnet.Spec.Pernodeports, "SnatSubnet.Status", found_snatsubnet.Status.Expandedsnatports)

	// Check if this SnatIP CR already exists
	found_snatip := &noironetworksv1.SnatIP{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: "foo1", Namespace: "default"}, found_snatip)
	if err != nil && errors.IsNotFound(err) {
		// Pod created successfully - don't requeue
		return reconcile.Result{}, err
	} else if err != nil {
		return reconcile.Result{}, err
	}

	reqLogger.Info("snat_allocation-2", "SnatSubnet.Name", found_snatip.Spec.Name, "SnatSubnet.Namespace", found_snatip.Spec.Namespace)
	reqLogger.Info("snat_allocation-3", "SnatSubnet.ResourceType", found_snatip.Spec.Resourcetype)

	return reconcile.Result{}, nil
}

// newSnatAllocationCR returns a SnatAllocationCR
func newSnatAllocationCR() *noironetworksv1.SnatAllocation {

	return &noironetworksv1.SnatAllocation{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "first",
			Namespace: "default",
		},
		Spec: noironetworksv1.SnatAllocationSpec{
			Podname: "gaurvpod",
		},
	}
}

func (r *ReconcileSnatAllocation) getAllSnatSubnets() (*noironetworksv1.SnatSubnetList, error) {
	snatSubnetList := &noironetworksv1.SnatSubnetList{}
	err := r.client.List(context.TODO(), &client.ListOptions{Namespace: ""}, snatSubnetList)
	if err != nil {
		log.Error(err, "failed to list existing snatsubnets")
		return &noironetworksv1.SnatSubnetList{}, err
	}
	return snatSubnetList, nil
}

func (r *ReconcileSnatAllocation) getAllSnatAllocations() (*noironetworksv1.SnatAllocationList, error) {
	snatAllocationList := &noironetworksv1.SnatAllocationList{}
	err := r.client.List(context.TODO(), &client.ListOptions{Namespace: ""}, snatAllocationList)
	if err != nil {
		log.Error(err, "failed to list existing SnatAllocationList")
		return &noironetworksv1.SnatAllocationList{}, err
	}
	return snatAllocationList, nil
}

func (r *ReconcileSnatAllocation) getAllSnatIps() (*noironetworksv1.SnatIPList, error) {
	snatIpList := &noironetworksv1.SnatIPList{}
	err := r.client.List(context.TODO(), &client.ListOptions{Namespace: ""}, snatIpList)
	if err != nil {
		log.Error(err, "failed to list existing snatsubnets")
		return &noironetworksv1.SnatIPList{}, err
	}
	return snatIpList, nil
}

package snatallocation

import (
	"context"

	noironetworksv1 "github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
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

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

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
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileSnatAllocation) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling SnatAllocation")

	// Fetch the SnatAllocation instance
	instance := &noironetworksv1.SnatAllocation{}
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

	// Check if this SnatSubnet already exists
	found_snatsubnet := &noironetworksv1.SnatSubnet{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: "foo", Namespace: "default"}, found_snatsubnet)
	if err != nil && errors.IsNotFound(err) {
		// Pod created successfully - don't requeue
		return reconcile.Result{}, err
	} else if err != nil {
		return reconcile.Result{}, err
	}

	reqLogger.Info("snat_allocation", "SnatSubnet.PerNodePorts", found_snatsubnet.Spec.PerNodePorts, "SnatSubnet.SnatIpSubnets", found_snatsubnet.Spec.SnatIpSubnets)
	reqLogger.Info("snat_allocation-1", "SnatSubnet.SnatPorts", found_snatsubnet.Spec.SnatPorts)

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

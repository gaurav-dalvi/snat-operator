package snatallocation

import (
	"context"
	"strings"

	"github.com/gaurav-dalvi/snat-operator/cmd/manager/utils"
	aciv1 "github.com/gaurav-dalvi/snat-operator/pkg/apis/aci/v1"
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
	err = c.Watch(&source.Kind{Type: &aciv1.SnatAllocation{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestsFromMapFunc{ToRequests: HandlePodsForPodsMapper(mgr.GetClient(), []predicate.Predicate{})})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &corev1.Node{}}, &handler.EnqueueRequestsFromMapFunc{ToRequests: HandlePodsForPodsMapper(mgr.GetClient(), []predicate.Predicate{})})
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

	// If pod belongs to any resource in the snatip
	if strings.HasPrefix(request.Name, "snat-") {
		result, err := r.handlePodEvent(request)
		return result, err
	} else {
		// Fetch the SnatAllocation instance
		instance := &aciv1.SnatAllocation{}
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

		reqLogger.Info("Instance found", "Instance name", instance.Name)
	}

	return reconcile.Result{}, nil
}

// newSnatAllocationCR returns a SnatAllocationCR
func newSnatAllocationCR(alloc aciv1.SnatAllocationSpec) *aciv1.SnatAllocation {

	return &aciv1.SnatAllocation{
		ObjectMeta: metav1.ObjectMeta{
			Name:      alloc.Podname + "-snat-alloc",
			Namespace: "default",
		},
		Spec: alloc,
	}
}

// This function handles Pod events which are triggering snatallocation's reconcile loop
func (r *ReconcileSnatAllocation) handlePodEvent(request reconcile.Request) (reconcile.Result, error) {
	// Podname: name of the pod for which loop was triggered
	// resourceType: type of snatip resource, in which this pod belongs
	// resourceName: name of the snatip resource, in which this pod belongs
	podName, resourceType, resourceName := utils.GetPodNameFromReoncileRequest(request.Name)

	// Query this pod using k8s client
	found_pod := &corev1.Pod{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: podName, Namespace: request.Namespace}, found_pod)
	if err != nil && errors.IsNotFound(err) {
		log.Info("Pod deleted", "PodName:", request.Name)
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}
	log.Info("********POD found********", "Pod name", found_pod.ObjectMeta.Name)

	// Get snatsubnet resource
	snatsubnetItem, err := utils.GetAllSnatSubnets(r.client)
	if err != nil {
		log.Error(err, "snatsubnetItem could not be found, resulting an err")
		return reconcile.Result{}, err
	}
	log.Info("SnatSubnet found", "snatSubnet:", snatsubnetItem)

	// Get the snatip resource in which this pod belongs
	snatipItem, err := utils.SearchSnatIPByName(resourceName, resourceType, r.client)
	if err != nil {
		log.Error(err, "snatip item could not be found, resulting an err")
		return reconcile.Result{}, err
	}
	log.Info("SnatIp found", "snatIp", snatipItem)

	ip, portRange, uid := utils.GetIPPortRangeForPod(*snatipItem, snatsubnetItem, r.client)
	// Create snatallocation CR object only when pod is in `Running` state
	if found_pod.Status.Phase == "Running" {
		// Create snatAllocation CR
		spec := aciv1.SnatAllocationSpec{
			Podname:       found_pod.ObjectMeta.Name,
			Poduid:        string(found_pod.ObjectMeta.UID),
			Nodename:      found_pod.Spec.NodeName,
			Snatportrange: portRange,
			Snatip:        ip,
			Namespace:     found_pod.ObjectMeta.Namespace,
			Snatipuid:     uid,
			Macaddress:    "f0:18:98:83:4a:8b",
		}
		cr := newSnatAllocationCR(spec)
		err = r.client.Create(context.TODO(), cr)
		if err != nil {
			log.Error(err, "failed to create a snat allocation cr")
			return reconcile.Result{}, err
		}
		log.Info("Created snatallocation object", "Snatallocation", cr)
	}

	return reconcile.Result{}, nil
}

package snatglobalinfo

import (
	"context"
	"os"
	"strings"

	"github.com/gaurav-dalvi/snat-operator/cmd/manager/utils"
	aciv1 "github.com/gaurav-dalvi/snat-operator/pkg/apis/aci/v1"
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

var log = logf.Log.WithName("controller_snatglobalinfo")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new SnatGlobalInfo Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileSnatGlobalInfo{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("snatglobalinfo-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource SnatGlobalInfo
	err = c.Watch(&source.Kind{Type: &aciv1.SnatGlobalInfo{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &aciv1.SnatLocalInfo{}}, &handler.EnqueueRequestsFromMapFunc{ToRequests: HandleLocalInfosMapper(mgr.GetClient(), []predicate.Predicate{})})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileSnatGlobalInfo implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileSnatGlobalInfo{}

// ReconcileSnatGlobalInfo reconciles a SnatGlobalInfo object
type ReconcileSnatGlobalInfo struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a SnatGlobalInfo object and makes changes based on the state read
// and what is in the SnatGlobalInfo.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileSnatGlobalInfo) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling SnatGlobalInfo")
	if strings.Contains(request.Name, "snat-localinfo-") {
		localInfoName := request.Name[len("snat-localinfo-"):]
		result, err := r.handleLocainfoEvent(localInfoName)
		return result, err
	} else {
		// Fetch the SnatGlobalInfo instance
		instance := &aciv1.SnatGlobalInfo{}
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
	}

	return reconcile.Result{}, nil
}

// This function handles Localinfo events which are triggering snatGlobalInfo's reconcile loop
func (r *ReconcileSnatGlobalInfo) handleLocainfoEvent(name string) (reconcile.Result, error) {

	// Fetch the SnatLocainfo instance
	instance := &aciv1.SnatLocalInfo{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: os.Getenv("ACI_SNAT_NAMESPACE")}, instance)
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

	// Get SnatGlobalInfo instance
	globalInfo := &aciv1.SnatGlobalInfo{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: os.Getenv("ACI_SNAGLOBALINFO_NAME"),
		Namespace: os.Getenv("ACI_SNAT_NAMESPACE")}, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Create SnatGlobalInfo Object
			globalInfos := []aciv1.GlobalInfo{}
			// get Mac Address
			nodeinfo, _ := utils.GetNodeInfoCRObject(r.client, instance.ObjectMeta.Name)
			temp := aciv1.GlobalInfo{
				MacAddress: nodeinfo.Spec.MacAddress,
				PortRanges: utils.GetNextPortRange(),
				NodeName:   instance.ObjectMeta.Name,
				SnatIpUid:  "some uid",
				Protocols:  []string{"tcp", "udp", "icmp"},
			}
			globalInfos = append(globalInfos, temp)
			tempMap := make(map[string][]aciv1.GlobalInfo)
			tempMap["10.20.30.40"] = globalInfos
			spec := aciv1.SnatGlobalInfoSpec{
				GlobalInfos: tempMap,
			}
			return utils.CreateSnatGlobalInfoCR(r.client, spec)
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	} else {
		// update snatGlobalInfo object
		// GlobalInfo CR is already present, Append GlobalInfo object into Spec's map  and update Globalinfo
		return utils.UpdateGlobalInfoCR(r.client, *globalInfo)
	}

}

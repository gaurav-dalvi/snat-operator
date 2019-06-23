package utils

import (
	"context"
	"os"
	"strings"

	// nodeinfo "github.com/noironetworks/aci-containers/pkg/nodeinfo/apis/aci.nodeinfo/v1"
	aciv1 "github.com/gaurav-dalvi/snat-operator/pkg/apis/aci/v1"
	"github.com/prometheus/common/log"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// Check if given pod belongs to given deployment or not
func CheckIfPodForDeployment(c client.Client, pod corev1.Pod, deploymentName, deploymentNamespace string) (bool, error) {

	// Get the deployment
	deployment := &appsv1.Deployment{}
	err := c.Get(context.TODO(), types.NamespacedName{Name: deploymentName, Namespace: deploymentNamespace}, deployment)
	if err != nil {
		UtilLog.Error(err, "Deployment deleted, name: "+deploymentName)
		return false, err
	}

	// Check if  any of the deployment lable is present in pod's label or not
	UtilLog.Info("Deployment labels", "Label", deployment.ObjectMeta.Labels)
	UtilLog.Info("Pod labels", "Label", pod.ObjectMeta.Labels)
	for dKey, dVal := range deployment.ObjectMeta.Labels {
		for pKey, pVal := range pod.ObjectMeta.Labels {
			if dKey == pKey && dVal == pVal {
				// Match found.
				UtilLog.Info("Labels matched", "Deployment label", dKey+"="+dVal, "PodLabel", pKey+"="+pVal)
				return true, nil
			}
		}
	}

	return false, nil
}

// Check if given pod belongs to given service or not
func CheckIfPodForService(corev1.Pod, corev1.Service) bool {
	return true
}

// Given a reconcile request name, it extracts out pod name by omiiting snat-policy- from it
// eg: snat-policy-foo-podname -> podname, foo
func GetPodNameFromReoncileRequest(requestName string) (string, string) {

	temp := strings.Split(requestName, "-")
	if len(temp) != 4 {
		UtilLog.Info("Length should be 4", "input string:", requestName, "lengthGot", len(temp))
		return "", ""
	}
	snatPolicyName, podName := temp[2], temp[3]
	return podName, snatPolicyName
}

// Get nodeinfo object matching given name of the node
// Optimization can be done here:
// if we know namespace of this nodeinfo object then we can type request.NamespacedName{Name: , Namespace:}
// in Get call and directly get the object instead of doing List and iterating.
// But for that namespace has to be knowen. We can push aci-containers-system / kube-system inserted as ENV var
// in this container then we can refer to that.
func GetNodeInfoCRObject(c client.Client, nodeName string) (aciv1.NodeInfo, error) {
	nodeinfoList := &aciv1.NodeInfoList{}
	err := c.List(context.TODO(), &client.ListOptions{Namespace: ""}, nodeinfoList)
	if err != nil && errors.IsNotFound(err) {
		UtilLog.Error(err, "Cound not find nodeinfo object")
		return aciv1.NodeInfo{}, err
	}

	for _, item := range nodeinfoList.Items {
		if item.ObjectMeta.Name == nodeName {
			UtilLog.Info("Nodeinfo object found", "For NodeName:", item.ObjectMeta.Name)
			return item, nil
		}
	}
	return aciv1.NodeInfo{}, err

}

// Given a reconcile request name, it extracts out node name by omiiting node-event- from it
func GetNodeNameFromReoncileRequest(requestName string) string {
	if strings.HasPrefix(requestName, "node-event-") {
		return requestName[len("node-event-"):]
	}
	return requestName
}

// Given a nodeName, return LocalInfo CR object if present
func GetLocalInfoCR(c client.Client, nodeName, namespace string) (aciv1.SnatLocalInfo, error) {

	foundLocalIfo := &aciv1.SnatLocalInfo{}
	err := c.Get(context.TODO(), types.NamespacedName{Name: nodeName, Namespace: namespace}, foundLocalIfo)
	if err != nil && errors.IsNotFound(err) {
		log.Info("LocalIfo not present ", "foundLocalIfo:", nodeName)
		return aciv1.SnatLocalInfo{}, nil
	} else if err != nil {
		return aciv1.SnatLocalInfo{}, err
	}

	return *foundLocalIfo, nil
}

// Given a policyName, return SnatPolicy CR object if present
func GetSnatPolicyCR(c client.Client, policyName string) (aciv1.SnatPolicy, error) {

	foundSnatPolicy := &aciv1.SnatPolicy{}
	err := c.Get(context.TODO(), types.NamespacedName{Name: policyName, Namespace: os.Getenv("ACI_SNAT_NAMESPACE")}, foundSnatPolicy)
	if err != nil && errors.IsNotFound(err) {
		log.Info("SnatPolicy not present", "foundSnatPolicy:", policyName)
		return aciv1.SnatPolicy{}, nil
	} else if err != nil {
		return aciv1.SnatPolicy{}, err
	}

	return *foundSnatPolicy, nil
}

// createSnatLocalInfoCR Creates a SnatLocalInfo CR
func CreateLocalInfoCR(c client.Client, localInfoSpec aciv1.SnatLocalInfoSpec, nodeName string) (reconcile.Result, error) {

	obj := &aciv1.SnatLocalInfo{
		ObjectMeta: metav1.ObjectMeta{
			Name:      nodeName,
			Namespace: os.Getenv("ACI_SNAT_NAMESPACE"),
		},
		Spec: localInfoSpec,
	}
	err := c.Create(context.TODO(), obj)
	if err != nil {
		log.Error(err, "failed to create a snat locainfo cr")
		return reconcile.Result{}, err
	}
	log.Info("Created localinfo object", "SnatLocalInfo", obj)
	return reconcile.Result{}, nil
}

// UpdateSnatLocalInfoCR Updates a SnatLocalInfo CR
func UpdateLocalInfoCR(c client.Client, localInfo aciv1.SnatLocalInfo) (reconcile.Result, error) {

	err := c.Update(context.TODO(), &localInfo)
	if err != nil {
		log.Error(err, "failed to update a snat locainfo cr")
		return reconcile.Result{}, err
	}
	log.Info("Updated localinfo object", "SnatLocalInfo", localInfo)
	return reconcile.Result{}, nil
}

// // Get all SnatSubnet CRs from k8s clustner
// func GetAllSnatSubnets(c client.Client) (aciv1.SnatSubnet, error) {
// 	snatSubnetList := &aciv1.SnatSubnetList{}
// 	err := c.List(context.TODO(), &client.ListOptions{Namespace: ""}, snatSubnetList)
// 	if err != nil {
// 		UtilLog.Error(err, "failed to list existing snatsubnets")
// 		return aciv1.SnatSubnet{}, err
// 	}

// 	// We are making sure that there will always be one instance of snatsubnet in the system.
// 	return snatSubnetList.Items[0], nil
// }

// // Get all SnatAllocation CRs from k8s clustner
// func GetAllSnatAllocations(c client.Client) (aciv1.SnatAllocationList, error) {
// 	snatAllocationList := &aciv1.SnatAllocationList{}
// 	err := c.List(context.TODO(), &client.ListOptions{Namespace: ""}, snatAllocationList)
// 	if err != nil {
// 		UtilLog.Error(err, "failed to list existing SnatAllocationList")
// 		return aciv1.SnatAllocationList{}, err
// 	}
// 	return *snatAllocationList, nil
// }

// // Get all SnatIP CRs from k8s clustner
// func GetAllSnatIps(c client.Client) (aciv1.SnatIPList, error) {
// 	snatIpList := &aciv1.SnatIPList{}
// 	err := c.List(context.TODO(), &client.ListOptions{Namespace: ""}, snatIpList)
// 	if err != nil {
// 		UtilLog.Error(err, "failed to list existing snatsubnets")
// 		return aciv1.SnatIPList{}, err
// 	}
// 	return *snatIpList, nil
// }

// // Delete respective snat-allocation cr object for given pod
// func DeleteSnatAllocationCR(podName, nameSpace string, c client.Client) error {

// 	// Get all snatallocation CR objects
// 	allocList, err := GetAllSnatAllocations(c)
// 	if len(allocList.Items) == 0 {
// 		// This can not happen. There has to be one entry matching for this pod
// 		UtilLog.Error(err, "This can not happen. There has to be one entry matching for this pod:", "PodName/Namespace", podName+"/"+nameSpace)
// 		return err
// 	}

// 	for _, item := range allocList.Items {
// 		if item.Spec.Podname == podName && item.Spec.Namespace == nameSpace {
// 			// Found snatalloc item, deleting it
// 			err = c.Delete(context.TODO(), &item)
// 			if err != nil {
// 				UtilLog.Error(err, "failed to delete a snatallocation item : "+item.ObjectMeta.Name)
// 				return err
// 			}
// 			break
// 		}
// 	}

// 	return nil
// }

// // Delete respective snatip cr object for given name
// func DeleteSnatIPCR(name string, c client.Client) error {

// 	// Get all snatip CR objects
// 	snatIPList, err := GetAllSnatIps(c)
// 	if len(snatIPList.Items) == 0 {
// 		UtilLog.Error(err, "Could not get list of snatIPs")
// 		return err
// 	}

// 	for _, item := range snatIPList.Items {
// 		if item.ObjectMeta.Name == name {
// 			// Found snatip item, deleting it
// 			err = c.Delete(context.TODO(), &item)
// 			if err != nil {
// 				UtilLog.Error(err, "failed to delete a snatip item : "+item.ObjectMeta.Name)
// 				return err
// 			}
// 			break
// 		}
// 	}

// 	return nil
// }

// // Get IP and port for pod for which notification has come to reconcile loop
// func GetIPPortRangeForPod(snatIpItem aciv1.SnatIP,
// 	snatSubnetItem aciv1.SnatSubnet, c client.Client) (string, snattypes.PortRange, string) {

// 	if len(snatIpItem.Status.Allips) <= 0 || len(snatSubnetItem.Status.Expandedsnatports) <= 0 {
// 		UtilLog.Info("Allips can not be empty. Resulting to error")
// 		return "", snattypes.PortRange{}, ""
// 	}

// 	allocList, _ := GetAllSnatAllocations(c)
// 	if len(allocList.Items) == 0 {
// 		// No allocation has been done so do first allocation
// 		return snatIpItem.Status.Allips[0], snatSubnetItem.Status.Expandedsnatports[0], string(uuid.NewUUID())
// 	}

// 	return "", snattypes.PortRange{}, ""
// }

// // Given a name, this function finds snatIP object
// func SearchSnatIPByName(name, resourceType string, c client.Client) (*aciv1.SnatIP, error) {
// 	instance := &aciv1.SnatIP{}
// 	snatipList, err := GetAllSnatIps(c)
// 	if err != nil {
// 		UtilLog.Error(err, "failed to list of all snatsubnets")
// 		return &aciv1.SnatIP{}, err
// 	}

// 	// Search for `name`
// 	for _, item := range snatipList.Items {
// 		if item.Spec.Name == name && item.Spec.Resourcetype == resourceType {
// 			instance = &item
// 			return instance, nil
// 		}
// 	}

// 	// Could not find snatip with name, so erroring it out
// 	UtilLog.Error(err, "Could not find snatip item for", "name", name)
// 	return instance, err
// }

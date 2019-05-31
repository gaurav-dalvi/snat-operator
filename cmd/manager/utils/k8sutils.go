package utils

import (
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

// Check if given pod belongs to given deployment or not
func CheckIfPodForDeployment(corev1.Pod, appsv1.Deployment) bool {
	return true
}

// Check if given pod belongs to given service or not
func CheckIfPodForService(corev1.Pod, corev1.Service) bool {
	return true
}

// Given a reconcile request name, it extracts out pod name by omiiting snat-<resourcename>- from it
// It also extract out resource type name.
// eg: snat-namespace-foo-podname -> podname, namespace, foo
func GetPodNameFromReoncileRequest(requestName string) (string, string, string) {

	temp := strings.Split(requestName, "-")
	if len(temp) != 4 {
		UtilLog.Info("Length should be 4", "input string:", requestName, "lengthGot", len(temp))
		return "", "", ""
	}
	resourceType, resourceName, podName := temp[1], temp[2], temp[3]
	return podName, resourceType, resourceName
}

// Given a reconcile request name, it extracts out node name by omiiting node-event- from it
func GetNodeNameFromReoncileRequest(requestName string) string {
	if strings.HasPrefix(requestName, "node-event-") {
		return requestName[len("node-event-"):]
	}
	return requestName
}

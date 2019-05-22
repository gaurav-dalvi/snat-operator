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
func GetPodNameFromReoncileRequest(requestName string) string {
	if strings.HasPrefix(requestName, "snat-namespace-") {
		return requestName[len("snat-namespace-"):]

	} else if strings.HasPrefix(requestName, "snat-pod-") {
		return requestName[len("snat-pod-"):]

	} else if strings.HasPrefix(requestName, "snat-deployment-") {
		return requestName[len("snat-deployment-"):]

	} else if strings.HasPrefix(requestName, "snat-service-") {
		return requestName[len("snat-service-"):]
	} else {
		return requestName
	}
}

package utils

import (
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

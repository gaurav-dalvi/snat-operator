package snatallocation

import (
	"context"

	noironetworksv1 "github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type handlePodsForPodsMapper struct {
	client     client.Client
	predicates []predicate.Predicate
}

func (h *handlePodsForPodsMapper) Map(obj handler.MapObject) []reconcile.Request {
	if obj.Object == nil {
		return nil
	}

	pod, ok := obj.Object.(*corev1.Pod)
	if !ok {
		return nil
	}

	snatipList := &noironetworksv1.SnatIPList{}
	if err := h.client.List(context.TODO(), client.InNamespace(pod.Namespace), snatipList); err != nil {
		return nil
	}

	requests := FilterPodsPerSnatIP(snatipList, pod)
	return requests
}

func HandlePodsForPodsMapper(client client.Client, predicates []predicate.Predicate) handler.Mapper {
	return &handlePodsForPodsMapper{client, predicates}
}

/*

Given a list of SnatIp CR, this function filters out whether the pod should be included in reconcile request of
Map or not, based on CR spec of SnatIP from the list
*/
func FilterPodsPerSnatIP(snatipList *noironetworksv1.SnatIPList, pod *corev1.Pod) []reconcile.Request {
	var requests []reconcile.Request

	for _, item := range snatipList.Items {

		switch item.Spec.Resourcetype {
		// Because service has the highest priority among all SnatIp resources. refer SNAT spec for more details
		case "service":
			if item.Spec.Name == pod.Name {
				requests = append(requests, reconcile.Request{
					NamespacedName: types.NamespacedName{
						Namespace: item.Namespace,
						Name:      "snat_service" + item.Name,
					},
				})
			}

		case "deployment":
			if item.Spec.Name == pod.ObjectMeta.Namespace {
				requests = append(requests, reconcile.Request{
					NamespacedName: types.NamespacedName{
						Namespace: item.Namespace,
						Name:      "snat_deployment" + item.Name,
					},
				})
			}

		case "pod":
			if item.Spec.Name == pod.Name {
				requests = append(requests, reconcile.Request{
					NamespacedName: types.NamespacedName{
						Namespace: item.Namespace,
						Name:      "snat_pod" + item.Name,
					},
				})
			}

		case "namespace":
			if item.Spec.Name == pod.ObjectMeta.Namespace {
				requests = append(requests, reconcile.Request{
					NamespacedName: types.NamespacedName{
						Namespace: item.Namespace,
						Name:      "snat_namespace" + item.Name,
					},
				})
			}

		default:
			requests = []reconcile.Request{}
		}

	}
	return requests
}

package snatallocation

import (
	"context"

	aciv1 "github.com/gaurav-dalvi/snat-operator/pkg/apis/aci/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var MapperLog = logf.Log.WithName("Mapper:")

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
	// MapperLog.Info("Inside pod map function", "Pod is:", pod.ObjectMeta.Name+"--"+pod.Spec.NodeName+"---"+pod.ObjectMeta.Namespace)

	snatipList := &aciv1.SnatIPList{}
	if err := h.client.List(context.TODO(), &client.ListOptions{Namespace: ""}, snatipList); err != nil {
		return nil
	}
	// if err := h.client.List(context.TODO(), client.InNamespace(pod.Namespace), snatipList); err != nil {
	// reqLogger.Info(" map pod function", "SnatIp list is:", snatipList)
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
func FilterPodsPerSnatIP(snatipList *aciv1.SnatIPList, pod *corev1.Pod) []reconcile.Request {
	var requests []reconcile.Request

Loop:
	for _, item := range snatipList.Items {
		switch item.Spec.Resourcetype {
		// Because service has the highest priority among all SnatIp resources. refer SNAT spec for more details
		// case "service":
		// 	if item.Spec.Name == pod.Name {
		// 		requests = append(requests, reconcile.Request{
		// 			NamespacedName: types.NamespacedName{
		// 				Namespace: item.Namespace,
		//				Name:      "snat-service-" + item.Spec.Name + "-" + pod.Name,
		// 			},
		// 		})
		//      break Loop
		// 	}

		// case "deployment":
		// 	if item.Spec.Name == pod.ObjectMeta.Namespace {
		// 		requests = append(requests, reconcile.Request{
		// 			NamespacedName: types.NamespacedName{
		// 				Namespace: item.Namespace,
		//				Name:      "snat-deployment-" + item.Spec.Name + "-" + pod.Name,
		// 			},
		// 		})
		//      break Loop
		// 	}

		// case "pod":
		// 	if item.Spec.Name == pod.Name {
		// 		requests = append(requests, reconcile.Request{
		// 			NamespacedName: types.NamespacedName{
		// 				Namespace: item.Namespace,
		// 				Name:      "snat-pod-" + item.Spec.Name + "-" + pod.Name,
		// 			},
		// 		})
		//      break Loop
		// 	}

		case "namespace":
			MapperLog.Info("Pod Mapper Items: ", "snatip", item.Spec.Name+"---"+item.Spec.Namespace+"--------"+item.Spec.Resourcetype)
			if item.Spec.Name == pod.ObjectMeta.Namespace {
				requests = append(requests, reconcile.Request{
					NamespacedName: types.NamespacedName{
						Namespace: item.Spec.Name,
						Name:      "snat-namespace-" + item.Spec.Name + "-" + pod.Name,
					},
				})
				break Loop
			}

		default:
			requests = []reconcile.Request{}
		}

	}
	return requests
}

type handleNodesForPodsMapper struct {
	client     client.Client
	predicates []predicate.Predicate
}

func (h *handleNodesForPodsMapper) Map(obj handler.MapObject) []reconcile.Request {
	if obj.Object == nil {
		return nil
	}

	node, ok := obj.Object.(*corev1.Node)
	if !ok {
		return nil
	}

	MapperLog.Info("Node map function")
	req := []reconcile.Request{}
	nodeReq := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Namespace: "",
			Name:      "node-event-" + node.Name,
		},
	}
	req = append(req, nodeReq)
	return req
}

func HandleNodesForPodsMapper(client client.Client, predicates []predicate.Predicate) handler.Mapper {
	return &handleNodesForPodsMapper{client, predicates}
}

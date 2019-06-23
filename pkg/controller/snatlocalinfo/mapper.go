package snatlocalinfo

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

	// Get all the snatpolicies
	snatPolicyList := &aciv1.SnatPolicyList{}
	if err := h.client.List(context.TODO(), &client.ListOptions{Namespace: ""}, snatPolicyList); err != nil {
		return nil
	}

	requests := FilterPodsPerSnatPolicy(h.client, snatPolicyList, pod)
	return requests
}

func HandlePodsForPodsMapper(client client.Client, predicates []predicate.Predicate) handler.Mapper {
	return &handlePodsForPodsMapper{client, predicates}
}

/*

Given a list of SnatPolicy CR, this function filters out whether the pod should be included in reconcile request of
Map or not, based on CR spec of SnatPolicy from the list
*/
func FilterPodsPerSnatPolicy(c client.Client, snatPolicyList *aciv1.SnatPolicyList, pod *corev1.Pod) []reconcile.Request {
	var requests []reconcile.Request

Loop:
	for _, item := range snatPolicyList.Items {

		if false {
			// Need to check here how to handle this.
			//item.Spec.Selector == aciv1.SnatPolicy.Spec.Selector{} {
			MapperLog.Info("Cluster Scoped", "Needs special handling", item.Spec.SnatIp)
		} else {
			// Now need to match pod with correct snatPolicy item
			// According to priority:
			// 1: Labels of pod should match exactly with labels of podSelector
			// 2: Deployment of pod should match exactly with deployment of podSelector
			// 3: Namespace of pod should match exactly with namespace of podSelector
			// right now namespace approach is implemented.
			if item.Spec.Selector.Namespace == pod.ObjectMeta.Namespace {
				requests = append(requests, reconcile.Request{
					NamespacedName: types.NamespacedName{
						Namespace: pod.ObjectMeta.Namespace,
						Name:      "snat-policy-" + item.ObjectMeta.Name + "-" + pod.ObjectMeta.Name,
					},
				})
				break Loop
			}
		}
	}
	return requests
}

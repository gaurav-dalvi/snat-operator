package snatglobalinfo

import (
	aciv1 "github.com/gaurav-dalvi/snat-operator/pkg/apis/aci/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var MapperLog = logf.Log.WithName("Mapper:")

type handleLocalInfosMapper struct {
	client     client.Client
	predicates []predicate.Predicate
}

// Globalinfo reconcile loop has secondary resource as LocalInfo.
// When Locainfo change happens, we are pusing request to reconcile loop of GlobalInfo CR
// Request will be of this format
// Name: "snat-localinfo-" + <locainfo CR name>
func (h *handleLocalInfosMapper) Map(obj handler.MapObject) []reconcile.Request {
	if obj.Object == nil {
		return nil
	}

	localInfo, ok := obj.Object.(*aciv1.SnatLocalInfo)
	if !ok {
		return nil
	}

	var requests []reconcile.Request
	requests = append(requests, reconcile.Request{
		NamespacedName: types.NamespacedName{
			Namespace: localInfo.ObjectMeta.Namespace,
			Name:      "snat-localinfo-" + localInfo.ObjectMeta.Name,
		},
	})

	return requests
}

func HandleLocalInfosMapper(client client.Client, predicates []predicate.Predicate) handler.Mapper {
	return &handleLocalInfosMapper{client, predicates}
}

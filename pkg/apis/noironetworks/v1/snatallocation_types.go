package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SnatAllocationSpec defines the desired state of SnatAllocation
// +k8s:openapi-gen=true
type SnatAllocationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Podname       string    `json:"podname"`
	Poduid        types.UID `json:"poduid"`
	Nodename      string    `json:"nodename"`
	Snatportrange PortRange `json:"snatportrange"`
	Snatip        string    `json:"snatip"`
	Namespace     string    `json:"namespace"`
	Macaddress    string    `json:"macaddress"`
	Scope         string    `json:"string,omitempty"`
}

// SnatAllocationStatus defines the observed state of SnatAllocation
// +k8s:openapi-gen=true
type SnatAllocationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SnatAllocation is the Schema for the snatallocations API
// +k8s:openapi-gen=true
type SnatAllocation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SnatAllocationSpec   `json:"spec,omitempty"`
	Status SnatAllocationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SnatAllocationList contains a list of SnatAllocation
type SnatAllocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnatAllocation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SnatAllocation{}, &SnatAllocationList{})
}

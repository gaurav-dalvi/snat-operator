package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SnatIPSpec defines the desired state of SnatIP
// +k8s:openapi-gen=true
type SnatIPSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// +kubebuilder:validation:Enum=pod,deployment,service,namespace
	Resourcetype  string   `json:"resourcetype"`
	Name          string   `json:"name"`
	Namespace     string   `json:"namespace"`
	Snatipsubnets []string `json:"snatipsubnets"`
	// +kubebuilder:validation:Enum=tcp,udp,icmp
	Protocols []string `json:"protocols,omitempty"`
}

// SnatIPStatus defines the observed state of SnatIP
// +k8s:openapi-gen=true
type SnatIPStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Allips []string `json:"allips,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SnatIP is the Schema for the snatips API
// +k8s:openapi-gen=true
type SnatIP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SnatIPSpec   `json:"spec,omitempty"`
	Status SnatIPStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SnatIPList contains a list of SnatIP
type SnatIPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnatIP `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SnatIP{}, &SnatIPList{})
}

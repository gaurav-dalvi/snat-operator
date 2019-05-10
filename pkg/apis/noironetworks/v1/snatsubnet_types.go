package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SnatSubnetSpec defines the desired state of SnatSubnet
// +k8s:openapi-gen=true
type SnatSubnetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	PerNodePorts  int         `json:"per_node_ports"`
	SnatIpSubnets []string    `json:"snat_ip_subnets"`
	SnatPorts     []PortRange `json:"snat_ports"`
}

// SnatSubnetStatus defines the observed state of SnatSubnet
// +k8s:openapi-gen=true
type SnatSubnetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SnatSubnet is the Schema for the snatsubnets API
// +k8s:openapi-gen=true
type SnatSubnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SnatSubnetSpec   `json:"spec,omitempty"`
	Status SnatSubnetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SnatSubnetList contains a list of SnatSubnet
type SnatSubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnatSubnet `json:"items"`
}

type PortRange struct {
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SnatSubnet{}, &SnatSubnetList{})
}

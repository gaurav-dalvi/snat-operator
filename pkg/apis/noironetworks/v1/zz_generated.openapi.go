// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatAllocation":       schema_pkg_apis_noironetworks_v1_SnatAllocation(ref),
		"github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatAllocationSpec":   schema_pkg_apis_noironetworks_v1_SnatAllocationSpec(ref),
		"github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatAllocationStatus": schema_pkg_apis_noironetworks_v1_SnatAllocationStatus(ref),
		"github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatIP":               schema_pkg_apis_noironetworks_v1_SnatIP(ref),
		"github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatIPSpec":           schema_pkg_apis_noironetworks_v1_SnatIPSpec(ref),
		"github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatIPStatus":         schema_pkg_apis_noironetworks_v1_SnatIPStatus(ref),
		"github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatSubnet":           schema_pkg_apis_noironetworks_v1_SnatSubnet(ref),
		"github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatSubnetSpec":       schema_pkg_apis_noironetworks_v1_SnatSubnetSpec(ref),
		"github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatSubnetStatus":     schema_pkg_apis_noironetworks_v1_SnatSubnetStatus(ref),
	}
}

func schema_pkg_apis_noironetworks_v1_SnatAllocation(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SnatAllocation is the Schema for the snatallocations API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatAllocationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatAllocationStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatAllocationSpec", "github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatAllocationStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_noironetworks_v1_SnatAllocationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SnatAllocationSpec defines the desired state of SnatAllocation",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_noironetworks_v1_SnatAllocationStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SnatAllocationStatus defines the observed state of SnatAllocation",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_noironetworks_v1_SnatIP(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SnatIP is the Schema for the snatips API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatIPSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatIPStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatIPSpec", "github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatIPStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_noironetworks_v1_SnatIPSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SnatIPSpec defines the desired state of SnatIP",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_noironetworks_v1_SnatIPStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SnatIPStatus defines the observed state of SnatIP",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_noironetworks_v1_SnatSubnet(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SnatSubnet is the Schema for the snatsubnets API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatSubnetSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatSubnetStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatSubnetSpec", "github.com/gaurav-dalvi/snat-operator/pkg/apis/noironetworks/v1.SnatSubnetStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_noironetworks_v1_SnatSubnetSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SnatSubnetSpec defines the desired state of SnatSubnet",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_noironetworks_v1_SnatSubnetStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SnatSubnetStatus defines the observed state of SnatSubnet",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

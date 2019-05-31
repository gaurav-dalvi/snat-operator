// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortRange) DeepCopyInto(out *PortRange) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRange.
func (in *PortRange) DeepCopy() *PortRange {
	if in == nil {
		return nil
	}
	out := new(PortRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatAllocation) DeepCopyInto(out *SnatAllocation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatAllocation.
func (in *SnatAllocation) DeepCopy() *SnatAllocation {
	if in == nil {
		return nil
	}
	out := new(SnatAllocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnatAllocation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatAllocationList) DeepCopyInto(out *SnatAllocationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SnatAllocation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatAllocationList.
func (in *SnatAllocationList) DeepCopy() *SnatAllocationList {
	if in == nil {
		return nil
	}
	out := new(SnatAllocationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnatAllocationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatAllocationSpec) DeepCopyInto(out *SnatAllocationSpec) {
	*out = *in
	out.Snatportrange = in.Snatportrange
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatAllocationSpec.
func (in *SnatAllocationSpec) DeepCopy() *SnatAllocationSpec {
	if in == nil {
		return nil
	}
	out := new(SnatAllocationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatAllocationStatus) DeepCopyInto(out *SnatAllocationStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatAllocationStatus.
func (in *SnatAllocationStatus) DeepCopy() *SnatAllocationStatus {
	if in == nil {
		return nil
	}
	out := new(SnatAllocationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatIP) DeepCopyInto(out *SnatIP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatIP.
func (in *SnatIP) DeepCopy() *SnatIP {
	if in == nil {
		return nil
	}
	out := new(SnatIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnatIP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatIPList) DeepCopyInto(out *SnatIPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SnatIP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatIPList.
func (in *SnatIPList) DeepCopy() *SnatIPList {
	if in == nil {
		return nil
	}
	out := new(SnatIPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnatIPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatIPSpec) DeepCopyInto(out *SnatIPSpec) {
	*out = *in
	if in.Snatipsubnets != nil {
		in, out := &in.Snatipsubnets, &out.Snatipsubnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatIPSpec.
func (in *SnatIPSpec) DeepCopy() *SnatIPSpec {
	if in == nil {
		return nil
	}
	out := new(SnatIPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatIPStatus) DeepCopyInto(out *SnatIPStatus) {
	*out = *in
	if in.AllIps != nil {
		in, out := &in.AllIps, &out.AllIps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatIPStatus.
func (in *SnatIPStatus) DeepCopy() *SnatIPStatus {
	if in == nil {
		return nil
	}
	out := new(SnatIPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatSubnet) DeepCopyInto(out *SnatSubnet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatSubnet.
func (in *SnatSubnet) DeepCopy() *SnatSubnet {
	if in == nil {
		return nil
	}
	out := new(SnatSubnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnatSubnet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatSubnetList) DeepCopyInto(out *SnatSubnetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SnatSubnet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatSubnetList.
func (in *SnatSubnetList) DeepCopy() *SnatSubnetList {
	if in == nil {
		return nil
	}
	out := new(SnatSubnetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnatSubnetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatSubnetSpec) DeepCopyInto(out *SnatSubnetSpec) {
	*out = *in
	if in.Snatipsubnets != nil {
		in, out := &in.Snatipsubnets, &out.Snatipsubnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Snatports != nil {
		in, out := &in.Snatports, &out.Snatports
		*out = make([]PortRange, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatSubnetSpec.
func (in *SnatSubnetSpec) DeepCopy() *SnatSubnetSpec {
	if in == nil {
		return nil
	}
	out := new(SnatSubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatSubnetStatus) DeepCopyInto(out *SnatSubnetStatus) {
	*out = *in
	if in.Expandedsnatports != nil {
		in, out := &in.Expandedsnatports, &out.Expandedsnatports
		*out = make([]PortRange, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatSubnetStatus.
func (in *SnatSubnetStatus) DeepCopy() *SnatSubnetStatus {
	if in == nil {
		return nil
	}
	out := new(SnatSubnetStatus)
	in.DeepCopyInto(out)
	return out
}

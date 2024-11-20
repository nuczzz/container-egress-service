//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Kube-OVN CES Controller Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationMatch) DeepCopyInto(out *DestinationMatch) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationMatch.
func (in *DestinationMatch) DeepCopy() *DestinationMatch {
	if in == nil {
		return nil
	}
	out := new(DestinationMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalIPRule) DeepCopyInto(out *ExternalIPRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalIPRule.
func (in *ExternalIPRule) DeepCopy() *ExternalIPRule {
	if in == nil {
		return nil
	}
	out := new(ExternalIPRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalIPRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalIPRuleList) DeepCopyInto(out *ExternalIPRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExternalIPRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalIPRuleList.
func (in *ExternalIPRuleList) DeepCopy() *ExternalIPRuleList {
	if in == nil {
		return nil
	}
	out := new(ExternalIPRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalIPRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalIPRuleSpec) DeepCopyInto(out *ExternalIPRuleSpec) {
	*out = *in
	if in.ExternalAddresses != nil {
		in, out := &in.ExternalAddresses, &out.ExternalAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DestinationMatches != nil {
		in, out := &in.DestinationMatches, &out.DestinationMatches
		*out = make([]DestinationMatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalIPRuleSpec.
func (in *ExternalIPRuleSpec) DeepCopy() *ExternalIPRuleSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalIPRuleSpec)
	in.DeepCopyInto(out)
	return out
}
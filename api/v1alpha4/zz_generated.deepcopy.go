// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha4

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
	"sigs.k8s.io/cluster-api/errors"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerNetwork) DeepCopyInto(out *DockerNetwork) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerNetwork.
func (in *DockerNetwork) DeepCopy() *DockerNetwork {
	if in == nil {
		return nil
	}
	out := new(DockerNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostVolume) DeepCopyInto(out *HostVolume) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostVolume.
func (in *HostVolume) DeepCopy() *HostVolume {
	if in == nil {
		return nil
	}
	out := new(HostVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMCluster) DeepCopyInto(out *KVMCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMCluster.
func (in *KVMCluster) DeepCopy() *KVMCluster {
	if in == nil {
		return nil
	}
	out := new(KVMCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KVMCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMClusterList) DeepCopyInto(out *KVMClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KVMCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMClusterList.
func (in *KVMClusterList) DeepCopy() *KVMClusterList {
	if in == nil {
		return nil
	}
	out := new(KVMClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KVMClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMClusterSpec) DeepCopyInto(out *KVMClusterSpec) {
	*out = *in
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMClusterSpec.
func (in *KVMClusterSpec) DeepCopy() *KVMClusterSpec {
	if in == nil {
		return nil
	}
	out := new(KVMClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMClusterStatus) DeepCopyInto(out *KVMClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1alpha4.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMClusterStatus.
func (in *KVMClusterStatus) DeepCopy() *KVMClusterStatus {
	if in == nil {
		return nil
	}
	out := new(KVMClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMMachine) DeepCopyInto(out *KVMMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMMachine.
func (in *KVMMachine) DeepCopy() *KVMMachine {
	if in == nil {
		return nil
	}
	out := new(KVMMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KVMMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMMachineList) DeepCopyInto(out *KVMMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KVMMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMMachineList.
func (in *KVMMachineList) DeepCopy() *KVMMachineList {
	if in == nil {
		return nil
	}
	out := new(KVMMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KVMMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMMachineSpec) DeepCopyInto(out *KVMMachineSpec) {
	*out = *in
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	if in.HostVolumes != nil {
		in, out := &in.HostVolumes, &out.HostVolumes
		*out = make([]HostVolume, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMMachineSpec.
func (in *KVMMachineSpec) DeepCopy() *KVMMachineSpec {
	if in == nil {
		return nil
	}
	out := new(KVMMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMMachineStatus) DeepCopyInto(out *KVMMachineStatus) {
	*out = *in
	if in.MachineState != nil {
		in, out := &in.MachineState, &out.MachineState
		*out = new(MachineState)
		**out = **in
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1alpha4.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMMachineStatus.
func (in *KVMMachineStatus) DeepCopy() *KVMMachineStatus {
	if in == nil {
		return nil
	}
	out := new(KVMMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMMachineTemplate) DeepCopyInto(out *KVMMachineTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMMachineTemplate.
func (in *KVMMachineTemplate) DeepCopy() *KVMMachineTemplate {
	if in == nil {
		return nil
	}
	out := new(KVMMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KVMMachineTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMMachineTemplateList) DeepCopyInto(out *KVMMachineTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KVMMachineTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMMachineTemplateList.
func (in *KVMMachineTemplateList) DeepCopy() *KVMMachineTemplateList {
	if in == nil {
		return nil
	}
	out := new(KVMMachineTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KVMMachineTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMMachineTemplateResource) DeepCopyInto(out *KVMMachineTemplateResource) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMMachineTemplateResource.
func (in *KVMMachineTemplateResource) DeepCopy() *KVMMachineTemplateResource {
	if in == nil {
		return nil
	}
	out := new(KVMMachineTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMMachineTemplateSpec) DeepCopyInto(out *KVMMachineTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMMachineTemplateSpec.
func (in *KVMMachineTemplateSpec) DeepCopy() *KVMMachineTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(KVMMachineTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineResources) DeepCopyInto(out *MachineResources) {
	*out = *in
	out.CPU = in.CPU.DeepCopy()
	out.Disk = in.Disk.DeepCopy()
	out.Memory = in.Memory.DeepCopy()
	out.DockerVolumeSize = in.DockerVolumeSize.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineResources.
func (in *MachineResources) DeepCopy() *MachineResources {
	if in == nil {
		return nil
	}
	out := new(MachineResources)
	in.DeepCopyInto(out)
	return out
}

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

package v1alpha4

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KVMMachineTemplateSpec defines the desired state of KVMMachineTemplate
type KVMMachineTemplateSpec struct {
	Template KVMMachineTemplateResource `json:"template"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kvmmachinetemplates,scope=Namespaced,categories=giantswarm;kvm;cluster-api,shortName=kvmmt

// KVMMachineTemplate is the Schema for the kvmmachinetemplates API
type KVMMachineTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KVMMachineTemplateSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// KVMMachineTemplateList contains a list of KVMMachineTemplate.
type KVMMachineTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KVMMachineTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KVMMachineTemplate{}, &KVMMachineTemplateList{})
}

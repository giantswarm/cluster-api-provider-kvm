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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha4"
	"sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1alpha4"
	"sigs.k8s.io/cluster-api/errors"
)

// KVMMachineSpec defines the desired state of KVMMachine
type KVMMachineSpec struct {
	// ProviderID is the unique identifier of the machine assigned by the provider.
	ProviderID *string `json:"providerID,omitempty"`

	// HostPodRef is a reference to the pod in which the virtual machine is running.
	HostPodRef *corev1.ObjectReference `json:"hostPodRef,omitempty"`

	// DockerNetwork defines the network configuration to be used for the docker daemon.
	// +optional
	DockerNetwork DockerNetwork `json:"dockerNetwork"`

	// ResourceRequirements defines the compute and storage requirements of the machine.
	ResourceRequirements MachineResources `json:"resourceRequirements,omitempty"`

	// HostVolumes defines the host volumes (PV) to be mounted into the workload cluster machine.
	// +optional
	// +nullable
	HostVolumes []HostVolume `json:"hostVolumes,omitempty"`

	// Users is the list of users to be configured in the machine's operating system.
	// +optional
	// +nullable
	Users []v1alpha4.User `json:"users,omitempty"`
}

// KVMMachineStatus defines the observed state of KVMMachine
type KVMMachineStatus struct {
	// Ready is true when the provider resource is ready.
	// +optional
	Ready bool `json:"ready"`

	// MachineState is the state of the underlying virtual machine.
	// +optional
	MachineState *MachineState `json:"machineState,omitempty"`

	// FailureReason will be set in the event that there is a terminal problem
	// reconciling the Machine and will contain a succinct value suitable
	// for machine interpretation.
	//
	// This field should not be set for transitive errors that a controller
	// faces that are expected to be fixed automatically over
	// time (like service outages), but instead indicate that something is
	// fundamentally wrong with the Machine's spec or the configuration of
	// the controller, and that manual intervention is required. Examples
	// of terminal errors would be invalid combinations of settings in the
	// spec, values that are unsupported by the controller, or the
	// responsible controller itself being critically misconfigured.
	//
	// Any transient errors that occur during the reconciliation of Machines
	// can be added as events to the Machine object and/or logged in the
	// controller's output.
	// +optional
	FailureReason *errors.MachineStatusError `json:"failureReason,omitempty"`

	// FailureMessage will be set in the event that there is a terminal problem
	// reconciling the Machine and will contain a more verbose string suitable
	// for logging and human consumption.
	//
	// This field should not be set for transitive errors that a controller
	// faces that are expected to be fixed automatically over
	// time (like service outages), but instead indicate that something is
	// fundamentally wrong with the Machine's spec or the configuration of
	// the controller, and that manual intervention is required. Examples
	// of terminal errors would be invalid combinations of settings in the
	// spec, values that are unsupported by the controller, or the
	// responsible controller itself being critically misconfigured.
	//
	// Any transient errors that occur during the reconciliation of Machines
	// can be added as events to the Machine object and/or logged in the
	// controller's output.
	// +optional
	FailureMessage *string `json:"failureMessage,omitempty"`

	// Conditions defines current service state of the KVMMachine.
	// +optional
	Conditions clusterv1.Conditions `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
// +kubebuilder:resource:path=kvmmachines,scope=Namespaced,categories=giantswarm;kvm;cluster-api,shortName=kvmm
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Cluster",type="string",JSONPath=".metadata.labels.cluster\\.x-k8s\\.io/cluster-name",description="Cluster to which this KVMMachine belongs"
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.machineState",description="Virtual machine state"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.ready",description="Machine ready status"
// +kubebuilder:printcolumn:name="ProviderID",type="string",JSONPath=".spec.providerID",description="Provider ID"
// +kubebuilder:printcolumn:name="Machine",type="string",JSONPath=".metadata.ownerReferences[?(@.kind==\"Machine\")].name",description="Machine object which owns with this KVMMachine"

// KVMMachine is the Schema for the kvmmachines API
type KVMMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KVMMachineSpec   `json:"spec,omitempty"`
	Status KVMMachineStatus `json:"status,omitempty"`
}

// GetConditions returns the observations of the operational state of the KVMMachine resource.
func (r *KVMMachine) GetConditions() clusterv1.Conditions {
	return r.Status.Conditions
}

// SetConditions sets the underlying service state of the KVMMachine to the predescribed clusterv1.Conditions.
func (r *KVMMachine) SetConditions(conditions clusterv1.Conditions) {
	r.Status.Conditions = conditions
}

//+kubebuilder:object:root=true

// KVMMachineList contains a list of KVMMachine
type KVMMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KVMMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KVMMachine{}, &KVMMachineList{})
}

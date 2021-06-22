package v1alpha4

import (
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/sets"
)

type DockerNetwork struct {
	// CIDR is the subnet to be used by the docker bridge network on each machine.
	CIDR string `json:"cidr"`
}

// KVMMachineTemplateResource describes the data needed to create am KVMMachine from a template
type KVMMachineTemplateResource struct {
	// Spec is the specification of the desired behavior of the machine.
	Spec KVMMachineSpec `json:"spec"`
}

// MachineResources describes the compute, memory, and storage requirements for a KVM machine and its host container.
type MachineResources struct {
	// CPU is the amount of CPU (whole or partial) reserved for the virtual machine.
	CPU resource.Quantity `json:"cpu"`

	// Disk is the amount of disk storage reserved for the virtual machine.
	Disk resource.Quantity `json:"disk"`

	// Memory is the amount of memory reserved for the virtual machine.
	Memory resource.Quantity `json:"memory"`

	// DockerVolumeSize is the size of the docker volume used for for the virtual machine container.
	DockerVolumeSize resource.Quantity `json:"dockerVolumeSize"`
}

type HostVolume struct {
	// MountTag is the value of the mount tag label which will be used to select the PV to mount to the machine.
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=31
	MountTag string `json:"mountTag"`

	// HostPath is the path of the directory in the management cluster machine to be mounted into the workload cluster machine.
	// +kubebuilder:validation:MinLength=1
	HostPath string `json:"hostPath"`
}

// MachineState describes the state of a KVM virtual machine.
type MachineState string

var (
	// MachineStatePending is the string representing a virtual machine in a pending state.
	MachineStatePending = MachineState("pending")

	// MachineStateRunning is the string representing a virtual machine in a running state.
	MachineStateRunning = MachineState("running")

	// MachineStateShuttingDown is the string representing a virtual machine shutting down.
	MachineStateShuttingDown = MachineState("shutting-down")

	// MachineStateTerminated is the string representing a virtual machine that has been terminated.
	MachineStateTerminated = MachineState("terminated")

	// MachineRunningStates defines the set of states in which a virtual machine is
	// running or going to be running soon.
	MachineRunningStates = sets.NewString(
		string(MachineStatePending),
		string(MachineStateRunning),
	)
)

package scope

import (
	"github.com/pkg/errors"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha4"

	infrav1 "github.com/giantswarm/cluster-api-provider-kvm/api/v1alpha4"
)

// ClusterScopeParams defines the input parameters used to create a new Scope.
type ClusterScopeParams struct {
	Cluster        *clusterv1.Cluster
	KVMCluster     *infrav1.KVMCluster
	ControllerName string
}

// NewClusterScope creates a new Scope from the supplied parameters.
// This is meant to be called for each reconcile iteration.
func NewClusterScope(params ClusterScopeParams) (*ClusterScope, error) {
	if params.Cluster == nil {
		return nil, errors.New("failed to generate new scope from nil Cluster")
	}
	if params.KVMCluster == nil {
		return nil, errors.New("failed to generate new scope from nil AWSCluster")
	}

	clusterScope := &ClusterScope{
		Cluster:        params.Cluster,
		KVMCluster:     params.KVMCluster,
		controllerName: params.ControllerName,
	}

	return clusterScope, nil
}

// ClusterScope defines the basic context for an actuator to operate upon.
type ClusterScope struct {
	Cluster    *clusterv1.Cluster
	KVMCluster *infrav1.KVMCluster

	controllerName string
}

// Close closes the current scope persisting the cluster configuration and status.
func (s *ClusterScope) Close() error {
	// return s.PatchObject()
	return nil
}

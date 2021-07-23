package scope

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"k8s.io/klog/klogr"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha4"
	"sigs.k8s.io/cluster-api/util/patch"
	"sigs.k8s.io/controller-runtime/pkg/client"

	infrav1 "github.com/giantswarm/cluster-api-provider-kvm/api/v1alpha4"
)

// ClusterScopeParams defines the input parameters used to create a new Scope.
type ClusterScopeParams struct {
	Client client.Client
	Logger logr.Logger

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

	if params.Logger == nil {
		params.Logger = klogr.New()
	}

	helper, err := patch.NewHelper(params.KVMCluster, params.Client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init patch helper")
	}

	clusterScope := &ClusterScope{
		Logger:      params.Logger,
		client:      params.Client,
		patchHelper: helper,

		Cluster:        params.Cluster,
		KVMCluster:     params.KVMCluster,
		controllerName: params.ControllerName,
	}

	return clusterScope, nil
}

// ClusterScope defines the basic context for an actuator to operate upon.
type ClusterScope struct {
	logr.Logger
	client      client.Client
	patchHelper *patch.Helper

	Cluster    *clusterv1.Cluster
	KVMCluster *infrav1.KVMCluster

	controllerName string
}

// Close closes the current scope persisting the cluster configuration and status.
func (s *ClusterScope) Close() error {
	return s.PatchObject()
}

// PatchObject persists the cluster configuration and status.
func (s *ClusterScope) PatchObject() error {
	return s.patchHelper.Patch(
		context.TODO(),
		s.KVMCluster)
}

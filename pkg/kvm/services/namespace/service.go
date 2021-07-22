package namespace

import "github.com/giantswarm/cluster-api-provider-kvm/pkg/kvm/scope"

// Service holds a collection of interfaces.
// The interfaces are broken down like this to group functions together.
// One alternative is to have a large list of functions from the ec2 client.
type Service struct {
	scope *scope.ClusterScope
}

// NewService returns a new service.
func NewService(scope *scope.ClusterScope) *Service {
	return &Service{
		scope: scope,
	}
}

func (s *Service) ReconcileNamespace() error {
	s.scope.Info("Reconciling namespace")
	return nil
}

func (s *Service) DeleteNamespace() error {
	s.scope.Info("Deleting namespace")
	return nil
}

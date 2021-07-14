package namespace

// Service holds a collection of interfaces.
// The interfaces are broken down like this to group functions together.
// One alternative is to have a large list of functions from the ec2 client.
type Service struct {
}

// NewService returns a new service.
func NewService() *Service {
	return &Service{}
}

func (s *Service) ReconcileNamespace() error {
	return nil
}

func (s *Service) DeleteNamespace() error {
	return nil
}

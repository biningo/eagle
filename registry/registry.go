package registry

import "context"

/**
*@Author lyer
*@Date 7/19/21 13:38
*@Describe
**/

// Registrar is service registrar.
type Registrar interface {
	// Register the registration.
	Register(ctx context.Context, service *ServiceInstance) error
	// Deregister the registration.
	Deregister(ctx context.Context, service *ServiceInstance) error
}

// Discovery is service discovery.
type Discovery interface {
	// GetService return the service instances according to the service name.
	GetService(ctx context.Context, opts ...ServiceOption) ([]*ServiceInstance, error)
}

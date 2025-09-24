package resource

import "context"

// ResourceService defines the business logic for resource management.
type ResourceService interface {
	GetAllResources(ctx context.Context) ([]Resource, error)
	GetResourceByID(ctx context.Context, id string) (*Resource, error)
	CreateResource(ctx context.Context, resource *Resource) (*Resource, error)
	UpdateResource(ctx context.Context, resource *Resource) (*Resource, error)
	DeleteResource(ctx context.Context, id string) error
	UpdateInstallStatus(ctx context.Context, id string, isInstalled bool) (*Resource, error)
}

// TransactionManager defines the interface for database transaction management.
type TransactionManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
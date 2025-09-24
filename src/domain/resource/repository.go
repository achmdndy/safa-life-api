package resource

import "context"

// ResourceRepository defines the interface for Resource data access.
type ResourceRepository interface {
	GetAll(ctx context.Context) ([]Resource, error)
	GetByID(ctx context.Context, id string) (*Resource, error)
	Create(ctx context.Context, resource *Resource) (*Resource, error)
	Update(ctx context.Context, resource *Resource) (*Resource, error)
	Delete(ctx context.Context, id string) error
	TransactionManager
}
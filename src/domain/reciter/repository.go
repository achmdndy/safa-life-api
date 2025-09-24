package reciter

import "context"

// ReciterRepository defines the interface for Reciter data access.
type ReciterRepository interface {
	GetAll(ctx context.Context) ([]Reciter, error)
	GetByID(ctx context.Context, id string) (*Reciter, error)
	Create(ctx context.Context, reciter *Reciter) (*Reciter, error)
	Update(ctx context.Context, reciter *Reciter) (*Reciter, error)
	Delete(ctx context.Context, id string) error
	TransactionManager
}
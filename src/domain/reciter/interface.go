package reciter

import "context"

// ReciterService defines the business logic for reciter operations.
type ReciterService interface {
	GetAllReciters(ctx context.Context) ([]Reciter, error)
	GetReciterByID(ctx context.Context, id string) (*Reciter, error)
	CreateReciter(ctx context.Context, reciter *Reciter) (*Reciter, error)
	UpdateReciter(ctx context.Context, reciter *Reciter) (*Reciter, error)
	DeleteReciter(ctx context.Context, id string) error
}

// TransactionManager for consistency
type TransactionManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
package core

import "context"

// ContextTransactionManager provides context-based transaction management
type ContextTransactionManager interface {
	// WithTransaction executes a function within a database transaction using context
	// If the function returns an error, the transaction is rolled back
	// Otherwise, the transaction is committed
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

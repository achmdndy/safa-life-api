package core

import (
	"context"

	"github.com/safalife/core-api/src/domain/core"
	"github.com/safalife/core-api/src/infrastructure/configs"
	"gorm.io/gorm"
)

// GormContextTransactionManager implements ContextTransactionManager using GORM
type GormContextTransactionManager struct{}

// NewGormContextTransactionManager creates a new GORM context transaction manager
func NewGormContextTransactionManager() core.ContextTransactionManager {
	return &GormContextTransactionManager{}
}

// WithTransaction executes a function within a database transaction using context
func (tm *GormContextTransactionManager) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return configs.WithTransaction(func(tx *gorm.DB) error {
		// Create a new context with the transaction
		type txKey struct{}

		txCtx := context.WithValue(ctx, txKey{}, tx)
		return fn(txCtx)
	})
}

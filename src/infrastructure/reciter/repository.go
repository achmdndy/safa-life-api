package reciter

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/reciter"
	"gorm.io/gorm"
)

type reciterRepository struct {
	db *gorm.DB
}

// NewReciterRepository creates a new instance of ReciterRepository
func NewReciterRepository(db *gorm.DB) reciter.ReciterRepository {
	return &reciterRepository{db: db}
}

func (r *reciterRepository) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		type contextKey string
		const txRepoKey contextKey = "tx_repo_reciter"
		txCtx := context.WithValue(ctx, txRepoKey, &reciterRepository{db: tx})
		return fn(txCtx)
	})
}

func (r *reciterRepository) getDB(ctx context.Context) *gorm.DB {
	if txRepo, ok := ctx.Value("tx_repo_reciter").(*reciterRepository); ok {
		return txRepo.db.WithContext(ctx)
	}
	return r.db.WithContext(ctx)
}

func (r *reciterRepository) GetAll(ctx context.Context) ([]reciter.Reciter, error) {
	var models []ReciterModel
	if err := r.getDB(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	return ToReciterDomainSlice(models), nil
}

func (r *reciterRepository) GetByID(ctx context.Context, id string) (*reciter.Reciter, error) {
	var model ReciterModel
	if err := r.getDB(ctx).First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	domain := model.ToReciterDomain()
	return &domain, nil
}

func (r *reciterRepository) Create(ctx context.Context, rec *reciter.Reciter) (*reciter.Reciter, error) {
	model := FromReciterDomain(rec)
	if err := r.getDB(ctx).Create(model).Error; err != nil {
		return nil, err
	}
	result := model.ToReciterDomain()
	return &result, nil
}

func (r *reciterRepository) Update(ctx context.Context, rec *reciter.Reciter) (*reciter.Reciter, error) {
	model := FromReciterDomain(rec)
	if err := r.getDB(ctx).Save(model).Error; err != nil {
		return nil, err
	}
	result := model.ToReciterDomain()
	return &result, nil
}

func (r *reciterRepository) Delete(ctx context.Context, id string) error {
	return r.getDB(ctx).Delete(&ReciterModel{}, "id = ?", id).Error
}
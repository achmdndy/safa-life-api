package resource

import (
	"context"
	"github.com/achmdndy/safa-life-api/src/domain/resource"
	"gorm.io/gorm"
)

type resourceRepository struct {
	db *gorm.DB
}

// NewResourceRepository creates a new instance of ResourceRepository
func NewResourceRepository(db *gorm.DB) resource.ResourceRepository {
	return &resourceRepository{db: db}
}

func (r *resourceRepository) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		type contextKey string
		const txRepoKey contextKey = "tx_repo_resource"
		txCtx := context.WithValue(ctx, txRepoKey, &resourceRepository{db: tx})
		return fn(txCtx)
	})
}

func (r *resourceRepository) getDB(ctx context.Context) *gorm.DB {
	if txRepo, ok := ctx.Value("tx_repo_resource").(*resourceRepository); ok {
		return txRepo.db.WithContext(ctx)
	}
	return r.db.WithContext(ctx)
}

func (r *resourceRepository) GetAll(ctx context.Context) ([]resource.Resource, error) {
	var models []ResourceModel
	if err := r.getDB(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	return ToResourceDomainSlice(models), nil
}

func (r *resourceRepository) GetByID(ctx context.Context, id string) (*resource.Resource, error) {
	var model ResourceModel
	if err := r.getDB(ctx).First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	domain := model.ToResourceDomain()
	return &domain, nil
}

func (r *resourceRepository) Create(ctx context.Context, res *resource.Resource) (*resource.Resource, error) {
	model := FromResourceDomain(res)
	if err := r.getDB(ctx).Create(model).Error; err != nil {
		return nil, err
	}
	result := model.ToResourceDomain()
	return &result, nil
}

func (r *resourceRepository) Update(ctx context.Context, res *resource.Resource) (*resource.Resource, error) {
	model := FromResourceDomain(res)
	if err := r.getDB(ctx).Save(model).Error; err != nil {
		return nil, err
	}
	result := model.ToResourceDomain()
	return &result, nil
}

func (r *resourceRepository) Delete(ctx context.Context, id string) error {
	return r.getDB(ctx).Delete(&ResourceModel{}, "id = ?", id).Error
}
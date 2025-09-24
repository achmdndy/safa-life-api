package tajweed

import (
	"context"
	"fmt"

	"github.com/achmdndy/safa-life-api/src/domain/tajweed"
	"gorm.io/gorm"
)

// tajweedRepository implements the TajweedRepository interface.
type tajweedRepository struct {
	db *gorm.DB
}

// NewTajweedRepository creates a new instance of TajweedRepository.
func NewTajweedRepository(db *gorm.DB) tajweed.TajweedRepository {
	return &tajweedRepository{
		db: db,
	}
}

// WithTransaction executes a function within a database transaction.
func (r *tajweedRepository) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		type contextKey string
		const txRepoKey contextKey = "tx_repo_tajweed"
		txCtx := context.WithValue(ctx, txRepoKey, &tajweedRepository{db: tx})
		return fn(txCtx)
	})
}

func (r *tajweedRepository) getDB(ctx context.Context) *gorm.DB {
	if txRepo, ok := ctx.Value("tx_repo_tajweed").(*tajweedRepository); ok {
		return txRepo.db.WithContext(ctx)
	}
	return r.db.WithContext(ctx)
}

// AyahTajweed Repository Methods
func (r *tajweedRepository) GetAyahTajweed(ctx context.Context, tajweedID string, surahID, ayahID int) (*tajweed.AyahTajweed, error) {
	var model AyahTajweedModel
	err := r.getDB(ctx).First(&model, "tajweed_id = ? AND surah_id = ? AND ayah_id = ?", tajweedID, surahID, ayahID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("ayah tajweed for %s, %d:%d not found", tajweedID, surahID, ayahID)
		}
		return nil, fmt.Errorf("failed to get ayah tajweed: %w", err)
	}
	domain := model.ToAyahTajweedDomain()
	return &domain, nil
}

func (r *tajweedRepository) CreateAyahTajweed(ctx context.Context, ayahTajweed *tajweed.AyahTajweed) (*tajweed.AyahTajweed, error) {
	model := FromAyahTajweedDomain(ayahTajweed)
	if err := r.getDB(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create ayah tajweed: %w", err)
	}
	result := model.ToAyahTajweedDomain()
	return &result, nil
}

func (r *tajweedRepository) UpdateAyahTajweed(ctx context.Context, ayahTajweed *tajweed.AyahTajweed) (*tajweed.AyahTajweed, error) {
	model := FromAyahTajweedDomain(ayahTajweed)
	err := r.getDB(ctx).Where("tajweed_id = ? AND surah_id = ? AND ayah_id = ?", model.TajweedID, model.SurahID, model.AyahID).Updates(model).Error
	if err != nil {
		return nil, fmt.Errorf("failed to update ayah tajweed: %w", err)
	}
	
	var updatedModel AyahTajweedModel
	err = r.getDB(ctx).First(&updatedModel, "tajweed_id = ? AND surah_id = ? AND ayah_id = ?", model.TajweedID, model.SurahID, model.AyahID).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated ayah tajweed: %w", err)
	}

	result := updatedModel.ToAyahTajweedDomain()
	return &result, nil
}

func (r *tajweedRepository) DeleteAyahTajweed(ctx context.Context, tajweedID string, surahID, ayahID int) error {
	err := r.getDB(ctx).Delete(&AyahTajweedModel{}, "tajweed_id = ? AND surah_id = ? AND ayah_id = ?", tajweedID, surahID, ayahID).Error
	if err != nil {
		return fmt.Errorf("failed to delete ayah tajweed: %w", err)
	}
	return nil
}

// TajweedRule Repository Methods
func (r *tajweedRepository) GetAllTajweedRules(ctx context.Context) ([]tajweed.TajweedRule, error) {
	var models []TajweedRuleModel
	if err := r.getDB(ctx).Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get all tajweed rules: %w", err)
	}
	return ToTajweedRuleDomainSlice(models), nil
}

func (r *tajweedRepository) GetTajweedRuleByID(ctx context.Context, ruleID string) (*tajweed.TajweedRule, error) {
	var model TajweedRuleModel
	err := r.getDB(ctx).First(&model, "id = ?", ruleID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("tajweed rule with id '%s' not found", ruleID)
		}
		return nil, fmt.Errorf("failed to get tajweed rule by id: %w", err)
	}
	domain := model.ToTajweedRuleDomain()
	return &domain, nil
}

func (r *tajweedRepository) GetTajweedRuleByName(ctx context.Context, ruleName string) (*tajweed.TajweedRule, error) {
	var model TajweedRuleModel
	err := r.getDB(ctx).First(&model, "rule = ?", ruleName).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("tajweed rule with name '%s' not found", ruleName)
		}
		return nil, fmt.Errorf("failed to get tajweed rule by name: %w", err)
	}
	domain := model.ToTajweedRuleDomain()
	return &domain, nil
}

func (r *tajweedRepository) CreateTajweedRule(ctx context.Context, rule *tajweed.TajweedRule) (*tajweed.TajweedRule, error) {
	model := FromTajweedRuleDomain(rule)
	if err := r.getDB(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create tajweed rule: %w", err)
	}
	result := model.ToTajweedRuleDomain()
	return &result, nil
}

func (r *tajweedRepository) UpdateTajweedRule(ctx context.Context, rule *tajweed.TajweedRule) (*tajweed.TajweedRule, error) {
	model := FromTajweedRuleDomain(rule)
	if err := r.getDB(ctx).Save(model).Error; err != nil {
		return nil, fmt.Errorf("failed to update tajweed rule: %w", err)
	}
	result := model.ToTajweedRuleDomain()
	return &result, nil
}

func (r *tajweedRepository) DeleteTajweedRule(ctx context.Context, ruleID string) error {
	if err := r.getDB(ctx).Delete(&TajweedRuleModel{}, "id = ?", ruleID).Error; err != nil {
		return fmt.Errorf("failed to delete tajweed rule: %w", err)
	}
	return nil
}

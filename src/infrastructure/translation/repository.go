package translation

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/translation"
	"gorm.io/gorm"
)

type translationRepository struct {
	db *gorm.DB
}

// NewTranslationRepository creates a new instance of FullTranslationRepository
func NewTranslationRepository(db *gorm.DB) translation.FullTranslationRepository {
	return &translationRepository{db: db}
}

func (r *translationRepository) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		type contextKey string
		const txRepoKey contextKey = "tx_repo_translation"
		txCtx := context.WithValue(ctx, txRepoKey, &translationRepository{db: tx})
		return fn(txCtx)
	})
}

func (r *translationRepository) getDB(ctx context.Context) *gorm.DB {
	if txRepo, ok := ctx.Value("tx_repo_translation").(*translationRepository); ok {
		return txRepo.db.WithContext(ctx)
	}
	return r.db.WithContext(ctx)
}

// Translation (edition) methods
func (r *translationRepository) GetAll(ctx context.Context) ([]translation.Translation, error) {
	var models []TranslationModel
	if err := r.getDB(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	return ToTranslationDomainSlice(models), nil
}

func (r *translationRepository) GetByID(ctx context.Context, id string) (*translation.Translation, error) {
	var model TranslationModel
	if err := r.getDB(ctx).First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	domain := model.ToTranslationDomain()
	return &domain, nil
}

func (r *translationRepository) CreateTranslation(ctx context.Context, t *translation.Translation) (*translation.Translation, error) {
	model := FromTranslationDomain(t)
	if err := r.getDB(ctx).Create(model).Error; err != nil {
		return nil, err
	}
	result := model.ToTranslationDomain()
	return &result, nil
}

func (r *translationRepository) UpdateTranslation(ctx context.Context, t *translation.Translation) (*translation.Translation, error) {
	model := FromTranslationDomain(t)
	if err := r.getDB(ctx).Save(model).Error; err != nil {
		return nil, err
	}
	result := model.ToTranslationDomain()
	return &result, nil
}

func (r *translationRepository) DeleteTranslation(ctx context.Context, id string) error {
	return r.getDB(ctx).Delete(&TranslationModel{}, "id = ?", id).Error
}

// AyahTranslation methods
func (r *translationRepository) GetAyahTranslation(ctx context.Context, translationID string, surahID, ayahID int) (*translation.AyahTranslation, error) {
	var model AyahTranslationModel
	conds := "translation_id = ? AND surah_id = ? AND ayah_id = ?"
	if err := r.getDB(ctx).First(&model, conds, translationID, surahID, ayahID).Error; err != nil {
		return nil, err
	}
	domain := model.ToAyahTranslationDomain()
	return &domain, nil
}

func (r *translationRepository) GetByAyah(ctx context.Context, surahID, ayahID int) ([]translation.AyahTranslation, error) {
	var models []AyahTranslationModel
	if err := r.getDB(ctx).Find(&models, "surah_id = ? AND ayah_id = ?", surahID, ayahID).Error; err != nil {
		return nil, err
	}
	return ToAyahTranslationDomainSlice(models), nil
}

func (r *translationRepository) GetBySurah(ctx context.Context, translationID string, surahID int) ([]translation.AyahTranslation, error) {
	var models []AyahTranslationModel
	conds := "translation_id = ? AND surah_id = ?"
	if err := r.getDB(ctx).Find(&models, conds, translationID, surahID).Error; err != nil {
		return nil, err
	}
	return ToAyahTranslationDomainSlice(models), nil
}

func (r *translationRepository) CreateAyahTranslation(ctx context.Context, at *translation.AyahTranslation) (*translation.AyahTranslation, error) {
	model := FromAyahTranslationDomain(at)
	if err := r.getDB(ctx).Create(model).Error; err != nil {
		return nil, err
	}
	result := model.ToAyahTranslationDomain()
	return &result, nil
}

func (r *translationRepository) UpdateAyahTranslation(ctx context.Context, at *translation.AyahTranslation) (*translation.AyahTranslation, error) {
	model := FromAyahTranslationDomain(at)
	if err := r.getDB(ctx).Save(model).Error; err != nil {
		return nil, err
	}
	result := model.ToAyahTranslationDomain()
	return &result, nil
}

func (r *translationRepository) DeleteAyahTranslation(ctx context.Context, translationID string, surahID, ayahID int) error {
	conds := "translation_id = ? AND surah_id = ? AND ayah_id = ?"
	return r.getDB(ctx).Delete(&AyahTranslationModel{}, conds, translationID, surahID, ayahID).Error
}

package tafsir

import (
	"context"
	"github.com/achmdndy/safa-life-api/src/domain/tafsir"
	"gorm.io/gorm"
)

type tafsirRepository struct {
	db *gorm.DB
}

func NewTafsirRepository(db *gorm.DB) tafsir.FullTafsirRepository {
	return &tafsirRepository{db: db}
}

func (r *tafsirRepository) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		type contextKey string
		const txRepoKey contextKey = "tx_repo_tafsir"
		txCtx := context.WithValue(ctx, txRepoKey, &tafsirRepository{db: tx})
		return fn(txCtx)
	})
}

func (r *tafsirRepository) getDB(ctx context.Context) *gorm.DB {
	if txRepo, ok := ctx.Value("tx_repo_tafsir").(*tafsirRepository); ok {
		return txRepo.db.WithContext(ctx)
	}
	return r.db.WithContext(ctx)
}

// Tafsir (edition) methods
func (r *tafsirRepository) GetAll(ctx context.Context) ([]tafsir.Tafsir, error) {
	var models []TafsirModel
	if err := r.getDB(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	return ToTafsirDomainSlice(models), nil
}

func (r *tafsirRepository) GetByID(ctx context.Context, id string) (*tafsir.Tafsir, error) {
	var model TafsirModel
	if err := r.getDB(ctx).First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	domain := model.ToTafsirDomain()
	return &domain, nil
}

func (r *tafsirRepository) CreateTafsir(ctx context.Context, t *tafsir.Tafsir) (*tafsir.Tafsir, error) {
	model := FromTafsirDomain(t)
	if err := r.getDB(ctx).Create(model).Error; err != nil {
		return nil, err
	}
	result := model.ToTafsirDomain()
	return &result, nil
}

func (r *tafsirRepository) UpdateTafsir(ctx context.Context, t *tafsir.Tafsir) (*tafsir.Tafsir, error) {
	model := FromTafsirDomain(t)
	if err := r.getDB(ctx).Save(model).Error; err != nil {
		return nil, err
	}
	result := model.ToTafsirDomain()
	return &result, nil
}

func (r *tafsirRepository) DeleteTafsir(ctx context.Context, id string) error {
	return r.getDB(ctx).Delete(&TafsirModel{}, "id = ?", id).Error
}

// AyahTafsir methods
func (r *tafsirRepository) GetAyahTafsir(ctx context.Context, tafsirID string, surahID, ayahID int) (*tafsir.AyahTafsir, error) {
	var model AyahTafsirModel
	conds := "tafsir_id = ? AND surah_id = ? AND ayah_id = ?"
	if err := r.getDB(ctx).First(&model, conds, tafsirID, surahID, ayahID).Error; err != nil {
		return nil, err
	}
	domain := model.ToAyahTafsirDomain()
	return &domain, nil
}

func (r *tafsirRepository) GetByAyah(ctx context.Context, surahID, ayahID int) ([]tafsir.AyahTafsir, error) {
	var models []AyahTafsirModel
	if err := r.getDB(ctx).Find(&models, "surah_id = ? AND ayah_id = ?", surahID, ayahID).Error; err != nil {
		return nil, err
	}
	return ToAyahTafsirDomainSlice(models), nil
}

func (r *tafsirRepository) GetBySurah(ctx context.Context, tafsirID string, surahID int) ([]tafsir.AyahTafsir, error) {
	var models []AyahTafsirModel
	conds := "tafsir_id = ? AND surah_id = ?"
	if err := r.getDB(ctx).Find(&models, conds, tafsirID, surahID).Error; err != nil {
		return nil, err
	}
	return ToAyahTafsirDomainSlice(models), nil
}

func (r *tafsirRepository) CreateAyahTafsir(ctx context.Context, at *tafsir.AyahTafsir) (*tafsir.AyahTafsir, error) {
	model := FromAyahTafsirDomain(at)
	if err := r.getDB(ctx).Create(model).Error; err != nil {
		return nil, err
	}
	result := model.ToAyahTafsirDomain()
	return &result, nil
}

func (r *tafsirRepository) UpdateAyahTafsir(ctx context.Context, at *tafsir.AyahTafsir) (*tafsir.AyahTafsir, error) {
	model := FromAyahTafsirDomain(at)
	if err := r.getDB(ctx).Save(model).Error; err != nil {
		return nil, err
	}
	result := model.ToAyahTafsirDomain()
	return &result, nil
}

func (r *tafsirRepository) DeleteAyahTafsir(ctx context.Context, tafsirID string, surahID, ayahID int) error {
	conds := "tafsir_id = ? AND surah_id = ? AND ayah_id = ?"
	return r.getDB(ctx).Delete(&AyahTafsirModel{}, conds, tafsirID, surahID, ayahID).Error
}

package audio

import (
	"context"
	"github.com/achmdndy/safa-life-api/src/domain/audio"
	"gorm.io/gorm"
)

type ayahAudioRepository struct {
	db *gorm.DB
}

// NewAyahAudioRepository creates a new instance of AyahAudioRepository
func NewAyahAudioRepository(db *gorm.DB) audio.AyahAudioRepository {
	return &ayahAudioRepository{db: db}
}

func (r *ayahAudioRepository) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		type contextKey string
		const txRepoKey contextKey = "tx_repo_audio"
		txCtx := context.WithValue(ctx, txRepoKey, &ayahAudioRepository{db: tx})
		return fn(txCtx)
	})
}

func (r *ayahAudioRepository) getDB(ctx context.Context) *gorm.DB {
	if txRepo, ok := ctx.Value("tx_repo_audio").(*ayahAudioRepository); ok {
		return txRepo.db.WithContext(ctx)
	}
	return r.db.WithContext(ctx)
}

func (r *ayahAudioRepository) Get(ctx context.Context, reciterID string, surahID, ayahID int) (*audio.AyahAudio, error) {
	var model AyahAudioModel
	conds := "reciter_id = ? AND surah_id = ? AND ayah_id = ?"
	if err := r.getDB(ctx).First(&model, conds, reciterID, surahID, ayahID).Error; err != nil {
		return nil, err
	}
	domain := model.ToAyahAudioDomain()
	return &domain, nil
}

func (r *ayahAudioRepository) GetBySurah(ctx context.Context, reciterID string, surahID int) ([]audio.AyahAudio, error) {
	var models []AyahAudioModel
	conds := "reciter_id = ? AND surah_id = ?"
	if err := r.getDB(ctx).Find(&models, conds, reciterID, surahID).Error; err != nil {
		return nil, err
	}
	return ToAyahAudioDomainSlice(models), nil
}

func (r *ayahAudioRepository) Create(ctx context.Context, audio *audio.AyahAudio) (*audio.AyahAudio, error) {
	model := FromAyahAudioDomain(audio)
	if err := r.getDB(ctx).Create(model).Error; err != nil {
		return nil, err
	}
	result := model.ToAyahAudioDomain()
	return &result, nil
}

func (r *ayahAudioRepository) Update(ctx context.Context, audio *audio.AyahAudio) (*audio.AyahAudio, error) {
	model := FromAyahAudioDomain(audio)
	if err := r.getDB(ctx).Save(model).Error; err != nil {
		return nil, err
	}
	result := model.ToAyahAudioDomain()
	return &result, nil
}

func (r *ayahAudioRepository) Delete(ctx context.Context, reciterID string, surahID, ayahID int) error {
	conds := "reciter_id = ? AND surah_id = ? AND ayah_id = ?"
	return r.getDB(ctx).Delete(&AyahAudioModel{}, conds, reciterID, surahID, ayahID).Error
}
package quran

import (
	"context"
	"fmt"
	"strings"

	"github.com/achmdndy/safa-life-api/src/domain/quran"
	"gorm.io/gorm"
)

// quranRepository implements the QuranRepository interface
type quranRepository struct {
	db *gorm.DB
}

// NewQuranRepository creates a new instance of QuranRepository
func NewQuranRepository(db *gorm.DB) quran.QuranRepository {
	return &quranRepository{
		db: db,
	}
}

// WithTransaction executes a function within a database transaction
func (r *quranRepository) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Create a new repository instance with the transaction
		txRepo := &quranRepository{db: tx}

		// Store the transaction repository in context
		// Define a custom type for the context key to avoid collisions
		type contextKey string
		const txRepoKey contextKey = "tx_repo"

		txCtx := context.WithValue(ctx, txRepoKey, txRepo)

		return fn(txCtx)
	})
}

// getDB returns the appropriate database instance (transaction or regular)
func (r *quranRepository) getDB(ctx context.Context) *gorm.DB {
	if txRepo, ok := ctx.Value("tx_repo").(*quranRepository); ok {
		return txRepo.db.WithContext(ctx)
	}
	return r.db.WithContext(ctx)
}

// Surah Repository Methods
func (r *quranRepository) GetAllSurahs(ctx context.Context) ([]quran.Surah, error) {
	var models []SurahModel

	if err := r.getDB(ctx).Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get all surahs: %w", err)
	}

	return ToSurahDomainSlice(models), nil
}

func (r *quranRepository) GetSurahByID(ctx context.Context, id int) (*quran.Surah, error) {
	var model SurahModel

	if err := r.getDB(ctx).Where("surah_id = ?", id).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("surah with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to get surah by ID: %w", err)
	}

	domain := model.ToSurahDomain()
	return &domain, nil
}

func (r *quranRepository) SearchSurahs(ctx context.Context, query string) ([]quran.Surah, error) {
	var models []SurahModel
	searchQuery := "%" + strings.ToLower(query) + "%"

	if err := r.getDB(ctx).Where(
		"LOWER(name_ar) LIKE ? OR LOWER(name_en) LIKE ? OR LOWER(revelation_place) LIKE ?",
		searchQuery, searchQuery, searchQuery,
	).Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to search surahs: %w", err)
	}

	return ToSurahDomainSlice(models), nil
}

func (r *quranRepository) CreateSurah(ctx context.Context, surah *quran.Surah) (*quran.Surah, error) {
	model := FromSurahDomain(surah)

	if err := r.getDB(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create surah: %w", err)
	}

	result := model.ToSurahDomain()
	return &result, nil
}

func (r *quranRepository) UpdateSurah(ctx context.Context, surah *quran.Surah) (*quran.Surah, error) {
	model := FromSurahDomain(surah)

	if err := r.getDB(ctx).Where("surah_id = ?", surah.ID).Updates(model).Error; err != nil {
		return nil, fmt.Errorf("failed to update surah: %w", err)
	}

	// Fetch the updated record
	var updatedModel SurahModel
	if err := r.getDB(ctx).Where("surah_id = ?", surah.ID).First(&updatedModel).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch updated surah: %w", err)
	}

	result := updatedModel.ToSurahDomain()
	return &result, nil
}

func (r *quranRepository) DeleteSurah(ctx context.Context, id int) error {
	if err := r.getDB(ctx).Where("surah_id = ?", id).Delete(&SurahModel{}).Error; err != nil {
		return fmt.Errorf("failed to delete surah: %w", err)
	}

	return nil
}

// Ayah Repository Methods
func (r *quranRepository) GetAyahsBySurah(ctx context.Context, surahID int) ([]quran.Ayah, error) {
	var models []AyahModel

	if err := r.getDB(ctx).Where("surah_id = ?", surahID).Order("ayah_id ASC").Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get ayahs by surah: %w", err)
	}

	return ToAyahDomainSlice(models), nil
}

func (r *quranRepository) GetAyahByID(ctx context.Context, surahID, ayahID int) (*quran.Ayah, error) {
	var model AyahModel

	if err := r.getDB(ctx).Where("surah_id = ? AND ayah_id = ?", surahID, ayahID).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("ayah %d:%d not found", surahID, ayahID)
		}
		return nil, fmt.Errorf("failed to get ayah by ID: %w", err)
	}

	domain := model.ToAyahDomain()
	return &domain, nil
}

func (r *quranRepository) GetAyahsByPage(ctx context.Context, pageNumber int) ([]quran.Ayah, error) {
	var models []AyahModel

	if err := r.getDB(ctx).Where("page_number = ?", pageNumber).Order("surah_id ASC, ayah_id ASC").Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get ayahs by page: %w", err)
	}

	return ToAyahDomainSlice(models), nil
}

func (r *quranRepository) GetAyahsByJuz(ctx context.Context, juzNumber int) ([]quran.Ayah, error) {
	var models []AyahModel

	if err := r.getDB(ctx).Where("juz_number = ?", juzNumber).Order("surah_id ASC, ayah_id ASC").Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get ayahs by juz: %w", err)
	}

	return ToAyahDomainSlice(models), nil
}

func (r *quranRepository) SearchAyahs(ctx context.Context, query string) ([]quran.Ayah, error) {
	var models []AyahModel
	searchQuery := "%" + strings.ToLower(query) + "%"

	if err := r.getDB(ctx).Where("LOWER(text) LIKE ?", searchQuery).Order("surah_id ASC, ayah_id ASC").Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to search ayahs: %w", err)
	}

	return ToAyahDomainSlice(models), nil
}

func (r *quranRepository) CreateAyah(ctx context.Context, ayah *quran.Ayah) (*quran.Ayah, error) {
	model := FromAyahDomain(ayah)

	if err := r.getDB(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create ayah: %w", err)
	}

	result := model.ToAyahDomain()
	return &result, nil
}

func (r *quranRepository) UpdateAyah(ctx context.Context, ayah *quran.Ayah) (*quran.Ayah, error) {
	model := FromAyahDomain(ayah)

	if err := r.getDB(ctx).Where("surah_id = ? AND ayah_id = ?", ayah.SurahID, ayah.AyahID).Updates(model).Error; err != nil {
		return nil, fmt.Errorf("failed to update ayah: %w", err)
	}

	// Fetch the updated record
	var updatedModel AyahModel
	if err := r.getDB(ctx).Where("surah_id = ? AND ayah_id = ?", ayah.SurahID, ayah.AyahID).First(&updatedModel).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch updated ayah: %w", err)
	}

	result := updatedModel.ToAyahDomain()
	return &result, nil
}

func (r *quranRepository) DeleteAyah(ctx context.Context, surahID, ayahID int) error {
	if err := r.getDB(ctx).Where("surah_id = ? AND ayah_id = ?", surahID, ayahID).Delete(&AyahModel{}).Error; err != nil {
		return fmt.Errorf("failed to delete ayah: %w", err)
	}

	return nil
}

func (r *quranRepository) GetAyahsByRange(ctx context.Context, startSurah, startAyah, endSurah, endAyah int) ([]quran.Ayah, error) {
	var models []AyahModel

	query := r.getDB(ctx).Where(
		"(surah_id > ? OR (surah_id = ? AND ayah_id >= ?)) AND (surah_id < ? OR (surah_id = ? AND ayah_id <= ?))",
		startSurah, startSurah, startAyah, endSurah, endSurah, endAyah,
	).Order("surah_id ASC, ayah_id ASC")

	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get ayahs by range: %w", err)
	}

	return ToAyahDomainSlice(models), nil
}

// Juz Repository Methods
func (r *quranRepository) GetAllJuz(ctx context.Context) ([]quran.Juz, error) {
	var models []JuzModel

	if err := r.getDB(ctx).Order("juz_id ASC").Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get all juz: %w", err)
	}

	return ToJuzDomainSlice(models), nil
}

func (r *quranRepository) GetJuzByID(ctx context.Context, id int) (*quran.Juz, error) {
	var model JuzModel

	if err := r.getDB(ctx).Where("juz_id = ?", id).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("juz with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to get juz by ID: %w", err)
	}

	domain := model.ToJuzDomain()
	return &domain, nil
}

func (r *quranRepository) CreateJuz(ctx context.Context, juz *quran.Juz) (*quran.Juz, error) {
	model := FromJuzDomain(juz)

	if err := r.getDB(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create juz: %w", err)
	}

	result := model.ToJuzDomain()
	return &result, nil
}

func (r *quranRepository) UpdateJuz(ctx context.Context, juz *quran.Juz) (*quran.Juz, error) {
	model := FromJuzDomain(juz)

	if err := r.getDB(ctx).Where("juz_id = ?", juz.ID).Updates(model).Error; err != nil {
		return nil, fmt.Errorf("failed to update juz: %w", err)
	}

	// Fetch the updated record
	var updatedModel JuzModel
	if err := r.getDB(ctx).Where("juz_id = ?", juz.ID).First(&updatedModel).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch updated juz: %w", err)
	}

	result := updatedModel.ToJuzDomain()
	return &result, nil
}

func (r *quranRepository) DeleteJuz(ctx context.Context, id int) error {
	if err := r.getDB(ctx).Where("juz_id = ?", id).Delete(&JuzModel{}).Error; err != nil {
		return fmt.Errorf("failed to delete juz: %w", err)
	}

	return nil
}

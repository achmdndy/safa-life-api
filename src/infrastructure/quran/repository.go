package quran

import (
	"context"
	"fmt"

	"github.com/safalife/core-api/src/domain/core"
	"github.com/safalife/core-api/src/domain/quran"
	infraCore "github.com/safalife/core-api/src/infrastructure/core"
	"gorm.io/gorm"
)

// SurahRepositoryImpl implements the SurahRepositoryInterface using GORM
type SurahRepositoryImpl struct {
	db     *gorm.DB
	mapper *Mapper
}

// NewSurahRepository creates a new instance of SurahRepositoryImpl
func NewSurahRepository(db *gorm.DB, mapper *Mapper) quran.SurahRepositoryInterface {
	return &SurahRepositoryImpl{
		db:     db,
		mapper: mapper,
	}
}

// GetById retrieves a Surah by its ID
func (r *SurahRepositoryImpl) GetById(ctx context.Context, id core.UUID) (*quran.Surah, error) {
	googleUUID := infraCore.ToGoogleUUID(id)
	var model SurahModel
	if err := r.db.WithContext(ctx).Where("id = ?", googleUUID).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrSurahNotFound
		}
		return nil, fmt.Errorf("failed to get surah by id: %w", err)
	}
	return r.mapper.SurahModelToEntity(&model), nil
}

// GetByIdWithAyahs retrieves a Surah by its ID with all related Ayahs
func (r *SurahRepositoryImpl) GetByIdWithAyahs(ctx context.Context, id core.UUID) (*quran.SurahWithAyahs, error) {
	googleUUID := infraCore.ToGoogleUUID(id)
	var model SurahModel
	if err := r.db.WithContext(ctx).Preload("Ayahs").Where("id = ?", googleUUID).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrSurahNotFound
		}
		return nil, fmt.Errorf("failed to get surah by id with ayahs: %w", err)
	}
	return r.mapper.SurahModelToEntityWithAyahs(&model), nil
}

// GetByNumber retrieves a Surah by its number
func (r *SurahRepositoryImpl) GetByNumber(ctx context.Context, number int) (*quran.Surah, error) {
	var model SurahModel
	if err := r.db.WithContext(ctx).Where("revelation_order = ?", number).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrSurahNotFound
		}
		return nil, fmt.Errorf("failed to get surah by number: %w", err)
	}
	return r.mapper.SurahModelToEntity(&model), nil
}

// GetByNumberWithAyahs retrieves a Surah by its number with all related Ayahs
func (r *SurahRepositoryImpl) GetByNumberWithAyahs(ctx context.Context, number int) (*quran.SurahWithAyahs, error) {
	var model SurahModel
	if err := r.db.WithContext(ctx).Preload("Ayahs").Where("revelation_order = ?", number).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrSurahNotFound
		}
		return nil, fmt.Errorf("failed to get surah by number with ayahs: %w", err)
	}
	return r.mapper.SurahModelToEntityWithAyahs(&model), nil
}

// GetAll retrieves all Surahs with pagination
func (r *SurahRepositoryImpl) GetAll(ctx context.Context, limit, offset int) ([]*quran.Surah, error) {
	var models []SurahModel
	query := r.db.WithContext(ctx).Order("revelation_order ASC")

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get all surahs: %w", err)
	}

	entities := make([]*quran.Surah, len(models))
	for i, model := range models {
		entities[i] = r.mapper.SurahModelToEntity(&model)
	}
	return entities, nil
}

// GetAllWithAyahs retrieves all Surahs with their Ayahs using pagination
func (r *SurahRepositoryImpl) GetAllWithAyahs(ctx context.Context, limit, offset int) ([]*quran.SurahWithAyahs, error) {
	var models []SurahModel
	query := r.db.WithContext(ctx).Preload("Ayahs").Order("revelation_order ASC")

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get all surahs with ayahs: %w", err)
	}

	entities := make([]*quran.SurahWithAyahs, len(models))
	for i, model := range models {
		entities[i] = r.mapper.SurahModelToEntityWithAyahs(&model)
	}
	return entities, nil
}

// GetByRevelationPlace retrieves Surahs by revelation place with pagination
func (r *SurahRepositoryImpl) GetByRevelationPlace(ctx context.Context, place string, limit, offset int) ([]*quran.Surah, error) {
	var models []SurahModel
	query := r.db.WithContext(ctx).Where("revelation_place = ?", place).Order("revelation_order ASC")

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get surahs by revelation place: %w", err)
	}

	entities := make([]*quran.Surah, len(models))
	for i, model := range models {
		entities[i] = r.mapper.SurahModelToEntity(&model)
	}
	return entities, nil
}

// Create creates a new Surah
func (r *SurahRepositoryImpl) Create(ctx context.Context, surah *quran.Surah) (*quran.Surah, error) {
	model := r.mapper.SurahEntityToModel(surah)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create surah: %w", err)
	}
	return r.mapper.SurahModelToEntity(model), nil
}

// Update updates an existing Surah
func (r *SurahRepositoryImpl) Update(ctx context.Context, surah *quran.Surah) (*quran.Surah, error) {
	model := r.mapper.SurahEntityToModel(surah)
	if err := r.db.WithContext(ctx).Where("id = ?", surah.ID).Updates(model).Error; err != nil {
		return nil, fmt.Errorf("failed to update surah: %w", err)
	}
	return r.mapper.SurahModelToEntity(model), nil
}

// Delete soft deletes a Surah
func (r *SurahRepositoryImpl) Delete(ctx context.Context, id core.UUID) error {
	googleUUID := infraCore.ToGoogleUUID(id)
	if err := r.db.WithContext(ctx).Delete(&SurahModel{}, googleUUID).Error; err != nil {
		return fmt.Errorf("failed to delete surah: %w", err)
	}
	return nil
}

// Count returns the total number of Surahs
func (r *SurahRepositoryImpl) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&SurahModel{}).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count surahs: %w", err)
	}
	return count, nil
}

// CountByRevelationPlace returns the number of Surahs by revelation place
func (r *SurahRepositoryImpl) CountByRevelationPlace(ctx context.Context, place string) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&SurahModel{}).Where("revelation_place = ?", place).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count surahs by revelation place: %w", err)
	}
	return count, nil
}

// AyahRepositoryImpl implements the AyahRepositoryInterface using GORM
type AyahRepositoryImpl struct {
	db     *gorm.DB
	mapper *Mapper
}

// NewAyahRepository creates a new instance of AyahRepositoryImpl
func NewAyahRepository(db *gorm.DB, mapper *Mapper) quran.AyahRepositoryInterface {
	return &AyahRepositoryImpl{
		db:     db,
		mapper: mapper,
	}
}

// GetById retrieves an Ayah by its ID
func (r *AyahRepositoryImpl) GetById(ctx context.Context, id core.UUID) (*quran.Ayah, error) {
	googleUUID := infraCore.ToGoogleUUID(id)
	var model AyahModel
	if err := r.db.WithContext(ctx).Where("id = ?", googleUUID).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrAyahNotFound
		}
		return nil, fmt.Errorf("failed to get ayah by id: %w", err)
	}
	return r.mapper.AyahModelToEntity(&model), nil
}

// GetBySurahId retrieves Ayahs by Surah ID with pagination
func (r *AyahRepositoryImpl) GetBySurahId(ctx context.Context, surahId core.UUID, limit, offset int) ([]*quran.Ayah, error) {
	googleUUID := infraCore.ToGoogleUUID(surahId)
	var models []AyahModel
	query := r.db.WithContext(ctx).Where("surah_id = ?", googleUUID).Order("created_at ASC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get ayahs by surah id: %w", err)
	}

	ayahs := make([]*quran.Ayah, len(models))
	for i, model := range models {
		ayahs[i] = r.mapper.AyahModelToEntity(&model)
	}
	return ayahs, nil
}

// GetByIdWithSurah retrieves an Ayah by its ID with related Surah
func (r *AyahRepositoryImpl) GetByIdWithSurah(ctx context.Context, id core.UUID) (*quran.AyahWithSurah, error) {
	googleUUID := infraCore.ToGoogleUUID(id)
	var model AyahModel
	if err := r.db.WithContext(ctx).Preload("Surah").Where("id = ?", googleUUID).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrAyahNotFound
		}
		return nil, fmt.Errorf("failed to get ayah by id with surah: %w", err)
	}
	return r.mapper.AyahModelToEntityWithSurah(&model), nil
}

// GetBySurahIdWithSurah retrieves Ayahs by Surah ID with related Surah and pagination
func (r *AyahRepositoryImpl) GetBySurahIdWithSurah(ctx context.Context, surahId core.UUID, limit, offset int) ([]*quran.AyahWithSurah, error) {
	googleUUID := infraCore.ToGoogleUUID(surahId)
	var models []AyahModel
	query := r.db.WithContext(ctx).Preload("Surah").Where("surah_id = ?", googleUUID).Order("created_at ASC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get ayahs by surah id with surah: %w", err)
	}

	ayahs := make([]*quran.AyahWithSurah, len(models))
	for i, model := range models {
		ayahs[i] = r.mapper.AyahModelToEntityWithSurah(&model)
	}
	return ayahs, nil
}

// GetAllWithSurah retrieves all Ayahs with their Surah using pagination
func (r *AyahRepositoryImpl) GetAllWithSurah(ctx context.Context, limit, offset int) ([]*quran.AyahWithSurah, error) {
	var models []AyahModel
	query := r.db.WithContext(ctx).Preload("Surah").Order("surah_id ASC, created_at ASC")

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get all ayahs with surah: %w", err)
	}

	entities := make([]*quran.AyahWithSurah, len(models))
	for i, model := range models {
		entities[i] = r.mapper.AyahModelToEntityWithSurah(&model)
	}
	return entities, nil
}

// GetByJuzNumber retrieves all Ayahs in a Juz with pagination
func (r *AyahRepositoryImpl) GetByJuzNumber(ctx context.Context, juzNumber int, limit, offset int) ([]*quran.Ayah, error) {
	var models []AyahModel
	query := r.db.WithContext(ctx).Where("juz_number = ?", juzNumber).Order("created_at ASC")

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get ayahs by juz: %w", err)
	}

	entities := make([]*quran.Ayah, len(models))
	for i, model := range models {
		entities[i] = r.mapper.AyahModelToEntity(&model)
	}
	return entities, nil
}

// GetByPageNumber retrieves all Ayahs on a page with pagination
func (r *AyahRepositoryImpl) GetByPageNumber(ctx context.Context, pageNumber int, limit, offset int) ([]*quran.Ayah, error) {
	var models []AyahModel
	query := r.db.WithContext(ctx).Where("page_number = ?", pageNumber).Order("created_at ASC")

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get ayahs by page: %w", err)
	}

	entities := make([]*quran.Ayah, len(models))
	for i, model := range models {
		entities[i] = r.mapper.AyahModelToEntity(&model)
	}
	return entities, nil
}

// GetByHizbNumber retrieves all Ayahs in a Hizb with pagination
func (r *AyahRepositoryImpl) GetByHizbNumber(ctx context.Context, hizbNumber int, limit, offset int) ([]*quran.Ayah, error) {
	var models []AyahModel
	query := r.db.WithContext(ctx).Where("hizb_number = ?", hizbNumber).Order("created_at ASC")

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get ayahs by hizb: %w", err)
	}

	entities := make([]*quran.Ayah, len(models))
	for i, model := range models {
		entities[i] = r.mapper.AyahModelToEntity(&model)
	}
	return entities, nil
}

// GetByManzilNumber retrieves all Ayahs in a Manzil with pagination
func (r *AyahRepositoryImpl) GetByManzilNumber(ctx context.Context, manzilNumber int, limit, offset int) ([]*quran.Ayah, error) {
	var models []AyahModel
	query := r.db.WithContext(ctx).Where("manzil_number = ?", manzilNumber).Order("created_at ASC")

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get ayahs by manzil: %w", err)
	}

	entities := make([]*quran.Ayah, len(models))
	for i, model := range models {
		entities[i] = r.mapper.AyahModelToEntity(&model)
	}
	return entities, nil
}

// GetAll retrieves all Ayahs with pagination
func (r *AyahRepositoryImpl) GetAll(ctx context.Context, limit, offset int) ([]*quran.Ayah, error) {
	var models []AyahModel
	query := r.db.WithContext(ctx).Order("surah_id ASC, created_at ASC")

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get all ayahs: %w", err)
	}

	entities := make([]*quran.Ayah, len(models))
	for i, model := range models {
		entities[i] = r.mapper.AyahModelToEntity(&model)
	}
	return entities, nil
}

// Create creates a new Ayah
func (r *AyahRepositoryImpl) Create(ctx context.Context, ayah *quran.Ayah) (*quran.Ayah, error) {
	model := r.mapper.AyahEntityToModel(ayah)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create ayah: %w", err)
	}
	return r.mapper.AyahModelToEntity(model), nil
}

// Update updates an existing Ayah
func (r *AyahRepositoryImpl) Update(ctx context.Context, ayah *quran.Ayah) (*quran.Ayah, error) {
	model := r.mapper.AyahEntityToModel(ayah)
	if err := r.db.WithContext(ctx).Where("id = ?", ayah.ID).Updates(model).Error; err != nil {
		return nil, fmt.Errorf("failed to update ayah: %w", err)
	}
	return r.mapper.AyahModelToEntity(model), nil
}

// Delete soft deletes an Ayah
func (r *AyahRepositoryImpl) Delete(ctx context.Context, id core.UUID) error {
	googleUUID := infraCore.ToGoogleUUID(id)
	if err := r.db.WithContext(ctx).Delete(&AyahModel{}, googleUUID).Error; err != nil {
		return fmt.Errorf("failed to delete ayah: %w", err)
	}
	return nil
}

// Count returns the total number of Ayahs
func (r *AyahRepositoryImpl) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&AyahModel{}).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count ayahs: %w", err)
	}
	return count, nil
}

// CountBySurahId returns the number of Ayahs in a Surah
func (r *AyahRepositoryImpl) CountBySurahId(ctx context.Context, surahId core.UUID) (int64, error) {
	googleUUID := infraCore.ToGoogleUUID(surahId)
	var count int64
	if err := r.db.WithContext(ctx).Model(&AyahModel{}).Where("surah_id = ?", googleUUID).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count ayahs by surah: %w", err)
	}
	return count, nil
}

// CountByJuzNumber returns the number of Ayahs in a Juz
func (r *AyahRepositoryImpl) CountByJuzNumber(ctx context.Context, juzNumber int) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&AyahModel{}).Where("juz_number = ?", juzNumber).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count ayahs by juz: %w", err)
	}
	return count, nil
}

// CountByPageNumber returns the number of Ayahs on a page
func (r *AyahRepositoryImpl) CountByPageNumber(ctx context.Context, pageNumber int) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&AyahModel{}).Where("page_number = ?", pageNumber).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count ayahs by page: %w", err)
	}
	return count, nil
}

// JuzRepositoryImpl implements the JuzRepositoryInterface using GORM
type JuzRepositoryImpl struct {
	db     *gorm.DB
	mapper *Mapper
}

// NewJuzRepository creates a new instance of JuzRepositoryImpl
func NewJuzRepository(db *gorm.DB, mapper *Mapper) quran.JuzRepositoryInterface {
	return &JuzRepositoryImpl{
		db:     db,
		mapper: mapper,
	}
}

// GetById retrieves a Juz by its ID
func (r *JuzRepositoryImpl) GetById(ctx context.Context, id core.UUID) (*quran.Juz, error) {
	googleUUID := infraCore.ToGoogleUUID(id)
	var model JuzModel
	if err := r.db.WithContext(ctx).Where("id = ?", googleUUID).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrJuzNotFound
		}
		return nil, fmt.Errorf("failed to get juz by id: %w", err)
	}
	return r.mapper.JuzModelToEntity(&model), nil
}

// GetByNumber retrieves a Juz by its number
// Since JuzModel doesn't have a Number field, we find it by looking for Ayahs with the specified juz_number
func (r *JuzRepositoryImpl) GetByNumber(ctx context.Context, number int) (*quran.Juz, error) {
	// Find the Juz that contains Ayahs with the specified juz_number
	// We do this by finding a Juz whose start_ayah_id has the specified juz_number
	var model JuzModel
	if err := r.db.WithContext(ctx).
		Joins("JOIN ayahs ON ayahs.id = juz.start_ayah_id").
		Where("ayahs.juz_number = ?", number).
		First(&model).Error; err != nil {
		return nil, fmt.Errorf("failed to get juz by number %d: %w", number, err)
	}

	return r.mapper.JuzModelToEntity(&model), nil
}

// GetAll retrieves all Juzs with pagination
func (r *JuzRepositoryImpl) GetAll(ctx context.Context, limit, offset int) ([]*quran.Juz, error) {
	var models []JuzModel
	query := r.db.WithContext(ctx).Order("created_at ASC")

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get all juzs: %w", err)
	}

	entities := make([]*quran.Juz, len(models))
	for i, model := range models {
		entities[i] = r.mapper.JuzModelToEntity(&model)
	}
	return entities, nil
}

// Create creates a new Juz
func (r *JuzRepositoryImpl) Create(ctx context.Context, juz *quran.Juz) (*quran.Juz, error) {
	model := r.mapper.JuzEntityToModel(juz)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create juz: %w", err)
	}
	return r.mapper.JuzModelToEntity(model), nil
}

// Update updates an existing Juz
func (r *JuzRepositoryImpl) Update(ctx context.Context, juz *quran.Juz) (*quran.Juz, error) {
	model := r.mapper.JuzEntityToModel(juz)
	if err := r.db.WithContext(ctx).Where("id = ?", juz.ID).Updates(model).Error; err != nil {
		return nil, fmt.Errorf("failed to update juz: %w", err)
	}
	return r.mapper.JuzModelToEntity(model), nil
}

// Delete soft deletes a Juz
func (r *JuzRepositoryImpl) Delete(ctx context.Context, id core.UUID) error {
	googleUUID := infraCore.ToGoogleUUID(id)
	if err := r.db.WithContext(ctx).Delete(&JuzModel{}, googleUUID).Error; err != nil {
		return fmt.Errorf("failed to delete juz: %w", err)
	}
	return nil
}

// GetByIdWithRelations retrieves a Juz by its ID with all related entities
func (r *JuzRepositoryImpl) GetByIdWithRelations(ctx context.Context, id core.UUID) (*quran.JuzWithRelations, error) {
	googleUUID := infraCore.ToGoogleUUID(id)
	var model JuzModel
	if err := r.db.WithContext(ctx).
		Preload("StartSurah").
		Preload("EndSurah").
		Preload("StartAyah").
		Preload("EndAyah").
		Where("id = ?", googleUUID).
		First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrJuzNotFound
		}
		return nil, fmt.Errorf("failed to get juz by id with relations: %w", err)
	}
	return r.mapper.JuzModelToEntityWithRelations(&model), nil
}

// GetByNumberWithRelations retrieves a Juz by its number with all related entities
func (r *JuzRepositoryImpl) GetByNumberWithRelations(ctx context.Context, number int) (*quran.JuzWithRelations, error) {
	// Find the Juz that contains Ayahs with the specified juz_number with relations
	// We do this by finding a Juz whose start_ayah_id has the specified juz_number
	var model JuzModel
	if err := r.db.WithContext(ctx).
		Preload("StartSurah").
		Preload("EndSurah").
		Preload("StartAyah").
		Preload("EndAyah").
		Joins("JOIN ayahs ON ayahs.id = juz.start_ayah_id").
		Where("ayahs.juz_number = ?", number).
		First(&model).Error; err != nil {
		return nil, fmt.Errorf("failed to get juz by number %d with relations: %w", number, err)
	}

	return r.mapper.JuzModelToEntityWithRelations(&model), nil
}

// GetAllWithRelations retrieves all Juz with their relations using pagination
func (r *JuzRepositoryImpl) GetAllWithRelations(ctx context.Context, limit, offset int) ([]*quran.JuzWithRelations, error) {
	var models []JuzModel
	query := r.db.WithContext(ctx).
		Preload("StartSurah").
		Preload("EndSurah").
		Preload("StartAyah").
		Preload("EndAyah").
		Order("created_at ASC")

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get all juz with relations: %w", err)
	}

	entities := make([]*quran.JuzWithRelations, len(models))
	for i, model := range models {
		entities[i] = r.mapper.JuzModelToEntityWithRelations(&model)
	}
	return entities, nil
}

func (r *JuzRepositoryImpl) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&JuzModel{}).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count juzs: %w", err)
	}
	return count, nil
}

// TranslationEditionRepositoryImpl implements TranslationEditionRepositoryInterface using GORM
type TranslationEditionRepositoryImpl struct {
	db     *gorm.DB
	mapper *Mapper
}

// NewTranslationEditionRepository creates a new instance of TranslationEditionRepositoryImpl
func NewTranslationEditionRepository(db *gorm.DB, mapper *Mapper) quran.TranslationEditionRepositoryInterface {
	return &TranslationEditionRepositoryImpl{
		db:     db,
		mapper: mapper,
	}
}

// GetById retrieves a TranslationEdition by its ID
func (r *TranslationEditionRepositoryImpl) GetById(ctx context.Context, id core.UUID) (*quran.TranslationEdition, error) {
	googleUUID := infraCore.ToGoogleUUID(id)
	var model TranslationEditionModel
	if err := r.db.WithContext(ctx).Where("id = ?", googleUUID).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrTranslationEditionNotFound
		}
		return nil, fmt.Errorf("failed to get translation edition by id: %w", err)
	}
	return r.mapper.TranslationEditionModelToEntity(&model), nil
}

// GetAll retrieves all TranslationEditions with pagination
func (r *TranslationEditionRepositoryImpl) GetAll(ctx context.Context, limit, offset int) ([]*quran.TranslationEdition, error) {
	var models []TranslationEditionModel
	query := r.db.WithContext(ctx).Order("created_at ASC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get translation editions: %w", err)
	}
	outs := make([]*quran.TranslationEdition, len(models))
	for i := range models {
		outs[i] = r.mapper.TranslationEditionModelToEntity(&models[i])
	}
	return outs, nil
}

// GetByLanguage retrieves TranslationEditions filtered by language with pagination
func (r *TranslationEditionRepositoryImpl) GetByLanguage(ctx context.Context, language string, limit, offset int) ([]*quran.TranslationEdition, error) {
	var models []TranslationEditionModel
	query := r.db.WithContext(ctx).Where("language = ?", language).Order("created_at ASC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get translation editions by language: %w", err)
	}
	outs := make([]*quran.TranslationEdition, len(models))
	for i := range models {
		outs[i] = r.mapper.TranslationEditionModelToEntity(&models[i])
	}
	return outs, nil
}

// Create creates a new TranslationEdition
func (r *TranslationEditionRepositoryImpl) Create(ctx context.Context, edition *quran.TranslationEdition) (*quran.TranslationEdition, error) {
	model := r.mapper.TranslationEditionEntityToModel(edition)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create translation edition: %w", err)
	}
	return r.mapper.TranslationEditionModelToEntity(model), nil
}

// Update updates an existing TranslationEdition
func (r *TranslationEditionRepositoryImpl) Update(ctx context.Context, edition *quran.TranslationEdition) (*quran.TranslationEdition, error) {
	model := r.mapper.TranslationEditionEntityToModel(edition)
	if err := r.db.WithContext(ctx).Where("id = ?", infraCore.ToGoogleUUID(edition.ID)).Updates(model).Error; err != nil {
		return nil, fmt.Errorf("failed to update translation edition: %w", err)
	}
	return r.mapper.TranslationEditionModelToEntity(model), nil
}

// Delete soft deletes a TranslationEdition
func (r *TranslationEditionRepositoryImpl) Delete(ctx context.Context, id core.UUID) error {
	googleUUID := infraCore.ToGoogleUUID(id)
	if err := r.db.WithContext(ctx).Delete(&TranslationEditionModel{}, googleUUID).Error; err != nil {
		return fmt.Errorf("failed to delete translation edition: %w", err)
	}
	return nil
}

// Count returns the number of TranslationEdition records
func (r *TranslationEditionRepositoryImpl) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&TranslationEditionModel{}).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count translation editions: %w", err)
	}
	return count, nil
}

// CountByLanguage returns the number of TranslationEdition records by language
func (r *TranslationEditionRepositoryImpl) CountByLanguage(ctx context.Context, language string) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&TranslationEditionModel{}).Where("language = ?", language).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count translation editions by language: %w", err)
	}
	return count, nil
}

// AyahTranslationRepositoryImpl implements AyahTranslationRepositoryInterface using GORM
type AyahTranslationRepositoryImpl struct {
	db     *gorm.DB
	mapper *Mapper
}

// NewAyahTranslationRepository creates a new instance of AyahTranslationRepositoryImpl
func NewAyahTranslationRepository(db *gorm.DB, mapper *Mapper) quran.AyahTranslationRepositoryInterface {
	return &AyahTranslationRepositoryImpl{
		db:     db,
		mapper: mapper,
	}
}

// GetById retrieves an AyahTranslation by its ID
func (r *AyahTranslationRepositoryImpl) GetById(ctx context.Context, id core.UUID) (*quran.AyahTranslation, error) {
	googleUUID := infraCore.ToGoogleUUID(id)
	var model AyahTranslationModel
	if err := r.db.WithContext(ctx).Where("id = ?", googleUUID).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrAyahTranslationNotFound
		}
		return nil, fmt.Errorf("failed to get ayah translation by id: %w", err)
	}
	return r.mapper.AyahTranslationModelToEntity(&model), nil
}

// GetByAyahAndEdition retrieves an AyahTranslation by ayah and edition IDs
func (r *AyahTranslationRepositoryImpl) GetByAyahAndEdition(ctx context.Context, ayahId core.UUID, editionId core.UUID) (*quran.AyahTranslation, error) {
	googleAyah := infraCore.ToGoogleUUID(ayahId)
	googleEdition := infraCore.ToGoogleUUID(editionId)
	var model AyahTranslationModel
	if err := r.db.WithContext(ctx).
		Where("ayah_id = ? AND translation_edition_id = ?", googleAyah, googleEdition).
		First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrAyahTranslationNotFound
		}
		return nil, fmt.Errorf("failed to get ayah translation by ayah and edition: %w", err)
	}
	return r.mapper.AyahTranslationModelToEntity(&model), nil
}

// GetBySurahAndEdition retrieves AyahTranslations by surah and edition IDs with pagination
func (r *AyahTranslationRepositoryImpl) GetBySurahAndEdition(ctx context.Context, surahId core.UUID, editionId core.UUID, limit, offset int) ([]*quran.AyahTranslation, error) {
	googleSurah := infraCore.ToGoogleUUID(surahId)
	googleEdition := infraCore.ToGoogleUUID(editionId)
	var models []AyahTranslationModel
	query := r.db.WithContext(ctx).
		Where("surah_id = ? AND translation_edition_id = ?", googleSurah, googleEdition).
		Order("created_at ASC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get ayah translations by surah and edition: %w", err)
	}
	outs := make([]*quran.AyahTranslation, len(models))
	for i := range models {
		outs[i] = r.mapper.AyahTranslationModelToEntity(&models[i])
	}
	return outs, nil
}

// Create creates a new AyahTranslation
func (r *AyahTranslationRepositoryImpl) Create(ctx context.Context, t *quran.AyahTranslation) (*quran.AyahTranslation, error) {
	model := r.mapper.AyahTranslationEntityToModel(t)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create ayah translation: %w", err)
	}
	return r.mapper.AyahTranslationModelToEntity(model), nil
}

// Update updates an existing AyahTranslation
func (r *AyahTranslationRepositoryImpl) Update(ctx context.Context, t *quran.AyahTranslation) (*quran.AyahTranslation, error) {
	model := r.mapper.AyahTranslationEntityToModel(t)
	if err := r.db.WithContext(ctx).Where("id = ?", infraCore.ToGoogleUUID(t.ID)).Updates(model).Error; err != nil {
		return nil, fmt.Errorf("failed to update ayah translation: %w", err)
	}
	return r.mapper.AyahTranslationModelToEntity(model), nil
}

// Delete soft deletes an AyahTranslation
func (r *AyahTranslationRepositoryImpl) Delete(ctx context.Context, id core.UUID) error {
	googleUUID := infraCore.ToGoogleUUID(id)
	if err := r.db.WithContext(ctx).Delete(&AyahTranslationModel{}, googleUUID).Error; err != nil {
		return fmt.Errorf("failed to delete ayah translation: %w", err)
	}
	return nil
}

// Count returns the number of AyahTranslation records
func (r *AyahTranslationRepositoryImpl) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&AyahTranslationModel{}).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count ayah translations: %w", err)
	}
	return count, nil
}

// CountByEdition returns the number of AyahTranslation records for a given edition
func (r *AyahTranslationRepositoryImpl) CountByEdition(ctx context.Context, editionId core.UUID) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&AyahTranslationModel{}).
		Where("translation_edition_id = ?", infraCore.ToGoogleUUID(editionId)).
		Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count ayah translations by edition: %w", err)
	}
	return count, nil
}

// CountBySurahAndEdition returns the number of AyahTranslation records for a given surah and edition
func (r *AyahTranslationRepositoryImpl) CountBySurahAndEdition(ctx context.Context, surahId core.UUID, editionId core.UUID) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&AyahTranslationModel{}).
		Where("surah_id = ? AND translation_edition_id = ?", infraCore.ToGoogleUUID(surahId), infraCore.ToGoogleUUID(editionId)).
		Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count ayah translations by surah and edition: %w", err)
	}
	return count, nil
}

// ReciterRepositoryImpl implements ReciterRepositoryInterface using GORM
type ReciterRepositoryImpl struct {
	db     *gorm.DB
	mapper *Mapper
}

// NewReciterRepository creates a new instance of ReciterRepositoryImpl
func NewReciterRepository(db *gorm.DB, mapper *Mapper) quran.ReciterRepositoryInterface {
	return &ReciterRepositoryImpl{
		db:     db,
		mapper: mapper,
	}
}

// GetById retrieves a Reciter by its ID
func (r *ReciterRepositoryImpl) GetById(ctx context.Context, id core.UUID) (*quran.Reciter, error) {
	googleUUID := infraCore.ToGoogleUUID(id)
	var model ReciterModel
	if err := r.db.WithContext(ctx).Where("id = ?", googleUUID).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrReciterNotFound
		}
		return nil, fmt.Errorf("failed to get reciter by id: %w", err)
	}
	return r.mapper.ReciterModelToEntity(&model), nil
}

// GetByName retrieves a Reciter by its name
func (r *ReciterRepositoryImpl) GetByName(ctx context.Context, name string) (*quran.Reciter, error) {
	var model ReciterModel
	if err := r.db.WithContext(ctx).Where("name = ?", name).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrReciterNotFound
		}
		return nil, fmt.Errorf("failed to get reciter by name: %w", err)
	}
	return r.mapper.ReciterModelToEntity(&model), nil
}

// GetAll retrieves all Reciters with pagination
func (r *ReciterRepositoryImpl) GetAll(ctx context.Context, limit, offset int) ([]*quran.Reciter, error) {
	var models []ReciterModel
	query := r.db.WithContext(ctx).Order("created_at ASC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get reciters: %w", err)
	}
	outs := make([]*quran.Reciter, len(models))
	for i := range models {
		outs[i] = r.mapper.ReciterModelToEntity(&models[i])
	}
	return outs, nil
}

// Create creates a new Reciter
func (r *ReciterRepositoryImpl) Create(ctx context.Context, reciter *quran.Reciter) (*quran.Reciter, error) {
	model := r.mapper.ReciterEntityToModel(reciter)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create reciter: %w", err)
	}
	return r.mapper.ReciterModelToEntity(model), nil
}

// Update updates an existing Reciter
func (r *ReciterRepositoryImpl) Update(ctx context.Context, reciter *quran.Reciter) (*quran.Reciter, error) {
	model := r.mapper.ReciterEntityToModel(reciter)
	if err := r.db.WithContext(ctx).Where("id = ?", infraCore.ToGoogleUUID(reciter.ID)).Updates(model).Error; err != nil {
		return nil, fmt.Errorf("failed to update reciter: %w", err)
	}
	return r.mapper.ReciterModelToEntity(model), nil
}

// Delete soft deletes a Reciter
func (r *ReciterRepositoryImpl) Delete(ctx context.Context, id core.UUID) error {
	googleUUID := infraCore.ToGoogleUUID(id)
	if err := r.db.WithContext(ctx).Delete(&ReciterModel{}, googleUUID).Error; err != nil {
		return fmt.Errorf("failed to delete reciter: %w", err)
	}
	return nil
}

// Count returns the number of Reciter records
func (r *ReciterRepositoryImpl) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&ReciterModel{}).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count reciters: %w", err)
	}
	return count, nil
}

// AyahAudioFileRepositoryImpl implements AyahAudioFileRepositoryInterface using GORM
type AyahAudioFileRepositoryImpl struct {
	db     *gorm.DB
	mapper *Mapper
}

// NewAyahAudioFileRepository creates a new instance of AyahAudioFileRepositoryImpl
func NewAyahAudioFileRepository(db *gorm.DB, mapper *Mapper) quran.AyahAudioFileRepositoryInterface {
	return &AyahAudioFileRepositoryImpl{
		db:     db,
		mapper: mapper,
	}
}

// GetById retrieves an AyahAudioFile by its ID
func (r *AyahAudioFileRepositoryImpl) GetById(ctx context.Context, id core.UUID) (*quran.AyahAudioFile, error) {
	googleUUID := infraCore.ToGoogleUUID(id)
	var model AyahAudioFileModel
	if err := r.db.WithContext(ctx).Where("id = ?", googleUUID).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrAyahAudioFileNotFound
		}
		return nil, fmt.Errorf("failed to get ayah audio file by id: %w", err)
	}
	return r.mapper.AyahAudioFileModelToEntity(&model), nil
}

// GetByAyahAndReciter retrieves an AyahAudioFile by ayah and reciter IDs
func (r *AyahAudioFileRepositoryImpl) GetByAyahAndReciter(ctx context.Context, ayahId core.UUID, reciterId core.UUID) (*quran.AyahAudioFile, error) {
	var model AyahAudioFileModel
	if err := r.db.WithContext(ctx).
		Where("ayah_id = ? AND reciter_id = ?", infraCore.ToGoogleUUID(ayahId), infraCore.ToGoogleUUID(reciterId)).
		First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrAyahAudioFileNotFound
		}
		return nil, fmt.Errorf("failed to get ayah audio file by ayah and reciter: %w", err)
	}
	return r.mapper.AyahAudioFileModelToEntity(&model), nil
}

// GetBySurahAndReciter retrieves AyahAudioFiles by surah and reciter IDs with pagination
func (r *AyahAudioFileRepositoryImpl) GetBySurahAndReciter(ctx context.Context, surahId core.UUID, reciterId core.UUID, limit, offset int) ([]*quran.AyahAudioFile, error) {
	var models []AyahAudioFileModel
	query := r.db.WithContext(ctx).
		Where("surah_id = ? AND reciter_id = ?", infraCore.ToGoogleUUID(surahId), infraCore.ToGoogleUUID(reciterId)).
		Order("created_at ASC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get ayah audio files by surah and reciter: %w", err)
	}
	outs := make([]*quran.AyahAudioFile, len(models))
	for i := range models {
		outs[i] = r.mapper.AyahAudioFileModelToEntity(&models[i])
	}
	return outs, nil
}

// Create creates a new AyahAudioFile
func (r *AyahAudioFileRepositoryImpl) Create(ctx context.Context, audio *quran.AyahAudioFile) (*quran.AyahAudioFile, error) {
	model := r.mapper.AyahAudioFileEntityToModel(audio)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create ayah audio file: %w", err)
	}
	return r.mapper.AyahAudioFileModelToEntity(model), nil
}

// Update updates an existing AyahAudioFile
func (r *AyahAudioFileRepositoryImpl) Update(ctx context.Context, audio *quran.AyahAudioFile) (*quran.AyahAudioFile, error) {
	model := r.mapper.AyahAudioFileEntityToModel(audio)
	if err := r.db.WithContext(ctx).Where("id = ?", infraCore.ToGoogleUUID(audio.ID)).Updates(model).Error; err != nil {
		return nil, fmt.Errorf("failed to update ayah audio file: %w", err)
	}
	return r.mapper.AyahAudioFileModelToEntity(model), nil
}

// Delete soft deletes an AyahAudioFile
func (r *AyahAudioFileRepositoryImpl) Delete(ctx context.Context, id core.UUID) error {
	googleUUID := infraCore.ToGoogleUUID(id)
	if err := r.db.WithContext(ctx).Delete(&AyahAudioFileModel{}, googleUUID).Error; err != nil {
		return fmt.Errorf("failed to delete ayah audio file: %w", err)
	}
	return nil
}

// Count returns the number of AyahAudioFile records
func (r *AyahAudioFileRepositoryImpl) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&AyahAudioFileModel{}).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count ayah audio files: %w", err)
	}
	return count, nil
}

// CountBySurahAndReciter returns the number of AyahAudioFile records by surah and reciter
func (r *AyahAudioFileRepositoryImpl) CountBySurahAndReciter(ctx context.Context, surahId core.UUID, reciterId core.UUID) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&AyahAudioFileModel{}).
		Where("surah_id = ? AND reciter_id = ?", infraCore.ToGoogleUUID(surahId), infraCore.ToGoogleUUID(reciterId)).
		Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count ayah audio files by surah and reciter: %w", err)
	}
	return count, nil
}

// BookmarkAyahRepositoryImpl implements BookmarkAyahRepositoryInterface using GORM
type BookmarkAyahRepositoryImpl struct {
	db     *gorm.DB
	mapper *Mapper
}

// NewBookmarkAyahRepository creates a new instance of BookmarkAyahRepositoryImpl
func NewBookmarkAyahRepository(db *gorm.DB, mapper *Mapper) quran.BookmarkAyahRepositoryInterface {
	return &BookmarkAyahRepositoryImpl{db: db, mapper: mapper}
}

// GetById retrieves a BookmarkAyah by its ID
func (r *BookmarkAyahRepositoryImpl) GetById(ctx context.Context, id core.UUID) (*quran.BookmarkAyah, error) {
	var model BookmarkAyahModel
	if err := r.db.WithContext(ctx).Where("id = ?", infraCore.ToGoogleUUID(id)).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrBookmarkAyahNotFound
		}
		return nil, fmt.Errorf("failed to get bookmark ayah by id: %w", err)
	}
	return r.mapper.BookmarkAyahModelToEntity(&model), nil
}

// GetByUser retrieves BookmarkAyahs by user with pagination
func (r *BookmarkAyahRepositoryImpl) GetByUser(ctx context.Context, userId core.UUID, limit, offset int) ([]*quran.BookmarkAyah, error) {
	var models []BookmarkAyahModel
	query := r.db.WithContext(ctx).Where("user_id = ?", infraCore.ToGoogleUUID(userId)).Order("created_at ASC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get bookmark ayahs by user: %w", err)
	}
	outs := make([]*quran.BookmarkAyah, len(models))
	for i := range models {
		outs[i] = r.mapper.BookmarkAyahModelToEntity(&models[i])
	}
	return outs, nil
}

// GetByUserAndAyah retrieves a BookmarkAyah by user and ayah
func (r *BookmarkAyahRepositoryImpl) GetByUserAndAyah(ctx context.Context, userId core.UUID, ayahId core.UUID) (*quran.BookmarkAyah, error) {
	var model BookmarkAyahModel
	if err := r.db.WithContext(ctx).
		Where("user_id = ? AND ayah_id = ?", infraCore.ToGoogleUUID(userId), infraCore.ToGoogleUUID(ayahId)).
		First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrBookmarkAyahNotFound
		}
		return nil, fmt.Errorf("failed to get bookmark ayah by user and ayah: %w", err)
	}
	return r.mapper.BookmarkAyahModelToEntity(&model), nil
}

// Create creates a new BookmarkAyah
func (r *BookmarkAyahRepositoryImpl) Create(ctx context.Context, bookmark *quran.BookmarkAyah) (*quran.BookmarkAyah, error) {
	model := r.mapper.BookmarkAyahEntityToModel(bookmark)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create bookmark ayah: %w", err)
	}
	return r.mapper.BookmarkAyahModelToEntity(model), nil
}

// Delete soft deletes a BookmarkAyah
func (r *BookmarkAyahRepositoryImpl) Delete(ctx context.Context, id core.UUID) error {
	if err := r.db.WithContext(ctx).Delete(&BookmarkAyahModel{}, infraCore.ToGoogleUUID(id)).Error; err != nil {
		return fmt.Errorf("failed to delete bookmark ayah: %w", err)
	}
	return nil
}

// CountByUser returns the count of BookmarkAyahs for a user
func (r *BookmarkAyahRepositoryImpl) CountByUser(ctx context.Context, userId core.UUID) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&BookmarkAyahModel{}).
		Where("user_id = ?", infraCore.ToGoogleUUID(userId)).
		Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count bookmark ayahs by user: %w", err)
	}
	return count, nil
}

// GetByIdWithAyah retrieves a BookmarkAyah with Ayah relation by ID
func (r *BookmarkAyahRepositoryImpl) GetByIdWithAyah(ctx context.Context, id core.UUID) (*quran.BookmarkAyahWithAyah, error) {
	var model BookmarkAyahModel
	if err := r.db.WithContext(ctx).Preload("Ayah").Where("id = ?", infraCore.ToGoogleUUID(id)).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrBookmarkAyahNotFound
		}
		return nil, fmt.Errorf("failed to get bookmark ayah by id with ayah: %w", err)
	}
	return r.mapper.BookmarkAyahModelToEntityWithAyah(&model), nil
}

// LastReadRepositoryImpl implements LastReadRepositoryInterface using GORM
type LastReadRepositoryImpl struct {
	db     *gorm.DB
	mapper *Mapper
}

// NewLastReadRepository creates a new instance of LastReadRepositoryImpl
func NewLastReadRepository(db *gorm.DB, mapper *Mapper) quran.LastReadRepositoryInterface {
	return &LastReadRepositoryImpl{db: db, mapper: mapper}
}

// GetById retrieves a LastRead by its ID
func (r *LastReadRepositoryImpl) GetById(ctx context.Context, id core.UUID) (*quran.LastRead, error) {
	var model LastReadModel
	if err := r.db.WithContext(ctx).Where("id = ?", infraCore.ToGoogleUUID(id)).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrLastReadNotFound
		}
		return nil, fmt.Errorf("failed to get last read by id: %w", err)
	}
	return r.mapper.LastReadModelToEntity(&model), nil
}

// GetByUser retrieves LastRead records by user with pagination
func (r *LastReadRepositoryImpl) GetByUser(ctx context.Context, userId core.UUID, limit, offset int) ([]*quran.LastRead, error) {
	var models []LastReadModel
	query := r.db.WithContext(ctx).Where("user_id = ?", infraCore.ToGoogleUUID(userId)).Order("last_read_at DESC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get last reads by user: %w", err)
	}
	outs := make([]*quran.LastRead, len(models))
	for i := range models {
		outs[i] = r.mapper.LastReadModelToEntity(&models[i])
	}
	return outs, nil
}

// GetByUserAndSurah retrieves a LastRead by user and surah
func (r *LastReadRepositoryImpl) GetByUserAndSurah(ctx context.Context, userId core.UUID, surahId core.UUID) (*quran.LastRead, error) {
	var model LastReadModel
	if err := r.db.WithContext(ctx).
		Where("user_id = ? AND surah_id = ?", infraCore.ToGoogleUUID(userId), infraCore.ToGoogleUUID(surahId)).
		First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrLastReadNotFound
		}
		return nil, fmt.Errorf("failed to get last read by user and surah: %w", err)
	}
	return r.mapper.LastReadModelToEntity(&model), nil
}

// Create creates a new LastRead
func (r *LastReadRepositoryImpl) Create(ctx context.Context, lr *quran.LastRead) (*quran.LastRead, error) {
	model := r.mapper.LastReadEntityToModel(lr)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create last read: %w", err)
	}
	return r.mapper.LastReadModelToEntity(model), nil
}

// Update updates an existing LastRead
func (r *LastReadRepositoryImpl) Update(ctx context.Context, lr *quran.LastRead) (*quran.LastRead, error) {
	model := r.mapper.LastReadEntityToModel(lr)
	if err := r.db.WithContext(ctx).Where("id = ?", infraCore.ToGoogleUUID(lr.ID)).Updates(model).Error; err != nil {
		return nil, fmt.Errorf("failed to update last read: %w", err)
	}
	return r.mapper.LastReadModelToEntity(model), nil
}

// Delete soft deletes a LastRead
func (r *LastReadRepositoryImpl) Delete(ctx context.Context, id core.UUID) error {
	if err := r.db.WithContext(ctx).Delete(&LastReadModel{}, infraCore.ToGoogleUUID(id)).Error; err != nil {
		return fmt.Errorf("failed to delete last read: %w", err)
	}
	return nil
}

// CountByUser returns the count of LastRead records for a user
func (r *LastReadRepositoryImpl) CountByUser(ctx context.Context, userId core.UUID) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&LastReadModel{}).
		Where("user_id = ?", infraCore.ToGoogleUUID(userId)).
		Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count last reads by user: %w", err)
	}
	return count, nil
}

// GetByIdWithRelations retrieves a LastRead with Surah and Ayah relations by ID
func (r *LastReadRepositoryImpl) GetByIdWithRelations(ctx context.Context, id core.UUID) (*quran.LastReadWithRelations, error) {
	var model LastReadModel
	if err := r.db.WithContext(ctx).Preload("Surah").Preload("Ayah").Where("id = ?", infraCore.ToGoogleUUID(id)).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrLastReadNotFound
		}
		return nil, fmt.Errorf("failed to get last read by id with relations: %w", err)
	}
	return r.mapper.LastReadModelToEntityWithRelations(&model), nil
}

// GetByUserWithRelations retrieves LastRead records by user with relations and pagination
func (r *LastReadRepositoryImpl) GetByUserWithRelations(ctx context.Context, userId core.UUID, limit, offset int) ([]*quran.LastReadWithRelations, error) {
	var models []LastReadModel
	query := r.db.WithContext(ctx).Preload("Surah").Preload("Ayah").Where("user_id = ?", infraCore.ToGoogleUUID(userId)).Order("last_read_at DESC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get last reads by user with relations: %w", err)
	}
	outs := make([]*quran.LastReadWithRelations, len(models))
	for i := range models {
		outs[i] = r.mapper.LastReadModelToEntityWithRelations(&models[i])
	}
	return outs, nil
}

// ProgressHatamRepositoryImpl implements ProgressHatamRepositoryInterface using GORM
type ProgressHatamRepositoryImpl struct {
	db     *gorm.DB
	mapper *Mapper
}

// NewProgressHatamRepository creates a new instance of ProgressHatamRepositoryImpl
func NewProgressHatamRepository(db *gorm.DB, mapper *Mapper) quran.ProgressHatamRepositoryInterface {
	return &ProgressHatamRepositoryImpl{db: db, mapper: mapper}
}

// GetById retrieves a ProgressHatam by its ID
func (r *ProgressHatamRepositoryImpl) GetById(ctx context.Context, id core.UUID) (*quran.ProgressHatam, error) {
	var model ProgressHatamModel
	if err := r.db.WithContext(ctx).Where("id = ?", infraCore.ToGoogleUUID(id)).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrProgressHatamNotFound
		}
		return nil, fmt.Errorf("failed to get progress hatam by id: %w", err)
	}
	return r.mapper.ProgressHatamModelToEntity(&model), nil
}

// GetByUser retrieves ProgressHatam records by user with pagination
func (r *ProgressHatamRepositoryImpl) GetByUser(ctx context.Context, userId core.UUID, limit, offset int) ([]*quran.ProgressHatam, error) {
	var models []ProgressHatamModel
	query := r.db.WithContext(ctx).Where("user_id = ?", infraCore.ToGoogleUUID(userId)).Order("started_at DESC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get progress hatam by user: %w", err)
	}
	outs := make([]*quran.ProgressHatam, len(models))
	for i := range models {
		outs[i] = r.mapper.ProgressHatamModelToEntity(&models[i])
	}
	return outs, nil
}

// GetByUserAndJuz retrieves a ProgressHatam by user and juz
func (r *ProgressHatamRepositoryImpl) GetByUserAndJuz(ctx context.Context, userId core.UUID, juzId core.UUID) (*quran.ProgressHatam, error) {
	var model ProgressHatamModel
	if err := r.db.WithContext(ctx).
		Where("user_id = ? AND juz_id = ?", infraCore.ToGoogleUUID(userId), infraCore.ToGoogleUUID(juzId)).
		First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrProgressHatamNotFound
		}
		return nil, fmt.Errorf("failed to get progress hatam by user and juz: %w", err)
	}
	return r.mapper.ProgressHatamModelToEntity(&model), nil
}

// Create creates a new ProgressHatam
func (r *ProgressHatamRepositoryImpl) Create(ctx context.Context, p *quran.ProgressHatam) (*quran.ProgressHatam, error) {
	model := r.mapper.ProgressHatamEntityToModel(p)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create progress hatam: %w", err)
	}
	return r.mapper.ProgressHatamModelToEntity(model), nil
}

// Update updates an existing ProgressHatam
func (r *ProgressHatamRepositoryImpl) Update(ctx context.Context, p *quran.ProgressHatam) (*quran.ProgressHatam, error) {
	model := r.mapper.ProgressHatamEntityToModel(p)
	if err := r.db.WithContext(ctx).Where("id = ?", infraCore.ToGoogleUUID(p.ID)).Updates(model).Error; err != nil {
		return nil, fmt.Errorf("failed to update progress hatam: %w", err)
	}
	return r.mapper.ProgressHatamModelToEntity(model), nil
}

// Delete soft deletes a ProgressHatam
func (r *ProgressHatamRepositoryImpl) Delete(ctx context.Context, id core.UUID) error {
	if err := r.db.WithContext(ctx).Delete(&ProgressHatamModel{}, infraCore.ToGoogleUUID(id)).Error; err != nil {
		return fmt.Errorf("failed to delete progress hatam: %w", err)
	}
	return nil
}

// CountByUser returns the count of ProgressHatam records for a user
func (r *ProgressHatamRepositoryImpl) CountByUser(ctx context.Context, userId core.UUID) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&ProgressHatamModel{}).
		Where("user_id = ?", infraCore.ToGoogleUUID(userId)).
		Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count progress hatam by user: %w", err)
	}
	return count, nil
}

// GetByIdWithRelations retrieves a ProgressHatam with relations by ID
func (r *ProgressHatamRepositoryImpl) GetByIdWithRelations(ctx context.Context, id core.UUID) (*quran.ProgressHatamWithRelations, error) {
	var model ProgressHatamModel
	if err := r.db.WithContext(ctx).
		Preload("Juz").Preload("StartAyah").Preload("LastAyah").
		Where("id = ?", infraCore.ToGoogleUUID(id)).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, quran.ErrProgressHatamNotFound
		}
		return nil, fmt.Errorf("failed to get progress hatam by id with relations: %w", err)
	}
	return r.mapper.ProgressHatamModelToEntityWithRelations(&model), nil
}

// GetByUserWithRelations retrieves ProgressHatam records by user with relations and pagination
func (r *ProgressHatamRepositoryImpl) GetByUserWithRelations(ctx context.Context, userId core.UUID, limit, offset int) ([]*quran.ProgressHatamWithRelations, error) {
	var models []ProgressHatamModel
	query := r.db.WithContext(ctx).
		Preload("Juz").Preload("StartAyah").Preload("LastAyah").
		Where("user_id = ?", infraCore.ToGoogleUUID(userId)).Order("started_at DESC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get progress hatam by user with relations: %w", err)
	}
	outs := make([]*quran.ProgressHatamWithRelations, len(models))
	for i := range models {
		outs[i] = r.mapper.ProgressHatamModelToEntityWithRelations(&models[i])
	}
	return outs, nil
}

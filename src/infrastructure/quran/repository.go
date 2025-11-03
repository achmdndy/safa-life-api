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
	if err := r.db.WithContext(ctx).Where("surah_id = ?", googleUUID).Limit(limit).Offset(offset).Find(&models).Error; err != nil {
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
	if err := r.db.WithContext(ctx).Preload("Surah").Where("surah_id = ?", googleUUID).Limit(limit).Offset(offset).Find(&models).Error; err != nil {
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

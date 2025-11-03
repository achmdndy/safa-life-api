package quran

import (
	"github.com/google/uuid"
	"github.com/safalife/core-api/src/domain/core"
	"github.com/safalife/core-api/src/domain/quran"
	infraCore "github.com/safalife/core-api/src/infrastructure/core"
)

// Mapper handles conversion between models and entities
type Mapper struct{}

// NewMapper creates a new mapper instance
func NewMapper() *Mapper {
	return &Mapper{}
}

// SurahModelToEntity converts SurahModel to Surah entity
func (m *Mapper) SurahModelToEntity(model *SurahModel) *quran.Surah {
	if model == nil {
		return nil
	}

	return &quran.Surah{
		ID:              model.ID,
		NameArabic:      model.NameArabic,
		NameEnglish:     model.NameEnglish,
		RevelationPlace: model.RevelationPlace,
		RevelationOrder: model.RevelationOrder,
		AyahCount:       model.AyahCount,
		CreatedBy:       model.CreatedBy,
		UpdatedBy:       model.UpdatedBy,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
	}
}

// SurahModelToEntityWithAyahs converts SurahModel with Ayahs to SurahWithAyahs entity
func (m *Mapper) SurahModelToEntityWithAyahs(model *SurahModel) *quran.SurahWithAyahs {
	if model == nil {
		return nil
	}

	ayahs := make([]*quran.Ayah, len(model.Ayahs))
	for i, ayahModel := range model.Ayahs {
		ayahs[i] = m.AyahModelToEntity(&ayahModel)
	}

	return &quran.SurahWithAyahs{
		ID:              model.ID,
		NameArabic:      model.NameArabic,
		NameEnglish:     model.NameEnglish,
		RevelationPlace: model.RevelationPlace,
		RevelationOrder: model.RevelationOrder,
		AyahCount:       model.AyahCount,
		CreatedBy:       model.CreatedBy,
		UpdatedBy:       model.UpdatedBy,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
		Ayahs:           ayahs,
	}
}

// AyahModelToEntityWithSurah converts AyahModel with Surah to AyahWithSurah entity
func (m *Mapper) AyahModelToEntityWithSurah(model *AyahModel) *quran.AyahWithSurah {
	if model == nil {
		return nil
	}

	ayah := &quran.AyahWithSurah{
		ID:           model.ID,
		SurahID:      model.SurahID,
		Text:         model.Text,
		PageNumber:   model.PageNumber,
		JuzNumber:    model.JuzNumber,
		HizbNumber:   model.HizbNumber,
		ManzilNumber: model.ManzilNumber,
		CreatedBy:    model.CreatedBy,
		UpdatedBy:    model.UpdatedBy,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
	}

	// Convert related Surah if present
	if model.Surah.ID != uuid.Nil {
		ayah.Surah = m.SurahModelToEntity(&model.Surah)
	}

	return ayah
}

// JuzModelToEntityWithRelations converts JuzModel with all relations to JuzWithRelations entity
func (m *Mapper) JuzModelToEntityWithRelations(model *JuzModel) *quran.JuzWithRelations {
	if model == nil {
		return nil
	}

	juz := &quran.JuzWithRelations{
		ID:           model.ID,
		StartSurahID: model.StartSurahID,
		EndSurahID:   model.EndSurahID,
		StartAyahID:  model.StartAyahID,
		EndAyahID:    model.EndAyahID,
		CreatedBy:    model.CreatedBy,
		UpdatedBy:    model.UpdatedBy,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
	}

	// Convert related entities if present
	if model.StartSurah.ID != uuid.Nil {
		juz.StartSurah = m.SurahModelToEntity(&model.StartSurah)
	}
	if model.EndSurah.ID != uuid.Nil {
		juz.EndSurah = m.SurahModelToEntity(&model.EndSurah)
	}
	if model.StartAyah.ID != uuid.Nil {
		juz.StartAyah = m.AyahModelToEntity(&model.StartAyah)
	}
	if model.EndAyah.ID != uuid.Nil {
		juz.EndAyah = m.AyahModelToEntity(&model.EndAyah)
	}

	return juz
}

// SurahEntityToModel converts Surah entity to SurahModel
func (m *Mapper) SurahEntityToModel(entity *quran.Surah) *SurahModel {
	if entity == nil {
		return nil
	}

	return &SurahModel{
		ID:              m.coreUUIDToGoogleUUID(entity.ID),
		NameArabic:      entity.NameArabic,
		NameEnglish:     entity.NameEnglish,
		RevelationPlace: entity.RevelationPlace,
		RevelationOrder: entity.RevelationOrder,
		AyahCount:       entity.AyahCount,
		CreatedBy:       entity.CreatedBy,
		UpdatedBy:       entity.UpdatedBy,
		CreatedAt:       entity.CreatedAt,
		UpdatedAt:       entity.UpdatedAt,
	}
}

// AyahModelToEntity converts AyahModel to Ayah entity
func (m *Mapper) AyahModelToEntity(model *AyahModel) *quran.Ayah {
	if model == nil {
		return nil
	}

	return &quran.Ayah{
		ID:           model.ID,
		SurahID:      model.SurahID,
		Text:         model.Text,
		PageNumber:   model.PageNumber,
		JuzNumber:    model.JuzNumber,
		HizbNumber:   model.HizbNumber,
		ManzilNumber: model.ManzilNumber,
		CreatedBy:    model.CreatedBy,
		UpdatedBy:    model.UpdatedBy,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
	}
}

// AyahEntityToModel converts Ayah entity to AyahModel
func (m *Mapper) AyahEntityToModel(entity *quran.Ayah) *AyahModel {
	if entity == nil {
		return nil
	}

	return &AyahModel{
		ID:           m.coreUUIDToGoogleUUID(entity.ID),
		SurahID:      m.coreUUIDToGoogleUUID(entity.SurahID),
		Text:         entity.Text,
		PageNumber:   entity.PageNumber,
		JuzNumber:    entity.JuzNumber,
		HizbNumber:   entity.HizbNumber,
		ManzilNumber: entity.ManzilNumber,
		CreatedBy:    entity.CreatedBy,
		UpdatedBy:    entity.UpdatedBy,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}
}

// JuzModelToEntity converts JuzModel to Juz entity
func (m *Mapper) JuzModelToEntity(model *JuzModel) *quran.Juz {
	if model == nil {
		return nil
	}

	return &quran.Juz{
		ID:           model.ID,
		StartSurahID: model.StartSurahID,
		EndSurahID:   model.EndSurahID,
		StartAyahID:  model.StartAyahID,
		EndAyahID:    model.EndAyahID,
		CreatedBy:    model.CreatedBy,
		UpdatedBy:    model.UpdatedBy,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
	}
}

// JuzEntityToModel converts Juz entity to JuzModel
func (m *Mapper) JuzEntityToModel(entity *quran.Juz) *JuzModel {
	if entity == nil {
		return nil
	}

	return &JuzModel{
		ID:           m.coreUUIDToGoogleUUID(entity.ID),
		StartSurahID: m.coreUUIDToGoogleUUID(entity.StartSurahID),
		EndSurahID:   m.coreUUIDToGoogleUUID(entity.EndSurahID),
		StartAyahID:  m.coreUUIDToGoogleUUID(entity.StartAyahID),
		EndAyahID:    m.coreUUIDToGoogleUUID(entity.EndAyahID),
		CreatedBy:    entity.CreatedBy,
		UpdatedBy:    entity.UpdatedBy,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}
}

// SurahModelsToEntities converts slice of SurahModel to slice of Surah entities
func (m *Mapper) SurahModelsToEntities(models []*SurahModel) []*quran.Surah {
	if models == nil {
		return nil
	}

	entities := make([]*quran.Surah, len(models))
	for i, model := range models {
		entities[i] = m.SurahModelToEntity(model)
	}
	return entities
}

// SurahEntitiesToModels converts slice of Surah entities to slice of SurahModel
func (m *Mapper) SurahEntitiesToModels(entities []*quran.Surah) []*SurahModel {
	if entities == nil {
		return nil
	}

	models := make([]*SurahModel, len(entities))
	for i, entity := range entities {
		models[i] = m.SurahEntityToModel(entity)
	}
	return models
}

// AyahModelsToEntities converts slice of AyahModel to slice of Ayah entities
func (m *Mapper) AyahModelsToEntities(models []*AyahModel) []*quran.Ayah {
	if models == nil {
		return nil
	}

	entities := make([]*quran.Ayah, len(models))
	for i, model := range models {
		entities[i] = m.AyahModelToEntity(model)
	}
	return entities
}

// AyahEntitiesToModels converts slice of Ayah entities to slice of AyahModel
func (m *Mapper) AyahEntitiesToModels(entities []*quran.Ayah) []*AyahModel {
	if entities == nil {
		return nil
	}

	models := make([]*AyahModel, len(entities))
	for i, entity := range entities {
		models[i] = m.AyahEntityToModel(entity)
	}
	return models
}

// JuzModelsToEntities converts slice of JuzModel to slice of Juz entities
func (m *Mapper) JuzModelsToEntities(models []*JuzModel) []*quran.Juz {
	if models == nil {
		return nil
	}

	entities := make([]*quran.Juz, len(models))
	for i, model := range models {
		entities[i] = m.JuzModelToEntity(model)
	}
	return entities
}

// JuzEntitiesToModels converts slice of Juz entities to slice of JuzModel
func (m *Mapper) JuzEntitiesToModels(entities []*quran.Juz) []*JuzModel {
	if entities == nil {
		return nil
	}

	models := make([]*JuzModel, len(entities))
	for i, entity := range entities {
		models[i] = m.JuzEntityToModel(entity)
	}
	return models
}

// Helper functions for UUID conversion

// coreUUIDToGoogleUUID converts core.UUID to google/uuid.UUID
func (m *Mapper) coreUUIDToGoogleUUID(coreUUID core.UUID) uuid.UUID {
	return infraCore.ToGoogleUUID(coreUUID)
}

// googleUUIDToCoreUUID converts google/uuid.UUID to core.UUID
func (m *Mapper) googleUUIDToCoreUUID(googleUUID uuid.UUID) core.UUID {
	return infraCore.FromGoogleUUID(googleUUID)
}

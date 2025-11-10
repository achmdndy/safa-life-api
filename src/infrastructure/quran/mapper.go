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

// TranslationEditionModelToEntity converts TranslationEditionModel to domain entity
func (m *Mapper) TranslationEditionModelToEntity(model *TranslationEditionModel) *quran.TranslationEdition {
	if model == nil {
		return nil
	}
	return &quran.TranslationEdition{
		ID:        model.ID,
		Name:      model.Name,
		Author:    model.Author,
		Language:  model.Language,
		CreatedBy: model.CreatedBy,
		UpdatedBy: model.UpdatedBy,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

// TranslationEditionEntityToModel converts domain entity to TranslationEditionModel
func (m *Mapper) TranslationEditionEntityToModel(entity *quran.TranslationEdition) *TranslationEditionModel {
	if entity == nil {
		return nil
	}
	return &TranslationEditionModel{
		ID:        m.coreUUIDToGoogleUUID(entity.ID),
		Name:      entity.Name,
		Author:    entity.Author,
		Language:  entity.Language,
		CreatedBy: entity.CreatedBy,
		UpdatedBy: entity.UpdatedBy,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

// TranslationEditionModelsToEntities converts slice of models to slice of entities
func (m *Mapper) TranslationEditionModelsToEntities(models []*TranslationEditionModel) []*quran.TranslationEdition {
	if models == nil {
		return nil
	}
	outs := make([]*quran.TranslationEdition, len(models))
	for i, model := range models {
		outs[i] = m.TranslationEditionModelToEntity(model)
	}
	return outs
}

// AyahTranslationModelToEntity converts AyahTranslationModel to domain entity
func (m *Mapper) AyahTranslationModelToEntity(model *AyahTranslationModel) *quran.AyahTranslation {
	if model == nil {
		return nil
	}
	return &quran.AyahTranslation{
		ID:                   model.ID,
		TranslationEditionID: model.TranslationEditionID,
		SurahID:              model.SurahID,
		AyahID:               model.AyahID,
		Text:                 model.Text,
		CreatedBy:            model.CreatedBy,
		UpdatedBy:            model.UpdatedBy,
		CreatedAt:            model.CreatedAt,
		UpdatedAt:            model.UpdatedAt,
	}
}

// AyahTranslationEntityToModel converts domain entity to AyahTranslationModel
func (m *Mapper) AyahTranslationEntityToModel(entity *quran.AyahTranslation) *AyahTranslationModel {
	if entity == nil {
		return nil
	}
	return &AyahTranslationModel{
		ID:                   m.coreUUIDToGoogleUUID(entity.ID),
		TranslationEditionID: m.coreUUIDToGoogleUUID(entity.TranslationEditionID),
		SurahID:              m.coreUUIDToGoogleUUID(entity.SurahID),
		AyahID:               m.coreUUIDToGoogleUUID(entity.AyahID),
		Text:                 entity.Text,
		CreatedBy:            entity.CreatedBy,
		UpdatedBy:            entity.UpdatedBy,
		CreatedAt:            entity.CreatedAt,
		UpdatedAt:            entity.UpdatedAt,
	}
}

// AyahTranslationModelsToEntities converts slice of models to slice of entities
func (m *Mapper) AyahTranslationModelsToEntities(models []*AyahTranslationModel) []*quran.AyahTranslation {
	if models == nil {
		return nil
	}
	outs := make([]*quran.AyahTranslation, len(models))
	for i, model := range models {
		outs[i] = m.AyahTranslationModelToEntity(model)
	}
	return outs
}

// ReciterModelToEntity converts ReciterModel to domain entity
func (m *Mapper) ReciterModelToEntity(model *ReciterModel) *quran.Reciter {
	if model == nil {
		return nil
	}
	return &quran.Reciter{
		ID:        model.ID,
		Name:      model.Name,
		Style:     model.Style,
		Place:     model.Place,
		Picture:   model.Picture,
		CreatedBy: model.CreatedBy,
		UpdatedBy: model.UpdatedBy,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

// ReciterEntityToModel converts domain entity to ReciterModel
func (m *Mapper) ReciterEntityToModel(entity *quran.Reciter) *ReciterModel {
	if entity == nil {
		return nil
	}
	return &ReciterModel{
		ID:        m.coreUUIDToGoogleUUID(entity.ID),
		Name:      entity.Name,
		Style:     entity.Style,
		Place:     entity.Place,
		Picture:   entity.Picture,
		CreatedBy: entity.CreatedBy,
		UpdatedBy: entity.UpdatedBy,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

// ReciterModelsToEntities converts slice of models to slice of entities
func (m *Mapper) ReciterModelsToEntities(models []*ReciterModel) []*quran.Reciter {
	if models == nil {
		return nil
	}
	outs := make([]*quran.Reciter, len(models))
	for i, model := range models {
		outs[i] = m.ReciterModelToEntity(model)
	}
	return outs
}

// AyahAudioFileModelToEntity converts AyahAudioFileModel to domain entity
func (m *Mapper) AyahAudioFileModelToEntity(model *AyahAudioFileModel) *quran.AyahAudioFile {
	if model == nil {
		return nil
	}
	return &quran.AyahAudioFile{
		ID:        model.ID,
		ReciterID: model.ReciterID,
		SurahID:   model.SurahID,
		AyahID:    model.AyahID,
		FilePath:  model.FilePath,
		Duration:  model.Duration,
		ByteSize:  model.ByteSize,
		CreatedBy: model.CreatedBy,
		UpdatedBy: model.UpdatedBy,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

// AyahAudioFileEntityToModel converts domain entity to AyahAudioFileModel
func (m *Mapper) AyahAudioFileEntityToModel(entity *quran.AyahAudioFile) *AyahAudioFileModel {
	if entity == nil {
		return nil
	}
	return &AyahAudioFileModel{
		ID:        m.coreUUIDToGoogleUUID(entity.ID),
		ReciterID: m.coreUUIDToGoogleUUID(entity.ReciterID),
		SurahID:   m.coreUUIDToGoogleUUID(entity.SurahID),
		AyahID:    m.coreUUIDToGoogleUUID(entity.AyahID),
		FilePath:  entity.FilePath,
		Duration:  entity.Duration,
		ByteSize:  entity.ByteSize,
		CreatedBy: entity.CreatedBy,
		UpdatedBy: entity.UpdatedBy,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

// AyahAudioFileModelsToEntities converts slice of models to slice of entities
func (m *Mapper) AyahAudioFileModelsToEntities(models []*AyahAudioFileModel) []*quran.AyahAudioFile {
	if models == nil {
		return nil
	}
	outs := make([]*quran.AyahAudioFile, len(models))
	for i, model := range models {
		outs[i] = m.AyahAudioFileModelToEntity(model)
	}
	return outs
}

// BookmarkAyahModelToEntity converts BookmarkAyahModel to domain entity
func (m *Mapper) BookmarkAyahModelToEntity(model *BookmarkAyahModel) *quran.BookmarkAyah {
	if model == nil {
		return nil
	}
	return &quran.BookmarkAyah{
		ID:        model.ID,
		UserID:    model.UserID,
		AyahID:    model.AyahID,
		CreatedBy: model.CreatedBy,
		UpdatedBy: model.UpdatedBy,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

// BookmarkAyahModelToEntityWithAyah converts model with Ayah relation to BookmarkAyahWithAyah
func (m *Mapper) BookmarkAyahModelToEntityWithAyah(model *BookmarkAyahModel) *quran.BookmarkAyahWithAyah {
	if model == nil {
		return nil
	}
	out := &quran.BookmarkAyahWithAyah{
		ID:        model.ID,
		UserID:    model.UserID,
		AyahID:    model.AyahID,
		CreatedBy: model.CreatedBy,
		UpdatedBy: model.UpdatedBy,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
	if model.Ayah.ID != uuid.Nil {
		out.Ayah = m.AyahModelToEntity(&model.Ayah)
	}
	return out
}

// BookmarkAyahEntityToModel converts domain entity to BookmarkAyahModel
func (m *Mapper) BookmarkAyahEntityToModel(entity *quran.BookmarkAyah) *BookmarkAyahModel {
	if entity == nil {
		return nil
	}
	return &BookmarkAyahModel{
		ID:        m.coreUUIDToGoogleUUID(entity.ID),
		UserID:    entity.UserID,
		AyahID:    m.coreUUIDToGoogleUUID(entity.AyahID),
		CreatedBy: entity.CreatedBy,
		UpdatedBy: entity.UpdatedBy,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

// BookmarkAyahModelsToEntities converts slice of models to slice of entities
func (m *Mapper) BookmarkAyahModelsToEntities(models []*BookmarkAyahModel) []*quran.BookmarkAyah {
	if models == nil {
		return nil
	}
	outs := make([]*quran.BookmarkAyah, len(models))
	for i, model := range models {
		outs[i] = m.BookmarkAyahModelToEntity(model)
	}
	return outs
}

// BookmarkAyahEntitiesToModels converts slice of entities to slice of models
func (m *Mapper) BookmarkAyahEntitiesToModels(entities []*quran.BookmarkAyah) []*BookmarkAyahModel {
	if entities == nil {
		return nil
	}
	outs := make([]*BookmarkAyahModel, len(entities))
	for i, entity := range entities {
		outs[i] = m.BookmarkAyahEntityToModel(entity)
	}
	return outs
}

// LastReadModelToEntity converts LastReadModel to domain entity
func (m *Mapper) LastReadModelToEntity(model *LastReadModel) *quran.LastRead {
	if model == nil {
		return nil
	}
	return &quran.LastRead{
		ID:          model.ID,
		UserID:      model.UserID,
		SurahID:     model.SurahID,
		AyahID:      model.AyahID,
		AyahNumber:  model.AyahNumber,
		ProgressPct: model.ProgressPct,
		LastReadAt:  model.LastReadAt,
		CreatedBy:   model.CreatedBy,
		UpdatedBy:   model.UpdatedBy,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

// LastReadModelToEntityWithRelations converts model with relations to LastReadWithRelations
func (m *Mapper) LastReadModelToEntityWithRelations(model *LastReadModel) *quran.LastReadWithRelations {
	if model == nil {
		return nil
	}
	out := &quran.LastReadWithRelations{
		ID:          model.ID,
		UserID:      model.UserID,
		SurahID:     model.SurahID,
		AyahID:      model.AyahID,
		AyahNumber:  model.AyahNumber,
		ProgressPct: model.ProgressPct,
		LastReadAt:  model.LastReadAt,
		CreatedBy:   model.CreatedBy,
		UpdatedBy:   model.UpdatedBy,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
	if model.Surah.ID != uuid.Nil {
		out.Surah = m.SurahModelToEntity(&model.Surah)
	}
	if model.Ayah.ID != uuid.Nil {
		out.Ayah = m.AyahModelToEntity(&model.Ayah)
	}
	return out
}

// LastReadEntityToModel converts domain entity to LastReadModel
func (m *Mapper) LastReadEntityToModel(entity *quran.LastRead) *LastReadModel {
	if entity == nil {
		return nil
	}
	return &LastReadModel{
		ID:          m.coreUUIDToGoogleUUID(entity.ID),
		UserID:      entity.UserID,
		SurahID:     m.coreUUIDToGoogleUUID(entity.SurahID),
		AyahID:      m.coreUUIDToGoogleUUID(entity.AyahID),
		AyahNumber:  entity.AyahNumber,
		ProgressPct: entity.ProgressPct,
		LastReadAt:  entity.LastReadAt,
		CreatedBy:   entity.CreatedBy,
		UpdatedBy:   entity.UpdatedBy,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

// LastReadModelsToEntities converts slice of models to slice of entities
func (m *Mapper) LastReadModelsToEntities(models []*LastReadModel) []*quran.LastRead {
	if models == nil {
		return nil
	}
	outs := make([]*quran.LastRead, len(models))
	for i, model := range models {
		outs[i] = m.LastReadModelToEntity(model)
	}
	return outs
}

// LastReadEntitiesToModels converts slice of entities to slice of models
func (m *Mapper) LastReadEntitiesToModels(entities []*quran.LastRead) []*LastReadModel {
	if entities == nil {
		return nil
	}
	outs := make([]*LastReadModel, len(entities))
	for i, entity := range entities {
		outs[i] = m.LastReadEntityToModel(entity)
	}
	return outs
}

// ProgressHatamModelToEntity converts ProgressHatamModel to domain entity
func (m *Mapper) ProgressHatamModelToEntity(model *ProgressHatamModel) *quran.ProgressHatam {
	if model == nil {
		return nil
	}
	return &quran.ProgressHatam{
		ID:          model.ID,
		UserID:      model.UserID,
		JuzID:       model.JuzID,
		StartAyahID: model.StartAyahID,
		LastAyahID:  model.LastAyahID,
		ProgressPct: model.ProgressPct,
		IsCompleted: model.IsCompleted,
		StartedAt:   model.StartedAt,
		CompletedAt: model.CompletedAt,
		CreatedBy:   model.CreatedBy,
		UpdatedBy:   model.UpdatedBy,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

// ProgressHatamModelToEntityWithRelations converts model with relations to ProgressHatamWithRelations
func (m *Mapper) ProgressHatamModelToEntityWithRelations(model *ProgressHatamModel) *quran.ProgressHatamWithRelations {
	if model == nil {
		return nil
	}
	out := &quran.ProgressHatamWithRelations{
		ID:          model.ID,
		UserID:      model.UserID,
		JuzID:       model.JuzID,
		StartAyahID: model.StartAyahID,
		LastAyahID:  model.LastAyahID,
		ProgressPct: model.ProgressPct,
		IsCompleted: model.IsCompleted,
		StartedAt:   model.StartedAt,
		CompletedAt: model.CompletedAt,
		CreatedBy:   model.CreatedBy,
		UpdatedBy:   model.UpdatedBy,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
	if model.Juz.ID != uuid.Nil {
		out.Juz = m.JuzModelToEntity(&model.Juz)
	}
	if model.StartAyah.ID != uuid.Nil {
		out.StartAyah = m.AyahModelToEntity(&model.StartAyah)
	}
	if model.LastAyah != nil && model.LastAyah.ID != uuid.Nil {
		out.LastAyah = m.AyahModelToEntity(model.LastAyah)
	}
	return out
}

// ProgressHatamEntityToModel converts domain entity to ProgressHatamModel
func (m *Mapper) ProgressHatamEntityToModel(entity *quran.ProgressHatam) *ProgressHatamModel {
	if entity == nil {
		return nil
	}
	return &ProgressHatamModel{
		ID:          m.coreUUIDToGoogleUUID(entity.ID),
		UserID:      entity.UserID,
		JuzID:       m.coreUUIDToGoogleUUID(entity.JuzID),
		StartAyahID: m.coreUUIDToGoogleUUID(entity.StartAyahID),
		LastAyahID:  m.coreUUIDToGoogleUUID(entity.LastAyahID),
		ProgressPct: entity.ProgressPct,
		IsCompleted: entity.IsCompleted,
		StartedAt:   entity.StartedAt,
		CompletedAt: entity.CompletedAt,
		CreatedBy:   entity.CreatedBy,
		UpdatedBy:   entity.UpdatedBy,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

// ProgressHatamModelsToEntities converts slice of models to slice of entities
func (m *Mapper) ProgressHatamModelsToEntities(models []*ProgressHatamModel) []*quran.ProgressHatam {
	if models == nil {
		return nil
	}
	outs := make([]*quran.ProgressHatam, len(models))
	for i, model := range models {
		outs[i] = m.ProgressHatamModelToEntity(model)
	}
	return outs
}

// ProgressHatamEntitiesToModels converts slice of entities to slice of models
func (m *Mapper) ProgressHatamEntitiesToModels(entities []*quran.ProgressHatam) []*ProgressHatamModel {
	if entities == nil {
		return nil
	}
	outs := make([]*ProgressHatamModel, len(entities))
	for i, entity := range entities {
		outs[i] = m.ProgressHatamEntityToModel(entity)
	}
	return outs
}

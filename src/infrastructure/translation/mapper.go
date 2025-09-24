package translation

import "github.com/achmdndy/safa-life-api/src/domain/translation"

// ToTranslationDomain converts TranslationModel to domain.Translation
func (m *TranslationModel) ToTranslationDomain() translation.Translation {
	return translation.Translation{
		ID:        m.ID,
		Name:      m.Name,
		Author:    m.Author,
		Language:  m.Language,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

// FromTranslationDomain converts domain.Translation to TranslationModel
func FromTranslationDomain(d *translation.Translation) *TranslationModel {
	return &TranslationModel{
		ID:        d.ID,
		Name:      d.Name,
		Author:    d.Author,
		Language:  d.Language,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

// ToTranslationDomainSlice converts a slice of TranslationModel to a slice of domain.Translation
func ToTranslationDomainSlice(models []TranslationModel) []translation.Translation {
	domains := make([]translation.Translation, len(models))
	for i, m := range models {
		domains[i] = m.ToTranslationDomain()
	}
	return domains
}

// ToAyahTranslationDomain converts AyahTranslationModel to domain.AyahTranslation
func (m *AyahTranslationModel) ToAyahTranslationDomain() translation.AyahTranslation {
	return translation.AyahTranslation{
		TranslationID: m.TranslationID,
		SurahID:       m.SurahID,
		AyahID:        m.AyahID,
		Text:          m.Text,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
	}
}

// FromAyahTranslationDomain converts domain.AyahTranslation to AyahTranslationModel
func FromAyahTranslationDomain(d *translation.AyahTranslation) *AyahTranslationModel {
	return &AyahTranslationModel{
		TranslationID: d.TranslationID,
		SurahID:       d.SurahID,
		AyahID:        d.AyahID,
		Text:          d.Text,
		CreatedAt:     d.CreatedAt,
		UpdatedAt:     d.UpdatedAt,
	}
}

// ToAyahTranslationDomainSlice converts a slice of AyahTranslationModel to a slice of domain.AyahTranslation
func ToAyahTranslationDomainSlice(models []AyahTranslationModel) []translation.AyahTranslation {
	domains := make([]translation.AyahTranslation, len(models))
	for i, m := range models {
		domains[i] = m.ToAyahTranslationDomain()
	}
	return domains
}

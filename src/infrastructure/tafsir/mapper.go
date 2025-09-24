package tafsir

import "github.com/achmdndy/safa-life-api/src/domain/tafsir"

// ToTafsirDomain converts TafsirModel to domain.Tafsir
func (m *TafsirModel) ToTafsirDomain() tafsir.Tafsir {
	return tafsir.Tafsir{
		ID:        m.ID,
		Name:      m.Name,
		Author:    m.Author,
		Language:  m.Language,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

// FromTafsirDomain converts domain.Tafsir to TafsirModel
func FromTafsirDomain(d *tafsir.Tafsir) *TafsirModel {
	return &TafsirModel{
		ID:        d.ID,
		Name:      d.Name,
		Author:    d.Author,
		Language:  d.Language,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

// ToTafsirDomainSlice converts a slice of TafsirModel to a slice of domain.Tafsir
func ToTafsirDomainSlice(models []TafsirModel) []tafsir.Tafsir {
	domains := make([]tafsir.Tafsir, len(models))
	for i, m := range models {
		domains[i] = m.ToTafsirDomain()
	}
	return domains
}

// ToAyahTafsirDomain converts AyahTafsirModel to domain.AyahTafsir
func (m *AyahTafsirModel) ToAyahTafsirDomain() tafsir.AyahTafsir {
	return tafsir.AyahTafsir{
		TafsirID:  m.TafsirID,
		SurahID:   m.SurahID,
		AyahID:    m.AyahID,
		Text:      m.Text,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

// FromAyahTafsirDomain converts domain.AyahTafsir to AyahTafsirModel
func FromAyahTafsirDomain(d *tafsir.AyahTafsir) *AyahTafsirModel {
	return &AyahTafsirModel{
		TafsirID:  d.TafsirID,
		SurahID:   d.SurahID,
		AyahID:    d.AyahID,
		Text:      d.Text,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

// ToAyahTafsirDomainSlice converts a slice of AyahTafsirModel to a slice of domain.AyahTafsir
func ToAyahTafsirDomainSlice(models []AyahTafsirModel) []tafsir.AyahTafsir {
	domains := make([]tafsir.AyahTafsir, len(models))
	for i, m := range models {
		domains[i] = m.ToAyahTafsirDomain()
	}
	return domains
}
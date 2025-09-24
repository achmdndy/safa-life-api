package quran

import (
	"github.com/achmdndy/safa-life-api/src/domain/quran"
)

// ToSurahDomain converts SurahModel to domain Surah
func (s *SurahModel) ToSurahDomain() quran.Surah {
	return quran.Surah{
		ID:              s.ID,
		NameArabic:      s.NameArabic,
		NameEnglish:     s.NameEnglish,
		RevelationPlace: s.RevelationPlace,
		RevelationOrder: s.RevelationOrder,
		AyahCount:       s.AyahCount,
		CreatedAt:       s.CreatedAt,
		UpdatedAt:       s.UpdatedAt,
	}
}

// ToAyahDomain converts AyahModel to domain Ayah
func (a *AyahModel) ToAyahDomain() quran.Ayah {
	return quran.Ayah{
		SurahID:      a.SurahID,
		AyahID:       a.AyahID,
		Text:         a.Text,
		PageNumber:   a.PageNumber,
		JuzNumber:    a.JuzNumber,
		HizbNumber:   a.HizbNumber,
		ManzilNumber: a.ManzilNumber,
		CreatedAt:    a.CreatedAt,
		UpdatedAt:    a.UpdatedAt,
	}
}

// ToJuzDomain converts JuzModel to domain Juz
func (j *JuzModel) ToJuzDomain() quran.Juz {
	return quran.Juz{
		ID:         j.ID,
		StartSurah: j.StartSurah,
		StartAyah:  j.StartAyah,
		EndSurah:   j.EndSurah,
		EndAyah:    j.EndAyah,
		CreatedAt:  j.CreatedAt,
		UpdatedAt:  j.UpdatedAt,
	}
}

// FromSurahDomain converts domain Surah to SurahModel
func FromSurahDomain(s *quran.Surah) *SurahModel {
	return &SurahModel{
		ID:              s.ID,
		NameArabic:      s.NameArabic,
		NameEnglish:     s.NameEnglish,
		RevelationPlace: s.RevelationPlace,
		RevelationOrder: s.RevelationOrder,
		AyahCount:       s.AyahCount,
		CreatedAt:       s.CreatedAt,
		UpdatedAt:       s.UpdatedAt,
	}
}

// FromAyahDomain converts domain Ayah to AyahModel
func FromAyahDomain(a *quran.Ayah) *AyahModel {
	return &AyahModel{
		SurahID:      a.SurahID,
		AyahID:       a.AyahID,
		Text:         a.Text,
		PageNumber:   a.PageNumber,
		JuzNumber:    a.JuzNumber,
		HizbNumber:   a.HizbNumber,
		ManzilNumber: a.ManzilNumber,
		CreatedAt:    a.CreatedAt,
		UpdatedAt:    a.UpdatedAt,
	}
}

// FromJuzDomain converts domain Juz to JuzModel
func FromJuzDomain(j *quran.Juz) *JuzModel {
	return &JuzModel{
		ID:         j.ID,
		StartSurah: j.StartSurah,
		StartAyah:  j.StartAyah,
		EndSurah:   j.EndSurah,
		EndAyah:    j.EndAyah,
		CreatedAt:  j.CreatedAt,
		UpdatedAt:  j.UpdatedAt,
	}
}

// ToSurahDomainSlice converts slice of SurahModel to slice of domain Surah
func ToSurahDomainSlice(models []SurahModel) []quran.Surah {
	result := make([]quran.Surah, len(models))
	for i, model := range models {
		result[i] = model.ToSurahDomain()
	}
	return result
}

// ToAyahDomainSlice converts slice of AyahModel to slice of domain Ayah
func ToAyahDomainSlice(models []AyahModel) []quran.Ayah {
	result := make([]quran.Ayah, len(models))
	for i, model := range models {
		result[i] = model.ToAyahDomain()
	}
	return result
}

// ToJuzDomainSlice converts slice of JuzModel to slice of domain Juz
func ToJuzDomainSlice(models []JuzModel) []quran.Juz {
	result := make([]quran.Juz, len(models))
	for i, model := range models {
		result[i] = model.ToJuzDomain()
	}
	return result
}
package dto

import (
	"time"

	"github.com/safalife/core-api/src/domain/quran"
)

// Surah Response DTOs
type SurahResponse struct {
	ID              string    `json:"id"`
	NameArabic      string    `json:"nameArabic"`
	NameEnglish     string    `json:"nameEnglish"`
	RevelationPlace string    `json:"revelationPlace"`
	RevelationOrder int       `json:"revelationOrder"`
	AyahCount       int       `json:"ayahCount"`
	CreatedBy       string    `json:"createdBy"`
	UpdatedBy       string    `json:"updatedBy"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// ToSurahResponse converts a domain Surah to its DTO.
func ToSurahResponse(d *quran.Surah) *SurahResponse {
	if d == nil {
		return nil
	}
	return &SurahResponse{
		ID:              d.ID.String(),
		NameArabic:      d.NameArabic,
		NameEnglish:     d.NameEnglish,
		RevelationPlace: d.RevelationPlace,
		RevelationOrder: d.RevelationOrder,
		AyahCount:       d.AyahCount,
		CreatedBy:       d.CreatedBy,
		UpdatedBy:       d.UpdatedBy,
		CreatedAt:       d.CreatedAt,
		UpdatedAt:       d.UpdatedAt,
	}
}

// ToSurahResponseSlice converts a slice of domain Surahs to a slice of DTOs.
func ToSurahResponseSlice(ds []*quran.Surah) []*SurahResponse {
	if ds == nil {
		return nil
	}
	outs := make([]*SurahResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToSurahResponse(d)
	}
	return outs
}

// Ayah Response DTOs
type AyahResponse struct {
	ID           string    `json:"id"`
	SurahID      string    `json:"surahId"`
	Text         string    `json:"text"`
	PageNumber   int       `json:"pageNumber"`
	JuzNumber    int       `json:"juzNumber"`
	HizbNumber   int       `json:"hizbNumber"`
	ManzilNumber int       `json:"manzilNumber"`
	CreatedBy    string    `json:"createdBy"`
	UpdatedBy    string    `json:"updatedBy"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// ToAyahResponse converts a domain Ayah to its DTO.
func ToAyahResponse(d *quran.Ayah) *AyahResponse {
	if d == nil {
		return nil
	}
	return &AyahResponse{
		ID:           d.ID.String(),
		SurahID:      d.SurahID.String(),
		Text:         d.Text,
		PageNumber:   d.PageNumber,
		JuzNumber:    d.JuzNumber,
		HizbNumber:   d.HizbNumber,
		ManzilNumber: d.ManzilNumber,
		CreatedBy:    d.CreatedBy,
		UpdatedBy:    d.UpdatedBy,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
	}
}

// ToAyahResponseSlice converts a slice of domain Ayahs to a slice of DTOs.
func ToAyahResponseSlice(ds []*quran.Ayah) []*AyahResponse {
	if ds == nil {
		return nil
	}
	outs := make([]*AyahResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToAyahResponse(d)
	}
	return outs
}

// Juz Response DTOs
type JuzResponse struct {
	ID           string    `json:"id"`
	StartSurahID string    `json:"startSurahId"`
	EndSurahID   string    `json:"endSurahId"`
	StartAyahID  string    `json:"startAyahId"`
	EndAyahID    string    `json:"endAyahId"`
	CreatedBy    string    `json:"createdBy"`
	UpdatedBy    string    `json:"updatedBy"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// ToJuzResponse converts a domain Juz to its DTO.
func ToJuzResponse(d *quran.Juz) *JuzResponse {
	if d == nil {
		return nil
	}
	return &JuzResponse{
		ID:           d.ID.String(),
		StartSurahID: d.StartSurahID.String(),
		EndSurahID:   d.EndSurahID.String(),
		StartAyahID:  d.StartAyahID.String(),
		EndAyahID:    d.EndAyahID.String(),
		CreatedBy:    d.CreatedBy,
		UpdatedBy:    d.UpdatedBy,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
	}
}

// ToJuzResponseSlice converts a slice of domain Juz to a slice of DTOs.
func ToJuzResponseSlice(ds []*quran.Juz) []*JuzResponse {
	if ds == nil {
		return nil
	}
	outs := make([]*JuzResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToJuzResponse(d)
	}
	return outs
}

// Pagination Response
type PaginationResponse struct {
	Total  int64 `json:"total"`
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
}

// Surah List Response with Pagination
type SurahListResponse struct {
	Data       []*SurahResponse    `json:"data"`
	Pagination *PaginationResponse `json:"pagination"`
}

// Ayah List Response with Pagination
type AyahListResponse struct {
	Data       []*AyahResponse     `json:"data"`
	Pagination *PaginationResponse `json:"pagination"`
}

// Juz List Response with Pagination
type JuzListResponse struct {
	Data       []*JuzResponse      `json:"data"`
	Pagination *PaginationResponse `json:"pagination"`
}

// Count Response
type CountResponse struct {
	Count int64 `json:"count"`
}

// Surah with Ayahs Response DTO
type SurahWithAyahsResponse struct {
	ID              string          `json:"id"`
	NameArabic      string          `json:"nameArabic"`
	NameEnglish     string          `json:"nameEnglish"`
	RevelationPlace string          `json:"revelationPlace"`
	RevelationOrder int             `json:"revelationOrder"`
	AyahCount       int             `json:"ayahCount"`
	CreatedBy       string          `json:"createdBy"`
	UpdatedBy       string          `json:"updatedBy"`
	CreatedAt       time.Time       `json:"createdAt"`
	UpdatedAt       time.Time       `json:"updatedAt"`
	Ayahs           []*AyahResponse `json:"ayahs,omitempty"`
}

// ToSurahWithAyahsResponse converts a domain SurahWithAyahs to its DTO.
func ToSurahWithAyahsResponse(d *quran.SurahWithAyahs) *SurahWithAyahsResponse {
	if d == nil {
		return nil
	}
	return &SurahWithAyahsResponse{
		ID:              d.ID.String(),
		NameArabic:      d.NameArabic,
		NameEnglish:     d.NameEnglish,
		RevelationPlace: d.RevelationPlace,
		RevelationOrder: d.RevelationOrder,
		AyahCount:       d.AyahCount,
		CreatedBy:       d.CreatedBy,
		UpdatedBy:       d.UpdatedBy,
		CreatedAt:       d.CreatedAt,
		UpdatedAt:       d.UpdatedAt,
		Ayahs:           ToAyahResponseSlice(d.Ayahs),
	}
}

// ToSurahWithAyahsResponseSlice converts a slice of domain SurahWithAyahs to a slice of DTOs.
func ToSurahWithAyahsResponseSlice(ds []*quran.SurahWithAyahs) []*SurahWithAyahsResponse {
	if ds == nil {
		return nil
	}
	outs := make([]*SurahWithAyahsResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToSurahWithAyahsResponse(d)
	}
	return outs
}

// Ayah with Surah Response DTO
type AyahWithSurahResponse struct {
	ID           string         `json:"id"`
	SurahID      string         `json:"surahId"`
	Text         string         `json:"text"`
	PageNumber   int            `json:"pageNumber"`
	JuzNumber    int            `json:"juzNumber"`
	HizbNumber   int            `json:"hizbNumber"`
	ManzilNumber int            `json:"manzilNumber"`
	CreatedBy    string         `json:"createdBy"`
	UpdatedBy    string         `json:"updatedBy"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	Surah        *SurahResponse `json:"surah,omitempty"`
}

// ToAyahWithSurahResponse converts a domain AyahWithSurah to its DTO.
func ToAyahWithSurahResponse(d *quran.AyahWithSurah) *AyahWithSurahResponse {
	if d == nil {
		return nil
	}
	return &AyahWithSurahResponse{
		ID:           d.ID.String(),
		SurahID:      d.SurahID.String(),
		Text:         d.Text,
		PageNumber:   d.PageNumber,
		JuzNumber:    d.JuzNumber,
		HizbNumber:   d.HizbNumber,
		ManzilNumber: d.ManzilNumber,
		CreatedBy:    d.CreatedBy,
		UpdatedBy:    d.UpdatedBy,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
		Surah:        ToSurahResponse(d.Surah),
	}
}

// ToAyahWithSurahResponseSlice converts a slice of domain AyahWithSurah to a slice of DTOs.
func ToAyahWithSurahResponseSlice(ds []*quran.AyahWithSurah) []*AyahWithSurahResponse {
	if ds == nil {
		return nil
	}
	outs := make([]*AyahWithSurahResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToAyahWithSurahResponse(d)
	}
	return outs
}

// Juz with Relations Response DTO
type JuzWithRelationsResponse struct {
	ID           string         `json:"id"`
	StartSurahID string         `json:"startSurahId"`
	EndSurahID   string         `json:"endSurahId"`
	StartAyahID  string         `json:"startAyahId"`
	EndAyahID    string         `json:"endAyahId"`
	CreatedBy    string         `json:"createdBy"`
	UpdatedBy    string         `json:"updatedBy"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	StartSurah   *SurahResponse `json:"startSurah,omitempty"`
	EndSurah     *SurahResponse `json:"endSurah,omitempty"`
	StartAyah    *AyahResponse  `json:"startAyah,omitempty"`
	EndAyah      *AyahResponse  `json:"endAyah,omitempty"`
}

// ToJuzWithRelationsResponse converts a domain JuzWithRelations to its DTO.
func ToJuzWithRelationsResponse(d *quran.JuzWithRelations) *JuzWithRelationsResponse {
	if d == nil {
		return nil
	}
	return &JuzWithRelationsResponse{
		ID:           d.ID.String(),
		StartSurahID: d.StartSurahID.String(),
		EndSurahID:   d.EndSurahID.String(),
		StartAyahID:  d.StartAyahID.String(),
		EndAyahID:    d.EndAyahID.String(),
		CreatedBy:    d.CreatedBy,
		UpdatedBy:    d.UpdatedBy,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
		StartSurah:   ToSurahResponse(d.StartSurah),
		EndSurah:     ToSurahResponse(d.EndSurah),
		StartAyah:    ToAyahResponse(d.StartAyah),
		EndAyah:      ToAyahResponse(d.EndAyah),
	}
}

// ToJuzWithRelationsResponseSlice converts a slice of domain JuzWithRelations to a slice of DTOs.
func ToJuzWithRelationsResponseSlice(ds []*quran.JuzWithRelations) []*JuzWithRelationsResponse {
	if ds == nil {
		return nil
	}
	outs := make([]*JuzWithRelationsResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToJuzWithRelationsResponse(d)
	}
	return outs
}

// List responses with relations
type SurahWithAyahsListResponse struct {
	Data       []*SurahWithAyahsResponse `json:"data"`
	Pagination *PaginationResponse       `json:"pagination"`
}

type AyahWithSurahListResponse struct {
	Data       []*AyahWithSurahResponse `json:"data"`
	Pagination *PaginationResponse      `json:"pagination"`
}

type JuzWithRelationsListResponse struct {
	Data       []*JuzWithRelationsResponse `json:"data"`
	Pagination *PaginationResponse         `json:"pagination"`
}
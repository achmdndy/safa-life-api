package dto

import (
	"time"
	"github.com/achmdndy/safa-life-api/src/domain/quran"
)

// SurahResponse represents the response for a surah
type SurahResponse struct {
	ID              int       `json:"surah_id"`
	NameArabic      string    `json:"name_ar"`
	NameEnglish     string    `json:"name_en"`
	RevelationPlace string    `json:"revelation_place"`
	RevelationOrder int       `json:"revelation_order"`
	AyahCount       int       `json:"ayah_count"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// AyahResponse represents the response for an ayah
type AyahResponse struct {
	SurahID      int       `json:"surah_id"`
	AyahID       int       `json:"ayah_id"`
	Text         string    `json:"text"`
	PageNumber   int       `json:"page_number"`
	JuzNumber    int       `json:"juz_number"`
	HizbNumber   int       `json:"hizb_number"`
	ManzilNumber int       `json:"manzil_number"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// JuzResponse represents the response for a juz
type JuzResponse struct {
	ID         int       `json:"juz_id"`
	StartSurah int       `json:"start_surah"`
	StartAyah  int       `json:"start_ayah"`
	EndSurah   int       `json:"end_surah"`
	EndAyah    int       `json:"end_ayah"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// SurahWithAyahsResponse represents the response for a surah with its ayahs
type SurahWithAyahsResponse struct {
	SurahResponse
	Ayahs []AyahResponse `json:"ayahs"`
}

// JuzWithContentResponse represents the response for a juz with content details
type JuzWithContentResponse struct {
	JuzResponse
	StartSurahName string `json:"start_surah_name"`
	EndSurahName   string `json:"end_surah_name"`
	TotalAyahs     int    `json:"total_ayahs"`
}

// PaginatedResponse represents a paginated response
type PaginatedResponse[T any] struct {
	Data       []T                `json:"data"`
	Pagination PaginationMetadata `json:"pagination"`
}

// PaginationMetadata represents pagination metadata
type PaginationMetadata struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// SearchResponse represents the search response
type SearchResponse struct {
	Query   string        `json:"query"`
	Type    string        `json:"type"`
	Results interface{}   `json:"results"`
	Count   int           `json:"count"`
}

// APIResponse represents a standard API response
type APIResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

// Conversion functions from domain to response DTOs

// ToSurahResponse converts domain Surah to SurahResponse
func ToSurahResponse(surah quran.Surah) SurahResponse {
	return SurahResponse{
		ID:              surah.ID,
		NameArabic:      surah.NameArabic,
		NameEnglish:     surah.NameEnglish,
		RevelationPlace: surah.RevelationPlace,
		RevelationOrder: surah.RevelationOrder,
		AyahCount:       surah.AyahCount,
		CreatedAt:       surah.CreatedAt,
		UpdatedAt:       surah.UpdatedAt,
	}
}

// ToAyahResponse converts domain Ayah to AyahResponse
func ToAyahResponse(ayah quran.Ayah) AyahResponse {
	return AyahResponse{
		SurahID:      ayah.SurahID,
		AyahID:       ayah.AyahID,
		Text:         ayah.Text,
		PageNumber:   ayah.PageNumber,
		JuzNumber:    ayah.JuzNumber,
		HizbNumber:   ayah.HizbNumber,
		ManzilNumber: ayah.ManzilNumber,
		CreatedAt:    ayah.CreatedAt,
		UpdatedAt:    ayah.UpdatedAt,
	}
}

// ToJuzResponse converts domain Juz to JuzResponse
func ToJuzResponse(juz quran.Juz) JuzResponse {
	return JuzResponse{
		ID:         juz.ID,
		StartSurah: juz.StartSurah,
		StartAyah:  juz.StartAyah,
		EndSurah:   juz.EndSurah,
		EndAyah:    juz.EndAyah,
		CreatedAt:  juz.CreatedAt,
		UpdatedAt:  juz.UpdatedAt,
	}
}

// ToSurahWithAyahsResponse converts domain SurahWithAyahs to SurahWithAyahsResponse
func ToSurahWithAyahsResponse(surahWithAyahs quran.SurahWithAyahs) SurahWithAyahsResponse {
	ayahs := make([]AyahResponse, len(surahWithAyahs.Ayahs))
	for i, ayah := range surahWithAyahs.Ayahs {
		ayahs[i] = ToAyahResponse(ayah)
	}

	return SurahWithAyahsResponse{
		SurahResponse: ToSurahResponse(surahWithAyahs.Surah),
		Ayahs:         ayahs,
	}
}

// ToJuzWithContentResponse converts domain JuzWithContent to JuzWithContentResponse
func ToJuzWithContentResponse(juzWithContent quran.JuzWithContent) JuzWithContentResponse {
	return JuzWithContentResponse{
		JuzResponse:    ToJuzResponse(juzWithContent.Juz),
		StartSurahName: juzWithContent.StartSurahName,
		EndSurahName:   juzWithContent.EndSurahName,
		TotalAyahs:     juzWithContent.TotalAyahs,
	}
}

// ToSurahResponseSlice converts slice of domain Surah to slice of SurahResponse
func ToSurahResponseSlice(surahs []quran.Surah) []SurahResponse {
	result := make([]SurahResponse, len(surahs))
	for i, surah := range surahs {
		result[i] = ToSurahResponse(surah)
	}
	return result
}

// ToAyahResponseSlice converts slice of domain Ayah to slice of AyahResponse
func ToAyahResponseSlice(ayahs []quran.Ayah) []AyahResponse {
	result := make([]AyahResponse, len(ayahs))
	for i, ayah := range ayahs {
		result[i] = ToAyahResponse(ayah)
	}
	return result
}

// ToJuzResponseSlice converts slice of domain Juz to slice of JuzResponse
func ToJuzResponseSlice(juzs []quran.Juz) []JuzResponse {
	result := make([]JuzResponse, len(juzs))
	for i, juz := range juzs {
		result[i] = ToJuzResponse(juz)
	}
	return result
}

// NewSuccessResponse creates a successful API response
func NewSuccessResponse[T any](message string, data T) APIResponse[T] {
	return APIResponse[T]{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// NewErrorResponse creates an error API response
func NewErrorResponse[T any](message string, err error) APIResponse[T] {
	errorMsg := ""
	if err != nil {
		errorMsg = err.Error()
	}
	
	return APIResponse[T]{
		Success: false,
		Message: message,
		Error:   errorMsg,
	}
}
package dto

import (
	"time"

	"github.com/achmdndy/safa-life-api/src/domain/translation"
)

// TranslationResponse represents a translation edition.
type TranslationResponse struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Author   string    `json:"author"`
	Language string    `json:"language"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AyahTranslationResponse represents the translation of a single Ayah.
type AyahTranslationResponse struct {
	TranslationID string    `json:"translation_id"`
	SurahID       int       `json:"surah_id"`
	AyahID        int       `json:"ayah_id"`
	Text          string    `json:"text"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// ToTranslationResponse converts a domain Translation to its DTO.
func ToTranslationResponse(d translation.Translation) TranslationResponse {
	return TranslationResponse{
		ID:        d.ID,
		Name:      d.Name,
		Author:    d.Author,
		Language:  d.Language,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

// ToTranslationResponseSlice converts a slice of domain Translations to a slice of DTOs.
func ToTranslationResponseSlice(ds []translation.Translation) []TranslationResponse {
	outs := make([]TranslationResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToTranslationResponse(d)
	}
	return outs
}

// ToAyahTranslationResponse converts a domain AyahTranslation to its DTO.
func ToAyahTranslationResponse(d translation.AyahTranslation) AyahTranslationResponse {
	return AyahTranslationResponse{
		TranslationID: d.TranslationID,
		SurahID:       d.SurahID,
		AyahID:        d.AyahID,
		Text:          d.Text,
		CreatedAt:     d.CreatedAt,
		UpdatedAt:     d.UpdatedAt,
	}
}

// ToAyahTranslationResponseSlice converts a slice of domain AyahTranslations to a slice of DTOs.
func ToAyahTranslationResponseSlice(ds []translation.AyahTranslation) []AyahTranslationResponse {
	outs := make([]AyahTranslationResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToAyahTranslationResponse(d)
	}
	return outs
}

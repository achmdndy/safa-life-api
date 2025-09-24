package dto

import (
	"time"

	"github.com/achmdndy/safa-life-api/src/domain/tafsir"
)

// TafsirResponse represents a tafsir edition.
type TafsirResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Author    string    `json:"author"`
	Language  string    `json:"language"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AyahTafsirResponse represents the tafsir of a single Ayah.
type AyahTafsirResponse struct {
	TafsirID  string    `json:"tafsir_id"`
	SurahID   int       `json:"surah_id"`
	AyahID    int       `json:"ayah_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToTafsirResponse converts a domain Tafsir to its DTO.
func ToTafsirResponse(d tafsir.Tafsir) TafsirResponse {
	return TafsirResponse{
		ID:        d.ID,
		Name:      d.Name,
		Author:    d.Author,
		Language:  d.Language,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

// ToTafsirResponseSlice converts a slice of domain Tafsirs to a slice of DTOs.
func ToTafsirResponseSlice(ds []tafsir.Tafsir) []TafsirResponse {
	outs := make([]TafsirResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToTafsirResponse(d)
	}
	return outs
}

// ToAyahTafsirResponse converts a domain AyahTafsir to its DTO.
func ToAyahTafsirResponse(d tafsir.AyahTafsir) AyahTafsirResponse {
	return AyahTafsirResponse{
		TafsirID:  d.TafsirID,
		SurahID:   d.SurahID,
		AyahID:    d.AyahID,
		Text:      d.Text,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

// ToAyahTafsirResponseSlice converts a slice of domain AyahTafsirs to a slice of DTOs.
func ToAyahTafsirResponseSlice(ds []tafsir.AyahTafsir) []AyahTafsirResponse {
	outs := make([]AyahTafsirResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToAyahTafsirResponse(d)
	}
	return outs
}

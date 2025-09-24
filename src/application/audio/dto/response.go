package dto

import (
	"time"

	"github.com/achmdndy/safa-life-api/src/domain/audio"
)

// AyahAudioResponse represents the audio file for a specific Ayah.
type AyahAudioResponse struct {
	ReciterID string    `json:"reciter_id"`
	SurahID   int       `json:"surah_id"`
	AyahID    int       `json:"ayah_id"`
	FilePath  string    `json:"file_path"`
	Duration  float64   `json:"duration"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToAyahAudioResponse converts a domain AyahAudio to its DTO.
func ToAyahAudioResponse(d audio.AyahAudio) AyahAudioResponse {
	return AyahAudioResponse{
		ReciterID: d.ReciterID,
		SurahID:   d.SurahID,
		AyahID:    d.AyahID,
		FilePath:  d.FilePath,
		Duration:  d.Duration,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

// ToAyahAudioResponseSlice converts a slice of domain AyahAudios to a slice of DTOs.
func ToAyahAudioResponseSlice(ds []audio.AyahAudio) []AyahAudioResponse {
	outs := make([]AyahAudioResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToAyahAudioResponse(d)
	}
	return outs
}

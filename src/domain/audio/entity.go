package audio

import "time"

// AyahAudio represents the audio file for a specific Ayah by a specific reciter.
type AyahAudio struct {
	ReciterID string    `json:"reciter_id"`
	SurahID   int       `json:"surah_id"`
	AyahID    int       `json:"ayah_id"`
	FilePath  string    `json:"file_path"`
	Duration  float64   `json:"duration"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
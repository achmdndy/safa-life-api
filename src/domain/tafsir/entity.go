package tafsir

import "time"

// Tafsir represents a tafsir work/edition.
type Tafsir struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Author    string    `json:"author"`
	Language  string    `json:"language"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AyahTafsir represents the tafsir of a single Ayah.
type AyahTafsir struct {
	TafsirID  string    `json:"tafsir_id"`
	SurahID   int       `json:"surah_id"`
	AyahID    int       `json:"ayah_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
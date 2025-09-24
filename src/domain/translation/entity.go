package translation

import "time"

// Translation represents a translation edition.
type Translation struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Author    string    `json:"author"`
	Language  string    `json:"language"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AyahTranslation represents the translation of a single Ayah.
type AyahTranslation struct {
	TranslationID string    `json:"translation_id"`
	SurahID       int       `json:"surah_id"`
	AyahID        int       `json:"ayah_id"`
	Text          string    `json:"text"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

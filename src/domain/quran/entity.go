package quran

import (
	"time"
)

// Surah represents a chapter in the Quran
type Surah struct {
	ID              int       `json:"surah_id"`
	NameArabic      string    `json:"name_ar"`
	NameEnglish     string    `json:"name_en"`
	RevelationPlace string    `json:"revelation_place"`
	RevelationOrder int       `json:"revelation_order"`
	AyahCount       int       `json:"ayah_count"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Ayah represents a verse in the Quran
type Ayah struct {
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

// Juz represents a section of the Quran (Para)
type Juz struct {
	ID         int       `json:"juz_id"`
	StartSurah int       `json:"start_surah"`
	StartAyah  int       `json:"start_ayah"`
	EndSurah   int       `json:"end_surah"`
	EndAyah    int       `json:"end_ayah"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// SurahWithAyahs represents a Surah with its Ayahs
type SurahWithAyahs struct {
	Surah
	Ayahs []Ayah `json:"ayahs"`
}

// JuzWithContent represents a Juz with its content details
type JuzWithContent struct {
	Juz
	StartSurahName string `json:"start_surah_name"`
	EndSurahName   string `json:"end_surah_name"`
	TotalAyahs     int    `json:"total_ayahs"`
}
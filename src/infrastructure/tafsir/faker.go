package tafsir

import (
	"math/rand"
	"time"
)

// TafsirEditionFaker generates fake Tafsir data
type TafsirEditionFaker struct{}

// NewTafsirEditionFaker creates a new TafsirEditionFaker
func NewTafsirEditionFaker() *TafsirEditionFaker {
	return &TafsirEditionFaker{}
}

// Generate creates a fake TafsirModel
func (f *TafsirEditionFaker) Generate() TafsirModel {
	editions := []struct {
		ID       string
		Name     string
		Author   string
		Language string
	}{
		{"en_ibn_kathir", "Tafsir Ibn Kathir", "Ibn Kathir", "English"},
		{"id_jalalayn", "Tafsir al-Jalalayn", "Jalal al-Din al-Mahalli and Jalal al-Din al-Suyuti", "Indonesian"},
	}
	edition := editions[rand.Intn(len(editions))]
	return TafsirModel{
		ID:        edition.ID,
		Name:      edition.Name,
		Author:    edition.Author,
		Language:  edition.Language,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// AyahTafsirFaker generates fake AyahTafsir data
type AyahTafsirFaker struct{}

// NewAyahTafsirFaker creates a new AyahTafsirFaker
func NewAyahTafsirFaker() *AyahTafsirFaker {
	return &AyahTafsirFaker{}
}

// Generate creates a fake AyahTafsirModel
func (f *AyahTafsirFaker) Generate(tafsirID string, surahID, ayahID int) AyahTafsirModel {
	texts := []string{
		"This is a detailed explanation of the verse.",
		"The historical context of this verse is as follows...",
	}
	return AyahTafsirModel{
		TafsirID:  tafsirID,
		SurahID:   surahID,
		AyahID:    ayahID,
		Text:      texts[rand.Intn(len(texts))],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// TafsirFaker combines all faker functionality
type TafsirFaker struct {
	EditionFaker *TafsirEditionFaker
	AyahFaker    *AyahTafsirFaker
}

// NewTafsirFaker creates a new TafsirFaker instance
func NewTafsirFaker() *TafsirFaker {
	rand.Seed(time.Now().UnixNano())
	return &TafsirFaker{
		EditionFaker: NewTafsirEditionFaker(),
		AyahFaker:    NewAyahTafsirFaker(),
	}
}
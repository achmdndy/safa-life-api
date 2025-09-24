package translation

import (
	"math/rand"
	"time"
)

// TranslationEditionFaker generates fake Translation data
type TranslationEditionFaker struct{}

// NewTranslationEditionFaker creates a new TranslationEditionFaker
func NewTranslationEditionFaker() *TranslationEditionFaker {
	return &TranslationEditionFaker{}
}

// Generate creates a fake TranslationModel
func (f *TranslationEditionFaker) Generate() TranslationModel {
	editions := []struct {
		ID       string
		Name     string
		Author   string
		Language string
	}{
		{"en_asad", "The Message of The Qur'an", "Muhammad Asad", "English"},
		{"en_sahih", "Sahih International", "Sahih International", "English"},
		{"id_indonesian", "Indonesian Ministry of Religious Affairs", "Kemenag RI", "Indonesian"},
	}
	edition := editions[rand.Intn(len(editions))]
	return TranslationModel{
		ID:        edition.ID,
		Name:      edition.Name,
		Author:    edition.Author,
		Language:  edition.Language,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// AyahTranslationFaker generates fake AyahTranslation data
type AyahTranslationFaker struct{}

// NewAyahTranslationFaker creates a new AyahTranslationFaker
func NewAyahTranslationFaker() *AyahTranslationFaker {
	return &AyahTranslationFaker{}
}

// Generate creates a fake AyahTranslationModel
func (f *AyahTranslationFaker) Generate(translationID string, surahID, ayahID int) AyahTranslationModel {
	texts := []string{
		"In the name of Allah, the Entirely Merciful, the Especially Merciful.",
		"[All] praise is [due] to Allah, Lord of the worlds -",
		"The Entirely Merciful, the Especially Merciful,",
		"Sovereign of the Day of Recompense.",
	}
	return AyahTranslationModel{
		TranslationID: translationID,
		SurahID:       surahID,
		AyahID:        ayahID,
		Text:          texts[rand.Intn(len(texts))],
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

// TranslationFaker combines all faker functionality
type TranslationFaker struct {
	EditionFaker *TranslationEditionFaker
	AyahFaker    *AyahTranslationFaker
}

// NewTranslationFaker creates a new TranslationFaker instance
func NewTranslationFaker() *TranslationFaker {
	rand.Seed(time.Now().UnixNano())
	return &TranslationFaker{
		EditionFaker: NewTranslationEditionFaker(),
		AyahFaker:    NewAyahTranslationFaker(),
	}
}
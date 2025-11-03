package quran

import (
	"time"

	"github.com/safalife/core-api/src/domain/core"
)

// Surah represents a surah (chapter) entity in the domain
type Surah struct {
	ID              core.UUID `json:"id"`
	NameArabic      string    `json:"nameArabic"`
	NameEnglish     string    `json:"nameEnglish"`
	RevelationPlace string    `json:"revelationPlace"`
	RevelationOrder int       `json:"revelationOrder"`
	AyahCount       int       `json:"ayahCount"`
	CreatedBy       string    `json:"createdBy"`
	UpdatedBy       string    `json:"updatedBy"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// SurahWithAyahs represents a surah entity with its related ayahs
type SurahWithAyahs struct {
	ID              core.UUID `json:"id"`
	NameArabic      string    `json:"nameArabic"`
	NameEnglish     string    `json:"nameEnglish"`
	RevelationPlace string    `json:"revelationPlace"`
	RevelationOrder int       `json:"revelationOrder"`
	AyahCount       int       `json:"ayahCount"`
	CreatedBy       string    `json:"createdBy"`
	UpdatedBy       string    `json:"updatedBy"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Ayahs           []*Ayah   `json:"ayahs,omitempty"`
}

// Ayah represents an ayah (verse) entity in the domain
type Ayah struct {
	ID           core.UUID `json:"id"`
	SurahID      core.UUID `json:"surahId"`
	Text         string    `json:"text"`
	PageNumber   int       `json:"pageNumber"`
	JuzNumber    int       `json:"juzNumber"`
	HizbNumber   int       `json:"hizbNumber"`
	ManzilNumber int       `json:"manzilNumber"`
	CreatedBy    string    `json:"createdBy"`
	UpdatedBy    string    `json:"updatedBy"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// AyahWithSurah represents an ayah entity with its related surah
type AyahWithSurah struct {
	ID           core.UUID `json:"id"`
	SurahID      core.UUID `json:"surahId"`
	Text         string    `json:"text"`
	PageNumber   int       `json:"pageNumber"`
	JuzNumber    int       `json:"juzNumber"`
	HizbNumber   int       `json:"hizbNumber"`
	ManzilNumber int       `json:"manzilNumber"`
	CreatedBy    string    `json:"createdBy"`
	UpdatedBy    string    `json:"updatedBy"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Surah        *Surah    `json:"surah,omitempty"`
}

// Juz represents a juz (part) entity in the domain
type Juz struct {
	ID           core.UUID `json:"id"`
	StartSurahID core.UUID `json:"startSurahId"`
	EndSurahID   core.UUID `json:"endSurahId"`
	StartAyahID  core.UUID `json:"startAyahId"`
	EndAyahID    core.UUID `json:"endAyahId"`
	CreatedBy    string    `json:"createdBy"`
	UpdatedBy    string    `json:"updatedBy"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// JuzWithRelations represents a juz entity with its related entities
type JuzWithRelations struct {
	ID           core.UUID `json:"id"`
	StartSurahID core.UUID `json:"startSurahId"`
	EndSurahID   core.UUID `json:"endSurahId"`
	StartAyahID  core.UUID `json:"startAyahId"`
	EndAyahID    core.UUID `json:"endAyahId"`
	CreatedBy    string    `json:"createdBy"`
	UpdatedBy    string    `json:"updatedBy"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	StartSurah   *Surah    `json:"startSurah,omitempty"`
	EndSurah     *Surah    `json:"endSurah,omitempty"`
	StartAyah    *Ayah     `json:"startAyah,omitempty"`
	EndAyah      *Ayah     `json:"endAyah,omitempty"`
}

// NewSurah creates a new Surah entity
func NewSurah(id core.UUID, nameArabic, nameEnglish, revelationPlace string, revelationOrder, ayahCount int, createdBy string) *Surah {
	now := time.Now()
	return &Surah{
		ID:              id,
		NameArabic:      nameArabic,
		NameEnglish:     nameEnglish,
		RevelationPlace: revelationPlace,
		RevelationOrder: revelationOrder,
		AyahCount:       ayahCount,
		CreatedBy:       createdBy,
		UpdatedBy:       createdBy,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}

// NewAyah creates a new Ayah entity
func NewAyah(id, surahID core.UUID, text string, pageNumber, juzNumber, hizbNumber, manzilNumber int, createdBy string) *Ayah {
	now := time.Now()
	return &Ayah{
		ID:           id,
		SurahID:      surahID,
		Text:         text,
		PageNumber:   pageNumber,
		JuzNumber:    juzNumber,
		HizbNumber:   hizbNumber,
		ManzilNumber: manzilNumber,
		CreatedBy:    createdBy,
		UpdatedBy:    createdBy,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

// NewJuz creates a new Juz entity
func NewJuz(id, startSurahID, endSurahID, startAyahID, endAyahID core.UUID, createdBy string) *Juz {
	now := time.Now()
	return &Juz{
		ID:           id,
		StartSurahID: startSurahID,
		EndSurahID:   endSurahID,
		StartAyahID:  startAyahID,
		EndAyahID:    endAyahID,
		CreatedBy:    createdBy,
		UpdatedBy:    createdBy,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

// Update updates the Surah entity
func (s *Surah) Update(nameArabic, nameEnglish, revelationPlace string, revelationOrder, ayahCount int, updatedBy string) {
	s.NameArabic = nameArabic
	s.NameEnglish = nameEnglish
	s.RevelationPlace = revelationPlace
	s.RevelationOrder = revelationOrder
	s.AyahCount = ayahCount
	s.UpdatedBy = updatedBy
	s.UpdatedAt = time.Now()
}

func (a *Ayah) Update(text string, pageNumber, juzNumber, hizbNumber, manzilNumber int, updatedBy string) {
	a.Text = text
	a.PageNumber = pageNumber
	a.JuzNumber = juzNumber
	a.HizbNumber = hizbNumber
	a.ManzilNumber = manzilNumber
	a.UpdatedBy = updatedBy
	a.UpdatedAt = time.Now()
}

func (j *Juz) Update(startSurahID, endSurahID, startAyahID, endAyahID core.UUID, updatedBy string) {
	j.StartSurahID = startSurahID
	j.EndSurahID = endSurahID
	j.StartAyahID = startAyahID
	j.EndAyahID = endAyahID
	j.UpdatedBy = updatedBy
	j.UpdatedAt = time.Now()
}

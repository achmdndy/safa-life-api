package quran

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// SurahFaker generates fake surah data
type SurahFaker struct{}

// NewSurahFaker creates a new surah faker
func NewSurahFaker() *SurahFaker {
	return &SurahFaker{}
}

// Generate creates a fake SurahModel
func (f *SurahFaker) Generate() *SurahModel {
	surahNames := []string{"Al-Fatihah", "Al-Baqarah", "Ali Imran", "An-Nisa", "Al-Maidah", "Al-Anam", "Al-Araf"}
	surahName := surahNames[rand.Intn(len(surahNames))]

	return &SurahModel{
		ID:              uuid.New(),
		NameArabic:      "سورة " + surahName,
		NameEnglish:     "Surah " + surahName,
		RevelationPlace: []string{"Mecca", "Medina"}[rand.Intn(2)],
		RevelationOrder: rand.Intn(114) + 1,
		AyahCount:       rand.Intn(284) + 3, // Between 3-286 ayahs
		CreatedBy:       fmt.Sprintf("user_%d", rand.Intn(1000)),
		UpdatedBy:       fmt.Sprintf("user_%d", rand.Intn(1000)),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

// GenerateWithNumber creates a fake SurahModel with specific number
func (f *SurahFaker) GenerateWithNumber(number int) *SurahModel {
	surahNames := map[int]string{
		1: "Al-Fatihah", 2: "Al-Baqarah", 3: "Ali 'Imran", 4: "An-Nisa", 5: "Al-Ma'idah",
		6: "Al-An'am", 7: "Al-A'raf", 8: "Al-Anfal", 9: "At-Tawbah", 10: "Yunus",
	}

	nameEnglish := surahNames[number]
	if nameEnglish == "" {
		nameEnglish = fmt.Sprintf("Surah %d", number)
	}

	return &SurahModel{
		ID:              uuid.New(),
		NameArabic:      "سورة " + nameEnglish,
		NameEnglish:     nameEnglish,
		RevelationPlace: []string{"Mecca", "Medina"}[rand.Intn(2)],
		RevelationOrder: number,
		AyahCount:       rand.Intn(284) + 3,
		CreatedBy:       fmt.Sprintf("user_%d", rand.Intn(1000)),
		UpdatedBy:       fmt.Sprintf("user_%d", rand.Intn(1000)),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

// GenerateWithPlace creates a fake SurahModel with specific revelation place
func (f *SurahFaker) GenerateWithPlace(place string) *SurahModel {
	surah := f.Generate()
	surah.RevelationPlace = place
	return surah
}

// AyahFaker generates fake ayah data
type AyahFaker struct{}

// NewAyahFaker creates a new ayah faker
func NewAyahFaker() *AyahFaker {
	return &AyahFaker{}
}

// Generate creates a fake AyahModel
func (f *AyahFaker) Generate(surahID uuid.UUID) *AyahModel {
	ayahTexts := []string{
		"بِسْمِ اللَّهِ الرَّحْمَٰنِ الرَّحِيمِ",
		"الْحَمْدُ لِلَّهِ رَبِّ الْعَالَمِينَ",
		"الرَّحْمَٰنِ الرَّحِيمِ",
		"مَالِكِ يَوْمِ الدِّينِ",
	}

	return &AyahModel{
		ID:           uuid.New(),
		SurahID:      surahID,
		Text:         ayahTexts[rand.Intn(len(ayahTexts))],
		PageNumber:   rand.Intn(604) + 1, // Quran has 604 pages
		JuzNumber:    rand.Intn(30) + 1,  // 30 Juz in Quran
		HizbNumber:   rand.Intn(60) + 1,  // 60 Hizb in Quran
		ManzilNumber: rand.Intn(7) + 1,   // 7 Manzil in Quran
		CreatedBy:    fmt.Sprintf("user_%d", rand.Intn(1000)),
		UpdatedBy:    fmt.Sprintf("user_%d", rand.Intn(1000)),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

// GenerateWithJuz creates a fake AyahModel with specific juz number
func (f *AyahFaker) GenerateWithJuz(surahID uuid.UUID, juzNumber int) *AyahModel {
	ayah := f.Generate(surahID)
	ayah.JuzNumber = juzNumber
	return ayah
}

// GenerateWithPage creates a fake AyahModel with specific page number
func (f *AyahFaker) GenerateWithPage(surahID uuid.UUID, pageNumber int) *AyahModel {
	ayah := f.Generate(surahID)
	ayah.PageNumber = pageNumber
	return ayah
}

// GenerateMultiple creates multiple fake AyahModel for a surah
func (f *AyahFaker) GenerateMultiple(surahID uuid.UUID, count int) []*AyahModel {
	ayahs := make([]*AyahModel, count)
	for i := 0; i < count; i++ {
		ayahs[i] = f.Generate(surahID)
	}
	return ayahs
}

// JuzFaker generates fake juz data
type JuzFaker struct{}

// NewJuzFaker creates a new juz faker
func NewJuzFaker() *JuzFaker {
	return &JuzFaker{}
}

// Generate creates a fake JuzModel
func (f *JuzFaker) Generate(startSurahID, endSurahID, startAyahID, endAyahID uuid.UUID) *JuzModel {
	return &JuzModel{
		ID:           uuid.New(),
		StartSurahID: startSurahID,
		EndSurahID:   endSurahID,
		StartAyahID:  startAyahID,
		EndAyahID:    endAyahID,
		CreatedBy:    fmt.Sprintf("user_%d", rand.Intn(1000)),
		UpdatedBy:    fmt.Sprintf("user_%d", rand.Intn(1000)),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

// GenerateWithNumber creates a fake JuzModel with specific number (1-30)
func (f *JuzFaker) GenerateWithNumber(number int, startSurahID, endSurahID, startAyahID, endAyahID uuid.UUID) *JuzModel {
	return &JuzModel{
		ID:           uuid.New(),
		StartSurahID: startSurahID,
		EndSurahID:   endSurahID,
		StartAyahID:  startAyahID,
		EndAyahID:    endAyahID,
		CreatedBy:    fmt.Sprintf("user_%d", rand.Intn(1000)),
		UpdatedBy:    fmt.Sprintf("user_%d", rand.Intn(1000)),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

// QuranFaker combines all fakers for Quran entities
type QuranFaker struct {
	SurahFaker *SurahFaker
	AyahFaker  *AyahFaker
	JuzFaker   *JuzFaker
}

// NewQuranFaker creates a new Quran faker
func NewQuranFaker() *QuranFaker {
	return &QuranFaker{
		SurahFaker: NewSurahFaker(),
		AyahFaker:  NewAyahFaker(),
		JuzFaker:   NewJuzFaker(),
	}
}

// GenerateCompleteSurah creates a complete surah with ayahs
func (f *QuranFaker) GenerateCompleteSurah(ayahCount int) (*SurahModel, []*AyahModel) {
	surah := f.SurahFaker.Generate()
	surah.AyahCount = ayahCount

	ayahs := make([]*AyahModel, ayahCount)
	for i := 0; i < ayahCount; i++ {
		ayahs[i] = f.AyahFaker.Generate(surah.ID)
	}

	return surah, ayahs
}

// GenerateQuranStructure creates a basic Quran structure with surahs, ayahs, and juz
func (f *QuranFaker) GenerateQuranStructure(surahCount, avgAyahPerSurah, juzCount int) ([]*SurahModel, []*AyahModel, []*JuzModel) {
	var surahs []*SurahModel
	var ayahs []*AyahModel
	var juzs []*JuzModel

	// Generate surahs and ayahs
	for i := 0; i < surahCount; i++ {
		surah := f.SurahFaker.GenerateWithNumber(i + 1)
		surahs = append(surahs, surah)

		ayahCount := avgAyahPerSurah + (rand.Intn(10) - 5) // Vary ayah count
		if ayahCount < 3 {
			ayahCount = 3
		}
		surah.AyahCount = ayahCount

		for j := 0; j < ayahCount; j++ {
			ayah := f.AyahFaker.Generate(surah.ID)
			ayahs = append(ayahs, ayah)
		}
	}

	// Generate juz
	for i := 0; i < juzCount && len(ayahs) > 0; i++ {
		startIdx := (i * len(ayahs)) / juzCount
		endIdx := ((i + 1) * len(ayahs)) / juzCount
		if endIdx > len(ayahs) {
			endIdx = len(ayahs)
		}

		if startIdx < len(ayahs) && endIdx > startIdx {
			startAyah := ayahs[startIdx]
			endAyah := ayahs[endIdx-1]

			juz := f.JuzFaker.GenerateWithNumber(i+1, startAyah.SurahID, endAyah.SurahID, startAyah.ID, endAyah.ID)
			juzs = append(juzs, juz)
		}
	}

	return surahs, ayahs, juzs
}

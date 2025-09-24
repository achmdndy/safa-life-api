package quran

import (
	"fmt"
	"math/rand"
	"time"
)

// SurahFaker generates fake Surah data
type SurahFaker struct{}

// NewSurahFaker creates a new SurahFaker instance
func NewSurahFaker() *SurahFaker {
	return &SurahFaker{}
}

// Generate creates a fake SurahModel
func (sf *SurahFaker) Generate(id int) SurahModel {
	surahNames := []struct {
		Arabic  string
		English string
	}{
		{"الفاتحة", "Al-Fatihah"},
		{"البقرة", "Al-Baqarah"},
		{"آل عمران", "Ali 'Imran"},
		{"النساء", "An-Nisa"},
		{"المائدة", "Al-Ma'idah"},
		{"الأنعام", "Al-An'am"},
		{"الأعراف", "Al-A'raf"},
		{"الأنفال", "Al-Anfal"},
		{"التوبة", "At-Tawbah"},
		{"يونس", "Yunus"},
		{"هود", "Hud"},
		{"يوسف", "Yusuf"},
		{"الرعد", "Ar-Ra'd"},
		{"إبراهيم", "Ibrahim"},
		{"الحجر", "Al-Hijr"},
		{"النحل", "An-Nahl"},
		{"الإسراء", "Al-Isra"},
		{"الكهف", "Al-Kahf"},
		{"مريم", "Maryam"},
		{"طه", "Taha"},
	}

	revelationPlaces := []string{"Mecca", "Medina"}
	
	// Use predefined names if available, otherwise generate
	var nameAr, nameEn string
	if id <= len(surahNames) {
		nameAr = surahNames[id-1].Arabic
		nameEn = surahNames[id-1].English
	} else {
		nameAr = fmt.Sprintf("سورة %d", id)
		nameEn = fmt.Sprintf("Surah %d", id)
	}

	return SurahModel{
		ID:              id,
		NameArabic:      nameAr,
		NameEnglish:     nameEn,
		RevelationPlace: revelationPlaces[rand.Intn(len(revelationPlaces))],
		RevelationOrder: rand.Intn(114) + 1,
		AyahCount:       rand.Intn(286) + 1, // Al-Baqarah has 286 verses (max)
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

// GenerateBatch creates multiple fake SurahModel instances
func (sf *SurahFaker) GenerateBatch(count int) []SurahModel {
	surahs := make([]SurahModel, count)
	for i := 0; i < count; i++ {
		surahs[i] = sf.Generate(i + 1)
	}
	return surahs
}

// AyahFaker generates fake Ayah data
type AyahFaker struct{}

// NewAyahFaker creates a new AyahFaker instance
func NewAyahFaker() *AyahFaker {
	return &AyahFaker{}
}

// Generate creates a fake AyahModel
func (af *AyahFaker) Generate(surahID, ayahID int) AyahModel {
	return AyahModel{
		SurahID:      surahID,
		AyahID:       ayahID,
		Text:         af.getRandomArabicText(),
		PageNumber:   rand.Intn(604) + 1, // Quran has 604 pages
		JuzNumber:    rand.Intn(30) + 1,  // Quran has 30 Juz
		HizbNumber:   rand.Intn(60) + 1,  // Quran has 60 Hizb
		ManzilNumber: rand.Intn(7) + 1,   // Quran has 7 Manzil
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

// GenerateBatch creates multiple fake AyahModel instances for a specific Surah
func (af *AyahFaker) GenerateBatch(surahID, count int) []AyahModel {
	ayahs := make([]AyahModel, count)
	for i := 0; i < count; i++ {
		ayahs[i] = af.Generate(surahID, i+1)
	}
	return ayahs
}

// GenerateForAllSurahs creates fake Ayahs for multiple Surahs
func (af *AyahFaker) GenerateForAllSurahs(surahs []SurahModel) []AyahModel {
	var allAyahs []AyahModel
	for _, surah := range surahs {
		ayahs := af.GenerateBatch(surah.ID, surah.AyahCount)
		allAyahs = append(allAyahs, ayahs...)
	}
	return allAyahs
}

// GenerateForAllSurahsWithValidJuz creates fake Ayahs for multiple Surahs with valid Juz references
func (af *AyahFaker) GenerateForAllSurahsWithValidJuz(surahs []SurahModel, juzs []JuzModel) []AyahModel {
	var allAyahs []AyahModel
	
	// Create a map of valid Juz IDs for quick lookup
	validJuzIDs := make([]int, len(juzs))
	for i, juz := range juzs {
		validJuzIDs[i] = juz.ID
	}
	
	for _, surah := range surahs {
		ayahs := make([]AyahModel, surah.AyahCount)
		for i := 0; i < surah.AyahCount; i++ {
			// Generate ayah with valid Juz reference
			juzID := validJuzIDs[rand.Intn(len(validJuzIDs))]
			
			ayahs[i] = AyahModel{
				SurahID:      surah.ID,
				AyahID:       i + 1,
				Text:         af.getRandomArabicText(),
				PageNumber:   rand.Intn(604) + 1,
				JuzNumber:    juzID,
				HizbNumber:   rand.Intn(60) + 1,
				ManzilNumber: rand.Intn(7) + 1,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}
		}
		allAyahs = append(allAyahs, ayahs...)
	}
	return allAyahs
}

// getRandomArabicText returns a random Arabic text for Ayah
func (af *AyahFaker) getRandomArabicText() string {
	arabicTexts := []string{
		"بِسْمِ اللَّهِ الرَّحْمَٰنِ الرَّحِيمِ",
		"الْحَمْدُ لِلَّهِ رَبِّ الْعَالَمِينَ",
		"الرَّحْمَٰنِ الرَّحِيمِ",
		"مَالِكِ يَوْمِ الدِّينِ",
		"إِيَّاكَ نَعْبُدُ وَإِيَّاكَ نَسْتَعِينُ",
		"اهْدِنَا الصِّرَاطَ الْمُسْتَقِيمَ",
		"صِرَاطَ الَّذِينَ أَنْعَمْتَ عَلَيْهِمْ غَيْرِ الْمَغْضُوبِ عَلَيْهِمْ وَلَا الضَّالِّينَ",
		"الم",
		"ذَٰلِكَ الْكِتَابُ لَا رَيْبَ فِيهِ هُدًى لِّلْمُتَّقِينَ",
		"الَّذِينَ يُؤْمِنُونَ بِالْغَيْبِ وَيُقِيمُونَ الصَّلَاةَ وَمِمَّا رَزَقْنَاهُمْ يُنفِقُونَ",
	}
	return arabicTexts[rand.Intn(len(arabicTexts))]
}

// JuzFaker generates fake Juz data
type JuzFaker struct{}

// NewJuzFaker creates a new JuzFaker instance
func NewJuzFaker() *JuzFaker {
	return &JuzFaker{}
}

// Generate creates a fake JuzModel
func (jf *JuzFaker) Generate(id int) JuzModel {
	// Realistic Juz boundaries (simplified)
	juzBoundaries := []struct {
		StartSurah, StartAyah, EndSurah, EndAyah int
	}{
		{1, 1, 2, 141},    // Juz 1
		{2, 142, 2, 252},  // Juz 2
		{2, 253, 3, 92},   // Juz 3
		{3, 93, 4, 23},    // Juz 4
		{4, 24, 4, 147},   // Juz 5
		{4, 148, 5, 81},   // Juz 6
		{5, 82, 6, 110},   // Juz 7
		{6, 111, 7, 87},   // Juz 8
		{7, 88, 8, 40},    // Juz 9
		{8, 41, 9, 92},    // Juz 10
	}

	var startSurah, startAyah, endSurah, endAyah int
	if id <= len(juzBoundaries) {
		boundary := juzBoundaries[id-1]
		startSurah = boundary.StartSurah
		startAyah = boundary.StartAyah
		endSurah = boundary.EndSurah
		endAyah = boundary.EndAyah
	} else {
		// Generate random boundaries for Juz beyond predefined ones
		startSurah = rand.Intn(114) + 1
		startAyah = rand.Intn(50) + 1
		endSurah = startSurah + rand.Intn(5) + 1
		if endSurah > 114 {
			endSurah = 114
		}
		endAyah = rand.Intn(100) + 1
	}

	return JuzModel{
		ID:         id,
		StartSurah: startSurah,
		StartAyah:  startAyah,
		EndSurah:   endSurah,
		EndAyah:    endAyah,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

// GenerateBatch creates multiple fake JuzModel instances
func (jf *JuzFaker) GenerateBatch(count int) []JuzModel {
	juzs := make([]JuzModel, count)
	for i := 0; i < count; i++ {
		juzs[i] = jf.Generate(i + 1)
	}
	return juzs
}

// GenerateBatchWithValidSurahs creates multiple fake JuzModel instances with valid Surah references
func (jf *JuzFaker) GenerateBatchWithValidSurahs(count int, surahs []SurahModel) []JuzModel {
	if len(surahs) == 0 {
		return jf.GenerateBatch(count)
	}
	
	juzs := make([]JuzModel, count)
	for i := 0; i < count; i++ {
		// Use realistic Juz boundaries for first 10, then generate valid random ones
		if i < 10 {
			juzs[i] = jf.Generate(i + 1)
		} else {
			// Generate with valid Surah references
			startSurahIdx := rand.Intn(len(surahs))
			endSurahIdx := startSurahIdx + rand.Intn(min(5, len(surahs)-startSurahIdx))
			if endSurahIdx >= len(surahs) {
				endSurahIdx = len(surahs) - 1
			}
			
			juzs[i] = JuzModel{
				ID:         i + 1,
				StartSurah: surahs[startSurahIdx].ID,
				StartAyah:  rand.Intn(50) + 1,
				EndSurah:   surahs[endSurahIdx].ID,
				EndAyah:    rand.Intn(100) + 1,
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			}
		}
	}
	return juzs
}

// min helper function
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// QuranFaker combines all faker functionality
type QuranFaker struct {
	SurahFaker *SurahFaker
	AyahFaker  *AyahFaker
	JuzFaker   *JuzFaker
}

// NewQuranFaker creates a new QuranFaker instance
func NewQuranFaker() *QuranFaker {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
	
	return &QuranFaker{
		SurahFaker: NewSurahFaker(),
		AyahFaker:  NewAyahFaker(),
		JuzFaker:   NewJuzFaker(),
	}
}

// GenerateCompleteQuranData creates a complete set of fake Quran data
func (qf *QuranFaker) GenerateCompleteQuranData(surahCount, juzCount int) ([]SurahModel, []AyahModel, []JuzModel) {
	// Generate Surahs
	surahs := qf.SurahFaker.GenerateBatch(surahCount)
	
	// Generate Ayahs for all Surahs
	ayahs := qf.AyahFaker.GenerateForAllSurahs(surahs)
	
	// Generate Juz
	juzs := qf.JuzFaker.GenerateBatch(juzCount)
	
	return surahs, ayahs, juzs
}
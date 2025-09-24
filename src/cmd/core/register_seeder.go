package core

import (
	"fmt"
	"gorm.io/gorm"

	"github.com/achmdndy/safa-life-api/src/infrastructure/audio"
	"github.com/achmdndy/safa-life-api/src/infrastructure/quran"
	"github.com/achmdndy/safa-life-api/src/infrastructure/reciter"
	"github.com/achmdndy/safa-life-api/src/infrastructure/resource"
	"github.com/achmdndy/safa-life-api/src/infrastructure/story"
	"github.com/achmdndy/safa-life-api/src/infrastructure/tafsir"
	"github.com/achmdndy/safa-life-api/src/infrastructure/tajweed"
	"github.com/achmdndy/safa-life-api/src/infrastructure/topic"
	"github.com/achmdndy/safa-life-api/src/infrastructure/translation"
)

// Seeder interface defines the contract for all seeders
type Seeder interface {
	Seed(db *gorm.DB) error
	GetName() string
}

// --- Quran Seeder ---
type QuranSeeder struct{ name string }

func NewQuranSeeder() *QuranSeeder { return &QuranSeeder{name: "quran"} }
func (s *QuranSeeder) GetName() string { return s.name }
func (s *QuranSeeder) Seed(db *gorm.DB) error {
	fmt.Println("🌱 Seeding Quran data...")
	
	// Check if surahs already exist to prevent duplicates
	var surahCount int64
	if err := db.Model(&quran.SurahModel{}).Count(&surahCount).Error; err != nil {
		return fmt.Errorf("failed to check existing surahs: %w", err)
	}
	if surahCount > 0 {
		fmt.Println("⚠️ Surahs already exist, skipping surah seeding...")
	} else {
		faker := quran.NewQuranFaker()
		surahs := faker.SurahFaker.GenerateBatch(20)
		if err := db.Create(&surahs).Error; err != nil {
			return fmt.Errorf("failed to seed surahs: %w", err)
		}
		fmt.Printf("✅ Surahs seeded successfully! (%d surahs)\n", len(surahs))
	}
	
	// Check if juzs already exist to prevent duplicates
	var juzCount int64
	if err := db.Model(&quran.JuzModel{}).Count(&juzCount).Error; err != nil {
		return fmt.Errorf("failed to check existing juzs: %w", err)
	}
	if juzCount > 0 {
		fmt.Println("⚠️ Juzs already exist, skipping juz seeding...")
	} else {
		// Get existing surahs for proper relationships
		var surahs []quran.SurahModel
		if err := db.Find(&surahs).Error; err != nil {
			return fmt.Errorf("failed to get surahs: %w", err)
		}
		if len(surahs) == 0 {
			return fmt.Errorf("no surahs found, cannot seed juzs")
		}
		
		faker := quran.NewQuranFaker()
		juzs := faker.JuzFaker.GenerateBatchWithValidSurahs(30, surahs)
		if err := db.Create(&juzs).Error; err != nil {
			return fmt.Errorf("failed to seed juzs: %w", err)
		}
		fmt.Printf("✅ Juzs seeded successfully! (%d juzs)\n", len(juzs))
	}
	
	// Check if ayahs already exist to prevent duplicates
	var ayahCount int64
	if err := db.Model(&quran.AyahModel{}).Count(&ayahCount).Error; err != nil {
		return fmt.Errorf("failed to check existing ayahs: %w", err)
	}
	if ayahCount > 0 {
		fmt.Println("⚠️ Ayahs already exist, skipping ayah seeding...")
	} else {
		// Get existing surahs and juzs for proper relationships
		var surahs []quran.SurahModel
		if err := db.Find(&surahs).Error; err != nil {
			return fmt.Errorf("failed to get surahs: %w", err)
		}
		var juzs []quran.JuzModel
		if err := db.Find(&juzs).Error; err != nil {
			return fmt.Errorf("failed to get juzs: %w", err)
		}
		if len(surahs) == 0 || len(juzs) == 0 {
			return fmt.Errorf("no surahs or juzs found, cannot seed ayahs")
		}
		
		faker := quran.NewQuranFaker()
		ayahs := faker.AyahFaker.GenerateForAllSurahsWithValidJuz(surahs, juzs)
		if err := db.CreateInBatches(ayahs, 1000).Error; err != nil {
			return fmt.Errorf("failed to seed ayahs: %w", err)
		}
		fmt.Printf("✅ Ayahs seeded successfully! (%d ayahs)\n", len(ayahs))
	}
	
	fmt.Println("✅ Quran data seeding completed!")
	return nil
}

// --- Reciter Seeder ---
type ReciterSeeder struct{ name string }

func NewReciterSeeder() *ReciterSeeder { return &ReciterSeeder{name: "reciter"} }
func (s *ReciterSeeder) GetName() string { return s.name }
func (s *ReciterSeeder) Seed(db *gorm.DB) error {
	fmt.Println("🌱 Seeding Reciter data...")
	
	// Check if reciters already exist to prevent duplicates
	var count int64
	if err := db.Model(&reciter.ReciterModel{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to check existing reciters: %w", err)
	}
	if count > 0 {
		fmt.Println("⚠️ Reciters already exist, skipping...")
		return nil
	}
	
	// Create predefined reciters to avoid duplicates
	reciters := []reciter.ReciterModel{
		{ID: "mishary_rashid", Name: "Mishary Rashid Alafasy", Style: "Hafs"},
		{ID: "abdul_basit", Name: "Abdul Basit Abdus Samad", Style: "Hafs"},
		{ID: "saud_shuraim", Name: "Saud Al-Shuraim", Style: "Hafs"},
		{ID: "maher_muaiqly", Name: "Maher Al Muaiqly", Style: "Hafs"},
		{ID: "saad_ghamdi", Name: "Saad Al Ghamdi", Style: "Hafs"},
	}
	
	if err := db.Create(&reciters).Error; err != nil {
		return fmt.Errorf("failed to seed reciters: %w", err)
	}
	fmt.Printf("✅ Reciter data seeded successfully! (%d reciters)\n", len(reciters))
	return nil
}

// --- Audio Seeder ---
type AudioSeeder struct{ name string }

func NewAudioSeeder() *AudioSeeder { return &AudioSeeder{name: "audio"} }
func (s *AudioSeeder) GetName() string { return s.name }
func (s *AudioSeeder) Seed(db *gorm.DB) error {
	fmt.Println("🌱 Seeding Audio data...")
	
	// Check if audio data already exists to prevent duplicates
	var count int64
	if err := db.Model(&audio.AyahAudioModel{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to check existing audio: %w", err)
	}
	if count > 0 {
		fmt.Println("⚠️ Audio data already exists, skipping...")
		return nil
	}
	
	// Get existing reciters and surahs for proper relationships
	var reciters []reciter.ReciterModel
	if err := db.Find(&reciters).Error; err != nil {
		return fmt.Errorf("failed to get reciters: %w", err)
	}
	if len(reciters) == 0 {
		return fmt.Errorf("no reciters found, please seed reciters first")
	}
	
	var surahs []quran.SurahModel
	if err := db.Limit(5).Find(&surahs).Error; err != nil {
		return fmt.Errorf("failed to get surahs: %w", err)
	}
	if len(surahs) == 0 {
		return fmt.Errorf("no surahs found, please seed quran first")
	}
	
	faker := audio.NewAyahAudioFaker()
	var audios []audio.AyahAudioModel
	
	// Generate audio for first few surahs and reciters
	for _, reciterModel := range reciters[:2] { // Use first 2 reciters
		for _, surah := range surahs[:3] { // Use first 3 surahs
			for ayahID := 1; ayahID <= min(surah.AyahCount, 5); ayahID++ { // First 5 ayahs per surah
				audios = append(audios, faker.Generate(reciterModel.ID, surah.ID, ayahID))
			}
		}
	}
	
	if err := db.CreateInBatches(audios, 100).Error; err != nil {
		return fmt.Errorf("failed to seed audio: %w", err)
	}
	fmt.Printf("✅ Audio data seeded successfully! (%d audio files)\n", len(audios))
	return nil
}

// --- Translation Seeder ---
type TranslationSeeder struct{ name string }

func NewTranslationSeeder() *TranslationSeeder { return &TranslationSeeder{name: "translation"} }
func (s *TranslationSeeder) GetName() string { return s.name }
func (s *TranslationSeeder) Seed(db *gorm.DB) error {
	fmt.Println("🌱 Seeding Translation data...")
	
	// Check if translation editions already exist to prevent duplicates
	var count int64
	if err := db.Model(&translation.TranslationModel{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to check existing translations: %w", err)
	}
	if count > 0 {
		fmt.Println("⚠️ Translation editions already exist, skipping...")
		return nil
	}
	
	// Create predefined translation editions
	editions := []translation.TranslationModel{
		{ID: "en_asad", Name: "The Message of The Qur'an", Author: "Muhammad Asad", Language: "English"},
		{ID: "en_sahih", Name: "Sahih International", Author: "Sahih International", Language: "English"},
		{ID: "id_indonesian", Name: "Indonesian Ministry of Religious Affairs", Author: "Kemenag RI", Language: "Indonesian"},
		{ID: "ar_jalalayn", Name: "Tafsir al-Jalalayn", Author: "Jalal al-Din al-Mahalli and Jalal al-Din al-Suyuti", Language: "Arabic"},
	}
	
	if err := db.Create(&editions).Error; err != nil {
		return fmt.Errorf("failed to seed translation editions: %w", err)
	}
	
	// Get existing surahs for proper relationships
	var surahs []quran.SurahModel
	if err := db.Limit(5).Find(&surahs).Error; err != nil {
		return fmt.Errorf("failed to get surahs: %w", err)
	}
	if len(surahs) == 0 {
		return fmt.Errorf("no surahs found, please seed quran first")
	}
	
	faker := translation.NewTranslationFaker()
	var ayahTranslations []translation.AyahTranslationModel
	
	// Generate translations for first few surahs and editions
	for _, edition := range editions[:2] { // Use first 2 editions
		for _, surah := range surahs[:3] { // Use first 3 surahs
			for ayahID := 1; ayahID <= min(surah.AyahCount, 7); ayahID++ { // First 7 ayahs per surah
				ayahTranslations = append(ayahTranslations, faker.AyahFaker.Generate(edition.ID, surah.ID, ayahID))
			}
		}
	}
	
	if err := db.CreateInBatches(ayahTranslations, 100).Error; err != nil {
		return fmt.Errorf("failed to seed ayah translations: %w", err)
	}
	
	fmt.Printf("✅ Translation data seeded successfully! (%d editions, %d ayah translations)\n", len(editions), len(ayahTranslations))
	return nil
}

// --- Tafsir Seeder ---
type TafsirSeeder struct{ name string }

func NewTafsirSeeder() *TafsirSeeder { return &TafsirSeeder{name: "tafsir"} }
func (s *TafsirSeeder) GetName() string { return s.name }
func (s *TafsirSeeder) Seed(db *gorm.DB) error {
	fmt.Println("🌱 Seeding Tafsir data...")
	
	// Check if tafsir editions already exist to prevent duplicates
	var count int64
	if err := db.Model(&tafsir.TafsirModel{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to check existing tafsirs: %w", err)
	}
	if count > 0 {
		fmt.Println("⚠️ Tafsir editions already exist, skipping...")
		return nil
	}
	
	// Create predefined tafsir editions
	editions := []tafsir.TafsirModel{
		{ID: "en_ibn_kathir", Name: "Tafsir Ibn Kathir", Author: "Ibn Kathir", Language: "English"},
		{ID: "id_jalalayn", Name: "Tafsir al-Jalalayn", Author: "Jalal al-Din al-Mahalli and Jalal al-Din al-Suyuti", Language: "Indonesian"},
		{ID: "ar_tabari", Name: "Tafsir al-Tabari", Author: "Muhammad ibn Jarir al-Tabari", Language: "Arabic"},
		{ID: "en_maududi", Name: "Tafhim al-Qur'an", Author: "Sayyid Abul Ala Maududi", Language: "English"},
	}
	
	if err := db.Create(&editions).Error; err != nil {
		return fmt.Errorf("failed to seed tafsir editions: %w", err)
	}
	
	// Get existing surahs for proper relationships
	var surahs []quran.SurahModel
	if err := db.Limit(5).Find(&surahs).Error; err != nil {
		return fmt.Errorf("failed to get surahs: %w", err)
	}
	if len(surahs) == 0 {
		return fmt.Errorf("no surahs found, please seed quran first")
	}
	
	faker := tafsir.NewTafsirFaker()
	var ayahTafsirs []tafsir.AyahTafsirModel
	
	// Generate tafsirs for first few surahs and editions
	for _, edition := range editions[:2] { // Use first 2 editions
		for _, surah := range surahs[:3] { // Use first 3 surahs
			for ayahID := 1; ayahID <= min(surah.AyahCount, 7); ayahID++ { // First 7 ayahs per surah
				ayahTafsirs = append(ayahTafsirs, faker.AyahFaker.Generate(edition.ID, surah.ID, ayahID))
			}
		}
	}
	
	if err := db.CreateInBatches(ayahTafsirs, 100).Error; err != nil {
		return fmt.Errorf("failed to seed ayah tafsirs: %w", err)
	}
	
	fmt.Printf("✅ Tafsir data seeded successfully! (%d editions, %d ayah tafsirs)\n", len(editions), len(ayahTafsirs))
	return nil
}

// --- Tajweed Seeder ---
type TajweedSeeder struct{ name string }

func NewTajweedSeeder() *TajweedSeeder { return &TajweedSeeder{name: "tajweed"} }
func (s *TajweedSeeder) GetName() string { return s.name }
func (s *TajweedSeeder) Seed(db *gorm.DB) error {
	fmt.Println("🌱 Seeding Tajweed data...")
	
	// Check if tajweed rules already exist to prevent duplicates
	var count int64
	if err := db.Model(&tajweed.TajweedRuleModel{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to check existing tajweed rules: %w", err)
	}
	if count > 0 {
		fmt.Println("⚠️ Tajweed rules already exist, skipping...")
		return nil
	}
	
	faker := tajweed.NewTajweedFaker()
	rules := faker.RuleFaker.GenerateBatch(9) // Generate all 9 predefined rules
	if err := db.Create(&rules).Error; err != nil {
		return fmt.Errorf("failed to seed tajweed rules: %w", err)
	}
	
	// Get existing surahs for proper relationships
	var surahs []quran.SurahModel
	if err := db.Limit(5).Find(&surahs).Error; err != nil {
		return fmt.Errorf("failed to get surahs: %w", err)
	}
	if len(surahs) == 0 {
		return fmt.Errorf("no surahs found, please seed quran first")
	}
	
	// Generate tajweed data for first few surahs
	var ayahTajweeds []tajweed.AyahTajweedModel
	for _, surah := range surahs[:3] { // Use first 3 surahs
		for ayahID := 1; ayahID <= min(surah.AyahCount, 5); ayahID++ { // First 5 ayahs per surah
			ayahTajweed := faker.AyahFaker.Generate("default", surah.ID, ayahID, 5)
			ayahTajweeds = append(ayahTajweeds, ayahTajweed)
		}
	}
	
	if err := db.CreateInBatches(ayahTajweeds, 100).Error; err != nil {
		return fmt.Errorf("failed to seed ayah tajweed: %w", err)
	}
	
	fmt.Printf("✅ Tajweed data seeded successfully! (%d rules, %d ayah tajweeds)\n", len(rules), len(ayahTajweeds))
	return nil
}

// --- Story Seeder ---
type StorySeeder struct{ name string }

func NewStorySeeder() *StorySeeder { return &StorySeeder{name: "story"} }
func (s *StorySeeder) GetName() string { return s.name }
func (s *StorySeeder) Seed(db *gorm.DB) error {
	fmt.Println("🌱 Seeding Story data...")
	
	// Check if stories already exist to prevent duplicates
	var count int64
	if err := db.Model(&story.StoryModel{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to check existing stories: %w", err)
	}
	if count > 0 {
		fmt.Println("⚠️ Stories already exist, skipping...")
		return nil
	}
	
	// Create predefined stories
	stories := []story.StoryModel{
		{ID: "prophet_musa", Title: "The Story of Prophet Musa", Summary: "The life and struggles of Prophet Musa (Moses)...", Characters: []string{"Musa", "Firaun", "Harun"}},
		{ID: "prophet_yusuf", Title: "The Story of Prophet Yusuf", Summary: "The story of Prophet Yusuf (Joseph) and his brothers.", Characters: []string{"Yusuf", "Yaqub", "Brothers"}},
		{ID: "prophet_ibrahim", Title: "The Story of Prophet Ibrahim", Summary: "The story of Prophet Ibrahim and his tests from Allah.", Characters: []string{"Ibrahim", "Ismail", "Hajar"}},
		{ID: "ashab_kahf", Title: "The People of the Cave", Summary: "The story of the young believers who slept in the cave.", Characters: []string{"Young Believers", "Dog"}},
	}
	
	if err := db.Create(&stories).Error; err != nil {
		return fmt.Errorf("failed to seed stories: %w", err)
	}
	
	// Get existing surahs for proper relationships
	var surahs []quran.SurahModel
	if err := db.Find(&surahs).Error; err != nil {
		return fmt.Errorf("failed to get surahs: %w", err)
	}
	if len(surahs) == 0 {
		return fmt.Errorf("no surahs found, please seed quran first")
	}
	
	faker := story.NewStoryFaker()
	var ayahLinks []story.StoryAyahModel
	
	// Generate story-ayah links for each story
	for _, storyModel := range stories {
		// Generate 3-5 ayah links per story
		linkCount := 3 + (len(storyModel.ID) % 3) // 3-5 links per story
		for i := 0; i < linkCount; i++ {
			ayahLinks = append(ayahLinks, faker.AyahFaker.Generate(storyModel.ID))
		}
	}
	
	if err := db.CreateInBatches(ayahLinks, 100).Error; err != nil {
		return fmt.Errorf("failed to seed story ayah links: %w", err)
	}
	
	fmt.Printf("✅ Story data seeded successfully! (%d stories, %d ayah links)\n", len(stories), len(ayahLinks))
	return nil
}

// --- Topic Seeder ---
type TopicSeeder struct{ name string }

func NewTopicSeeder() *TopicSeeder { return &TopicSeeder{name: "topic"} }
func (s *TopicSeeder) GetName() string { return s.name }
func (s *TopicSeeder) Seed(db *gorm.DB) error {
	fmt.Println("🌱 Seeding Topic data...")
	
	// Check if topics already exist to prevent duplicates
	var count int64
	if err := db.Model(&topic.TopicModel{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to check existing topics: %w", err)
	}
	if count > 0 {
		fmt.Println("⚠️ Topics already exist, skipping...")
		return nil
	}
	
	// Create predefined topics
	topics := []topic.TopicModel{
		{ID: "patience", Name: "Patience (Sabr)", Description: "Verses about patience and perseverance in the face of trials"},
		{ID: "charity", Name: "Charity (Sadaqa)", Description: "Verses about giving charity and helping the needy"},
		{ID: "prayer", Name: "Prayer (Salah)", Description: "Verses about the importance and method of prayer"},
		{ID: "forgiveness", Name: "Forgiveness (Maghfira)", Description: "Verses about Allah's forgiveness and forgiving others"},
		{ID: "justice", Name: "Justice (Adl)", Description: "Verses about justice and fairness in Islam"},
		{ID: "knowledge", Name: "Knowledge (Ilm)", Description: "Verses about seeking and valuing knowledge"},
	}
	
	if err := db.Create(&topics).Error; err != nil {
		return fmt.Errorf("failed to seed topics: %w", err)
	}
	
	// Get existing surahs for proper relationships
	var surahs []quran.SurahModel
	if err := db.Find(&surahs).Error; err != nil {
		return fmt.Errorf("failed to get surahs: %w", err)
	}
	if len(surahs) == 0 {
		return fmt.Errorf("no surahs found, please seed quran first")
	}
	
	faker := topic.NewTopicFaker()
	var ayahLinks []topic.TopicAyahModel
	
	// Generate topic-ayah links for each topic
	for _, topicModel := range topics {
		// Generate 4-7 ayah links per topic
		linkCount := 4 + (len(topicModel.ID) % 4) // 4-7 links per topic
		for i := 0; i < linkCount; i++ {
			ayahLinks = append(ayahLinks, faker.AyahFaker.Generate(topicModel.ID))
		}
	}
	
	if err := db.CreateInBatches(ayahLinks, 100).Error; err != nil {
		return fmt.Errorf("failed to seed topic ayah links: %w", err)
	}
	
	fmt.Printf("✅ Topic data seeded successfully! (%d topics, %d ayah links)\n", len(topics), len(ayahLinks))
	return nil
}

// --- Resource Seeder ---
type ResourceSeeder struct{ name string }

func NewResourceSeeder() *ResourceSeeder { return &ResourceSeeder{name: "resource"} }
func (s *ResourceSeeder) GetName() string { return s.name }
func (s *ResourceSeeder) Seed(db *gorm.DB) error {
	fmt.Println("🌱 Seeding Resource data...")
	
	// Check if resources already exist to prevent duplicates
	var count int64
	if err := db.Model(&resource.ResourceModel{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to check existing resources: %w", err)
	}
	if count > 0 {
		fmt.Println("⚠️ Resources already exist, skipping...")
		return nil
	}
	
	// Create predefined resources
	resources := []resource.ResourceModel{
		{
			ID:            "en_asad_trans",
			Type:          "translation",
			Name:          "The Message of The Qur'an by Muhammad Asad",
			Language:      "en",
			Version:       "1.0.0",
			DownloadURL:   "https://example.com/en_asad.zip",
			Size:          2048,
			IsInstalled:   false,
		},
		{
			ID:            "id_kemenag_trans",
			Type:          "translation",
			Name:          "Indonesian Ministry of Religious Affairs Translation",
			Language:      "id",
			Version:       "2.1.0",
			DownloadURL:   "https://example.com/id_kemenag.zip",
			Size:          1536,
			IsInstalled:   true,
		},
		{
			ID:            "mishary_audio",
			Type:          "audio",
			Name:          "Mishary Rashid Alafasy Audio Collection",
			Language:      "ar",
			Version:       "1.5.0",
			DownloadURL:   "https://example.com/mishary_audio.zip",
			Size:          512000,
			IsInstalled:   false,
		},
		{
			ID:            "ibn_kathir_tafsir",
			Type:          "tafsir",
			Name:          "Tafsir Ibn Kathir Complete",
			Language:      "en",
			Version:       "3.0.0",
			DownloadURL:   "https://example.com/ibn_kathir.zip",
			Size:          8192,
			IsInstalled:   false,
		},
	}
	
	if err := db.Create(&resources).Error; err != nil {
		return fmt.Errorf("failed to seed resources: %w", err)
	}
	
	fmt.Printf("✅ Resource data seeded successfully! (%d resources)\n", len(resources))
	return nil
}

// --- Seeder Registry ---
type SeederRegistry struct {
	seeders map[string]Seeder
}

func NewSeederRegistry() *SeederRegistry {
	registry := &SeederRegistry{
		seeders: make(map[string]Seeder),
	}

	registry.Register(NewQuranSeeder())
	registry.Register(NewReciterSeeder())
	registry.Register(NewAudioSeeder())
	registry.Register(NewTranslationSeeder())
	registry.Register(NewTafsirSeeder())
	registry.Register(NewTajweedSeeder())
	registry.Register(NewStorySeeder())
	registry.Register(NewTopicSeeder())
	registry.Register(NewResourceSeeder())

	return registry
}

func (sr *SeederRegistry) Register(seeder Seeder) {
	sr.seeders[seeder.GetName()] = seeder
}

func (sr *SeederRegistry) GetSeeder(name string) (Seeder, bool) {
	seeder, exists := sr.seeders[name]
	return seeder, exists
}

func (sr *SeederRegistry) GetSeederNames() []string {
	names := make([]string, 0, len(sr.seeders))
	for name := range sr.seeders {
		names = append(names, name)
	}
	return names
}

func (sr *SeederRegistry) SeedAll(db *gorm.DB) error {
	fmt.Println("🌱 Running all seeders...")
	// Define a seed order to handle dependencies
	seedOrder := []string{"quran", "reciter", "translation", "tafsir", "tajweed", "story", "topic", "audio", "resource"}

	for _, name := range seedOrder {
		seeder, exists := sr.GetSeeder(name)
		if !exists {
			fmt.Printf("⚠️ Seeder '%s' not found, skipping.\n", name)
			continue
		}
		fmt.Printf("🔄 Running seeder: %s\n", name)
		if err := seeder.Seed(db); err != nil {
			return fmt.Errorf("seeder %s failed: %w", name, err)
		}
	}

	fmt.Println("✅ All seeders completed successfully!")
	return nil
}

func (sr *SeederRegistry) SeedSpecific(db *gorm.DB, name string) error {
	seeder, exists := sr.GetSeeder(name)
	if !exists {
		return fmt.Errorf("seeder '%s' not found. Available seeders: %v", name, sr.GetSeederNames())
	}

	fmt.Printf("🌱 Running specific seeder: %s\n", name)
	if err := seeder.Seed(db); err != nil {
		return fmt.Errorf("seeder %s failed: %w", name, err)
	}

	fmt.Printf("✅ Seeder %s completed successfully!\n", name)
	return nil
}

package seeders

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	quran "github.com/safalife/core-api/src/infrastructure/quran"
)

// SurahData represents the JSON structure of surah files
type SurahData struct {
	Index string            `json:"index"`
	Name  string            `json:"name"`
	Verse map[string]string `json:"verse"`
	Count int               `json:"count"`
	Juz   []JuzInfo         `json:"juz"`
}

type JuzInfo struct {
	Index string    `json:"index"`
	Verse VerseInfo `json:"verse"`
}

type VerseInfo struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// SurahSeeder implements the Seeder interface for seeding Quran data from JSON files
type SurahSeeder struct{}

// GetName returns the name of this seeder
func (s *SurahSeeder) GetName() string {
	return "surah"
}

// Seed implements the Seeder interface to seed Quran data from JSON files
func (s *SurahSeeder) Seed(db *gorm.DB) error {
	return s.SeedWithForce(db, false)
}

// SeedWithForce seeds Quran data from JSON files with optional force parameter
func (s *SurahSeeder) SeedWithForce(db *gorm.DB, force bool) error {
	// Use the correct data directory path
	dataDir := "data/surah"

	// Check if data directory exists
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		return fmt.Errorf("data directory not found: %s", dataDir)
	}

	return SurahMigrateWithForce(db, dataDir, force)
}

func SurahMigrate(db *gorm.DB, dataDir string) error {
	return SurahMigrateWithForce(db, dataDir, false)
}

func SurahMigrateWithForce(db *gorm.DB, dataDir string, force bool) error {
	fmt.Println("Starting Surah migration from JSON files...")

	// Get all surah JSON files
	files, err := filepath.Glob(filepath.Join(dataDir, "surah_*.json"))
	if err != nil {
		return fmt.Errorf("failed to find surah files: %v", err)
	}

	if len(files) == 0 {
		return fmt.Errorf("no surah files found in directory: %s", dataDir)
	}

	// Sort files to ensure proper order
	sort.Strings(files)

	fmt.Printf("Found %d surah files to process\n", len(files))

	createdSurahs := 0
	createdAyahs := 0
	processedJuzs := make(map[string]bool) // Track processed juz to avoid duplicates

	for _, file := range files {
		fmt.Printf("Processing file: %s\n", filepath.Base(file))

		// Read JSON file
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Failed to read file %s: %v\n", file, err)
			continue
		}

		// Parse JSON
		var surahData SurahData
		if jsonErr := json.Unmarshal(data, &surahData); jsonErr != nil {
			fmt.Printf("Failed to parse JSON from %s: %v\n", file, jsonErr)
			continue
		}

		// Convert index to integer
		surahIndex, err := strconv.Atoi(surahData.Index)
		if err != nil {
			fmt.Printf("Invalid surah index '%s' in file %s: %v\n", surahData.Index, file, err)
			continue
		}

		// Check if surah already exists
		var existingSurah quran.SurahModel
		if err := db.Where("revelation_order = ?", surahIndex).First(&existingSurah).Error; err == nil {
			if !force {
				fmt.Printf("Surah %d (%s) already exists, skipping\n", surahIndex, surahData.Name)
				continue
			} else {
				fmt.Printf("Surah %d (%s) already exists, force deleting and recreating\n", surahIndex, surahData.Name)
				// Delete existing ayahs first (due to foreign key constraints)
				if err := db.Where("surah_id = ?", existingSurah.ID).Delete(&quran.AyahModel{}).Error; err != nil {
					fmt.Printf("Failed to delete existing ayahs for surah %d: %v\n", surahIndex, err)
					continue
				}
				// Delete the surah
				if err := db.Delete(&existingSurah).Error; err != nil {
					fmt.Printf("Failed to delete existing surah %d: %v\n", surahIndex, err)
					continue
				}
			}
		}

		// Create Surah record
		surah := &quran.SurahModel{
			ID:              uuid.New(),
			NameArabic:      surahData.Name, // We'll use the name from JSON
			NameEnglish:     surahData.Name,
			RevelationPlace: "Unknown", // This info is not in the JSON, can be updated later
			RevelationOrder: surahIndex,
			AyahCount:       surahData.Count,
			CreatedBy:       "system",
			UpdatedBy:       "system",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		// Insert Surah
		if err := db.Create(surah).Error; err != nil {
			fmt.Printf("Failed to create surah '%s': %v\n", surahData.Name, err)
			continue
		}

		fmt.Printf("Created surah: %s (Index: %d)\n", surahData.Name, surahIndex)
		createdSurahs++

		// Create Ayah records - handle both verse_0 (Bismillah) and regular verses
		ayahCount := 0

		// Get all verse keys and sort them properly
		verseKeys := make([]string, 0, len(surahData.Verse))
		for key := range surahData.Verse {
			verseKeys = append(verseKeys, key)
		}

		// Sort verse keys to ensure proper order (verse_0, verse_1, verse_2, etc.)
		sort.Slice(verseKeys, func(i, j int) bool {
			numI := extractVerseNumber(verseKeys[i])
			numJ := extractVerseNumber(verseKeys[j])
			return numI < numJ
		})

		for _, verseKey := range verseKeys {
			verseText := surahData.Verse[verseKey]

			// Extract verse number from key (verse_0, verse_1, verse_2, etc.)
			verseNum := extractVerseNumber(verseKey)

			// Find which juz this ayah belongs to
			juzNumber := findJuzForVerse(surahData.Juz, verseKey)

			ayah := &quran.AyahModel{
				ID:           uuid.New(),
				SurahID:      surah.ID,
				Text:         verseText,
				PageNumber:   1, // Default value, can be calculated later
				JuzNumber:    juzNumber,
				HizbNumber:   1, // Default value, can be calculated later
				ManzilNumber: 1, // Default value, can be calculated later
				CreatedBy:    "system",
				UpdatedBy:    "system",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}

			if err := db.Create(ayah).Error; err != nil {
				fmt.Printf("Failed to create ayah %d for surah %s: %v\n", verseNum, surahData.Name, err)
				continue
			}

			ayahCount++
		}

		fmt.Printf("Created %d ayahs for surah %s\n", ayahCount, surahData.Name)
		createdAyahs += ayahCount

		// Process Juz information (we'll handle this in a separate function to avoid duplicates)
		for _, juzInfo := range surahData.Juz {
			juzKey := fmt.Sprintf("%s_%s_%s", juzInfo.Index, juzInfo.Verse.Start, juzInfo.Verse.End)
			if !processedJuzs[juzKey] {
				processedJuzs[juzKey] = true
				// Note: Juz creation is complex as it requires start/end ayah IDs
				// This would need to be handled after all ayahs are created
			}
		}
	}

	fmt.Printf("Successfully seeded %d surahs and %d ayahs\n", createdSurahs, createdAyahs)
	return nil
}

// findJuzForVerse finds which juz a specific verse belongs to
func findJuzForVerse(juzList []JuzInfo, verseKey string) int {
	for _, juz := range juzList {
		startNum := extractVerseNumber(juz.Verse.Start)
		endNum := extractVerseNumber(juz.Verse.End)
		currentNum := extractVerseNumber(verseKey)

		if currentNum >= startNum && currentNum <= endNum {
			juzIndex, _ := strconv.Atoi(juz.Index)
			return juzIndex
		}
	}
	return 1 // Default to juz 1 if not found
}

// extractVerseNumber extracts the numeric part from verse keys like "verse_1"
func extractVerseNumber(verseKey string) int {
	numStr := strings.TrimPrefix(verseKey, "verse_")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0
	}
	return num
}

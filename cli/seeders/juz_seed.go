package seeders

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	quran "github.com/safalife/core-api/src/infrastructure/quran"
)

// JuzSeeder handles seeding of Juz data
type JuzSeeder struct{}

// Seed implements the Seeder interface for JuzSeeder
func (s *JuzSeeder) Seed(db *gorm.DB) error {
	return s.SeedWithForce(db, false)
}

// SeedWithForce seeds juz data with optional force flag
func (s *JuzSeeder) SeedWithForce(db *gorm.DB, force bool) error {
	return JuzMigrateWithForce(db, force)
}

// JuzMigrate seeds juz data without force
func JuzMigrate(db *gorm.DB) error {
	return JuzMigrateWithForce(db, false)
}

// SurahJuzData represents the structure of surah JSON files for juz processing
type SurahJuzData struct {
	Index string    `json:"index"`
	Name  string    `json:"name"`
	Juz   []JuzInfo `json:"juz"`
}

// JuzSegment represents a segment of a juz within a surah
type JuzSegment struct {
	SurahID     uuid.UUID
	StartAyahID uuid.UUID
	EndAyahID   uuid.UUID
}

// JuzMigrateWithForce seeds juz data with force option
func JuzMigrateWithForce(db *gorm.DB, force bool) error {
	fmt.Println("Starting Juz seeding...")

	// Get data directory path
	dataDir := "/Users/kira/Documents/Al-Quran/backend/rest-api/data/surah"

	// Read all JSON files
	files, err := os.ReadDir(dataDir)
	if err != nil {
		return fmt.Errorf("failed to read data directory: %w", err)
	}

	// Collect all juz information from all surahs
	juzMap := make(map[int][]JuzSegment) // juz number -> segments
	
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		filePath := filepath.Join(dataDir, file.Name())
		jsonData, jsonErr := os.ReadFile(filePath)
		if jsonErr != nil {
			fmt.Printf("Warning: failed to read file %s: %v\n", file.Name(), jsonErr)
			continue
		}

		var surahData SurahJuzData
		if jsonErr := json.Unmarshal(jsonData, &surahData); jsonErr != nil {
			fmt.Printf("Warning: failed to parse JSON from %s: %v\n", file.Name(), jsonErr)
			continue
		}

		// Convert surah index to integer
		surahIndex, parseErr := strconv.Atoi(surahData.Index)
		if parseErr != nil {
			fmt.Printf("Warning: invalid surah index %s in file %s\n", surahData.Index, file.Name())
			continue
		}

		// Get surah from database using revelation_order
		var surahModel quran.SurahModel
		if err := db.Where("revelation_order = ?", surahIndex).First(&surahModel).Error; err != nil {
			fmt.Printf("Warning: surah with revelation_order %d not found in database, skipping juz processing\n", surahIndex)
			continue
		}

		// Process juz information for this surah
		for _, juzInfo := range surahData.Juz {
			juzNumber, parseErr := strconv.Atoi(juzInfo.Index)
			if parseErr != nil {
				fmt.Printf("Warning: invalid juz index %s in surah %d\n", juzInfo.Index, surahIndex)
				continue
			}

			// Get start and end ayah numbers
			startAyahNumber := extractVerseNumber(juzInfo.Verse.Start)
			endAyahNumber := extractVerseNumber(juzInfo.Verse.End)

			if startAyahNumber == 0 || endAyahNumber == 0 {
				fmt.Printf("Warning: invalid verse numbers in juz %d for surah %d\n", juzNumber, surahIndex)
				continue
			}

			// Find ayahs by their position in the surah (we'll need to get them by order)
			var startAyah, endAyah quran.AyahModel
			
			// Get all ayahs for this surah and find the ones at the specified positions
			var allAyahs []quran.AyahModel
			if err := db.Where("surah_id = ?", surahModel.ID).Order("created_at ASC").Find(&allAyahs).Error; err != nil {
				fmt.Printf("Warning: failed to get ayahs for surah %d: %v\n", surahIndex, err)
				continue
			}

			if len(allAyahs) < endAyahNumber {
				fmt.Printf("Warning: not enough ayahs in surah %d (has %d, need %d)\n", surahIndex, len(allAyahs), endAyahNumber)
				continue
			}

			startAyah = allAyahs[startAyahNumber-1] // Convert to 0-based index
			endAyah = allAyahs[endAyahNumber-1]     // Convert to 0-based index

			segment := JuzSegment{
				SurahID:     surahModel.ID,
				StartAyahID: startAyah.ID,
				EndAyahID:   endAyah.ID,
			}

			juzMap[juzNumber] = append(juzMap[juzNumber], segment)
		}
	}

	// Create juz records
	createdJuzs := 0
	for juzNumber := 1; juzNumber <= 30; juzNumber++ {
		segments, exists := juzMap[juzNumber]
		if !exists || len(segments) == 0 {
			fmt.Printf("Warning: no segments found for juz %d\n", juzNumber)
			continue
		}

		// Check if juz already exists (we'll use a combination of start/end to identify)
		var existingJuz quran.JuzModel
		firstSegment := segments[0]
		juzExists := db.Where("start_surah_id = ? AND start_ayah_id = ?", firstSegment.SurahID, firstSegment.StartAyahID).First(&existingJuz).Error == nil

		if juzExists && !force {
			fmt.Printf("Juz %d already exists, skipping\n", juzNumber)
			continue
		}

		if juzExists && force {
			fmt.Printf("Force deleting existing juz %d\n", juzNumber)
			if err := db.Where("start_surah_id = ? AND start_ayah_id = ?", firstSegment.SurahID, firstSegment.StartAyahID).Delete(&quran.JuzModel{}).Error; err != nil {
				fmt.Printf("Warning: failed to delete existing juz %d: %v\n", juzNumber, err)
				continue
			}
		}

		// Sort segments to find overall start and end
		sort.Slice(segments, func(i, j int) bool {
			return segments[i].SurahID.String() < segments[j].SurahID.String()
		})

		lastSegment := segments[len(segments)-1]

		// Create new juz
		juzModel := &quran.JuzModel{
			ID:           uuid.New(),
			StartSurahID: firstSegment.SurahID,
			EndSurahID:   lastSegment.SurahID,
			StartAyahID:  firstSegment.StartAyahID,
			EndAyahID:    lastSegment.EndAyahID,
			CreatedBy:    "system",
			UpdatedBy:    "system",
		}

		if err := db.Create(juzModel).Error; err != nil {
			fmt.Printf("Error creating juz %d: %v\n", juzNumber, err)
			continue
		}

		createdJuzs++
		fmt.Printf("Created juz %d\n", juzNumber)
	}

	fmt.Printf("Successfully seeded %d juzs\n", createdJuzs)
	return nil
}
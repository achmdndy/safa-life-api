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
	quran "github.com/safalife/core-api/src/infrastructure/quran"
	"gorm.io/gorm"
)

// TranslationAyahSeeder implements the Seeder interface for seeding Ayah translations
type TranslationAyahSeeder struct{}

// GetName returns the name of this seeder
func (s *TranslationAyahSeeder) GetName() string {
	return "translation_ayah"
}

// Seed implements the Seeder interface to seed Ayah translations
func (s *TranslationAyahSeeder) Seed(db *gorm.DB) error {
	return s.SeedWithForce(db, false)
}

// SeedWithForce seeds Ayah translations with optional force parameter
func (s *TranslationAyahSeeder) SeedWithForce(db *gorm.DB, force bool) error {
	fmt.Println("Starting Translation Ayah seeding from JSON files...")

	// Directories per language
	langDirs := map[string]string{
		"ar": "data/translation/ar",
		"en": "data/translation/en",
		"id": "data/translation/id",
	}

	// Default edition names per language (used if JSON name is missing)
	defaultEditionNames := map[string]string{
		"ar": "Arabic Translation",
		"en": "English Translation",
		"id": "Indonesian Translation",
	}

	// Ensure directories exist
	for lang, dir := range langDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			return fmt.Errorf("translation data directory not found for %s: %s", lang, dir)
		}
	}

	// Optionally wipe existing translation editions and ayah translations for these languages
	if force {
		if err := s.forceWipeTranslations(db, []string{"ar", "en", "id"}); err != nil {
			return fmt.Errorf("failed wiping existing translations: %w", err)
		}
	}

	totalCreated := 0

	// Process each language
	for lang, dir := range langDirs {
		fmt.Printf("Processing language: %s\n", lang)

		// Read all translation files for this language
		pattern := filepath.Join(dir, fmt.Sprintf("%s_translation_*.json", lang))
		files, globErr := filepath.Glob(pattern)
		if globErr != nil {
			return fmt.Errorf("failed to find translation files for %s: %v", lang, globErr)
		}
		if len(files) == 0 {
			return fmt.Errorf("no %s translation files found in directory: %s", lang, dir)
		}

		// Sort files to ensure proper order
		sort.Strings(files)

		fmt.Printf("Found %d %s translation files to process\n", len(files), lang)

		// Parse first file to get edition name from JSON (if available)
		firstData, firstErr := parseTranslationFile(files[0])
		if firstErr != nil {
			fmt.Printf("Warning: failed to parse first translation file %s: %v\n", files[0], firstErr)
		}
		editionName := defaultEditionNames[lang]
		if firstData != nil && strings.TrimSpace(firstData.Name) != "" {
			editionName = strings.TrimSpace(firstData.Name)
		}

		// Ensure TranslationEdition exists (create if not found) using detected name
		edition, ensureErr := ensureTranslationEdition(db, lang, editionName, "system")
		if ensureErr != nil {
			return fmt.Errorf("failed ensuring translation edition for %s: %w", lang, ensureErr)
		}

		// Process each file
		for _, file := range files {
			fmt.Printf("Processing file: %s\n", filepath.Base(file))

			// Parse translation JSON file
			tData, parseErr := parseTranslationFile(file)
			if parseErr != nil {
				fmt.Printf("Failed to parse translation file %s: %v\n", file, parseErr)
				continue
			}

			// Locate Surah by revelation_order
			surahIndex := tData.Index
			var surah quran.SurahModel
			if surahErr := db.Where("revelation_order = ?", surahIndex).First(&surah).Error; surahErr != nil {
				fmt.Printf("Surah with index %d not found for %s (%s), skipping file: %v\n", surahIndex, lang, filepath.Base(file), surahErr)
				continue
			}

			// Determine if this Surah has verse_0 in original Surah JSON
			hasVerse0, hasErr := surahHasVerse0(surahIndex)
			if hasErr != nil {
				fmt.Printf("Failed to inspect surah_%d.json for verse_0 presence: %v\n", surahIndex, hasErr)
				// Proceed assuming verse_0 exists only if key is in Surah JSON; default to false to reduce misalignment
				hasVerse0 = false
			}

			// Fetch Ayahs for this Surah ordered by creation time to preserve seeding order
			var ayahs []quran.AyahModel
			if fetchErr := db.Where("surah_id = ?", surah.ID).Order("created_at ASC").Find(&ayahs).Error; fetchErr != nil {
				fmt.Printf("Failed to fetch ayahs for surah %d: %v\n", surahIndex, fetchErr)
				continue
			}
			if len(ayahs) == 0 {
				fmt.Printf("No ayahs found for surah %d; ensure Surah seeder ran first\n", surahIndex)
				continue
			}

			// Create Ayah translations
			createdForFile := 0
			verseKeys := make([]string, 0, len(tData.Verse))
			for k := range tData.Verse {
				verseKeys = append(verseKeys, k)
			}
			sort.Slice(verseKeys, func(i, j int) bool {
				return extractVerseNum(verseKeys[i]) < extractVerseNum(verseKeys[j])
			})

			now := time.Now()
			for _, vKey := range verseKeys {
				vNum := extractVerseNum(vKey)
				if vNum < 0 {
					continue
				}

				// Map verse number to ayah index in ayahs slice
				ayahIdx := vNum
				if !hasVerse0 {
					ayahIdx = vNum - 1
				}
				if ayahIdx < 0 || ayahIdx >= len(ayahs) {
					fmt.Printf("Warning: computed ayah index %d out of range for surah %d (hasVerse0=%t). Skipping verse %s\n", ayahIdx, surahIndex, hasVerse0, vKey)
					continue
				}

				a := ayahs[ayahIdx]
				text := tData.Verse[vKey]

				// Check if translation already exists to avoid duplicates
				var existing quran.AyahTranslationModel
				queryErr := db.Where("translation_edition_id = ? AND surah_id = ? AND ayah_id = ?",
					edition.ID, surah.ID, a.ID).First(&existing).Error
				if queryErr == nil {
					// Already exists; optionally update text if force
					if force {
						existing.Text = text
						existing.UpdatedBy = "system"
						existing.UpdatedAt = now
						if saveErr := db.Save(&existing).Error; saveErr != nil {
							fmt.Printf("Failed to update existing translation for surah %d verse %s: %v\n", surahIndex, vKey, saveErr)
						} else {
							createdForFile++
						}
					}
					continue
				}

				// Create new translation entry
				tr := &quran.AyahTranslationModel{
					ID:                   uuid.New(),
					TranslationEditionID: edition.ID,
					SurahID:              surah.ID,
					AyahID:               a.ID,
					Text:                 text,
					CreatedBy:            "system",
					UpdatedBy:            "system",
					CreatedAt:            now,
					UpdatedAt:            now,
				}
				if createErr := db.Create(tr).Error; createErr != nil {
					fmt.Printf("Failed to create translation for surah %d verse %s: %v\n", surahIndex, vKey, createErr)
					continue
				}
				createdForFile++
			}

			fmt.Printf("Created/updated %d translations for surah %d (%s) [%s]\n",
				createdForFile, surahIndex, surah.NameEnglish, lang)
			totalCreated += createdForFile
		}
	}

	fmt.Printf("Successfully processed translations. Total created/updated: %d\n", totalCreated)
	return nil
}

// Translation JSON structure (flexible index handling)
type translationJSON struct {
	Name  string            `json:"name"`
	Verse map[string]string `json:"verse"`
	Count int               `json:"count"`
	// Index can be string (e.g. "001") or number (e.g. 1)
	RawIndex json.RawMessage `json:"index"`
}

type translationData struct {
	Index int
	Name  string
	Verse map[string]string
	Count int
}

func parseTranslationFile(file string) (*translationData, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var tj translationJSON
	if err := json.Unmarshal(bytes, &tj); err != nil {
		return nil, err
	}

	// Parse index that may be string or number
	var idx int
	// Try as string first
	var idxStr string
	if len(tj.RawIndex) > 0 && tj.RawIndex[0] == '"' {
		if err := json.Unmarshal(tj.RawIndex, &idxStr); err == nil {
			i, convErr := strconv.Atoi(strings.TrimLeft(idxStr, "0"))
			if convErr != nil {
				// Fallback to raw Atoi
				i, convErr = strconv.Atoi(idxStr)
			}
			if convErr == nil {
				idx = i
			} else {
				return nil, fmt.Errorf("invalid index string '%s' in %s: %v", idxStr, file, convErr)
			}
		} else {
			return nil, fmt.Errorf("failed to parse index string in %s: %v", file, err)
		}
	} else {
		// Try as number
		var idxNum int
		if err := json.Unmarshal(tj.RawIndex, &idxNum); err == nil {
			idx = idxNum
		} else {
			// Sometimes numbers come as float64
			var idxFloat float64
			if err2 := json.Unmarshal(tj.RawIndex, &idxFloat); err2 == nil {
				idx = int(idxFloat)
			} else {
				return nil, fmt.Errorf("failed to parse index in %s: %v", file, err)
			}
		}
	}

	return &translationData{
		Index: idx,
		Name:  tj.Name,
		Verse: tj.Verse,
		Count: tj.Count,
	}, nil
}

// surahHasVerse0 checks original surah JSON for presence of verse_0
func surahHasVerse0(index int) (bool, error) {
	surahFile := filepath.Join("data/surah", fmt.Sprintf("surah_%d.json", index))
	if _, err := os.Stat(surahFile); os.IsNotExist(err) {
		return false, fmt.Errorf("surah file not found: %s", surahFile)
	}

	type surahJSON struct {
		Verse map[string]string `json:"verse"`
	}

	b, err := os.ReadFile(surahFile)
	if err != nil {
		return false, err
	}

	var sj surahJSON
	if err := json.Unmarshal(b, &sj); err != nil {
		return false, err
	}

	_, ok := sj.Verse["verse_0"]
	return ok, nil
}

// ensureTranslationEdition finds or creates a TranslationEdition for given language
func ensureTranslationEdition(db *gorm.DB, language, name, author string) (*quran.TranslationEditionModel, error) {
	var existing quran.TranslationEditionModel
	// Prefer existing edition by language, regardless of name variations across files
	if err := db.Where("language = ?", language).First(&existing).Error; err == nil {
		return &existing, nil
	}

	te := &quran.TranslationEditionModel{
		ID:        uuid.New(),
		Name:      name,
		Author:    author,
		Language:  language,
		CreatedBy: "system",
		UpdatedBy: "system",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := db.Create(te).Error; err != nil {
		return nil, err
	}
	return te, nil
}

// extractVerseNumber extracts the numeric part from verse keys like "verse_1"
func extractVerseNum(verseKey string) int {
	numStr := strings.TrimPrefix(verseKey, "verse_")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return -1
	}
	return num
}

// forceWipeTranslations deletes translations and editions for specified languages
func (s *TranslationAyahSeeder) forceWipeTranslations(db *gorm.DB, languages []string) error {
	for _, lang := range languages {
		var editions []quran.TranslationEditionModel
		if err := db.Where("language = ?", lang).Find(&editions).Error; err != nil {
			return err
		}
		for _, ed := range editions {
			// Delete translations for this edition
			if err := db.Where("translation_edition_id = ?", ed.ID).Delete(&quran.AyahTranslationModel{}).Error; err != nil {
				return err
			}
			// Delete the edition itself
			if err := db.Delete(&ed).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

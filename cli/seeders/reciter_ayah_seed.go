package seeders

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/safalife/core-api/cli/core"
	domStorage "github.com/safalife/core-api/src/domain/storage"
	quran "github.com/safalife/core-api/src/infrastructure/quran"
	infraStorage "github.com/safalife/core-api/src/infrastructure/storage"
	"gorm.io/gorm"
)

// ReciterAyahSeeder seeds Ayah audio files and uploads MP3s to S3
type ReciterAyahSeeder struct{}

// GetName returns the name of this seeder
func (s *ReciterAyahSeeder) GetName() string { return "reciter_ayah" }

// Seed implements the Seeder interface
func (s *ReciterAyahSeeder) Seed(db *gorm.DB, force bool) error {
	return s.SeedWithForce(db, force)
}

// SeedWithForce seeds ayah audio files with optional force overwrite
func (s *ReciterAyahSeeder) SeedWithForce(db *gorm.DB, force bool) error {
	fmt.Println("Starting Reciter Ayah audio seeding from JSON files and uploading to S3...")

	// Validate S3 config and initialize service
	s3cfg := core.Config.Storage.S3
	if !s3cfg.Enabled {
		return fmt.Errorf("S3 storage is disabled in configuration")
	}
	svc, err := infraStorage.NewS3StorageService(context.Background(), infraStorage.S3Config{
		Enabled:         s3cfg.Enabled,
		Bucket:          s3cfg.Bucket,
		Region:          s3cfg.Region,
		AccessKeyID:     s3cfg.AccessKeyID,
		SecretAccessKey: s3cfg.SecretAccessKey,
		Endpoint:        s3cfg.Endpoint,
		UsePathStyle:    s3cfg.UsePathStyle,
		PublicURLBase:   s3cfg.PublicURLBase,
	})
	if err != nil {
		return fmt.Errorf("failed to initialize S3 storage service: %w", err)
	}

	// Ensure default reciter exists
	reciter, err := ensureReciter(db, "Default Reciter", "Murattal")
	if err != nil {
		return fmt.Errorf("failed ensuring reciter: %w", err)
	}

	audioRoot := "data/audio"
	entries, readErr := os.ReadDir(audioRoot)
	if readErr != nil {
		return fmt.Errorf("failed to read audio root directory %s: %w", audioRoot, readErr)
	}

	// Sort directories to process in order
	dirNames := make([]string, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() {
			dirNames = append(dirNames, e.Name())
		}
	}
	sort.Strings(dirNames)

	totalProcessed := 0
	for _, dir := range dirNames {
		indexPath := filepath.Join(audioRoot, dir, "index.json")
		if _, statErr := os.Stat(indexPath); os.IsNotExist(statErr) {
			// Skip directories without index.json
			continue
		}

		// Parse index.json for this surah
		aData, parseErr := parseAudioIndexFile(indexPath)
		if parseErr != nil {
			fmt.Printf("Failed to parse audio index %s: %v\n", indexPath, parseErr)
			continue
		}

		// Locate Surah by revelation_order
		surahIndex := aData.Index
		var surah quran.SurahModel
		if surahErr := db.Where("revelation_order = ?", surahIndex).First(&surah).Error; surahErr != nil {
			fmt.Printf("Surah with index %d not found for %s, skipping: %v\n", surahIndex, filepath.Base(indexPath), surahErr)
			continue
		}

		// Determine if this Surah has verse_0 in original Surah JSON
		hasV0, v0Err := surahHasVerse0(surahIndex)
		if v0Err != nil {
			fmt.Printf("Warning: could not determine verse_0 presence for surah %d: %v\n", surahIndex, v0Err)
			hasV0 = false
		}

		// Fetch Ayahs for this Surah ordered by creation time
		var ayahs []quran.AyahModel
		if fetchErr := db.Where("surah_id = ?", surah.ID).Order("created_at ASC").Find(&ayahs).Error; fetchErr != nil {
			fmt.Printf("Failed to fetch ayahs for surah %d: %v\n", surahIndex, fetchErr)
			continue
		}
		if len(ayahs) == 0 {
			fmt.Printf("No ayahs found for surah %d; ensure Surah seeder ran first\n", surahIndex)
			continue
		}

		// Prepare ordered verse keys
		verseKeys := make([]string, 0, len(aData.Verse))
		for k := range aData.Verse {
			verseKeys = append(verseKeys, k)
		}
		sort.Slice(verseKeys, func(i, j int) bool { return extractVerseNum(verseKeys[i]) < extractVerseNum(verseKeys[j]) })

		now := time.Now()
		processedForSurah := 0
		for _, vKey := range verseKeys {
			vNum := extractVerseNum(vKey)
			if vNum < 0 {
				continue
			}

			// Map verse number to ayah index in ayahs slice
			ayahIdx := vNum
			if !hasV0 {
				ayahIdx = vNum - 1
			}
			if ayahIdx < 0 || ayahIdx >= len(ayahs) {
				fmt.Printf("Warning: computed ayah index %d out of range for surah %d (hasVerse0=%t). Skipping verse %s\n", ayahIdx, surahIndex, hasV0, vKey)
				continue
			}

			a := ayahs[ayahIdx]
			v := aData.Verse[vKey]
			localPath := filepath.Join(audioRoot, fmt.Sprintf("%03d", surahIndex), v.File)

			// Open local MP3 file
			f, openErr := os.Open(localPath)
			if openErr != nil {
				fmt.Printf("Failed to open audio file %s: %v\n", localPath, openErr)
				continue
			}
			// Get local file size
			stat, statErr := f.Stat()
			if statErr != nil {
				f.Close()
				fmt.Printf("Failed to stat audio file %s: %v\n", localPath, statErr)
				continue
			}

			// Compose S3 key and metadata
			s3Key := fmt.Sprintf("audio/%03d/%s", surahIndex, v.File)
			meta := map[string]string{
				"surah": fmt.Sprintf("%d", surahIndex),
				"ayah":  fmt.Sprintf("%d", vNum),
				"file":  v.File,
			}

			// Upload to S3
			obj, upErr := svc.Upload(context.Background(), domStorage.UploadInput{
				Key:         s3Key,
				Body:        f,
				ContentType: "audio/mpeg",
				PublicRead:  true,
				Metadata:    meta,
			})
			// Close file handle regardless of upload outcome
			f.Close()
			if upErr != nil {
				fmt.Printf("Failed to upload %s to S3: %v\n", localPath, upErr)
				continue
			}

			// Check if audio already exists for this ayah and reciter
			var existing quran.AyahAudioFileModel
			qErr := db.Where("ayah_id = ? AND reciter_id = ?", a.ID, reciter.ID).First(&existing).Error
			if qErr == nil {
				if force {
					existing.FilePath = obj.URL
					existing.Duration = 0
					existing.ByteSize = float64(stat.Size())
					existing.UpdatedBy = "system"
					existing.UpdatedAt = now
					if saveErr := db.Save(&existing).Error; saveErr != nil {
						fmt.Printf("Failed to update existing audio for surah %d verse %s: %v\n", surahIndex, vKey, saveErr)
					} else {
						processedForSurah++
					}
				}
				// Skip if not forcing update
				continue
			}

			// Create new AyahAudioFile entry
			audio := &quran.AyahAudioFileModel{
				ID:        uuid.New(),
				ReciterID: reciter.ID,
				SurahID:   surah.ID,
				AyahID:    a.ID,
				FilePath:  obj.URL,
				Duration:  0,
				ByteSize:  float64(stat.Size()),
				CreatedBy: "system",
				UpdatedBy: "system",
				CreatedAt: now,
				UpdatedAt: now,
			}
			if createErr := db.Create(audio).Error; createErr != nil {
				fmt.Printf("Failed to create audio entry for surah %d verse %s: %v\n", surahIndex, vKey, createErr)
				continue
			}
			processedForSurah++
		}

		fmt.Printf("Created/updated %d audio files for surah %d (%s)\n", processedForSurah, surahIndex, surah.NameEnglish)
		totalProcessed += processedForSurah
	}

	fmt.Printf("Successfully processed ayah audio files. Total created/updated: %d\n", totalProcessed)
	return nil
}

// Audio index JSON structures
type audioVerseJSON struct {
	File string `json:"file"`
	Size int    `json:"size"`
	Unit string `json:"unit"`
}

type audioIndexJSON struct {
	Verse    map[string]audioVerseJSON `json:"verse"`
	Count    int                       `json:"count"`
	RawIndex json.RawMessage           `json:"index"`
}

type audioIndexData struct {
	Index int
	Verse map[string]audioVerseJSON
	Count int
}

func parseAudioIndexFile(file string) (*audioIndexData, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var aj audioIndexJSON
	if err := json.Unmarshal(b, &aj); err != nil {
		return nil, err
	}

	var idx int
	// Parse index that may be string ("001") or number (1)
	if len(aj.RawIndex) > 0 && aj.RawIndex[0] == '"' {
		var idxStr string
		if err := json.Unmarshal(aj.RawIndex, &idxStr); err != nil {
			return nil, fmt.Errorf("failed to parse index string in %s: %v", file, err)
		}
		i, convErr := strconv.Atoi(strings.TrimLeft(idxStr, "0"))
		if convErr != nil {
			i, convErr = strconv.Atoi(idxStr)
			if convErr != nil {
				return nil, fmt.Errorf("invalid index string '%s' in %s: %v", idxStr, file, convErr)
			}
		}
		idx = i
	} else {
		var idxNum int
		if err := json.Unmarshal(aj.RawIndex, &idxNum); err == nil {
			idx = idxNum
		} else {
			var idxFloat float64
			if err2 := json.Unmarshal(aj.RawIndex, &idxFloat); err2 == nil {
				idx = int(idxFloat)
			} else {
				return nil, fmt.Errorf("failed to parse index in %s: %v", file, err)
			}
		}
	}

	return &audioIndexData{Index: idx, Verse: aj.Verse, Count: aj.Count}, nil
}

// ensureReciter finds or creates a Reciter
func ensureReciter(db *gorm.DB, name, style string) (*quran.ReciterModel, error) {
	var existing quran.ReciterModel
	if err := db.Where("name = ?", name).First(&existing).Error; err == nil {
		return &existing, nil
	}
	r := &quran.ReciterModel{
		ID:        uuid.New(),
		Name:      name,
		Style:     style,
		CreatedBy: "system",
		UpdatedBy: "system",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := db.Create(r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

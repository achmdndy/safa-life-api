package story

import "time"

// Story represents a story from the Quran.
type Story struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	Characters  []string  `json:"characters"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// StoryAyah links a story to a specific Ayah.
type StoryAyah struct {
	StoryID   string    `json:"story_id"`
	SurahID   int       `json:"surah_id"`
	AyahID    int       `json:"ayah_id"`
	CreatedAt time.Time `json:"created_at"`
}
package tajweed

import "time"

// TajweedWord represents a word in an Ayah and its associated Tajweed rule.
type TajweedWord struct {
	Word  string `json:"word"`
	Rule  string `json:"rule"`
	Color string `json:"color"`
}

// AyahTajweed holds the collection of Tajweed rules for all words in a specific Ayah.
type AyahTajweed struct {
	TajweedID string        `json:"tajweed_id"` // e.g., "default_tajweed", "advanced_tajweed"
	SurahID   int           `json:"surah_id"`
	AyahID    int           `json:"ayah_id"`
	Words     []TajweedWord `json:"words"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

// TajweedRule defines a specific tajweed rule.
type TajweedRule struct {
	ID          string    `json:"id"`
	Rule        string    `json:"rule"`
	Explanation string    `json:"explanation"`
	Color       string    `json:"color"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
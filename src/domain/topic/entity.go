package topic

import "time"

// Topic represents a topic in the Quran.
type Topic struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TopicAyah links a topic to a specific Ayah.
type TopicAyah struct {
	TopicID   string    `json:"topic_id"`
	SurahID   int       `json:"surah_id"`
	AyahID    int       `json:"ayah_id"`
	CreatedAt time.Time `json:"created_at"`
}
package story

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

// StringSlice is a custom type for string slices to be stored as JSON.
type StringSlice []string

// Value implements the driver.Valuer interface for JSON serialization.
func (s StringSlice) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// Scan implements the sql.Scanner interface for JSON deserialization.
func (s *StringSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, s)
}

// StoryModel represents the stories table.
type StoryModel struct {
	ID          string         `gorm:"primaryKey;column:id;size:255"`
	Title       string         `gorm:"column:title;not null;size:255"`
	Summary     string         `gorm:"column:summary;type:text"`
	Characters  StringSlice    `gorm:"column:characters;type:jsonb"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (StoryModel) TableName() string {
	return "stories"
}

// StoryAyahModel represents the story_ayahs table (many-to-many link).
type StoryAyahModel struct {
	StoryID   string    `gorm:"primaryKey;column:story_id;size:255"`
	SurahID   int       `gorm:"primaryKey;column:surah_id"`
	AyahID    int       `gorm:"primaryKey;column:ayah_id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (StoryAyahModel) TableName() string {
	return "story_ayahs"
}
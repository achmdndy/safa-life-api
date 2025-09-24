package topic

import (
	"time"

	"gorm.io/gorm"
)

// TopicModel represents the topics table.
type TopicModel struct {
	ID          string         `gorm:"primaryKey;column:id;size:255"`
	Name        string         `gorm:"column:name;not null;unique;size:255"`
	Description string         `gorm:"column:description;type:text"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (TopicModel) TableName() string {
	return "topics"
}

// TopicAyahModel represents the topic_ayahs table (many-to-many link).
type TopicAyahModel struct {
	TopicID   string    `gorm:"primaryKey;column:topic_id;size:255"`
	SurahID   int       `gorm:"primaryKey;column:surah_id"`
	AyahID    int       `gorm:"primaryKey;column:ayah_id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (TopicAyahModel) TableName() string {
	return "topic_ayahs"
}
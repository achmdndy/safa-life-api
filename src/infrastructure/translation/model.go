package translation

import (
	"time"
	"gorm.io/gorm"
)

// TranslationModel represents the translation_editions table.
type TranslationModel struct {
	ID        string         `gorm:"primaryKey;column:id;size:255"`
	Name      string         `gorm:"column:name;not null;size:255"`
	Author    string         `gorm:"column:author;not null;size:255"`
	Language  string         `gorm:"column:language;not null;size:50"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (TranslationModel) TableName() string {
	return "translation_editions"
}

// AyahTranslationModel represents the ayah_translations table.
type AyahTranslationModel struct {
	TranslationID string         `gorm:"primaryKey;column:translation_id;size:255"`
	SurahID       int            `gorm:"primaryKey;column:surah_id"`
	AyahID        int            `gorm:"primaryKey;column:ayah_id"`
	Text          string         `gorm:"column:text;not null;type:text"`
	CreatedAt     time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index"`

	// Relationships
	Translation TranslationModel `gorm:"foreignKey:TranslationID;references:ID"`
}

func (AyahTranslationModel) TableName() string {
	return "ayah_translations"
}
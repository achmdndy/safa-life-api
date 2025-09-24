package tafsir

import (
	"time"
	"gorm.io/gorm"
)

// TafsirModel represents the tafsir_editions table.
type TafsirModel struct {
	ID        string         `gorm:"primaryKey;column:id;size:255"`
	Name      string         `gorm:"column:name;not null;size:255"`
	Author    string         `gorm:"column:author;not null;size:255"`
	Language  string         `gorm:"column:language;not null;size:50"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (TafsirModel) TableName() string {
	return "tafsir_editions"
}

// AyahTafsirModel represents the ayah_tafsirs table.
type AyahTafsirModel struct {
	TafsirID  string         `gorm:"primaryKey;column:tafsir_id;size:255"`
	SurahID   int            `gorm:"primaryKey;column:surah_id"`
	AyahID    int            `gorm:"primaryKey;column:ayah_id"`
	Text      string         `gorm:"column:text;not null;type:text"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`

	// Relationships
	Tafsir TafsirModel `gorm:"foreignKey:TafsirID;references:ID"`
}

func (AyahTafsirModel) TableName() string {
	return "ayah_tafsirs"
}
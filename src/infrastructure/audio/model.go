package audio

import (
	"time"

	"gorm.io/gorm"
)

// AyahAudioModel represents the ayah_audio_files table.
type AyahAudioModel struct {
	ReciterID string         `gorm:"primaryKey;column:reciter_id;size:255"`
	SurahID   int            `gorm:"primaryKey;column:surah_id"`
	AyahID    int            `gorm:"primaryKey;column:ayah_id"`
	FilePath  string         `gorm:"column:file_path;not null;size:512"`
	Duration  float64        `gorm:"column:duration"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (AyahAudioModel) TableName() string {
	return "ayah_audio_files"
}
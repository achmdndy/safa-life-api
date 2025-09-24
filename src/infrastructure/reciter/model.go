package reciter

import (
	"time"
	"gorm.io/gorm"
)

// ReciterModel represents the reciters table.
type ReciterModel struct {
	ID        string         `gorm:"primaryKey;column:id;size:255"`
	Name      string         `gorm:"column:name;not null;unique;size:255"`
	Style     string         `gorm:"column:style;not null;size:100"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (ReciterModel) TableName() string {
	return "reciters"
}
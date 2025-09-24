package resource

import (
	"time"

	"gorm.io/gorm"
)

// ResourceModel represents the downloadable_resources table.
type ResourceModel struct {
	ID           string         `gorm:"primaryKey;column:id;size:255"`
	Type         string         `gorm:"column:type;not null;size:100"`
	Name         string         `gorm:"column:name;not null;size:255"`
	Language     string         `gorm:"column:language;size:50"`
	Version      string         `gorm:"column:version;size:50"`
	LastUpdatedAt time.Time      `gorm:"column:last_updated_at"`
	DownloadURL  string         `gorm:"column:download_url;not null;size:512"`
	Size         int64          `gorm:"column:size_kb"`
	IsInstalled  bool           `gorm:"column:is_installed;default:false"`
	InstalledAt  *time.Time     `gorm:"column:installed_at"`
	CreatedAt    time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (ResourceModel) TableName() string {
	return "downloadable_resources"
}
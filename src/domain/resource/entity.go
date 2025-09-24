package resource

import "time"

// Resource represents a downloadable resource like a translation or tafsir database.
type Resource struct {
	ID           string     `json:"id"`
	Type         string     `json:"type"`
	Name         string     `json:"name"`
	Language     string     `json:"language"`
	Version      string     `json:"version"`
	LastUpdatedAt time.Time  `json:"last_updated_at"`
	DownloadURL  string     `json:"download_url"`
	Size         int64      `json:"size_kb"`
	IsInstalled  bool       `json:"is_installed"`
	InstalledAt  *time.Time `json:"installed_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
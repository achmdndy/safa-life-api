package dto

import (
	"time"

	"github.com/achmdndy/safa-life-api/src/domain/resource"
)

// ResourceResponse represents a downloadable resource.
type ResourceResponse struct {
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

func ToResourceResponse(d resource.Resource) ResourceResponse {
	return ResourceResponse{
		ID:           d.ID,
		Type:         d.Type,
		Name:         d.Name,
		Language:     d.Language,
		Version:      d.Version,
		LastUpdatedAt: d.LastUpdatedAt,
		DownloadURL:  d.DownloadURL,
		Size:         d.Size,
		IsInstalled:  d.IsInstalled,
		InstalledAt:  d.InstalledAt,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
	}
}

func ToResourceResponseSlice(ds []resource.Resource) []ResourceResponse {
	outs := make([]ResourceResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToResourceResponse(d)
	}
	return outs
}

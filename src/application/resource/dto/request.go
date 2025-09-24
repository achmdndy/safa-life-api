package dto

import "time"

// GetResourceByIDRequest represents the request to get a resource by ID.
type GetResourceByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

// CreateResourceRequest represents the request to create a new resource.
type CreateResourceRequest struct {
	ID            string    `json:"id" binding:"required"`
	Type          string    `json:"type" binding:"required"`
	Name          string    `json:"name" binding:"required"`
	Language      string    `json:"language"`
	Version       string    `json:"version" binding:"required"`
	LastUpdatedAt time.Time `json:"last_updated_at" binding:"required"`
	DownloadURL   string    `json:"download_url" binding:"required"`
	Size          int64     `json:"size_kb" binding:"required"`
}

// UpdateResourceRequest represents the request to update a resource.
type UpdateResourceRequest struct {
	ID            string    `uri:"id" binding:"required"`
	Type          string    `json:"type" binding:"omitempty"`
	Name          string    `json:"name" binding:"omitempty"`
	Language      string    `json:"language" binding:"omitempty"`
	Version       string    `json:"version" binding:"omitempty"`
	LastUpdatedAt time.Time `json:"last_updated_at" binding:"omitempty"`
	DownloadURL   string    `json:"download_url" binding:"omitempty"`
	Size          int64     `json:"size_kb" binding:"omitempty"`
}

// UpdateInstallStatusRequest represents the request to update the install status of a resource.
type UpdateInstallStatusRequest struct {
	ID          string `uri:"id" binding:"required"`
	IsInstalled bool   `json:"is_installed"`
}

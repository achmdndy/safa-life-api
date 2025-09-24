package resource

import (
	"math/rand"
	"time"
)

// ResourceFaker generates fake Resource data
type ResourceFaker struct{}

// NewResourceFaker creates a new ResourceFaker
func NewResourceFaker() *ResourceFaker {
	rand.Seed(time.Now().UnixNano())
	return &ResourceFaker{}
}

// Generate creates a fake ResourceModel
func (f *ResourceFaker) Generate() ResourceModel {
	res := ResourceModel{
		ID:           "en_asad_trans",
		Type:         "translation",
		Name:         "The Message of The Qur'an by Muhammad Asad",
		Language:     "en",
		Version:      "1.0.0",
		LastUpdatedAt: time.Now().Add(-24 * time.Hour),
		DownloadURL:  "https://example.com/en_asad.zip",
		Size:         2048,
		IsInstalled:  false,
		InstalledAt:  nil,
		CreatedAt:    time.Now().Add(-48 * time.Hour),
		UpdatedAt:    time.Now().Add(-24 * time.Hour),
	}
	return res
}

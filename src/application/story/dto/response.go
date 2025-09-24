package dto

import (
	"time"

	"github.com/achmdndy/safa-life-api/src/domain/story"
)

// StoryResponse represents a story from the Quran.
type StoryResponse struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Summary    string    `json:"summary"`
	Characters []string  `json:"characters"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// StoryAyahResponse links a story to a specific Ayah.
type StoryAyahResponse struct {
	StoryID string `json:"story_id"`
	SurahID int    `json:"surah_id"`
	AyahID  int    `json:"ayah_id"`
}

func ToStoryResponse(d story.Story) StoryResponse {
	return StoryResponse{
		ID:         d.ID,
		Title:      d.Title,
		Summary:    d.Summary,
		Characters: d.Characters,
		CreatedAt:  d.CreatedAt,
		UpdatedAt:  d.UpdatedAt,
	}
}

func ToStoryResponseSlice(ds []story.Story) []StoryResponse {
	outs := make([]StoryResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToStoryResponse(d)
	}
	return outs
}

func ToStoryAyahResponse(d story.StoryAyah) StoryAyahResponse {
	return StoryAyahResponse{
		StoryID: d.StoryID,
		SurahID: d.SurahID,
		AyahID:  d.AyahID,
	}
}

func ToStoryAyahResponseSlice(ds []story.StoryAyah) []StoryAyahResponse {
	outs := make([]StoryAyahResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToStoryAyahResponse(d)
	}
	return outs
}

package dto

import (
	"time"

	"github.com/achmdndy/safa-life-api/src/domain/topic"
)

// TopicResponse represents a topic in the Quran.
type TopicResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TopicAyahResponse links a topic to a specific Ayah.
type TopicAyahResponse struct {
	TopicID string `json:"topic_id"`
	SurahID int    `json:"surah_id"`
	AyahID  int    `json:"ayah_id"`
}

func ToTopicResponse(d topic.Topic) TopicResponse {
	return TopicResponse{
		ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

func ToTopicResponseSlice(ds []topic.Topic) []TopicResponse {
	outs := make([]TopicResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToTopicResponse(d)
	}
	return outs
}

func ToTopicAyahResponse(d topic.TopicAyah) TopicAyahResponse {
	return TopicAyahResponse{
		TopicID: d.TopicID,
		SurahID: d.SurahID,
		AyahID:  d.AyahID,
	}
}

func ToTopicAyahResponseSlice(ds []topic.TopicAyah) []TopicAyahResponse {
	outs := make([]TopicAyahResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToTopicAyahResponse(d)
	}
	return outs
}

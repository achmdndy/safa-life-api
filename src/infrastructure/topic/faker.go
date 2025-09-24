package topic

import (
	"math/rand"
	"time"
)

// TopicInfoFaker generates fake Topic data
type TopicInfoFaker struct{}

// NewTopicInfoFaker creates a new TopicInfoFaker
func NewTopicInfoFaker() *TopicInfoFaker {
	return &TopicInfoFaker{}
}

// Generate creates a fake TopicModel
func (f *TopicInfoFaker) Generate() TopicModel {
	topics := []struct {
		ID          string
		Name        string
		Description string
	}{
		{"patience", "Patience (Sabr)", "Verses about patience and perseverance"},
		{"charity", "Charity (Sadaqa)", "Verses about giving charity"},
	}
	topic := topics[rand.Intn(len(topics))]
	return TopicModel{
		ID:          topic.ID,
		Name:        topic.Name,
		Description: topic.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// TopicAyahFaker generates fake TopicAyah data
type TopicAyahFaker struct{}

// NewTopicAyahFaker creates a new TopicAyahFaker
func NewTopicAyahFaker() *TopicAyahFaker {
	return &TopicAyahFaker{}
}

// Generate creates a fake TopicAyahModel
func (f *TopicAyahFaker) Generate(topicID string) TopicAyahModel {
	return TopicAyahModel{
		TopicID:   topicID,
		SurahID:   rand.Intn(114) + 1,
		AyahID:    rand.Intn(286) + 1,
		CreatedAt: time.Now(),
	}
}

// TopicFaker combines all faker functionality
type TopicFaker struct {
	InfoFaker *TopicInfoFaker
	AyahFaker *TopicAyahFaker
}

// NewTopicFaker creates a new TopicFaker instance
func NewTopicFaker() *TopicFaker {
	rand.Seed(time.Now().UnixNano())
	return &TopicFaker{
		InfoFaker: NewTopicInfoFaker(),
		AyahFaker: NewTopicAyahFaker(),
	}
}
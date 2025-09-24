package story

import (
	"math/rand"
	"time"
)

// StoryInfoFaker generates fake Story data
type StoryInfoFaker struct{}

// NewStoryInfoFaker creates a new StoryInfoFaker
func NewStoryInfoFaker() *StoryInfoFaker {
	return &StoryInfoFaker{}
}

// Generate creates a fake StoryModel
func (f *StoryInfoFaker) Generate() StoryModel {
	stories := []struct {
		ID         string
		Title      string
		Summary    string
		Characters []string
	}{
		{"prophet_musa", "The Story of Prophet Musa", "The life and struggles of Prophet Musa (Moses)...", []string{"Musa", "Firaun"}},
		{"prophet_yusuf", "The Story of Prophet Yusuf", "The story of Prophet Yusuf (Joseph) and his brothers.", []string{"Yusuf", "Yaqub", "Brothers"}},
	}
	story := stories[rand.Intn(len(stories))]
	return StoryModel{
		ID:          story.ID,
		Title:       story.Title,
		Summary:     story.Summary,
		Characters:  story.Characters,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// StoryAyahFaker generates fake StoryAyah data
type StoryAyahFaker struct{}

// NewStoryAyahFaker creates a new StoryAyahFaker
func NewStoryAyahFaker() *StoryAyahFaker {
	return &StoryAyahFaker{}
}

// Generate creates a fake StoryAyahModel
func (f *StoryAyahFaker) Generate(storyID string) StoryAyahModel {
	return StoryAyahModel{
		StoryID:   storyID,
		SurahID:   rand.Intn(114) + 1,
		AyahID:    rand.Intn(286) + 1,
		CreatedAt: time.Now(),
	}
}

// StoryFaker combines all faker functionality
type StoryFaker struct {
	InfoFaker *StoryInfoFaker
	AyahFaker *StoryAyahFaker
}

// NewStoryFaker creates a new StoryFaker instance
func NewStoryFaker() *StoryFaker {
	rand.Seed(time.Now().UnixNano())
	return &StoryFaker{
		InfoFaker: NewStoryInfoFaker(),
		AyahFaker: NewStoryAyahFaker(),
	}
}
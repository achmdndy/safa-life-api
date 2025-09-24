package story

import (
	"context"
	"fmt"
)

type storyService struct {
	repo FullStoryRepository
}

// NewStoryService creates a new instance of StoryService.
func NewStoryService(repo FullStoryRepository) StoryService {
	return &storyService{repo: repo}
}

func (s *storyService) GetAllStories(ctx context.Context) ([]Story, error) {
	return s.repo.GetAll(ctx)
}

func (s *storyService) GetStoryByID(ctx context.Context, id string) (*Story, error) {
	if id == "" {
		return nil, fmt.Errorf("story ID cannot be empty")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *storyService) CreateStory(ctx context.Context, story *Story) (*Story, error) {
	if story.ID == "" || story.Title == "" {
		return nil, fmt.Errorf("story ID and title cannot be empty")
	}
	return s.repo.CreateStory(ctx, story)
}

func (s *storyService) UpdateStory(ctx context.Context, story *Story) (*Story, error) {
	if story.ID == "" {
		return nil, fmt.Errorf("story ID cannot be empty")
	}
	return s.repo.UpdateStory(ctx, story)
}

func (s *storyService) DeleteStory(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("story ID cannot be empty")
	}
	return s.repo.DeleteStory(ctx, id)
}

func (s *storyService) GetAyahsForStory(ctx context.Context, storyID string) ([]StoryAyah, error) {
	if storyID == "" {
		return nil, fmt.Errorf("story ID cannot be empty")
	}
	return s.repo.GetByStoryID(ctx, storyID)
}

func (s *storyService) AddAyahToStory(ctx context.Context, storyAyah *StoryAyah) (*StoryAyah, error) {
	if storyAyah.StoryID == "" {
		return nil, fmt.Errorf("story ID cannot be empty")
	}
	return s.repo.CreateStoryAyah(ctx, storyAyah)
}

func (s *storyService) RemoveAyahFromStory(ctx context.Context, storyID string, surahID, ayahID int) error {
	if storyID == "" {
		return fmt.Errorf("story ID cannot be empty")
	}
	return s.repo.DeleteStoryAyah(ctx, storyID, surahID, ayahID)
}

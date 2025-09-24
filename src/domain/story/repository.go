package story

import "context"

// StoryRepository defines the interface for Story data access.
type StoryRepository interface {
	GetAll(ctx context.Context) ([]Story, error)
	GetByID(ctx context.Context, id string) (*Story, error)
	CreateStory(ctx context.Context, story *Story) (*Story, error)
	UpdateStory(ctx context.Context, story *Story) (*Story, error)
	DeleteStory(ctx context.Context, id string) error
}

// StoryAyahRepository defines the interface for StoryAyah data access.
type StoryAyahRepository interface {
	GetByStoryID(ctx context.Context, storyID string) ([]StoryAyah, error)
	CreateStoryAyah(ctx context.Context, storyAyah *StoryAyah) (*StoryAyah, error)
	DeleteStoryAyah(ctx context.Context, storyID string, surahID, ayahID int) error
}

// FullStoryRepository defines the aggregate repository interface.
type FullStoryRepository interface {
	StoryRepository
	StoryAyahRepository
	TransactionManager
}

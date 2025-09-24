package story

import "context"

// StoryService defines the business logic for story operations.
type StoryService interface {
	GetAllStories(ctx context.Context) ([]Story, error)
	GetStoryByID(ctx context.Context, id string) (*Story, error)
	CreateStory(ctx context.Context, story *Story) (*Story, error)
	UpdateStory(ctx context.Context, story *Story) (*Story, error)
	DeleteStory(ctx context.Context, id string) error

	GetAyahsForStory(ctx context.Context, storyID string) ([]StoryAyah, error)
	AddAyahToStory(ctx context.Context, storyAyah *StoryAyah) (*StoryAyah, error)
	RemoveAyahFromStory(ctx context.Context, storyID string, surahID, ayahID int) error
}

// TransactionManager defines the interface for database transaction management.
type TransactionManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
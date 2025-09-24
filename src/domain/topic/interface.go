package topic

import "context"

// TopicService defines the business logic for topic operations.
type TopicService interface {
	GetAllTopics(ctx context.Context) ([]Topic, error)
	GetTopicByID(ctx context.Context, id string) (*Topic, error)
	CreateTopic(ctx context.Context, topic *Topic) (*Topic, error)
	UpdateTopic(ctx context.Context, topic *Topic) (*Topic, error)
	DeleteTopic(ctx context.Context, id string) error

	GetAyahsForTopic(ctx context.Context, topicID string) ([]TopicAyah, error)
	AddAyahToTopic(ctx context.Context, topicAyah *TopicAyah) (*TopicAyah, error)
	RemoveAyahFromTopic(ctx context.Context, topicID string, surahID, ayahID int) error
}

// TransactionManager defines the interface for database transaction management.
type TransactionManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
package topic

import "context"

// TopicRepository defines the interface for Topic data access.
type TopicRepository interface {
	GetAll(ctx context.Context) ([]Topic, error)
	GetByID(ctx context.Context, id string) (*Topic, error)
	CreateTopic(ctx context.Context, topic *Topic) (*Topic, error)
	UpdateTopic(ctx context.Context, topic *Topic) (*Topic, error)
	DeleteTopic(ctx context.Context, id string) error
}

// TopicAyahRepository defines the interface for TopicAyah data access.
type TopicAyahRepository interface {
	GetByTopicID(ctx context.Context, topicID string) ([]TopicAyah, error)
	CreateTopicAyah(ctx context.Context, topicAyah *TopicAyah) (*TopicAyah, error)
	DeleteTopicAyah(ctx context.Context, topicID string, surahID, ayahID int) error
}

// FullTopicRepository defines the aggregate repository interface.
type FullTopicRepository interface {
	TopicRepository
	TopicAyahRepository
	TransactionManager
}

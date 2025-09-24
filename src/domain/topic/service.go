package topic

import (
	"context"
	"fmt"
)

type topicService struct {
	repo FullTopicRepository
}

// NewTopicService creates a new instance of TopicService.
func NewTopicService(repo FullTopicRepository) TopicService {
	return &topicService{repo: repo}
}

func (s *topicService) GetAllTopics(ctx context.Context) ([]Topic, error) {
	return s.repo.GetAll(ctx)
}

func (s *topicService) GetTopicByID(ctx context.Context, id string) (*Topic, error) {
	if id == "" {
		return nil, fmt.Errorf("topic ID cannot be empty")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *topicService) CreateTopic(ctx context.Context, topic *Topic) (*Topic, error) {
	if topic.ID == "" || topic.Name == "" {
		return nil, fmt.Errorf("topic ID and name cannot be empty")
	}
	return s.repo.CreateTopic(ctx, topic)
}

func (s *topicService) UpdateTopic(ctx context.Context, topic *Topic) (*Topic, error) {
	if topic.ID == "" {
		return nil, fmt.Errorf("topic ID cannot be empty")
	}
	return s.repo.UpdateTopic(ctx, topic)
}

func (s *topicService) DeleteTopic(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("topic ID cannot be empty")
	}
	return s.repo.DeleteTopic(ctx, id)
}

func (s *topicService) GetAyahsForTopic(ctx context.Context, topicID string) ([]TopicAyah, error) {
	if topicID == "" {
		return nil, fmt.Errorf("topic ID cannot be empty")
	}
	return s.repo.GetByTopicID(ctx, topicID)
}

func (s *topicService) AddAyahToTopic(ctx context.Context, topicAyah *TopicAyah) (*TopicAyah, error) {
	if topicAyah.TopicID == "" {
		return nil, fmt.Errorf("topic ID cannot be empty")
	}
	return s.repo.CreateTopicAyah(ctx, topicAyah)
}

func (s *topicService) RemoveAyahFromTopic(ctx context.Context, topicID string, surahID, ayahID int) error {
	if topicID == "" {
		return fmt.Errorf("topic ID cannot be empty")
	}
	return s.repo.DeleteTopicAyah(ctx, topicID, surahID, ayahID)
}

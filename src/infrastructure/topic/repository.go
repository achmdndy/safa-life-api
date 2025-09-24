package topic

import (
	"context"
	"github.com/achmdndy/safa-life-api/src/domain/topic"
	"gorm.io/gorm"
)

type topicRepository struct {
	db *gorm.DB
}

func NewTopicRepository(db *gorm.DB) topic.FullTopicRepository {
	return &topicRepository{db: db}
}

func (r *topicRepository) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		type contextKey string
		const txRepoKey contextKey = "tx_repo_topic"
		txCtx := context.WithValue(ctx, txRepoKey, &topicRepository{db: tx})
		return fn(txCtx)
	})
}

func (r *topicRepository) getDB(ctx context.Context) *gorm.DB {
	if txRepo, ok := ctx.Value("tx_repo_topic").(*topicRepository); ok {
		return txRepo.db.WithContext(ctx)
	}
	return r.db.WithContext(ctx)
}

// Topic methods
func (r *topicRepository) GetAll(ctx context.Context) ([]topic.Topic, error) {
	var models []TopicModel
	if err := r.getDB(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	return ToTopicDomainSlice(models), nil
}

func (r *topicRepository) GetByID(ctx context.Context, id string) (*topic.Topic, error) {
	var model TopicModel
	if err := r.getDB(ctx).First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	domain := model.ToTopicDomain()
	return &domain, nil
}

func (r *topicRepository) CreateTopic(ctx context.Context, t *topic.Topic) (*topic.Topic, error) {
	model := FromTopicDomain(t)
	if err := r.getDB(ctx).Create(model).Error; err != nil {
		return nil, err
	}
	result := model.ToTopicDomain()
	return &result, nil
}

func (r *topicRepository) UpdateTopic(ctx context.Context, t *topic.Topic) (*topic.Topic, error) {
	model := FromTopicDomain(t)
	if err := r.getDB(ctx).Save(model).Error; err != nil {
		return nil, err
	}
	result := model.ToTopicDomain()
	return &result, nil
}

func (r *topicRepository) DeleteTopic(ctx context.Context, id string) error {
	return r.getDB(ctx).Delete(&TopicModel{}, "id = ?", id).Error
}

// TopicAyah methods
func (r *topicRepository) GetByTopicID(ctx context.Context, topicID string) ([]topic.TopicAyah, error) {
	var models []TopicAyahModel
	if err := r.getDB(ctx).Find(&models, "topic_id = ?", topicID).Error; err != nil {
		return nil, err
	}
	return ToTopicAyahDomainSlice(models), nil
}

func (r *topicRepository) CreateTopicAyah(ctx context.Context, ta *topic.TopicAyah) (*topic.TopicAyah, error) {
	model := FromTopicAyahDomain(ta)
	if err := r.getDB(ctx).Create(model).Error; err != nil {
		return nil, err
	}
	result := model.ToTopicAyahDomain()
	return &result, nil
}

func (r *topicRepository) DeleteTopicAyah(ctx context.Context, topicID string, surahID, ayahID int) error {
	conds := "topic_id = ? AND surah_id = ? AND ayah_id = ?"
	return r.getDB(ctx).Delete(&TopicAyahModel{}, conds, topicID, surahID, ayahID).Error
}

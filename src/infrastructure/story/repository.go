package story

import (
	"context"
	"github.com/achmdndy/safa-life-api/src/domain/story"
	"gorm.io/gorm"
)

type storyRepository struct {
	db *gorm.DB
}

func NewStoryRepository(db *gorm.DB) story.FullStoryRepository {
	return &storyRepository{db: db}
}

func (r *storyRepository) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		type contextKey string
		const txRepoKey contextKey = "tx_repo_story"
		txCtx := context.WithValue(ctx, txRepoKey, &storyRepository{db: tx})
		return fn(txCtx)
	})
}

func (r *storyRepository) getDB(ctx context.Context) *gorm.DB {
	if txRepo, ok := ctx.Value("tx_repo_story").(*storyRepository); ok {
		return txRepo.db.WithContext(ctx)
	}
	return r.db.WithContext(ctx)
}

// Story methods
func (r *storyRepository) GetAll(ctx context.Context) ([]story.Story, error) {
	var models []StoryModel
	if err := r.getDB(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	return ToStoryDomainSlice(models), nil
}

func (r *storyRepository) GetByID(ctx context.Context, id string) (*story.Story, error) {
	var model StoryModel
	if err := r.getDB(ctx).First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	domain := model.ToStoryDomain()
	return &domain, nil
}

func (r *storyRepository) CreateStory(ctx context.Context, s *story.Story) (*story.Story, error) {
	model := FromStoryDomain(s)
	if err := r.getDB(ctx).Create(model).Error; err != nil {
		return nil, err
	}
	result := model.ToStoryDomain()
	return &result, nil
}

func (r *storyRepository) UpdateStory(ctx context.Context, s *story.Story) (*story.Story, error) {
	model := FromStoryDomain(s)
	if err := r.getDB(ctx).Save(model).Error; err != nil {
		return nil, err
	}
	result := model.ToStoryDomain()
	return &result, nil
}

func (r *storyRepository) DeleteStory(ctx context.Context, id string) error {
	return r.getDB(ctx).Delete(&StoryModel{}, "id = ?", id).Error
}

// StoryAyah methods
func (r *storyRepository) GetByStoryID(ctx context.Context, storyID string) ([]story.StoryAyah, error) {
	var models []StoryAyahModel
	if err := r.getDB(ctx).Find(&models, "story_id = ?", storyID).Error; err != nil {
		return nil, err
	}
	return ToStoryAyahDomainSlice(models), nil
}

func (r *storyRepository) CreateStoryAyah(ctx context.Context, sa *story.StoryAyah) (*story.StoryAyah, error) {
	model := FromStoryAyahDomain(sa)
	if err := r.getDB(ctx).Create(model).Error; err != nil {
		return nil, err
	}
	result := model.ToStoryAyahDomain()
	return &result, nil
}

func (r *storyRepository) DeleteStoryAyah(ctx context.Context, storyID string, surahID, ayahID int) error {
	conds := "story_id = ? AND surah_id = ? AND ayah_id = ?"
	return r.getDB(ctx).Delete(&StoryAyahModel{}, conds, storyID, surahID, ayahID).Error
}

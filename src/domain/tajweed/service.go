package tajweed

import (
	"context"
	"fmt"
)

// tajweedService implements the TajweedService interface.
type tajweedService struct {
	repo TajweedRepository
}

// NewTajweedService creates a new instance of TajweedService.
func NewTajweedService(repo TajweedRepository) TajweedService {
	return &tajweedService{
		repo: repo,
	}
}

// AyahTajweed operations
func (s *tajweedService) GetAyahTajweed(ctx context.Context, tajweedID string, surahID, ayahID int) (*AyahTajweed, error) {
	if tajweedID == "" {
		return nil, fmt.Errorf("tajweed ID cannot be empty")
	}
	if surahID <= 0 || surahID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", surahID)
	}
	if ayahID <= 0 {
		return nil, fmt.Errorf("invalid ayah ID: %d", ayahID)
	}
	return s.repo.GetAyahTajweed(ctx, tajweedID, surahID, ayahID)
}

func (s *tajweedService) CreateAyahTajweed(ctx context.Context, ayahTajweed *AyahTajweed) (*AyahTajweed, error) {
	// Basic validation
	if ayahTajweed.TajweedID == "" {
		return nil, fmt.Errorf("tajweed ID cannot be empty")
	}
	if ayahTajweed.SurahID <= 0 || ayahTajweed.SurahID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", ayahTajweed.SurahID)
	}
	if ayahTajweed.AyahID <= 0 {
		return nil, fmt.Errorf("invalid ayah ID: %d", ayahTajweed.AyahID)
	}
	if len(ayahTajweed.Words) == 0 {
		return nil, fmt.Errorf("tajweed words cannot be empty")
	}
	return s.repo.CreateAyahTajweed(ctx, ayahTajweed)
}

func (s *tajweedService) UpdateAyahTajweed(ctx context.Context, ayahTajweed *AyahTajweed) (*AyahTajweed, error) {
	// Basic validation
	if ayahTajweed.TajweedID == "" {
		return nil, fmt.Errorf("tajweed ID cannot be empty")
	}
	if ayahTajweed.SurahID <= 0 || ayahTajweed.SurahID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", ayahTajweed.SurahID)
	}
	if ayahTajweed.AyahID <= 0 {
		return nil, fmt.Errorf("invalid ayah ID: %d", ayahTajweed.AyahID)
	}
	if len(ayahTajweed.Words) == 0 {
		return nil, fmt.Errorf("tajweed words cannot be empty")
	}
	return s.repo.UpdateAyahTajweed(ctx, ayahTajweed)
}

func (s *tajweedService) DeleteAyahTajweed(ctx context.Context, tajweedID string, surahID, ayahID int) error {
	if tajweedID == "" {
		return fmt.Errorf("tajweed ID cannot be empty")
	}
	if surahID <= 0 || surahID > 114 {
		return fmt.Errorf("invalid surah ID: %d", surahID)
	}
	if ayahID <= 0 {
		return fmt.Errorf("invalid ayah ID: %d", ayahID)
	}
	return s.repo.DeleteAyahTajweed(ctx, tajweedID, surahID, ayahID)
}

// TajweedRule operations
func (s *tajweedService) GetAllTajweedRules(ctx context.Context) ([]TajweedRule, error) {
	return s.repo.GetAllTajweedRules(ctx)
}

func (s *tajweedService) GetTajweedRuleByID(ctx context.Context, ruleID string) (*TajweedRule, error) {
	if ruleID == "" {
		return nil, fmt.Errorf("rule ID cannot be empty")
	}
	return s.repo.GetTajweedRuleByID(ctx, ruleID)
}

func (s *tajweedService) GetTajweedRuleByName(ctx context.Context, ruleName string) (*TajweedRule, error) {
	if ruleName == "" {
		return nil, fmt.Errorf("rule name cannot be empty")
	}
	return s.repo.GetTajweedRuleByName(ctx, ruleName)
}

func (s *tajweedService) CreateTajweedRule(ctx context.Context, rule *TajweedRule) (*TajweedRule, error) {
	if rule.Rule == "" {
		return nil, fmt.Errorf("rule name cannot be empty")
	}
	if rule.Explanation == "" {
		return nil, fmt.Errorf("rule explanation cannot be empty")
	}
	return s.repo.CreateTajweedRule(ctx, rule)
}

func (s *tajweedService) UpdateTajweedRule(ctx context.Context, rule *TajweedRule) (*TajweedRule, error) {
	if rule.ID == "" {
		return nil, fmt.Errorf("rule ID cannot be empty")
	}
	if rule.Rule == "" {
		return nil, fmt.Errorf("rule name cannot be empty")
	}
	if rule.Explanation == "" {
		return nil, fmt.Errorf("rule explanation cannot be empty")
	}
	return s.repo.UpdateTajweedRule(ctx, rule)
}

func (s *tajweedService) DeleteTajweedRule(ctx context.Context, ruleID string) error {
	if ruleID == "" {
		return fmt.Errorf("rule ID cannot be empty")
	}
	return s.repo.DeleteTajweedRule(ctx, ruleID)
}

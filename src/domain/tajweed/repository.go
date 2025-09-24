package tajweed

import "context"

// AyahTajweedRepository defines the interface for AyahTajweed data access.
type AyahTajweedRepository interface {
	GetAyahTajweed(ctx context.Context, tajweedID string, surahID, ayahID int) (*AyahTajweed, error)
	CreateAyahTajweed(ctx context.Context, ayahTajweed *AyahTajweed) (*AyahTajweed, error)
	UpdateAyahTajweed(ctx context.Context, ayahTajweed *AyahTajweed) (*AyahTajweed, error)
	DeleteAyahTajweed(ctx context.Context, tajweedID string, surahID, ayahID int) error
}

// TajweedRuleRepository defines the interface for TajweedRule data access.
type TajweedRuleRepository interface {
	GetAllTajweedRules(ctx context.Context) ([]TajweedRule, error)
	GetTajweedRuleByID(ctx context.Context, ruleID string) (*TajweedRule, error)
	GetTajweedRuleByName(ctx context.Context, ruleName string) (*TajweedRule, error)
	CreateTajweedRule(ctx context.Context, rule *TajweedRule) (*TajweedRule, error)
	UpdateTajweedRule(ctx context.Context, rule *TajweedRule) (*TajweedRule, error)
	DeleteTajweedRule(ctx context.Context, ruleID string) error
}

// TajweedRepository defines the aggregate repository interface.
type TajweedRepository interface {
	AyahTajweedRepository
	TajweedRuleRepository
	TransactionManager
}

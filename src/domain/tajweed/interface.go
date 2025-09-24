package tajweed

import "context"

// TajweedService defines the business logic for Tajweed operations.
type TajweedService interface {
	// AyahTajweed read operations
	GetAyahTajweed(ctx context.Context, tajweedID string, surahID, ayahID int) (*AyahTajweed, error)
	
	// AyahTajweed write operations
	CreateAyahTajweed(ctx context.Context, ayahTajweed *AyahTajweed) (*AyahTajweed, error)
	UpdateAyahTajweed(ctx context.Context, ayahTajweed *AyahTajweed) (*AyahTajweed, error)
	DeleteAyahTajweed(ctx context.Context, tajweedID string, surahID, ayahID int) error

	// TajweedRule read operations
	GetAllTajweedRules(ctx context.Context) ([]TajweedRule, error)
	GetTajweedRuleByID(ctx context.Context, ruleID string) (*TajweedRule, error)
	GetTajweedRuleByName(ctx context.Context, ruleName string) (*TajweedRule, error)

	// TajweedRule write operations
	CreateTajweedRule(ctx context.Context, rule *TajweedRule) (*TajweedRule, error)
	UpdateTajweedRule(ctx context.Context, rule *TajweedRule) (*TajweedRule, error)
	DeleteTajweedRule(ctx context.Context, ruleID string) error
}

// TransactionManager defines the interface for database transaction management
type TransactionManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

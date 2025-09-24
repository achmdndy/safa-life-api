package translation

import "context"

// TranslationRepository defines the interface for Translation (edition) data access.
type TranslationRepository interface {
	GetAll(ctx context.Context) ([]Translation, error)
	GetByID(ctx context.Context, id string) (*Translation, error)
	CreateTranslation(ctx context.Context, translation *Translation) (*Translation, error)
	UpdateTranslation(ctx context.Context, translation *Translation) (*Translation, error)
	DeleteTranslation(ctx context.Context, id string) error
}

// AyahTranslationRepository defines the interface for AyahTranslation data access.
type AyahTranslationRepository interface {
	GetAyahTranslation(ctx context.Context, translationID string, surahID, ayahID int) (*AyahTranslation, error)
	GetByAyah(ctx context.Context, surahID, ayahID int) ([]AyahTranslation, error)
	GetBySurah(ctx context.Context, translationID string, surahID int) ([]AyahTranslation, error)
	CreateAyahTranslation(ctx context.Context, ayahTranslation *AyahTranslation) (*AyahTranslation, error)
	UpdateAyahTranslation(ctx context.Context, ayahTranslation *AyahTranslation) (*AyahTranslation, error)
	DeleteAyahTranslation(ctx context.Context, translationID string, surahID, ayahID int) error
}

// FullTranslationRepository defines the aggregate repository interface.
type FullTranslationRepository interface {
	TranslationRepository
	AyahTranslationRepository
	TransactionManager
}
package translation

import "context"

// TranslationService defines the business logic for translation operations.
type TranslationService interface {
	// Translation (edition) operations
	GetAllTranslations(ctx context.Context) ([]Translation, error)
	GetTranslationByID(ctx context.Context, id string) (*Translation, error)
	CreateTranslation(ctx context.Context, translation *Translation) (*Translation, error)
	UpdateTranslation(ctx context.Context, translation *Translation) (*Translation, error)
	DeleteTranslation(ctx context.Context, id string) error

	// AyahTranslation operations
	GetAyahTranslation(ctx context.Context, translationID string, surahID, ayahID int) (*AyahTranslation, error)
	GetTranslationsForAyah(ctx context.Context, surahID, ayahID int) ([]AyahTranslation, error)
	GetTranslationsForSurah(ctx context.Context, translationID string, surahID int) ([]AyahTranslation, error)
	CreateAyahTranslation(ctx context.Context, ayahTranslation *AyahTranslation) (*AyahTranslation, error)
	UpdateAyahTranslation(ctx context.Context, ayahTranslation *AyahTranslation) (*AyahTranslation, error)
	DeleteAyahTranslation(ctx context.Context, translationID string, surahID, ayahID int) error
}

// TransactionManager defines the interface for database transaction management
type TransactionManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
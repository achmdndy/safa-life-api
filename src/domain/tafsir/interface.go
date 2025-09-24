package tafsir

import "context"

// TafsirService defines the business logic for tafsir operations.
type TafsirService interface {
	GetAllTafsirs(ctx context.Context) ([]Tafsir, error)
	GetTafsirByID(ctx context.Context, id string) (*Tafsir, error)
	CreateTafsir(ctx context.Context, tafsir *Tafsir) (*Tafsir, error)
	UpdateTafsir(ctx context.Context, tafsir *Tafsir) (*Tafsir, error)
	DeleteTafsir(ctx context.Context, id string) error

	GetAyahTafsir(ctx context.Context, tafsirID string, surahID, ayahID int) (*AyahTafsir, error)
	GetTafsirsForAyah(ctx context.Context, surahID, ayahID int) ([]AyahTafsir, error)
	GetTafsirsForSurah(ctx context.Context, tafsirID string, surahID int) ([]AyahTafsir, error)
	CreateAyahTafsir(ctx context.Context, ayahTafsir *AyahTafsir) (*AyahTafsir, error)
	UpdateAyahTafsir(ctx context.Context, ayahTafsir *AyahTafsir) (*AyahTafsir, error)
	DeleteAyahTafsir(ctx context.Context, tafsirID string, surahID, ayahID int) error
}

// TransactionManager defines the interface for database transaction management.
type TransactionManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
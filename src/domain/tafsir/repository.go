package tafsir

import "context"

// TafsirRepository defines the interface for Tafsir (edition) data access.
type TafsirRepository interface {
	GetAll(ctx context.Context) ([]Tafsir, error)
	GetByID(ctx context.Context, id string) (*Tafsir, error)
	CreateTafsir(ctx context.Context, tafsir *Tafsir) (*Tafsir, error)
	UpdateTafsir(ctx context.Context, tafsir *Tafsir) (*Tafsir, error)
	DeleteTafsir(ctx context.Context, id string) error
}

// AyahTafsirRepository defines the interface for AyahTafsir data access.
type AyahTafsirRepository interface {
	GetAyahTafsir(ctx context.Context, tafsirID string, surahID, ayahID int) (*AyahTafsir, error)
	GetByAyah(ctx context.Context, surahID, ayahID int) ([]AyahTafsir, error)
	GetBySurah(ctx context.Context, tafsirID string, surahID int) ([]AyahTafsir, error)
	CreateAyahTafsir(ctx context.Context, ayahTafsir *AyahTafsir) (*AyahTafsir, error)
	UpdateAyahTafsir(ctx context.Context, ayahTafsir *AyahTafsir) (*AyahTafsir, error)
	DeleteAyahTafsir(ctx context.Context, tafsirID string, surahID, ayahID int) error
}

// FullTafsirRepository defines the aggregate repository interface.
type FullTafsirRepository interface {
	TafsirRepository
	AyahTafsirRepository
	TransactionManager
}

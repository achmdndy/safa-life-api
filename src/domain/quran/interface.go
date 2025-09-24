package quran

import (
	"context"
)

// QuranService defines the business logic interface for Quran operations
type QuranService interface {
	// Surah read operations
	GetAllSurahs(ctx context.Context) ([]Surah, error)
	GetSurahByID(ctx context.Context, id int) (*Surah, error)
	GetSurahWithAyahs(ctx context.Context, id int) (*SurahWithAyahs, error)
	
	// Surah write operations
	CreateSurah(ctx context.Context, surah *Surah) (*Surah, error)
	UpdateSurah(ctx context.Context, surah *Surah) (*Surah, error)
	DeleteSurah(ctx context.Context, id int) error
	
	// Ayah read operations
	GetAyahsBySurah(ctx context.Context, surahID int) ([]Ayah, error)
	GetAyahByID(ctx context.Context, surahID, ayahID int) (*Ayah, error)
	GetAyahsByPage(ctx context.Context, pageNumber int) ([]Ayah, error)
	GetAyahsByJuz(ctx context.Context, juzNumber int) ([]Ayah, error)
	
	// Ayah write operations
	CreateAyah(ctx context.Context, ayah *Ayah) (*Ayah, error)
	UpdateAyah(ctx context.Context, ayah *Ayah) (*Ayah, error)
	DeleteAyah(ctx context.Context, surahID, ayahID int) error
	
	// Juz read operations
	GetAllJuz(ctx context.Context) ([]Juz, error)
	GetJuzByID(ctx context.Context, id int) (*Juz, error)
	GetJuzWithContent(ctx context.Context, id int) (*JuzWithContent, error)
	
	// Juz write operations
	CreateJuz(ctx context.Context, juz *Juz) (*Juz, error)
	UpdateJuz(ctx context.Context, juz *Juz) (*Juz, error)
	DeleteJuz(ctx context.Context, id int) error
	
	// Search operations
	SearchAyahs(ctx context.Context, query string) ([]Ayah, error)
	SearchSurahs(ctx context.Context, query string) ([]Surah, error)
}

// TransactionManager defines the interface for database transaction management
type TransactionManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
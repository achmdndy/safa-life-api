package quran

import (
	"context"
)

// SurahRepository defines the interface for Surah data access
type SurahRepository interface {
	GetAllSurahs(ctx context.Context) ([]Surah, error)
	GetSurahByID(ctx context.Context, id int) (*Surah, error)
	SearchSurahs(ctx context.Context, query string) ([]Surah, error)
	CreateSurah(ctx context.Context, surah *Surah) (*Surah, error)
	UpdateSurah(ctx context.Context, surah *Surah) (*Surah, error)
	DeleteSurah(ctx context.Context, id int) error
}

// AyahRepository defines the interface for Ayah data access
type AyahRepository interface {
	GetAyahsBySurah(ctx context.Context, surahID int) ([]Ayah, error)
	GetAyahByID(ctx context.Context, surahID, ayahID int) (*Ayah, error)
	GetAyahsByPage(ctx context.Context, pageNumber int) ([]Ayah, error)
	GetAyahsByJuz(ctx context.Context, juzNumber int) ([]Ayah, error)
	SearchAyahs(ctx context.Context, query string) ([]Ayah, error)
	CreateAyah(ctx context.Context, ayah *Ayah) (*Ayah, error)
	UpdateAyah(ctx context.Context, ayah *Ayah) (*Ayah, error)
	DeleteAyah(ctx context.Context, surahID, ayahID int) error
	GetAyahsByRange(ctx context.Context, startSurah, startAyah, endSurah, endAyah int) ([]Ayah, error)
}

// JuzRepository defines the interface for Juz data access
type JuzRepository interface {
	GetAllJuz(ctx context.Context) ([]Juz, error)
	GetJuzByID(ctx context.Context, id int) (*Juz, error)
	CreateJuz(ctx context.Context, juz *Juz) (*Juz, error)
	UpdateJuz(ctx context.Context, juz *Juz) (*Juz, error)
	DeleteJuz(ctx context.Context, id int) error
}

// QuranRepository defines the aggregate repository interface
type QuranRepository interface {
	SurahRepository
	AyahRepository
	JuzRepository
	TransactionManager
}
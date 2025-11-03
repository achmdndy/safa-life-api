package quran

import (
	"context"

	"github.com/safalife/core-api/src/domain/core"
)

// SurahServiceInterface defines the business logic for surah operations.
type SurahServiceInterface interface {
	GetSurahById(ctx context.Context, id core.UUID) (*Surah, error)
	GetSurahByNumber(ctx context.Context, number int) (*Surah, error)
	GetAllSurahs(ctx context.Context, limit, offset int) ([]*Surah, error)
	GetSurahsByRevelationPlace(ctx context.Context, place string, limit, offset int) ([]*Surah, error)
	CreateSurah(ctx context.Context, surah *Surah) (*Surah, error)
	UpdateSurah(ctx context.Context, surah *Surah) (*Surah, error)
	DeleteSurah(ctx context.Context, id core.UUID) error
	CountSurahs(ctx context.Context) (int64, error)
	CountSurahsByRevelationPlace(ctx context.Context, place string) (int64, error)
	
	// Eager loading methods
	GetSurahByIdWithAyahs(ctx context.Context, id core.UUID) (*SurahWithAyahs, error)
	GetSurahByNumberWithAyahs(ctx context.Context, number int) (*SurahWithAyahs, error)
	GetAllSurahsWithAyahs(ctx context.Context, limit, offset int) ([]*SurahWithAyahs, error)
}

// AyahServiceInterface defines the business logic for ayah operations.
type AyahServiceInterface interface {
	GetAyahById(ctx context.Context, id core.UUID) (*Ayah, error)
	GetAyahsBySurahId(ctx context.Context, surahId core.UUID, limit, offset int) ([]*Ayah, error)
	GetAyahsByJuzNumber(ctx context.Context, juzNumber int, limit, offset int) ([]*Ayah, error)
	GetAyahsByPageNumber(ctx context.Context, pageNumber int, limit, offset int) ([]*Ayah, error)
	GetAyahsByHizbNumber(ctx context.Context, hizbNumber int, limit, offset int) ([]*Ayah, error)
	GetAyahsByManzilNumber(ctx context.Context, manzilNumber int, limit, offset int) ([]*Ayah, error)
	GetAllAyahs(ctx context.Context, limit, offset int) ([]*Ayah, error)
	CreateAyah(ctx context.Context, ayah *Ayah) (*Ayah, error)
	UpdateAyah(ctx context.Context, ayah *Ayah) (*Ayah, error)
	DeleteAyah(ctx context.Context, id core.UUID) error
	CountAyahs(ctx context.Context) (int64, error)
	CountAyahsBySurahId(ctx context.Context, surahId core.UUID) (int64, error)
	CountAyahsByJuzNumber(ctx context.Context, juzNumber int) (int64, error)
	CountAyahsByPageNumber(ctx context.Context, pageNumber int) (int64, error)
	
	// Eager loading methods
	GetAyahByIdWithSurah(ctx context.Context, id core.UUID) (*AyahWithSurah, error)
	GetAyahsBySurahIdWithSurah(ctx context.Context, surahId core.UUID, limit, offset int) ([]*AyahWithSurah, error)
	GetAllAyahsWithSurah(ctx context.Context, limit, offset int) ([]*AyahWithSurah, error)
}

// JuzServiceInterface defines the business logic for juz operations.
type JuzServiceInterface interface {
	GetJuzById(ctx context.Context, id core.UUID) (*Juz, error)
	GetJuzByNumber(ctx context.Context, number int) (*Juz, error)
	GetAllJuz(ctx context.Context, limit, offset int) ([]*Juz, error)
	CreateJuz(ctx context.Context, juz *Juz) (*Juz, error)
	UpdateJuz(ctx context.Context, juz *Juz) (*Juz, error)
	DeleteJuz(ctx context.Context, id core.UUID) error
	CountJuz(ctx context.Context) (int64, error)
	
	// Eager loading methods
	GetJuzByIdWithRelations(ctx context.Context, id core.UUID) (*JuzWithRelations, error)
	GetJuzByNumberWithRelations(ctx context.Context, number int) (*JuzWithRelations, error)
	GetAllJuzWithRelations(ctx context.Context, limit, offset int) ([]*JuzWithRelations, error)
}
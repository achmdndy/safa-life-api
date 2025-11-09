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

// TranslationEditionServiceInterface defines business logic for translation edition operations.
type TranslationEditionServiceInterface interface {
	GetTranslationEditionById(ctx context.Context, id core.UUID) (*TranslationEdition, error)
	GetAllTranslationEditions(ctx context.Context, limit, offset int) ([]*TranslationEdition, error)
	GetTranslationEditionsByLanguage(ctx context.Context, language string, limit, offset int) ([]*TranslationEdition, error)
	CreateTranslationEdition(ctx context.Context, edition *TranslationEdition) (*TranslationEdition, error)
	UpdateTranslationEdition(ctx context.Context, edition *TranslationEdition) (*TranslationEdition, error)
	DeleteTranslationEdition(ctx context.Context, id core.UUID) error
	CountTranslationEditions(ctx context.Context) (int64, error)
	CountTranslationEditionsByLanguage(ctx context.Context, language string) (int64, error)
}

// AyahTranslationServiceInterface defines business logic for ayah translation operations.
type AyahTranslationServiceInterface interface {
	GetAyahTranslationById(ctx context.Context, id core.UUID) (*AyahTranslation, error)
	GetAyahTranslationByAyahAndEdition(ctx context.Context, ayahId core.UUID, editionId core.UUID) (*AyahTranslation, error)
	GetAyahTranslationsBySurahAndEdition(ctx context.Context, surahId core.UUID, editionId core.UUID, limit, offset int) ([]*AyahTranslation, error)
	CreateAyahTranslation(ctx context.Context, t *AyahTranslation) (*AyahTranslation, error)
	UpdateAyahTranslation(ctx context.Context, t *AyahTranslation) (*AyahTranslation, error)
	DeleteAyahTranslation(ctx context.Context, id core.UUID) error
	CountAyahTranslations(ctx context.Context) (int64, error)
	CountAyahTranslationsByEdition(ctx context.Context, editionId core.UUID) (int64, error)
	CountAyahTranslationsBySurahAndEdition(ctx context.Context, surahId core.UUID, editionId core.UUID) (int64, error)
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

// ReciterServiceInterface defines business logic for reciter operations.
type ReciterServiceInterface interface {
	GetReciterById(ctx context.Context, id core.UUID) (*Reciter, error)
	GetReciterByName(ctx context.Context, name string) (*Reciter, error)
	GetAllReciters(ctx context.Context, limit, offset int) ([]*Reciter, error)
	CreateReciter(ctx context.Context, reciter *Reciter) (*Reciter, error)
	UpdateReciter(ctx context.Context, reciter *Reciter) (*Reciter, error)
	DeleteReciter(ctx context.Context, id core.UUID) error
	CountReciters(ctx context.Context) (int64, error)
}

// AyahAudioFileServiceInterface defines business logic for ayah audio file operations.
type AyahAudioFileServiceInterface interface {
	GetAyahAudioFileById(ctx context.Context, id core.UUID) (*AyahAudioFile, error)
	GetAyahAudioFileByAyahAndReciter(ctx context.Context, ayahId core.UUID, reciterId core.UUID) (*AyahAudioFile, error)
	GetAyahAudioFilesBySurahAndReciter(ctx context.Context, surahId core.UUID, reciterId core.UUID, limit, offset int) ([]*AyahAudioFile, error)
	CreateAyahAudioFile(ctx context.Context, audio *AyahAudioFile) (*AyahAudioFile, error)
	UpdateAyahAudioFile(ctx context.Context, audio *AyahAudioFile) (*AyahAudioFile, error)
	DeleteAyahAudioFile(ctx context.Context, id core.UUID) error
	CountAyahAudioFiles(ctx context.Context) (int64, error)
	CountAyahAudioFilesBySurahAndReciter(ctx context.Context, surahId core.UUID, reciterId core.UUID) (int64, error)
}

// BookmarkAyahServiceInterface defines business logic for ayah bookmarks.
type BookmarkAyahServiceInterface interface {
	GetBookmarkAyahById(ctx context.Context, id core.UUID) (*BookmarkAyah, error)
	GetBookmarkAyahsByUser(ctx context.Context, userId core.UUID, limit, offset int) ([]*BookmarkAyah, error)
	GetBookmarkAyahByUserAndAyah(ctx context.Context, userId core.UUID, ayahId core.UUID) (*BookmarkAyah, error)
	CreateBookmarkAyah(ctx context.Context, bookmark *BookmarkAyah) (*BookmarkAyah, error)
	DeleteBookmarkAyah(ctx context.Context, id core.UUID) error
	CountBookmarkAyahsByUser(ctx context.Context, userId core.UUID) (int64, error)

	// Eager loading
	GetBookmarkAyahByIdWithAyah(ctx context.Context, id core.UUID) (*BookmarkAyahWithAyah, error)
}

// LastReadServiceInterface defines business logic for tracking last read progress per user.
type LastReadServiceInterface interface {
	GetLastReadById(ctx context.Context, id core.UUID) (*LastRead, error)
	GetLastReadsByUser(ctx context.Context, userId core.UUID, limit, offset int) ([]*LastRead, error)
	GetLastReadByUserAndSurah(ctx context.Context, userId core.UUID, surahId core.UUID) (*LastRead, error)
	CreateLastRead(ctx context.Context, lr *LastRead) (*LastRead, error)
	UpdateLastRead(ctx context.Context, lr *LastRead) (*LastRead, error)
	DeleteLastRead(ctx context.Context, id core.UUID) error
	CountLastReadsByUser(ctx context.Context, userId core.UUID) (int64, error)

	// Eager loading
	GetLastReadByIdWithRelations(ctx context.Context, id core.UUID) (*LastReadWithRelations, error)
	GetLastReadsByUserWithRelations(ctx context.Context, userId core.UUID, limit, offset int) ([]*LastReadWithRelations, error)
}

// ProgressHatamServiceInterface defines business logic for hatam progress tracking.
type ProgressHatamServiceInterface interface {
	GetProgressHatamById(ctx context.Context, id core.UUID) (*ProgressHatam, error)
	GetProgressHatamByUser(ctx context.Context, userId core.UUID, limit, offset int) ([]*ProgressHatam, error)
	GetProgressHatamByUserAndJuz(ctx context.Context, userId core.UUID, juzId core.UUID) (*ProgressHatam, error)
	CreateProgressHatam(ctx context.Context, p *ProgressHatam) (*ProgressHatam, error)
	UpdateProgressHatam(ctx context.Context, p *ProgressHatam) (*ProgressHatam, error)
	DeleteProgressHatam(ctx context.Context, id core.UUID) error
	CountProgressHatamByUser(ctx context.Context, userId core.UUID) (int64, error)

	// Eager loading
	GetProgressHatamByIdWithRelations(ctx context.Context, id core.UUID) (*ProgressHatamWithRelations, error)
	GetProgressHatamByUserWithRelations(ctx context.Context, userId core.UUID, limit, offset int) ([]*ProgressHatamWithRelations, error)
}

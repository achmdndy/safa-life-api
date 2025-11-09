package quran

import (
	"context"

	"github.com/safalife/core-api/src/domain/core"
)

// SurahRepositoryInterface defines the data access layer for surah operations.
type SurahRepositoryInterface interface {
	GetById(ctx context.Context, id core.UUID) (*Surah, error)
	GetByNumber(ctx context.Context, number int) (*Surah, error)
	GetAll(ctx context.Context, limit, offset int) ([]*Surah, error)
	GetByRevelationPlace(ctx context.Context, place string, limit, offset int) ([]*Surah, error)
	Create(ctx context.Context, surah *Surah) (*Surah, error)
	Update(ctx context.Context, surah *Surah) (*Surah, error)
	Delete(ctx context.Context, id core.UUID) error
	Count(ctx context.Context) (int64, error)
	CountByRevelationPlace(ctx context.Context, place string) (int64, error)

	// Eager loading methods
	GetByIdWithAyahs(ctx context.Context, id core.UUID) (*SurahWithAyahs, error)
	GetByNumberWithAyahs(ctx context.Context, number int) (*SurahWithAyahs, error)
	GetAllWithAyahs(ctx context.Context, limit, offset int) ([]*SurahWithAyahs, error)
}

// AyahRepositoryInterface defines the data access layer for ayah operations.
type AyahRepositoryInterface interface {
	GetById(ctx context.Context, id core.UUID) (*Ayah, error)
	GetBySurahId(ctx context.Context, surahId core.UUID, limit, offset int) ([]*Ayah, error)
	GetByJuzNumber(ctx context.Context, juzNumber int, limit, offset int) ([]*Ayah, error)
	GetByPageNumber(ctx context.Context, pageNumber int, limit, offset int) ([]*Ayah, error)
	GetByHizbNumber(ctx context.Context, hizbNumber int, limit, offset int) ([]*Ayah, error)
	GetByManzilNumber(ctx context.Context, manzilNumber int, limit, offset int) ([]*Ayah, error)
	GetAll(ctx context.Context, limit, offset int) ([]*Ayah, error)
	Create(ctx context.Context, ayah *Ayah) (*Ayah, error)
	Update(ctx context.Context, ayah *Ayah) (*Ayah, error)
	Delete(ctx context.Context, id core.UUID) error
	Count(ctx context.Context) (int64, error)
	CountBySurahId(ctx context.Context, surahId core.UUID) (int64, error)
	CountByJuzNumber(ctx context.Context, juzNumber int) (int64, error)
	CountByPageNumber(ctx context.Context, pageNumber int) (int64, error)

	// Eager loading methods
	GetByIdWithSurah(ctx context.Context, id core.UUID) (*AyahWithSurah, error)
	GetBySurahIdWithSurah(ctx context.Context, surahId core.UUID, limit, offset int) ([]*AyahWithSurah, error)
	GetAllWithSurah(ctx context.Context, limit, offset int) ([]*AyahWithSurah, error)
}

// JuzRepositoryInterface defines the data access layer for juz operations.
type JuzRepositoryInterface interface {
	GetById(ctx context.Context, id core.UUID) (*Juz, error)
	GetByNumber(ctx context.Context, number int) (*Juz, error)
	GetAll(ctx context.Context, limit, offset int) ([]*Juz, error)
	Create(ctx context.Context, juz *Juz) (*Juz, error)
	Update(ctx context.Context, juz *Juz) (*Juz, error)
	Delete(ctx context.Context, id core.UUID) error
	Count(ctx context.Context) (int64, error)

	// Eager loading methods
	GetByIdWithRelations(ctx context.Context, id core.UUID) (*JuzWithRelations, error)
	GetByNumberWithRelations(ctx context.Context, number int) (*JuzWithRelations, error)
	GetAllWithRelations(ctx context.Context, limit, offset int) ([]*JuzWithRelations, error)
}

// TranslationEditionRepositoryInterface defines the data access layer for translation editions.
type TranslationEditionRepositoryInterface interface {
	GetById(ctx context.Context, id core.UUID) (*TranslationEdition, error)
	GetAll(ctx context.Context, limit, offset int) ([]*TranslationEdition, error)
	GetByLanguage(ctx context.Context, language string, limit, offset int) ([]*TranslationEdition, error)
	Create(ctx context.Context, edition *TranslationEdition) (*TranslationEdition, error)
	Update(ctx context.Context, edition *TranslationEdition) (*TranslationEdition, error)
	Delete(ctx context.Context, id core.UUID) error
	Count(ctx context.Context) (int64, error)
	CountByLanguage(ctx context.Context, language string) (int64, error)
}

// AyahTranslationRepositoryInterface defines the data access layer for ayah translations.
type AyahTranslationRepositoryInterface interface {
	GetById(ctx context.Context, id core.UUID) (*AyahTranslation, error)
	GetByAyahAndEdition(ctx context.Context, ayahId core.UUID, editionId core.UUID) (*AyahTranslation, error)
	GetBySurahAndEdition(ctx context.Context, surahId core.UUID, editionId core.UUID, limit, offset int) ([]*AyahTranslation, error)
	Create(ctx context.Context, t *AyahTranslation) (*AyahTranslation, error)
	Update(ctx context.Context, t *AyahTranslation) (*AyahTranslation, error)
	Delete(ctx context.Context, id core.UUID) error
	Count(ctx context.Context) (int64, error)
	CountByEdition(ctx context.Context, editionId core.UUID) (int64, error)
	CountBySurahAndEdition(ctx context.Context, surahId core.UUID, editionId core.UUID) (int64, error)
}

// ReciterRepositoryInterface defines the data access layer for reciters.
type ReciterRepositoryInterface interface {
	GetById(ctx context.Context, id core.UUID) (*Reciter, error)
	GetByName(ctx context.Context, name string) (*Reciter, error)
	GetAll(ctx context.Context, limit, offset int) ([]*Reciter, error)
	Create(ctx context.Context, reciter *Reciter) (*Reciter, error)
	Update(ctx context.Context, reciter *Reciter) (*Reciter, error)
	Delete(ctx context.Context, id core.UUID) error
	Count(ctx context.Context) (int64, error)
}

// AyahAudioFileRepositoryInterface defines the data access layer for ayah audio files.
type AyahAudioFileRepositoryInterface interface {
	GetById(ctx context.Context, id core.UUID) (*AyahAudioFile, error)
	GetByAyahAndReciter(ctx context.Context, ayahId core.UUID, reciterId core.UUID) (*AyahAudioFile, error)
	GetBySurahAndReciter(ctx context.Context, surahId core.UUID, reciterId core.UUID, limit, offset int) ([]*AyahAudioFile, error)
	Create(ctx context.Context, audio *AyahAudioFile) (*AyahAudioFile, error)
	Update(ctx context.Context, audio *AyahAudioFile) (*AyahAudioFile, error)
	Delete(ctx context.Context, id core.UUID) error
	Count(ctx context.Context) (int64, error)
	CountBySurahAndReciter(ctx context.Context, surahId core.UUID, reciterId core.UUID) (int64, error)
}

// BookmarkAyahRepositoryInterface defines the data access layer for ayah bookmarks.
type BookmarkAyahRepositoryInterface interface {
	GetById(ctx context.Context, id core.UUID) (*BookmarkAyah, error)
	GetByUser(ctx context.Context, userId core.UUID, limit, offset int) ([]*BookmarkAyah, error)
	GetByUserAndAyah(ctx context.Context, userId core.UUID, ayahId core.UUID) (*BookmarkAyah, error)
	Create(ctx context.Context, bookmark *BookmarkAyah) (*BookmarkAyah, error)
	Delete(ctx context.Context, id core.UUID) error
	CountByUser(ctx context.Context, userId core.UUID) (int64, error)

	// Eager loading
	GetByIdWithAyah(ctx context.Context, id core.UUID) (*BookmarkAyahWithAyah, error)
}

// LastReadRepositoryInterface defines the data access layer for tracking last read per user.
type LastReadRepositoryInterface interface {
	GetById(ctx context.Context, id core.UUID) (*LastRead, error)
	GetByUser(ctx context.Context, userId core.UUID, limit, offset int) ([]*LastRead, error)
	GetByUserAndSurah(ctx context.Context, userId core.UUID, surahId core.UUID) (*LastRead, error)
	Create(ctx context.Context, lr *LastRead) (*LastRead, error)
	Update(ctx context.Context, lr *LastRead) (*LastRead, error)
	Delete(ctx context.Context, id core.UUID) error
	CountByUser(ctx context.Context, userId core.UUID) (int64, error)

	// Eager loading
	GetByIdWithRelations(ctx context.Context, id core.UUID) (*LastReadWithRelations, error)
	GetByUserWithRelations(ctx context.Context, userId core.UUID, limit, offset int) ([]*LastReadWithRelations, error)
}

// ProgressHatamRepositoryInterface defines the data access layer for hatam progress.
type ProgressHatamRepositoryInterface interface {
	GetById(ctx context.Context, id core.UUID) (*ProgressHatam, error)
	GetByUser(ctx context.Context, userId core.UUID, limit, offset int) ([]*ProgressHatam, error)
	GetByUserAndJuz(ctx context.Context, userId core.UUID, juzId core.UUID) (*ProgressHatam, error)
	Create(ctx context.Context, p *ProgressHatam) (*ProgressHatam, error)
	Update(ctx context.Context, p *ProgressHatam) (*ProgressHatam, error)
	Delete(ctx context.Context, id core.UUID) error
	CountByUser(ctx context.Context, userId core.UUID) (int64, error)

	// Eager loading
	GetByIdWithRelations(ctx context.Context, id core.UUID) (*ProgressHatamWithRelations, error)
	GetByUserWithRelations(ctx context.Context, userId core.UUID, limit, offset int) ([]*ProgressHatamWithRelations, error)
}

// QuranRepositoryInterface defines the aggregate repository interface for all Quran entities.
type QuranRepositoryInterface interface {
	SurahRepository() SurahRepositoryInterface
	AyahRepository() AyahRepositoryInterface
	JuzRepository() JuzRepositoryInterface
	// Note: other repositories exposed separately via containers
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

// QuranRepository implements QuranRepositoryInterface with transaction support.
type QuranRepository struct {
	surahRepo          SurahRepositoryInterface
	ayahRepo           AyahRepositoryInterface
	juzRepo            JuzRepositoryInterface
	transactionManager core.ContextTransactionManager
}

// NewQuranRepository creates a new QuranRepository instance.
func NewQuranRepository(
	surahRepo SurahRepositoryInterface,
	ayahRepo AyahRepositoryInterface,
	juzRepo JuzRepositoryInterface,
	transactionManager core.ContextTransactionManager,
) QuranRepositoryInterface {
	return &QuranRepository{
		surahRepo:          surahRepo,
		ayahRepo:           ayahRepo,
		juzRepo:            juzRepo,
		transactionManager: transactionManager,
	}
}

// SurahRepository returns the surah repository.
func (r *QuranRepository) SurahRepository() SurahRepositoryInterface {
	return r.surahRepo
}

// AyahRepository returns the ayah repository.
func (r *QuranRepository) AyahRepository() AyahRepositoryInterface {
	return r.ayahRepo
}

// JuzRepository returns the juz repository.
func (r *QuranRepository) JuzRepository() JuzRepositoryInterface {
	return r.juzRepo
}

// WithTransaction executes a function within a database transaction.
func (r *QuranRepository) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.transactionManager.WithTransaction(ctx, fn)
}

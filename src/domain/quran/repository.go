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

// QuranRepositoryInterface defines the aggregate repository interface for all Quran entities.
type QuranRepositoryInterface interface {
	SurahRepository() SurahRepositoryInterface
	AyahRepository() AyahRepositoryInterface
	JuzRepository() JuzRepositoryInterface
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

// QuranRepository implements QuranRepositoryInterface with transaction support.
type QuranRepository struct {
	surahRepo         SurahRepositoryInterface
	ayahRepo          AyahRepositoryInterface
	juzRepo           JuzRepositoryInterface
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
		surahRepo:         surahRepo,
		ayahRepo:          ayahRepo,
		juzRepo:           juzRepo,
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
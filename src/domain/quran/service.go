package quran

import (
	"context"

	"github.com/safalife/core-api/src/domain/core"
)

// SurahService implements SurahServiceInterface.
type SurahService struct {
	surahRepo SurahRepositoryInterface
}

// NewSurahService creates a new SurahService instance.
func NewSurahService(surahRepo SurahRepositoryInterface) SurahServiceInterface {
	return &SurahService{
		surahRepo: surahRepo,
	}
}

func (s *SurahService) GetSurahById(ctx context.Context, id core.UUID) (*Surah, error) {
	return s.surahRepo.GetById(ctx, id)
}

func (s *SurahService) GetSurahByNumber(ctx context.Context, number int) (*Surah, error) {
	return s.surahRepo.GetByNumber(ctx, number)
}

func (s *SurahService) GetAllSurahs(ctx context.Context, limit, offset int) ([]*Surah, error) {
	return s.surahRepo.GetAll(ctx, limit, offset)
}

func (s *SurahService) GetSurahsByRevelationPlace(ctx context.Context, place string, limit, offset int) ([]*Surah, error) {
	return s.surahRepo.GetByRevelationPlace(ctx, place, limit, offset)
}

func (s *SurahService) CreateSurah(ctx context.Context, surah *Surah) (*Surah, error) {
	return s.surahRepo.Create(ctx, surah)
}

func (s *SurahService) UpdateSurah(ctx context.Context, surah *Surah) (*Surah, error) {
	return s.surahRepo.Update(ctx, surah)
}

func (s *SurahService) DeleteSurah(ctx context.Context, id core.UUID) error {
	return s.surahRepo.Delete(ctx, id)
}

func (s *SurahService) CountSurahs(ctx context.Context) (int64, error) {
	return s.surahRepo.Count(ctx)
}

func (s *SurahService) CountSurahsByRevelationPlace(ctx context.Context, place string) (int64, error) {
	return s.surahRepo.CountByRevelationPlace(ctx, place)
}

// Eager loading methods
func (s *SurahService) GetSurahByIdWithAyahs(ctx context.Context, id core.UUID) (*SurahWithAyahs, error) {
	return s.surahRepo.GetByIdWithAyahs(ctx, id)
}

func (s *SurahService) GetSurahByNumberWithAyahs(ctx context.Context, number int) (*SurahWithAyahs, error) {
	return s.surahRepo.GetByNumberWithAyahs(ctx, number)
}

func (s *SurahService) GetAllSurahsWithAyahs(ctx context.Context, limit, offset int) ([]*SurahWithAyahs, error) {
	return s.surahRepo.GetAllWithAyahs(ctx, limit, offset)
}

// AyahService implements AyahServiceInterface.
type AyahService struct {
	ayahRepo AyahRepositoryInterface
}

// NewAyahService creates a new AyahService instance.
func NewAyahService(ayahRepo AyahRepositoryInterface) AyahServiceInterface {
	return &AyahService{
		ayahRepo: ayahRepo,
	}
}

func (s *AyahService) GetAyahById(ctx context.Context, id core.UUID) (*Ayah, error) {
	return s.ayahRepo.GetById(ctx, id)
}

func (s *AyahService) GetAyahsBySurahId(ctx context.Context, surahId core.UUID, limit, offset int) ([]*Ayah, error) {
	return s.ayahRepo.GetBySurahId(ctx, surahId, limit, offset)
}

func (s *AyahService) GetAyahsByJuzNumber(ctx context.Context, juzNumber int, limit, offset int) ([]*Ayah, error) {
	return s.ayahRepo.GetByJuzNumber(ctx, juzNumber, limit, offset)
}

func (s *AyahService) GetAyahsByPageNumber(ctx context.Context, pageNumber int, limit, offset int) ([]*Ayah, error) {
	return s.ayahRepo.GetByPageNumber(ctx, pageNumber, limit, offset)
}

func (s *AyahService) GetAyahsByHizbNumber(ctx context.Context, hizbNumber int, limit, offset int) ([]*Ayah, error) {
	return s.ayahRepo.GetByHizbNumber(ctx, hizbNumber, limit, offset)
}

func (s *AyahService) GetAyahsByManzilNumber(ctx context.Context, manzilNumber int, limit, offset int) ([]*Ayah, error) {
	return s.ayahRepo.GetByManzilNumber(ctx, manzilNumber, limit, offset)
}

func (s *AyahService) GetAllAyahs(ctx context.Context, limit, offset int) ([]*Ayah, error) {
	return s.ayahRepo.GetAll(ctx, limit, offset)
}

func (s *AyahService) CreateAyah(ctx context.Context, ayah *Ayah) (*Ayah, error) {
	return s.ayahRepo.Create(ctx, ayah)
}

func (s *AyahService) UpdateAyah(ctx context.Context, ayah *Ayah) (*Ayah, error) {
	return s.ayahRepo.Update(ctx, ayah)
}

func (s *AyahService) DeleteAyah(ctx context.Context, id core.UUID) error {
	return s.ayahRepo.Delete(ctx, id)
}

func (s *AyahService) CountAyahs(ctx context.Context) (int64, error) {
	return s.ayahRepo.Count(ctx)
}

func (s *AyahService) CountAyahsBySurahId(ctx context.Context, surahId core.UUID) (int64, error) {
	return s.ayahRepo.CountBySurahId(ctx, surahId)
}

func (s *AyahService) CountAyahsByJuzNumber(ctx context.Context, juzNumber int) (int64, error) {
	return s.ayahRepo.CountByJuzNumber(ctx, juzNumber)
}

func (s *AyahService) CountAyahsByPageNumber(ctx context.Context, pageNumber int) (int64, error) {
	return s.ayahRepo.CountByPageNumber(ctx, pageNumber)
}

// Eager loading methods
func (s *AyahService) GetAyahByIdWithSurah(ctx context.Context, id core.UUID) (*AyahWithSurah, error) {
	return s.ayahRepo.GetByIdWithSurah(ctx, id)
}

func (s *AyahService) GetAyahsBySurahIdWithSurah(ctx context.Context, surahId core.UUID, limit, offset int) ([]*AyahWithSurah, error) {
	return s.ayahRepo.GetBySurahIdWithSurah(ctx, surahId, limit, offset)
}

func (s *AyahService) GetAllAyahsWithSurah(ctx context.Context, limit, offset int) ([]*AyahWithSurah, error) {
	return s.ayahRepo.GetAllWithSurah(ctx, limit, offset)
}

// JuzService implements JuzServiceInterface.
type JuzService struct {
	juzRepo JuzRepositoryInterface
}

// NewJuzService creates a new JuzService instance.
func NewJuzService(juzRepo JuzRepositoryInterface) JuzServiceInterface {
	return &JuzService{
		juzRepo: juzRepo,
	}
}

func (s *JuzService) GetJuzById(ctx context.Context, id core.UUID) (*Juz, error) {
	return s.juzRepo.GetById(ctx, id)
}

func (s *JuzService) GetJuzByNumber(ctx context.Context, number int) (*Juz, error) {
	return s.juzRepo.GetByNumber(ctx, number)
}

func (s *JuzService) GetAllJuz(ctx context.Context, limit, offset int) ([]*Juz, error) {
	return s.juzRepo.GetAll(ctx, limit, offset)
}

func (s *JuzService) CreateJuz(ctx context.Context, juz *Juz) (*Juz, error) {
	return s.juzRepo.Create(ctx, juz)
}

func (s *JuzService) UpdateJuz(ctx context.Context, juz *Juz) (*Juz, error) {
	return s.juzRepo.Update(ctx, juz)
}

func (s *JuzService) DeleteJuz(ctx context.Context, id core.UUID) error {
	return s.juzRepo.Delete(ctx, id)
}

func (s *JuzService) CountJuz(ctx context.Context) (int64, error) {
	return s.juzRepo.Count(ctx)
}

// Eager loading methods
func (s *JuzService) GetJuzByIdWithRelations(ctx context.Context, id core.UUID) (*JuzWithRelations, error) {
	return s.juzRepo.GetByIdWithRelations(ctx, id)
}

func (s *JuzService) GetJuzByNumberWithRelations(ctx context.Context, number int) (*JuzWithRelations, error) {
	return s.juzRepo.GetByNumberWithRelations(ctx, number)
}

func (s *JuzService) GetAllJuzWithRelations(ctx context.Context, limit, offset int) ([]*JuzWithRelations, error) {
	return s.juzRepo.GetAllWithRelations(ctx, limit, offset)
}

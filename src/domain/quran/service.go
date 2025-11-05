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

// TranslationEditionService implements TranslationEditionServiceInterface.
type TranslationEditionService struct {
	editionRepo TranslationEditionRepositoryInterface
}

// NewTranslationEditionService creates a new TranslationEditionService instance.
func NewTranslationEditionService(editionRepo TranslationEditionRepositoryInterface) TranslationEditionServiceInterface {
	return &TranslationEditionService{
		editionRepo: editionRepo,
	}
}

func (s *TranslationEditionService) GetTranslationEditionById(ctx context.Context, id core.UUID) (*TranslationEdition, error) {
	return s.editionRepo.GetById(ctx, id)
}

func (s *TranslationEditionService) GetAllTranslationEditions(ctx context.Context, limit, offset int) ([]*TranslationEdition, error) {
	return s.editionRepo.GetAll(ctx, limit, offset)
}

func (s *TranslationEditionService) GetTranslationEditionsByLanguage(ctx context.Context, language string, limit, offset int) ([]*TranslationEdition, error) {
	return s.editionRepo.GetByLanguage(ctx, language, limit, offset)
}

func (s *TranslationEditionService) CreateTranslationEdition(ctx context.Context, edition *TranslationEdition) (*TranslationEdition, error) {
	return s.editionRepo.Create(ctx, edition)
}

func (s *TranslationEditionService) UpdateTranslationEdition(ctx context.Context, edition *TranslationEdition) (*TranslationEdition, error) {
	return s.editionRepo.Update(ctx, edition)
}

func (s *TranslationEditionService) DeleteTranslationEdition(ctx context.Context, id core.UUID) error {
	return s.editionRepo.Delete(ctx, id)
}

func (s *TranslationEditionService) CountTranslationEditions(ctx context.Context) (int64, error) {
	return s.editionRepo.Count(ctx)
}

func (s *TranslationEditionService) CountTranslationEditionsByLanguage(ctx context.Context, language string) (int64, error) {
	return s.editionRepo.CountByLanguage(ctx, language)
}

// AyahTranslationService implements AyahTranslationServiceInterface.
type AyahTranslationService struct {
	translationRepo AyahTranslationRepositoryInterface
}

// NewAyahTranslationService creates a new AyahTranslationService instance.
func NewAyahTranslationService(translationRepo AyahTranslationRepositoryInterface) AyahTranslationServiceInterface {
	return &AyahTranslationService{
		translationRepo: translationRepo,
	}
}

func (s *AyahTranslationService) GetAyahTranslationById(ctx context.Context, id core.UUID) (*AyahTranslation, error) {
	return s.translationRepo.GetById(ctx, id)
}

func (s *AyahTranslationService) GetAyahTranslationByAyahAndEdition(ctx context.Context, ayahId core.UUID, editionId core.UUID) (*AyahTranslation, error) {
	return s.translationRepo.GetByAyahAndEdition(ctx, ayahId, editionId)
}

func (s *AyahTranslationService) GetAyahTranslationsBySurahAndEdition(ctx context.Context, surahId core.UUID, editionId core.UUID, limit, offset int) ([]*AyahTranslation, error) {
	return s.translationRepo.GetBySurahAndEdition(ctx, surahId, editionId, limit, offset)
}

func (s *AyahTranslationService) CreateAyahTranslation(ctx context.Context, t *AyahTranslation) (*AyahTranslation, error) {
	return s.translationRepo.Create(ctx, t)
}

func (s *AyahTranslationService) UpdateAyahTranslation(ctx context.Context, t *AyahTranslation) (*AyahTranslation, error) {
	return s.translationRepo.Update(ctx, t)
}

func (s *AyahTranslationService) DeleteAyahTranslation(ctx context.Context, id core.UUID) error {
	return s.translationRepo.Delete(ctx, id)
}

func (s *AyahTranslationService) CountAyahTranslations(ctx context.Context) (int64, error) {
	return s.translationRepo.Count(ctx)
}

func (s *AyahTranslationService) CountAyahTranslationsByEdition(ctx context.Context, editionId core.UUID) (int64, error) {
	return s.translationRepo.CountByEdition(ctx, editionId)
}

func (s *AyahTranslationService) CountAyahTranslationsBySurahAndEdition(ctx context.Context, surahId core.UUID, editionId core.UUID) (int64, error) {
	return s.translationRepo.CountBySurahAndEdition(ctx, surahId, editionId)
}

// ReciterService implements ReciterServiceInterface.
type ReciterService struct {
	reciterRepo ReciterRepositoryInterface
}

// NewReciterService creates a new ReciterService instance.
func NewReciterService(reciterRepo ReciterRepositoryInterface) ReciterServiceInterface {
	return &ReciterService{reciterRepo: reciterRepo}
}

func (s *ReciterService) GetReciterById(ctx context.Context, id core.UUID) (*Reciter, error) {
	return s.reciterRepo.GetById(ctx, id)
}

func (s *ReciterService) GetReciterByName(ctx context.Context, name string) (*Reciter, error) {
	return s.reciterRepo.GetByName(ctx, name)
}

func (s *ReciterService) GetAllReciters(ctx context.Context, limit, offset int) ([]*Reciter, error) {
	return s.reciterRepo.GetAll(ctx, limit, offset)
}

func (s *ReciterService) CreateReciter(ctx context.Context, reciter *Reciter) (*Reciter, error) {
	return s.reciterRepo.Create(ctx, reciter)
}

func (s *ReciterService) UpdateReciter(ctx context.Context, reciter *Reciter) (*Reciter, error) {
	return s.reciterRepo.Update(ctx, reciter)
}

func (s *ReciterService) DeleteReciter(ctx context.Context, id core.UUID) error {
	return s.reciterRepo.Delete(ctx, id)
}

func (s *ReciterService) CountReciters(ctx context.Context) (int64, error) {
	return s.reciterRepo.Count(ctx)
}

// AyahAudioFileService implements AyahAudioFileServiceInterface.
type AyahAudioFileService struct {
	audioRepo AyahAudioFileRepositoryInterface
}

// NewAyahAudioFileService creates a new AyahAudioFileService instance.
func NewAyahAudioFileService(audioRepo AyahAudioFileRepositoryInterface) AyahAudioFileServiceInterface {
	return &AyahAudioFileService{audioRepo: audioRepo}
}

func (s *AyahAudioFileService) GetAyahAudioFileById(ctx context.Context, id core.UUID) (*AyahAudioFile, error) {
	return s.audioRepo.GetById(ctx, id)
}

func (s *AyahAudioFileService) GetAyahAudioFileByAyahAndReciter(ctx context.Context, ayahId core.UUID, reciterId core.UUID) (*AyahAudioFile, error) {
	return s.audioRepo.GetByAyahAndReciter(ctx, ayahId, reciterId)
}

func (s *AyahAudioFileService) GetAyahAudioFilesBySurahAndReciter(ctx context.Context, surahId core.UUID, reciterId core.UUID, limit, offset int) ([]*AyahAudioFile, error) {
	return s.audioRepo.GetBySurahAndReciter(ctx, surahId, reciterId, limit, offset)
}

func (s *AyahAudioFileService) CreateAyahAudioFile(ctx context.Context, audio *AyahAudioFile) (*AyahAudioFile, error) {
	return s.audioRepo.Create(ctx, audio)
}

func (s *AyahAudioFileService) UpdateAyahAudioFile(ctx context.Context, audio *AyahAudioFile) (*AyahAudioFile, error) {
	return s.audioRepo.Update(ctx, audio)
}

func (s *AyahAudioFileService) DeleteAyahAudioFile(ctx context.Context, id core.UUID) error {
	return s.audioRepo.Delete(ctx, id)
}

func (s *AyahAudioFileService) CountAyahAudioFiles(ctx context.Context) (int64, error) {
	return s.audioRepo.Count(ctx)
}

func (s *AyahAudioFileService) CountAyahAudioFilesBySurahAndReciter(ctx context.Context, surahId core.UUID, reciterId core.UUID) (int64, error) {
	return s.audioRepo.CountBySurahAndReciter(ctx, surahId, reciterId)
}

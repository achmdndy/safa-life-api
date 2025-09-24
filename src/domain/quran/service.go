package quran

import (
	"context"
	"fmt"
)

// quranService implements the QuranService interface
type quranService struct {
	repo QuranRepository
}

// NewQuranService creates a new instance of QuranService
func NewQuranService(repo QuranRepository) QuranService {
	return &quranService{
		repo: repo,
	}
}

// Surah operations
func (s *quranService) GetAllSurahs(ctx context.Context) ([]Surah, error) {
	return s.repo.GetAllSurahs(ctx)
}

func (s *quranService) GetSurahByID(ctx context.Context, id int) (*Surah, error) {
	if id <= 0 || id > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d, must be between 1 and 114", id)
	}
	return s.repo.GetSurahByID(ctx, id)
}

func (s *quranService) GetSurahWithAyahs(ctx context.Context, id int) (*SurahWithAyahs, error) {
	surah, err := s.GetSurahByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	ayahs, err := s.repo.GetAyahsBySurah(ctx, id)
	if err != nil {
		return nil, err
	}
	
	return &SurahWithAyahs{
		Surah: *surah,
		Ayahs: ayahs,
	}, nil
}

// Ayah operations
func (s *quranService) GetAyahsBySurah(ctx context.Context, surahID int) ([]Ayah, error) {
	if surahID <= 0 || surahID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d, must be between 1 and 114", surahID)
	}
	return s.repo.GetAyahsBySurah(ctx, surahID)
}

func (s *quranService) GetAyahByID(ctx context.Context, surahID, ayahID int) (*Ayah, error) {
	if surahID <= 0 || surahID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d, must be between 1 and 114", surahID)
	}
	if ayahID <= 0 {
		return nil, fmt.Errorf("invalid ayah ID: %d, must be greater than 0", ayahID)
	}
	return s.repo.GetAyahByID(ctx, surahID, ayahID)
}

func (s *quranService) GetAyahsByPage(ctx context.Context, pageNumber int) ([]Ayah, error) {
	if pageNumber <= 0 || pageNumber > 604 {
		return nil, fmt.Errorf("invalid page number: %d, must be between 1 and 604", pageNumber)
	}
	return s.repo.GetAyahsByPage(ctx, pageNumber)
}

func (s *quranService) GetAyahsByJuz(ctx context.Context, juzNumber int) ([]Ayah, error) {
	if juzNumber <= 0 || juzNumber > 30 {
		return nil, fmt.Errorf("invalid juz number: %d, must be between 1 and 30", juzNumber)
	}
	return s.repo.GetAyahsByJuz(ctx, juzNumber)
}

// Juz operations
func (s *quranService) GetAllJuz(ctx context.Context) ([]Juz, error) {
	return s.repo.GetAllJuz(ctx)
}

func (s *quranService) GetJuzByID(ctx context.Context, id int) (*Juz, error) {
	if id <= 0 || id > 30 {
		return nil, fmt.Errorf("invalid juz ID: %d, must be between 1 and 30", id)
	}
	return s.repo.GetJuzByID(ctx, id)
}

func (s *quranService) GetJuzWithContent(ctx context.Context, id int) (*JuzWithContent, error) {
	juz, err := s.GetJuzByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	// Get start and end surah names
	startSurah, err := s.repo.GetSurahByID(ctx, juz.StartSurah)
	if err != nil {
		return nil, err
	}
	
	endSurah, err := s.repo.GetSurahByID(ctx, juz.EndSurah)
	if err != nil {
		return nil, err
	}
	
	// Get ayahs in this juz to count total
	ayahs, err := s.repo.GetAyahsByRange(ctx, juz.StartSurah, juz.StartAyah, juz.EndSurah, juz.EndAyah)
	if err != nil {
		return nil, err
	}
	
	return &JuzWithContent{
		Juz:            *juz,
		StartSurahName: startSurah.NameEnglish,
		EndSurahName:   endSurah.NameEnglish,
		TotalAyahs:     len(ayahs),
	}, nil
}

// Search operations
func (s *quranService) SearchAyahs(ctx context.Context, query string) ([]Ayah, error) {
	if query == "" {
		return nil, fmt.Errorf("search query cannot be empty")
	}
	return s.repo.SearchAyahs(ctx, query)
}

func (s *quranService) SearchSurahs(ctx context.Context, query string) ([]Surah, error) {
	if query == "" {
		return nil, fmt.Errorf("search query cannot be empty")
	}
	return s.repo.SearchSurahs(ctx, query)
}

// Write operations for Surah
func (s *quranService) CreateSurah(ctx context.Context, surah *Surah) (*Surah, error) {
	if surah.ID < 1 || surah.ID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", surah.ID)
	}

	if surah.NameEnglish == "" {
		return nil, fmt.Errorf("surah english name cannot be empty")
	}

	if surah.NameArabic == "" {
		return nil, fmt.Errorf("surah arabic name cannot be empty")
	}

	if surah.AyahCount < 1 {
		return nil, fmt.Errorf("ayah count must be greater than 0")
	}

	return s.repo.CreateSurah(ctx, surah)
}

func (s *quranService) UpdateSurah(ctx context.Context, surah *Surah) (*Surah, error) {
	if surah.ID < 1 || surah.ID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", surah.ID)
	}

	if surah.NameEnglish == "" {
		return nil, fmt.Errorf("surah english name cannot be empty")
	}

	if surah.NameArabic == "" {
		return nil, fmt.Errorf("surah arabic name cannot be empty")
	}

	if surah.AyahCount < 1 {
		return nil, fmt.Errorf("ayah count must be greater than 0")
	}

	return s.repo.UpdateSurah(ctx, surah)
}

func (s *quranService) DeleteSurah(ctx context.Context, id int) error {
	if id < 1 || id > 114 {
		return fmt.Errorf("invalid surah ID: %d", id)
	}

	return s.repo.DeleteSurah(ctx, id)
}

// Write operations for Ayah
func (s *quranService) CreateAyah(ctx context.Context, ayah *Ayah) (*Ayah, error) {
	if ayah.SurahID < 1 || ayah.SurahID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", ayah.SurahID)
	}

	if ayah.AyahID < 1 {
		return nil, fmt.Errorf("invalid ayah ID: %d", ayah.AyahID)
	}

	if ayah.Text == "" {
		return nil, fmt.Errorf("ayah text cannot be empty")
	}

	if ayah.PageNumber < 1 || ayah.PageNumber > 604 {
		return nil, fmt.Errorf("invalid page number: %d", ayah.PageNumber)
	}

	if ayah.JuzNumber < 1 || ayah.JuzNumber > 30 {
		return nil, fmt.Errorf("invalid juz number: %d", ayah.JuzNumber)
	}

	return s.repo.CreateAyah(ctx, ayah)
}

func (s *quranService) UpdateAyah(ctx context.Context, ayah *Ayah) (*Ayah, error) {
	if ayah.SurahID < 1 || ayah.SurahID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", ayah.SurahID)
	}

	if ayah.AyahID < 1 {
		return nil, fmt.Errorf("invalid ayah ID: %d", ayah.AyahID)
	}

	if ayah.Text == "" {
		return nil, fmt.Errorf("ayah text cannot be empty")
	}

	if ayah.PageNumber < 1 || ayah.PageNumber > 604 {
		return nil, fmt.Errorf("invalid page number: %d", ayah.PageNumber)
	}

	if ayah.JuzNumber < 1 || ayah.JuzNumber > 30 {
		return nil, fmt.Errorf("invalid juz number: %d", ayah.JuzNumber)
	}

	return s.repo.UpdateAyah(ctx, ayah)
}

func (s *quranService) DeleteAyah(ctx context.Context, surahID, ayahID int) error {
	if surahID < 1 || surahID > 114 {
		return fmt.Errorf("invalid surah ID: %d", surahID)
	}

	if ayahID < 1 {
		return fmt.Errorf("invalid ayah ID: %d", ayahID)
	}

	return s.repo.DeleteAyah(ctx, surahID, ayahID)
}

// Write operations for Juz
func (s *quranService) CreateJuz(ctx context.Context, juz *Juz) (*Juz, error) {
	if juz.ID < 1 || juz.ID > 30 {
		return nil, fmt.Errorf("invalid juz ID: %d", juz.ID)
	}

	if juz.StartSurah < 1 || juz.StartSurah > 114 {
		return nil, fmt.Errorf("invalid start surah: %d", juz.StartSurah)
	}

	if juz.EndSurah < 1 || juz.EndSurah > 114 {
		return nil, fmt.Errorf("invalid end surah: %d", juz.EndSurah)
	}

	return s.repo.CreateJuz(ctx, juz)
}

func (s *quranService) UpdateJuz(ctx context.Context, juz *Juz) (*Juz, error) {
	if juz.ID < 1 || juz.ID > 30 {
		return nil, fmt.Errorf("invalid juz ID: %d", juz.ID)
	}

	if juz.StartSurah < 1 || juz.StartSurah > 114 {
		return nil, fmt.Errorf("invalid start surah: %d", juz.StartSurah)
	}

	if juz.EndSurah < 1 || juz.EndSurah > 114 {
		return nil, fmt.Errorf("invalid end surah: %d", juz.EndSurah)
	}

	return s.repo.UpdateJuz(ctx, juz)
}

func (s *quranService) DeleteJuz(ctx context.Context, id int) error {
	if id < 1 || id > 30 {
		return fmt.Errorf("invalid juz ID: %d", id)
	}

	return s.repo.DeleteJuz(ctx, id)
}
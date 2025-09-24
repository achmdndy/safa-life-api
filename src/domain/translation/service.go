package translation

import (
	"context"
	"fmt"
)

type translationService struct {
	repo FullTranslationRepository
}

// NewTranslationService creates a new instance of TranslationService.
func NewTranslationService(repo FullTranslationRepository) TranslationService {
	return &translationService{repo: repo}
}

// Translation (edition) operations
func (s *translationService) GetAllTranslations(ctx context.Context) ([]Translation, error) {
	return s.repo.GetAll(ctx)
}

func (s *translationService) GetTranslationByID(ctx context.Context, id string) (*Translation, error) {
	if id == "" {
		return nil, fmt.Errorf("translation ID cannot be empty")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *translationService) CreateTranslation(ctx context.Context, translation *Translation) (*Translation, error) {
	if translation.ID == "" || translation.Name == "" || translation.Author == "" || translation.Language == "" {
		return nil, fmt.Errorf("translation fields cannot be empty")
	}
	return s.repo.CreateTranslation(ctx, translation)
}

func (s *translationService) UpdateTranslation(ctx context.Context, translation *Translation) (*Translation, error) {
	if translation.ID == "" {
		return nil, fmt.Errorf("translation ID cannot be empty")
	}
	return s.repo.UpdateTranslation(ctx, translation)
}

func (s *translationService) DeleteTranslation(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("translation ID cannot be empty")
	}
	return s.repo.DeleteTranslation(ctx, id)
}

// AyahTranslation operations
func (s *translationService) GetAyahTranslation(ctx context.Context, translationID string, surahID, ayahID int) (*AyahTranslation, error) {
	if translationID == "" {
		return nil, fmt.Errorf("translation ID cannot be empty")
	}
	if surahID <= 0 || surahID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", surahID)
	}
	if ayahID <= 0 {
		return nil, fmt.Errorf("invalid ayah ID: %d", ayahID)
	}
	return s.repo.GetAyahTranslation(ctx, translationID, surahID, ayahID)
}

func (s *translationService) GetTranslationsForAyah(ctx context.Context, surahID, ayahID int) ([]AyahTranslation, error) {
	if surahID <= 0 || surahID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", surahID)
	}
	if ayahID <= 0 {
		return nil, fmt.Errorf("invalid ayah ID: %d", ayahID)
	}
	return s.repo.GetByAyah(ctx, surahID, ayahID)
}

func (s *translationService) GetTranslationsForSurah(ctx context.Context, translationID string, surahID int) ([]AyahTranslation, error) {
	if translationID == "" {
		return nil, fmt.Errorf("translation ID cannot be empty")
	}
	if surahID <= 0 || surahID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", surahID)
	}
	return s.repo.GetBySurah(ctx, translationID, surahID)
}

func (s *translationService) CreateAyahTranslation(ctx context.Context, ayahTranslation *AyahTranslation) (*AyahTranslation, error) {
	if ayahTranslation.TranslationID == "" || ayahTranslation.Text == "" {
		return nil, fmt.Errorf("translation ID and text cannot be empty")
	}
	if ayahTranslation.SurahID <= 0 || ayahTranslation.SurahID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", ayahTranslation.SurahID)
	}
	if ayahTranslation.AyahID <= 0 {
		return nil, fmt.Errorf("invalid ayah ID: %d", ayahTranslation.AyahID)
	}
	return s.repo.CreateAyahTranslation(ctx, ayahTranslation)
}

func (s *translationService) UpdateAyahTranslation(ctx context.Context, ayahTranslation *AyahTranslation) (*AyahTranslation, error) {
	if ayahTranslation.TranslationID == "" {
		return nil, fmt.Errorf("translation ID cannot be empty")
	}
	if ayahTranslation.SurahID <= 0 || ayahTranslation.SurahID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", ayahTranslation.SurahID)
	}
	if ayahTranslation.AyahID <= 0 {
		return nil, fmt.Errorf("invalid ayah ID: %d", ayahTranslation.AyahID)
	}
	return s.repo.UpdateAyahTranslation(ctx, ayahTranslation)
}

func (s *translationService) DeleteAyahTranslation(ctx context.Context, translationID string, surahID, ayahID int) error {
	if translationID == "" {
		return fmt.Errorf("translation ID cannot be empty")
	}
	if surahID <= 0 || surahID > 114 {
		return fmt.Errorf("invalid surah ID: %d", surahID)
	}
	if ayahID <= 0 {
		return fmt.Errorf("invalid ayah ID: %d", ayahID)
	}
	return s.repo.DeleteAyahTranslation(ctx, translationID, surahID, ayahID)
}
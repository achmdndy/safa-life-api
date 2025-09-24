package tafsir

import (
	"context"
	"fmt"
)

type tafsirService struct {
	repo FullTafsirRepository
}

// NewTafsirService creates a new instance of TafsirService.
func NewTafsirService(repo FullTafsirRepository) TafsirService {
	return &tafsirService{repo: repo}
}

func (s *tafsirService) GetAllTafsirs(ctx context.Context) ([]Tafsir, error) {
	return s.repo.GetAll(ctx)
}

func (s *tafsirService) GetTafsirByID(ctx context.Context, id string) (*Tafsir, error) {
	if id == "" {
		return nil, fmt.Errorf("tafsir ID cannot be empty")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *tafsirService) CreateTafsir(ctx context.Context, tafsir *Tafsir) (*Tafsir, error) {
	if tafsir.ID == "" || tafsir.Name == "" {
		return nil, fmt.Errorf("tafsir ID and name cannot be empty")
	}
	return s.repo.CreateTafsir(ctx, tafsir)
}

func (s *tafsirService) UpdateTafsir(ctx context.Context, tafsir *Tafsir) (*Tafsir, error) {
	if tafsir.ID == "" {
		return nil, fmt.Errorf("tafsir ID cannot be empty")
	}
	return s.repo.UpdateTafsir(ctx, tafsir)
}

func (s *tafsirService) DeleteTafsir(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("tafsir ID cannot be empty")
	}
	return s.repo.DeleteTafsir(ctx, id)
}

func (s *tafsirService) GetAyahTafsir(ctx context.Context, tafsirID string, surahID, ayahID int) (*AyahTafsir, error) {
	if tafsirID == "" {
		return nil, fmt.Errorf("tafsir ID cannot be empty")
	}
	return s.repo.GetAyahTafsir(ctx, tafsirID, surahID, ayahID)
}

func (s *tafsirService) GetTafsirsForAyah(ctx context.Context, surahID, ayahID int) ([]AyahTafsir, error) {
	return s.repo.GetByAyah(ctx, surahID, ayahID)
}

func (s *tafsirService) GetTafsirsForSurah(ctx context.Context, tafsirID string, surahID int) ([]AyahTafsir, error) {
	return s.repo.GetBySurah(ctx, tafsirID, surahID)
}

func (s *tafsirService) CreateAyahTafsir(ctx context.Context, ayahTafsir *AyahTafsir) (*AyahTafsir, error) {
	return s.repo.CreateAyahTafsir(ctx, ayahTafsir)
}

func (s *tafsirService) UpdateAyahTafsir(ctx context.Context, ayahTafsir *AyahTafsir) (*AyahTafsir, error) {
	return s.repo.UpdateAyahTafsir(ctx, ayahTafsir)
}

func (s *tafsirService) DeleteAyahTafsir(ctx context.Context, tafsirID string, surahID, ayahID int) error {
	return s.repo.DeleteAyahTafsir(ctx, tafsirID, surahID, ayahID)
}

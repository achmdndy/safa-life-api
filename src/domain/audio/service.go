package audio

import (
	"context"
	"fmt"
)

type audioService struct {
	repo AyahAudioRepository
}

// NewAudioService creates a new instance of AudioService.
func NewAudioService(repo AyahAudioRepository) AudioService {
	return &audioService{repo: repo}
}

func (s *audioService) GetAyahAudio(ctx context.Context, reciterID string, surahID, ayahID int) (*AyahAudio, error) {
	if reciterID == "" {
		return nil, fmt.Errorf("reciter ID cannot be empty")
	}
	return s.repo.Get(ctx, reciterID, surahID, ayahID)
}

func (s *audioService) GetAudioFilesForSurah(ctx context.Context, reciterID string, surahID int) ([]AyahAudio, error) {
	if reciterID == "" {
		return nil, fmt.Errorf("reciter ID cannot be empty")
	}
	return s.repo.GetBySurah(ctx, reciterID, surahID)
}

func (s *audioService) CreateAyahAudio(ctx context.Context, audio *AyahAudio) (*AyahAudio, error) {
	if audio.ReciterID == "" || audio.FilePath == "" {
		return nil, fmt.Errorf("reciter ID and file path cannot be empty")
	}
	return s.repo.Create(ctx, audio)
}

func (s *audioService) UpdateAyahAudio(ctx context.Context, audio *AyahAudio) (*AyahAudio, error) {
	if audio.ReciterID == "" {
		return nil, fmt.Errorf("reciter ID cannot be empty")
	}
	return s.repo.Update(ctx, audio)
}

func (s *audioService) DeleteAyahAudio(ctx context.Context, reciterID string, surahID, ayahID int) error {
	if reciterID == "" {
		return fmt.Errorf("reciter ID cannot be empty")
	}
	return s.repo.Delete(ctx, reciterID, surahID, ayahID)
}
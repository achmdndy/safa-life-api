package command

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/audio"
)

type CreateAyahAudioCommand struct {
	ReciterID string
	SurahID   int
	AyahID    int
	FilePath  string
	Duration  float64
}

func (h *CommandHandler) CreateAyahAudio(ctx context.Context, cmd CreateAyahAudioCommand) (*audio.AyahAudio, error) {
	domain := &audio.AyahAudio{
		ReciterID: cmd.ReciterID,
		SurahID:   cmd.SurahID,
		AyahID:    cmd.AyahID,
		FilePath:  cmd.FilePath,
		Duration:  cmd.Duration,
	}

	var result *audio.AyahAudio
	err := h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		result, err = h.audioService.CreateAyahAudio(ctx, domain)
		return err
	})

	return result, err
}

package command

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/audio"
)

type UpdateAyahAudioCommand struct {
	ReciterID string
	SurahID   int
	AyahID    int
	FilePath  *string
	Duration  *float64
}

func (h *CommandHandler) UpdateAyahAudio(ctx context.Context, cmd UpdateAyahAudioCommand) (*audio.AyahAudio, error) {
	existing, err := h.audioService.GetAyahAudio(ctx, cmd.ReciterID, cmd.SurahID, cmd.AyahID)
	if err != nil {
		return nil, err
	}

	if cmd.FilePath != nil {
		existing.FilePath = *cmd.FilePath
	}
	if cmd.Duration != nil {
		existing.Duration = *cmd.Duration
	}

	var result *audio.AyahAudio
	err = h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		result, err = h.audioService.UpdateAyahAudio(ctx, existing)
		return err
	})

	return result, err
}

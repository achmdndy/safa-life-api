package command

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type UpdateAyahAudioFileCommand struct {
	ID        string
	FilePath  string
	Duration  float64
	ByteSize  float64
	UpdatedBy string
}

func (h *CommandHandler) UpdateAyahAudioFile(ctx context.Context, command UpdateAyahAudioFileCommand) (*dto.AyahAudioFileResponse, error) {
	id, err := h.uuidGenerator.Parse(command.ID)
	if err != nil {
		return nil, err
	}

	existing, err := h.audioService.GetAyahAudioFileById(ctx, id)
	if err != nil {
		return nil, err
	}

	existing.Update(command.FilePath, command.Duration, command.ByteSize, command.UpdatedBy)

	var updated *quran.AyahAudioFile
	err = h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		var txErr error
		updated, txErr = h.audioService.UpdateAyahAudioFile(ctx, existing)
		return txErr
	})
	if err != nil {
		return nil, err
	}

	return dto.ToAyahAudioFileResponse(updated), nil
}

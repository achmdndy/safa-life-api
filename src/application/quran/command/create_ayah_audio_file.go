package command

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type CreateAyahAudioFileCommand struct {
	ReciterID string
	SurahID   string
	AyahID    string
	FilePath  string
	Duration  float64
	ByteSize  float64
	CreatedBy string
}

func (h *CommandHandler) CreateAyahAudioFile(ctx context.Context, command CreateAyahAudioFileCommand) (*dto.AyahAudioFileResponse, error) {
	id := h.uuidGenerator.New()

	reciterID, err := h.uuidGenerator.Parse(command.ReciterID)
	if err != nil {
		return nil, err
	}
	surahID, err := h.uuidGenerator.Parse(command.SurahID)
	if err != nil {
		return nil, err
	}
	ayahID, err := h.uuidGenerator.Parse(command.AyahID)
	if err != nil {
		return nil, err
	}

	audio := quran.NewAyahAudioFile(
		id,
		reciterID,
		surahID,
		ayahID,
		command.FilePath,
		command.Duration,
		command.ByteSize,
		command.CreatedBy,
	)

	var created *quran.AyahAudioFile
	err = h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		var txErr error
		created, txErr = h.audioService.CreateAyahAudioFile(ctx, audio)
		return txErr
	})
	if err != nil {
		return nil, err
	}

	return dto.ToAyahAudioFileResponse(created), nil
}

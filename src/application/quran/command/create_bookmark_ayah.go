package command

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type CreateBookmarkAyahCommand struct {
	UserID    string
	AyahID    string
	CreatedBy string
}

func (h *CommandHandler) CreateBookmarkAyah(ctx context.Context, command CreateBookmarkAyahCommand) (*dto.BookmarkAyahResponse, error) {
	id := h.uuidGenerator.New()
	ayahID, err := h.uuidGenerator.Parse(command.AyahID)
	if err != nil {
		return nil, err
	}

	bookmark := quran.NewBookmarkAyah(
		id,
		command.UserID,
		ayahID,
		command.CreatedBy,
	)

	var created *quran.BookmarkAyah
	err = h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		var txErr error
		created, txErr = h.bookmarkService.CreateBookmarkAyah(ctx, bookmark)
		return txErr
	})
	if err != nil {
		return nil, err
	}

	return dto.ToBookmarkAyahResponse(created), nil
}

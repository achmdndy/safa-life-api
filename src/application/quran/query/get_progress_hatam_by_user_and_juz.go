package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type GetProgressHatamByUserAndJuzQuery struct {
	UserID  string
	JuzID   string
	Include string
}

func (h *QueryHandler) GetProgressHatamByUserAndJuz(ctx context.Context, query GetProgressHatamByUserAndJuzQuery) (interface{}, error) {
	juzID, err := h.uuidGenerator.Parse(query.JuzID)
	if err != nil {
		return nil, err
	}

	var p *quran.ProgressHatam
	p, err = h.progressService.GetProgressHatamByUserAndJuz(ctx, query.UserID, juzID)
	if err != nil {
		return nil, err
	}

	if query.Include == "relations" && p != nil {
		var pWith *quran.ProgressHatamWithRelations
		pWith, err = h.progressService.GetProgressHatamByIdWithRelations(ctx, p.ID)
		if err != nil {
			return nil, err
		}
		return dto.ToProgressHatamWithRelationsResponse(pWith), nil
	}

	return dto.ToProgressHatamResponse(p), nil
}

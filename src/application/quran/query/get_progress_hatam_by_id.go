package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type GetProgressHatamByIdQuery struct {
	ID      string
	Include string
}

func (h *QueryHandler) GetProgressHatamById(ctx context.Context, query GetProgressHatamByIdQuery) (interface{}, error) {
	id, err := h.uuidGenerator.Parse(query.ID)
	if err != nil {
		return nil, err
	}

	if query.Include == "relations" {
		var pWith *quran.ProgressHatamWithRelations
		pWith, err = h.progressService.GetProgressHatamByIdWithRelations(ctx, id)
		if err != nil {
			return nil, err
		}
		return dto.ToProgressHatamWithRelationsResponse(pWith), nil
	}

	var p *quran.ProgressHatam
	p, err = h.progressService.GetProgressHatamById(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ToProgressHatamResponse(p), nil
}

package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetAllJuzWithUserProgressQuery struct {
	Limit   int
	Offset  int
	Include string
	UserID  string
}

// GetAllJuzWithUserProgress returns all Juz along with the authenticated user's ProgressHatam per Juz.
func (h *QueryHandler) GetAllJuzWithUserProgress(ctx context.Context, query GetAllJuzWithUserProgressQuery) (interface{}, error) {
	// Relations variant
	if query.Include == "relations" {
		juzList, err := h.juzService.GetAllJuzWithRelations(ctx, query.Limit, query.Offset)
		if err != nil {
			return nil, err
		}

		// Fetch all progress for user with relations and map by JuzID
		phList, err := h.progressService.GetProgressHatamByUserWithRelations(ctx, query.UserID, 0, 0)
		if err != nil {
			return nil, err
		}
		phByJuzID := make(map[string]*dto.ProgressHatamWithRelationsResponse, len(phList))
		for _, p := range phList {
			if p == nil {
				continue
			}
			phByJuzID[p.JuzID.String()] = dto.ToProgressHatamWithRelationsResponse(p)
		}

		// Build response slice by pairing Juz with user's progress for that Juz
		out := make([]*dto.JuzWithProgressWithRelationsResponse, len(juzList))
		for i, j := range juzList {
			var phResp *dto.ProgressHatamWithRelationsResponse
			if j != nil {
				phResp = phByJuzID[j.ID.String()]
			}
			out[i] = &dto.JuzWithProgressWithRelationsResponse{
				Juz:           dto.ToJuzWithRelationsResponse(j),
				ProgressHatam: phResp,
			}
		}

		count, err := h.juzService.CountJuz(ctx)
		if err != nil {
			return nil, err
		}

		return &dto.JuzWithProgressWithRelationsListResponse{
			Data: out,
			Pagination: &dto.PaginationResponse{
				Limit:  query.Limit,
				Offset: query.Offset,
				Total:  count,
			},
		}, nil
	}

	// Basic variant without relations
	juzList, err := h.juzService.GetAllJuz(ctx, query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}

	phList, err := h.progressService.GetProgressHatamByUser(ctx, query.UserID, 0, 0)
	if err != nil {
		return nil, err
	}
	phByJuzID := make(map[string]*dto.ProgressHatamResponse, len(phList))
	for _, p := range phList {
		if p == nil {
			continue
		}
		phByJuzID[p.JuzID.String()] = dto.ToProgressHatamResponse(p)
	}

	out := make([]*dto.JuzWithProgressBasicResponse, len(juzList))
	for i, j := range juzList {
		var phResp *dto.ProgressHatamResponse
		if j != nil {
			phResp = phByJuzID[j.ID.String()]
		}
		out[i] = &dto.JuzWithProgressBasicResponse{
			Juz:           dto.ToJuzResponse(j),
			ProgressHatam: phResp,
		}
	}

	count, err := h.juzService.CountJuz(ctx)
	if err != nil {
		return nil, err
	}

	return &dto.JuzWithProgressBasicListResponse{
		Data: out,
		Pagination: &dto.PaginationResponse{
			Limit:  query.Limit,
			Offset: query.Offset,
			Total:  count,
		},
	}, nil
}

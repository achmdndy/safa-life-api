package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetTranslationEditionsQuery struct {
	Language string
	Limit    int
	Offset   int
}

func (h *QueryHandler) GetTranslationEditions(ctx context.Context, query GetTranslationEditionsQuery) (*dto.TranslationEditionListResponse, error) {
	var (
		editions []*dto.TranslationEditionResponse
		total    int64
	)

	if query.Language != "" {
		items, ierr := h.translationEditionService.GetTranslationEditionsByLanguage(ctx, query.Language, query.Limit, query.Offset)
		if ierr != nil {
			return nil, ierr
		}
		count, cerr := h.translationEditionService.CountTranslationEditionsByLanguage(ctx, query.Language)
		if cerr != nil {
			return nil, cerr
		}
		editions = dto.ToTranslationEditionResponseSlice(items)
		total = count
	} else {
		items, ierr := h.translationEditionService.GetAllTranslationEditions(ctx, query.Limit, query.Offset)
		if ierr != nil {
			return nil, ierr
		}
		count, cerr := h.translationEditionService.CountTranslationEditions(ctx)
		if cerr != nil {
			return nil, cerr
		}
		editions = dto.ToTranslationEditionResponseSlice(items)
		total = count
	}

	return &dto.TranslationEditionListResponse{
		Data: editions,
		Pagination: &dto.PaginationResponse{
			Limit:  query.Limit,
			Offset: query.Offset,
			Total:  total,
		},
	}, nil
}

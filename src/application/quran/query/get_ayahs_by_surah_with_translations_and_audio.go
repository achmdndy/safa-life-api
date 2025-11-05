package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetAyahsBySurahWithTranslationsQuery struct {
	SurahID   string
	EditionID string
	Limit     int
	Offset    int
	ReciterID string
}

// GetAyahsBySurahWithTranslations returns surah details with a list of ayahs,
// each enriched with its translation for the specified edition.
func (h *QueryHandler) GetAyahsBySurahWithTranslations(ctx context.Context, query GetAyahsBySurahWithTranslationsQuery) (*dto.SurahAyahsWithTranslationsResponse, error) {
	surahID, err := h.uuidGenerator.Parse(query.SurahID)
	if err != nil {
		return nil, err
	}

	editionID, err := h.uuidGenerator.Parse(query.EditionID)
	if err != nil {
		return nil, err
	}

	surah, err := h.surahService.GetSurahById(ctx, surahID)
	if err != nil {
		return nil, err
	}

	ayahs, err := h.ayahService.GetAyahsBySurahId(ctx, surahID, query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}

	count, err := h.ayahService.CountAyahsBySurahId(ctx, surahID)
	if err != nil {
		return nil, err
	}

	translations, err := h.ayahTranslationService.GetAyahTranslationsBySurahAndEdition(ctx, surahID, editionID, 0, 0)
	if err != nil {
		return nil, err
	}

	trByAyah := make(map[string]*dto.AyahTranslationResponse, len(translations))
	for _, tr := range translations {
		trByAyah[tr.AyahID.String()] = dto.ToAyahTranslationResponse(tr)
	}

	// Optionally load reciter and per-ayah audio if ReciterID provided
	var reciterResp *dto.ReciterResponse
	audioByAyah := make(map[string]*dto.AyahAudioFileResponse)
	if query.ReciterID != "" {
		reciterID, err := h.uuidGenerator.Parse(query.ReciterID)
		if err != nil {
			return nil, err
		}
		reciter, err := h.reciterService.GetReciterById(ctx, reciterID)
		if err != nil {
			return nil, err
		}
		reciterResp = dto.ToReciterResponse(reciter)

		audioFiles, err := h.audioService.GetAyahAudioFilesBySurahAndReciter(ctx, surahID, reciterID, 0, 0)
		if err != nil {
			return nil, err
		}
		for _, af := range audioFiles {
			audioByAyah[af.AyahID.String()] = dto.ToAyahAudioFileResponse(af)
		}
	}

	outAyahs := make([]*dto.AyahWithTranslationResponse, len(ayahs))
	for i, a := range ayahs {
		var tr *dto.AyahTranslationResponse
		if t, ok := trByAyah[a.ID.String()]; ok {
			tr = t
		}
		var audio *dto.AyahAudioFileResponse
		if af, ok := audioByAyah[a.ID.String()]; ok {
			audio = af
		}
		outAyahs[i] = &dto.AyahWithTranslationResponse{
			ID:           a.ID.String(),
			SurahID:      a.SurahID.String(),
			Text:         a.Text,
			PageNumber:   a.PageNumber,
			JuzNumber:    a.JuzNumber,
			HizbNumber:   a.HizbNumber,
			ManzilNumber: a.ManzilNumber,
			CreatedBy:    a.CreatedBy,
			UpdatedBy:    a.UpdatedBy,
			CreatedAt:    a.CreatedAt,
			UpdatedAt:    a.UpdatedAt,
			Translation:  tr,
			Audio:        audio,
		}
	}

	return &dto.SurahAyahsWithTranslationsResponse{
		Surah: dto.ToSurahResponse(surah),
		Ayahs: outAyahs,
		Pagination: &dto.PaginationResponse{
			Limit:  query.Limit,
			Offset: query.Offset,
			Total:  count,
		},
		Reciter: reciterResp,
	}, nil
}

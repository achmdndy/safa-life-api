package query

import (
	"github.com/achmdndy/safa-life-api/src/domain/quran"
)

type QueryHandler struct {
	quranService quran.QuranService
}

func NewQueryHandler(quranService quran.QuranService) *QueryHandler {
	return &QueryHandler{
		quranService: quranService,
	}
}

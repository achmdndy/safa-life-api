package query

import (
	"github.com/safalife/core-api/src/domain/core"
	"github.com/safalife/core-api/src/domain/quran"
)

type QueryHandler struct {
	surahService  quran.SurahServiceInterface
	ayahService   quran.AyahServiceInterface
	juzService    quran.JuzServiceInterface
	uuidGenerator core.UUIDGenerator
}

func NewQueryHandler(
	surahService quran.SurahServiceInterface,
	ayahService quran.AyahServiceInterface,
	juzService quran.JuzServiceInterface,
	uuidGenerator core.UUIDGenerator,
) *QueryHandler {
	return &QueryHandler{
		surahService:  surahService,
		ayahService:   ayahService,
		juzService:    juzService,
		uuidGenerator: uuidGenerator,
	}
}

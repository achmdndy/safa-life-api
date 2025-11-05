package query

import (
	"github.com/safalife/core-api/src/domain/core"
	"github.com/safalife/core-api/src/domain/quran"
)

type QueryHandler struct {
	surahService              quran.SurahServiceInterface
	ayahService               quran.AyahServiceInterface
	juzService                quran.JuzServiceInterface
	reciterService            quran.ReciterServiceInterface
	audioService              quran.AyahAudioFileServiceInterface
	translationEditionService quran.TranslationEditionServiceInterface
	ayahTranslationService    quran.AyahTranslationServiceInterface
	uuidGenerator             core.UUIDGenerator
}

func NewQueryHandler(
	surahService quran.SurahServiceInterface,
	ayahService quran.AyahServiceInterface,
	juzService quran.JuzServiceInterface,
	reciterService quran.ReciterServiceInterface,
	audioService quran.AyahAudioFileServiceInterface,
	translationEditionService quran.TranslationEditionServiceInterface,
	ayahTranslationService quran.AyahTranslationServiceInterface,
	uuidGenerator core.UUIDGenerator,
) *QueryHandler {
	return &QueryHandler{
		surahService:              surahService,
		ayahService:               ayahService,
		juzService:                juzService,
		reciterService:            reciterService,
		audioService:              audioService,
		translationEditionService: translationEditionService,
		ayahTranslationService:    ayahTranslationService,
		uuidGenerator:             uuidGenerator,
	}
}

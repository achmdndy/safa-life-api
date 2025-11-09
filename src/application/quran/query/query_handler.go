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
	bookmarkService           quran.BookmarkAyahServiceInterface
	lastReadService           quran.LastReadServiceInterface
	progressService           quran.ProgressHatamServiceInterface
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
	bookmarkService quran.BookmarkAyahServiceInterface,
	lastReadService quran.LastReadServiceInterface,
	progressService quran.ProgressHatamServiceInterface,
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
		bookmarkService:           bookmarkService,
		lastReadService:           lastReadService,
		progressService:           progressService,
		uuidGenerator:             uuidGenerator,
	}
}

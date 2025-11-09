package command

import (
	"github.com/safalife/core-api/src/domain/core"
	"github.com/safalife/core-api/src/domain/quran"
)

type CommandHandler struct {
	surahService    quran.SurahServiceInterface
	ayahService     quran.AyahServiceInterface
	juzService      quran.JuzServiceInterface
	reciterService  quran.ReciterServiceInterface
	audioService    quran.AyahAudioFileServiceInterface
	bookmarkService quran.BookmarkAyahServiceInterface
	lastReadService quran.LastReadServiceInterface
	progressService quran.ProgressHatamServiceInterface
	uuidGenerator   core.UUIDGenerator
	transactionMgr  core.ContextTransactionManager
}

func NewCommandHandler(
	surahService quran.SurahServiceInterface,
	ayahService quran.AyahServiceInterface,
	juzService quran.JuzServiceInterface,
	reciterService quran.ReciterServiceInterface,
	audioService quran.AyahAudioFileServiceInterface,
	bookmarkService quran.BookmarkAyahServiceInterface,
	lastReadService quran.LastReadServiceInterface,
	progressService quran.ProgressHatamServiceInterface,
	uuidGenerator core.UUIDGenerator,
	transactionMgr core.ContextTransactionManager,
) *CommandHandler {
	return &CommandHandler{
		surahService:    surahService,
		ayahService:     ayahService,
		juzService:      juzService,
		reciterService:  reciterService,
		audioService:    audioService,
		bookmarkService: bookmarkService,
		lastReadService: lastReadService,
		progressService: progressService,
		uuidGenerator:   uuidGenerator,
		transactionMgr:  transactionMgr,
	}
}

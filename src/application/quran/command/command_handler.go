package command

import (
	"github.com/safalife/core-api/src/domain/core"
	"github.com/safalife/core-api/src/domain/quran"
)

type CommandHandler struct {
	surahService   quran.SurahServiceInterface
	ayahService    quran.AyahServiceInterface
	juzService     quran.JuzServiceInterface
	uuidGenerator  core.UUIDGenerator
	transactionMgr core.ContextTransactionManager
}

func NewCommandHandler(
	surahService quran.SurahServiceInterface,
	ayahService quran.AyahServiceInterface,
	juzService quran.JuzServiceInterface,
	uuidGenerator core.UUIDGenerator,
	transactionMgr core.ContextTransactionManager,
) *CommandHandler {
	return &CommandHandler{
		surahService:   surahService,
		ayahService:    ayahService,
		juzService:     juzService,
		uuidGenerator:  uuidGenerator,
		transactionMgr: transactionMgr,
	}
}
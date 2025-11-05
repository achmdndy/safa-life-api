package container

import (
	"github.com/safalife/core-api/src/domain/core"
	"github.com/safalife/core-api/src/domain/monitoring"
	"github.com/safalife/core-api/src/domain/quran"
	domStorage "github.com/safalife/core-api/src/domain/storage"
)

// DomainContainer holds all domain services
type DomainContainer struct {
	MonitoringService  monitoring.MonitoringService
	MonitoringHandlers monitoring.MonitoringHandlers

	// Core services
	TransactionManager core.ContextTransactionManager
	IDGenerator        core.UUIDGenerator

	// Storage
	StorageService domStorage.StorageServiceInterface

	// Quran services
	SurahService              quran.SurahServiceInterface
	AyahService               quran.AyahServiceInterface
	JuzService                quran.JuzServiceInterface
	TranslationEditionService quran.TranslationEditionServiceInterface
	AyahTranslationService    quran.AyahTranslationServiceInterface
	ReciterService            quran.ReciterServiceInterface
	AyahAudioFileService      quran.AyahAudioFileServiceInterface
}

// NewDomainContainer creates a new domain container with dependency injection
func NewDomainContainer(
	monitoringService interface{},
	transactionManager core.ContextTransactionManager,
	idGenerator core.UUIDGenerator,
	storageService domStorage.StorageServiceInterface,
	surahRepo quran.SurahRepositoryInterface,
	ayahRepo quran.AyahRepositoryInterface,
	juzRepo quran.JuzRepositoryInterface,
	translationEditionRepo quran.TranslationEditionRepositoryInterface,
	ayahTranslationRepo quran.AyahTranslationRepositoryInterface,
	reciterRepo quran.ReciterRepositoryInterface,
	ayahAudioFileRepo quran.AyahAudioFileRepositoryInterface,
) *DomainContainer {
	// Create Quran services
	surahService := quran.NewSurahService(surahRepo)
	ayahService := quran.NewAyahService(ayahRepo)
	juzService := quran.NewJuzService(juzRepo)
	translationEditionService := quran.NewTranslationEditionService(translationEditionRepo)
	ayahTranslationService := quran.NewAyahTranslationService(ayahTranslationRepo)
	reciterService := quran.NewReciterService(reciterRepo)
	ayahAudioFileService := quran.NewAyahAudioFileService(ayahAudioFileRepo)

	return &DomainContainer{
		MonitoringService:  monitoringService.(monitoring.MonitoringService),
		MonitoringHandlers: monitoringService.(monitoring.MonitoringHandlers),

		// Core services
		TransactionManager: transactionManager,
		IDGenerator:        idGenerator,

		// Storage
		StorageService: storageService,

		// Quran services
		SurahService:              surahService,
		AyahService:               ayahService,
		JuzService:                juzService,
		TranslationEditionService: translationEditionService,
		AyahTranslationService:    ayahTranslationService,
		ReciterService:            reciterService,
		AyahAudioFileService:      ayahAudioFileService,
	}
}

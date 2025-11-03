package container

import (
	"github.com/safalife/core-api/src/domain/core"
	"github.com/safalife/core-api/src/domain/monitoring"
	"github.com/safalife/core-api/src/domain/quran"
)

// DomainContainer holds all domain services
type DomainContainer struct {
	MonitoringService  monitoring.MonitoringService
	MonitoringHandlers monitoring.MonitoringHandlers

	// Core services
	TransactionManager core.ContextTransactionManager
	IDGenerator        core.UUIDGenerator

	// Quran services
	SurahService quran.SurahServiceInterface
	AyahService  quran.AyahServiceInterface
	JuzService   quran.JuzServiceInterface
}

// NewDomainContainer creates a new domain container with dependency injection
func NewDomainContainer(
	monitoringService interface{},
	transactionManager core.ContextTransactionManager,
	idGenerator core.UUIDGenerator,
	surahRepo quran.SurahRepositoryInterface,
	ayahRepo quran.AyahRepositoryInterface,
	juzRepo quran.JuzRepositoryInterface,
) *DomainContainer {
	// Create Quran services
	surahService := quran.NewSurahService(surahRepo)
	ayahService := quran.NewAyahService(ayahRepo)
	juzService := quran.NewJuzService(juzRepo)

	return &DomainContainer{
		MonitoringService:  monitoringService.(monitoring.MonitoringService),
		MonitoringHandlers: monitoringService.(monitoring.MonitoringHandlers),

		// Core services
		TransactionManager: transactionManager,
		IDGenerator:        idGenerator,

		// Quran services
		SurahService: surahService,
		AyahService:  ayahService,
		JuzService:   juzService,
	}
}

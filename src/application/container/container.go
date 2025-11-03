package container

import (
	quranCommand "github.com/safalife/core-api/src/application/quran/command"
	quranQuery "github.com/safalife/core-api/src/application/quran/query"
	domainContainer "github.com/safalife/core-api/src/domain/container"
	"github.com/safalife/core-api/src/domain/monitoring"
)

// ApplicationContainer holds application dependencies
type ApplicationContainer struct {
	// Quran handlers
	QuranCommandHandler *quranCommand.CommandHandler
	QuranQueryHandler   *quranQuery.QueryHandler

	Monitoring interface{} // Single field for both MonitoringService and MonitoringHandlers
}

// NewApplicationContainer creates a new application container with dependency injection
func NewApplicationContainer(domainContainer *domainContainer.DomainContainer) *ApplicationContainer {
	// Create Quran command and query handlers
	quranCommandHandler := quranCommand.NewCommandHandler(
		domainContainer.SurahService,
		domainContainer.AyahService,
		domainContainer.JuzService,
		domainContainer.IDGenerator,
		domainContainer.TransactionManager,
	)

	quranQueryHandler := quranQuery.NewQueryHandler(
		domainContainer.SurahService,
		domainContainer.AyahService,
		domainContainer.JuzService,
		domainContainer.IDGenerator,
	)

	return &ApplicationContainer{
		// Quran handlers
		QuranCommandHandler: quranCommandHandler,
		QuranQueryHandler:   quranQueryHandler,

		Monitoring: domainContainer.MonitoringService,
	}
}

// GetMonitoringService returns the monitoring service interface
func (c *ApplicationContainer) GetMonitoringService() monitoring.MonitoringService {
	return c.Monitoring.(monitoring.MonitoringService)
}

// GetMonitoringHandlers returns the monitoring handlers interface
func (c *ApplicationContainer) GetMonitoringHandlers() monitoring.MonitoringHandlers {
	return c.Monitoring.(monitoring.MonitoringHandlers)
}

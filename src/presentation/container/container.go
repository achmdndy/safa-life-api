package container

import (
	appContainer "github.com/safalife/core-api/src/application/container"
	"github.com/safalife/core-api/src/domain/monitoring"
	"github.com/safalife/core-api/src/presentation/handlers/quran"
)

// PresentationContainer holds presentation dependencies
type PresentationContainer struct {
	QuranHandler *quran.Handler
	Monitoring   interface{} // Single field for both MonitoringService and MonitoringHandlers
}

// NewPresentationContainer creates a new presentation container
func NewPresentationContainer(appContainer *appContainer.ApplicationContainer) *PresentationContainer {
	quranHandler := quran.NewHandler(
		appContainer.QuranCommandHandler,
		appContainer.QuranQueryHandler,
	)

	return &PresentationContainer{
		QuranHandler: quranHandler,
		Monitoring:   appContainer.Monitoring,
	}
}

// GetMonitoringService returns the monitoring service interface
func (c *PresentationContainer) GetMonitoringService() monitoring.MonitoringService {
	return c.Monitoring.(monitoring.MonitoringService)
}

// GetMonitoringHandlers returns the monitoring handlers interface
func (c *PresentationContainer) GetMonitoringHandlers() monitoring.MonitoringHandlers {
	return c.Monitoring.(monitoring.MonitoringHandlers)
}

package di

import (
	appContainer "github.com/achmdndy/safa-life-api/src/application/container"
	"github.com/achmdndy/safa-life-api/src/presentation/bootstrap"
	presentationContainer "github.com/achmdndy/safa-life-api/src/presentation/container"
)

// PresentationDI manages dependency injection for presentation layer
// This only handles presentation-specific DI and receives application factory from outside
type PresentationDI struct {
	appFactory appContainer.ContainerFactory
}

// NewPresentationDI creates a new presentation dependency injection container
func NewPresentationDI(appFactory appContainer.ContainerFactory) *PresentationDI {
	return &PresentationDI{
		appFactory: appFactory,
	}
}

// CreatePresentationFactory creates a presentation factory with injected application factory
func (pdi *PresentationDI) CreatePresentationFactory() *presentationContainer.PresentationContainerFactory {
	return presentationContainer.NewPresentationContainerFactory(pdi.appFactory)
}

// Bootstrap starts the application with proper dependency injection
func (pdi *PresentationDI) Bootstrap(config bootstrap.Config) error {
	// Create the presentation factory with properly injected application factory
	presentationFactory := pdi.CreatePresentationFactory()
	
	// Use the existing bootstrap function with the properly injected factory
	return bootstrap.Bootstrap(config, presentationFactory)
}

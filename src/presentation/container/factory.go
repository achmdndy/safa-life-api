package container

import (
	"github.com/achmdndy/safa-life-api/src/application/container"
)

// PresentationContainerFactory implements the ContainerFactory interface for presentation layer
// This factory delegates to the injected infrastructure factory to maintain Clean Architecture
type PresentationContainerFactory struct {
	appContainerFactory container.ContainerFactory
}

// NewPresentationContainerFactory creates a new presentation container factory
// The application container factory is injected to maintain Clean Architecture
func NewPresentationContainerFactory(appContainerFactory container.ContainerFactory) *PresentationContainerFactory {
	return &PresentationContainerFactory{
		appContainerFactory: appContainerFactory,
	}
}

// GetApplicationContainerFactory returns the injected application container factory
func (f *PresentationContainerFactory) GetApplicationContainerFactory() container.ContainerFactory {
	return f.appContainerFactory
}
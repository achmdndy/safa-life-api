package monitoring

import "github.com/achmdndy/safa-life-api/src/domain/shared"

// TracingFactory implements the shared.TracingFactory interface
type TracingFactory struct{}

// NewTracingFactory creates a new tracing factory
func NewTracingFactory() shared.TracingFactory {
	return &TracingFactory{}
}

// CreateTracingProvider creates a new Jaeger tracing provider
func (f *TracingFactory) CreateTracingProvider() shared.TracingProvider {
	return NewJaegerTracingProvider()
}
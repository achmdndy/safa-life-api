package monitoring

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// ConcreteTracingProvider is a concrete implementation of TracingProvider
// This will be injected with the actual tracing logic from outside
type ConcreteTracingProvider struct {
	initializeFunc func(config TracingConfig) (func(), error)
}

// NewConcreteTracingProvider creates a new concrete tracing provider
func NewConcreteTracingProvider(initFunc func(config TracingConfig) (func(), error)) TracingProvider {
	return &ConcreteTracingProvider{
		initializeFunc: initFunc,
	}
}

// Initialize implements TracingProvider interface
func (c *ConcreteTracingProvider) Initialize(config TracingConfig) (func(), error) {
	if c.initializeFunc != nil {
		return c.initializeFunc(config)
	}
	return func() {}, nil
}

// ConcreteMonitoringMiddleware is a concrete implementation of MonitoringMiddleware
type ConcreteMonitoringMiddleware struct {
	setupFunc func(router *gin.Engine, serviceName string)
}

// NewConcreteMonitoringMiddleware creates a new concrete monitoring middleware
func NewConcreteMonitoringMiddleware(setupFunc func(router *gin.Engine, serviceName string)) MonitoringMiddleware {
	return &ConcreteMonitoringMiddleware{
		setupFunc: setupFunc,
	}
}

// Setup implements MonitoringMiddleware interface
func (c *ConcreteMonitoringMiddleware) Setup(router *gin.Engine, serviceName string) {
	if c.setupFunc != nil {
		c.setupFunc(router, serviceName)
	}
}

// DefaultMonitoringMiddleware provides a default implementation with OpenTelemetry tracing
type DefaultMonitoringMiddleware struct{}

// NewDefaultMonitoringMiddleware creates a new default monitoring middleware
func NewDefaultMonitoringMiddleware() MonitoringMiddleware {
	return &DefaultMonitoringMiddleware{}
}

// Setup implements MonitoringMiddleware interface with OpenTelemetry tracing
func (d *DefaultMonitoringMiddleware) Setup(router *gin.Engine, serviceName string) {
	// Add OpenTelemetry tracing middleware FIRST
	router.Use(otelgin.Middleware(serviceName))
}
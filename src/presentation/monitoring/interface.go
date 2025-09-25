package monitoring

import "github.com/gin-gonic/gin"

// TracingConfig holds configuration for tracing initialization
type TracingConfig struct {
	ServiceName string
	Endpoint    string
}

// TracingProvider defines the interface for tracing providers
type TracingProvider interface {
	// Initialize sets up the tracing provider with the given configuration
	// Returns a cleanup function that should be called when shutting down
	Initialize(config TracingConfig) (cleanup func(), err error)
}

// MonitoringMiddleware defines the interface for monitoring middleware
type MonitoringMiddleware interface {
	// Setup configures monitoring middleware on the given router
	Setup(router *gin.Engine, serviceName string)
}
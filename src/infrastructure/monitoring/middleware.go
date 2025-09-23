package monitoring

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// MonitoringMiddleware combines Prometheus metrics and OpenTelemetry tracing
type MonitoringMiddleware struct {
	prometheus *PrometheusMetrics
}

// NewMonitoringMiddleware creates a new monitoring middleware
func NewMonitoringMiddleware() *MonitoringMiddleware {
	return &MonitoringMiddleware{
		prometheus: NewPrometheusMetrics(),
	}
}

// Setup configures all monitoring middleware for the Gin engine
func (m *MonitoringMiddleware) Setup(engine *gin.Engine, serviceName string) {
	// Add OpenTelemetry tracing middleware
	engine.Use(otelgin.Middleware(serviceName))
	
	// Add Prometheus metrics middleware
	engine.Use(m.prometheus.PrometheusMiddleware())
	
	// Add metrics endpoint
	engine.GET("/metrics", m.prometheus.Handler())
}

// GetPrometheusMetrics returns the Prometheus metrics instance
func (m *MonitoringMiddleware) GetPrometheusMetrics() *PrometheusMetrics {
	return m.prometheus
}
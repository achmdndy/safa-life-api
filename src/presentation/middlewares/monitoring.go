package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// MonitoringMiddleware provides monitoring capabilities for the presentation layer
type MonitoringMiddleware struct {
	requestDuration *prometheus.HistogramVec
	requestCounter  *prometheus.CounterVec
}

// NewMonitoringMiddleware creates a new monitoring middleware
func NewMonitoringMiddleware() *MonitoringMiddleware {
	requestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of HTTP requests in seconds",
		},
		[]string{"method", "path", "status"},
	)

	requestCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(requestCounter)

	return &MonitoringMiddleware{
		requestDuration: requestDuration,
		requestCounter:  requestCounter,
	}
}

// Setup configures monitoring middleware on the router
func (m *MonitoringMiddleware) Setup(router *gin.Engine, serviceName string) {
	// Add metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	
	// Add monitoring middleware
	router.Use(m.MetricsMiddleware())
}

// MetricsMiddleware returns a gin middleware that records metrics
func (m *MonitoringMiddleware) MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		
		c.Next()
		
		duration := time.Since(start).Seconds()
		status := string(rune(c.Writer.Status()))
		
		m.requestDuration.WithLabelValues(c.Request.Method, c.FullPath(), status).Observe(duration)
		m.requestCounter.WithLabelValues(c.Request.Method, c.FullPath(), status).Inc()
	}
}
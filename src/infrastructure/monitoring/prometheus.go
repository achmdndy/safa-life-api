package monitoring

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// PrometheusHandler creates a Gin handler for Prometheus metrics endpoint
func PrometheusHandler(registry *prometheus.Registry) gin.HandlerFunc {
	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		EnableOpenMetrics: true,
	})

	return gin.WrapH(handler)
}

// InitPrometheus initializes Prometheus metrics collection
func InitPrometheus(config PrometheusConfig) (*prometheus.Registry, error) {
	if !config.Enabled {
		return nil, nil
	}

	// Create a new registry
	registry := prometheus.NewRegistry()

	// Add Go runtime metrics using new collectors
	registry.MustRegister(collectors.NewGoCollector())
	registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	return registry, nil
}

// HealthHandler returns a simple health check for monitoring
func HealthHandler(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": serviceName,
		})
	}
}

// PrometheusHTTPHandler returns an http.Handler for Prometheus metrics
func PrometheusHTTPHandler(registry *prometheus.Registry) http.Handler {
	return promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		EnableOpenMetrics: true,
	})
}

// HealthHTTPHandler returns an http.Handler for health checks
func HealthHTTPHandler(serviceName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := fmt.Sprintf(`{"status":"healthy","service":"%s"}`, serviceName)
		if _, err := w.Write([]byte(response)); err != nil {
			log.Printf("Error writing health check response: %v", err)
		}
	})
}

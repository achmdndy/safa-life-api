package monitoring

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"strconv"
	"time"
)

// PrometheusMetrics holds all the Prometheus metrics
type PrometheusMetrics struct {
	httpRequestsTotal     *prometheus.CounterVec
	httpRequestDuration   *prometheus.HistogramVec
	httpRequestsInFlight  prometheus.Gauge
	applicationInfo       *prometheus.GaugeVec
}

// NewPrometheusMetrics creates and registers Prometheus metrics
func NewPrometheusMetrics() *PrometheusMetrics {
	metrics := &PrometheusMetrics{
		httpRequestsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests",
			},
			[]string{"method", "endpoint", "status_code"},
		),
		httpRequestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_request_duration_seconds",
				Help:    "Duration of HTTP requests in seconds",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"method", "endpoint"},
		),
		httpRequestsInFlight: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "http_requests_in_flight",
				Help: "Number of HTTP requests currently being processed",
			},
		),
		applicationInfo: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "application_info",
				Help: "Application information",
			},
			[]string{"version", "name"},
		),
	}

	// Register metrics with Prometheus
	prometheus.MustRegister(
		metrics.httpRequestsTotal,
		metrics.httpRequestDuration,
		metrics.httpRequestsInFlight,
		metrics.applicationInfo,
	)

	// Set application info
	metrics.applicationInfo.WithLabelValues("1.0.0", "safa-life-api").Set(1)

	return metrics
}

// PrometheusMiddleware returns a Gin middleware that collects Prometheus metrics
func (m *PrometheusMetrics) PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip metrics endpoint to avoid self-monitoring
		if c.Request.URL.Path == "/metrics" {
			c.Next()
			return
		}

		start := time.Now()
		m.httpRequestsInFlight.Inc()

		c.Next()

		duration := time.Since(start).Seconds()
		statusCode := strconv.Itoa(c.Writer.Status())

		m.httpRequestsTotal.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
			statusCode,
		).Inc()

		m.httpRequestDuration.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
		).Observe(duration)

		m.httpRequestsInFlight.Dec()
	}
}

// Handler returns the Prometheus metrics handler
func (m *PrometheusMetrics) Handler() gin.HandlerFunc {
	return gin.WrapH(promhttp.Handler())
}
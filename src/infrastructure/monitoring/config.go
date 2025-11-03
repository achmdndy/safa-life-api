package monitoring

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"

	"github.com/safalife/core-api/src/domain/monitoring"
)

// MonitoringConfig holds all monitoring configuration
type MonitoringConfig struct {
	Jaeger     JaegerConfig          `yaml:"jaeger"`
	Prometheus PrometheusConfig      `yaml:"prometheus"`
	Database   DatabaseTracingConfig `yaml:"database"`
}

// JaegerConfig holds Jaeger tracing configuration
type JaegerConfig struct {
	Endpoint       string `yaml:"endpoint"`
	ServiceName    string `yaml:"service_name"`
	ServiceVersion string `yaml:"service_version"`
	Environment    string `yaml:"environment"`
}

// PrometheusConfig holds Prometheus metrics configuration
type PrometheusConfig struct {
	Enabled bool `yaml:"enabled"`
}

// DatabaseTracingConfig holds database tracing configuration
type DatabaseTracingConfig struct {
	Enabled        bool   `yaml:"enabled"`
	ServiceName    string `yaml:"service_name"`
	ServiceVersion string `yaml:"service_version"`
	Environment    string `yaml:"environment"`
	IncludeParams  bool   `yaml:"include_params"`
}

// DefaultMonitoringConfig returns default monitoring configuration
func DefaultMonitoringConfig() MonitoringConfig {
	return MonitoringConfig{
		Jaeger: JaegerConfig{
			Endpoint:       "http://localhost:14268/api/traces",
			ServiceName:    "safalife-api",
			ServiceVersion: "1.0.0",
			Environment:    "development",
		},
		Prometheus: PrometheusConfig{
			Enabled: true,
		},
		Database: DatabaseTracingConfig{
			Enabled:        true,
			ServiceName:    "safalife-api",
			ServiceVersion: "1.0.0",
			Environment:    "development",
			IncludeParams:  false, // Set to false for security by default
		},
	}
}

// MonitoringService provides monitoring capabilities
type MonitoringService struct {
	tracerProvider *sdktrace.TracerProvider
	tracer         trace.Tracer
	registry       *prometheus.Registry
	metrics        *Metrics
	serviceName    string
}

// Metrics holds all Prometheus metrics
type Metrics struct {
	HTTPRequestsTotal    *prometheus.CounterVec
	HTTPRequestDuration  *prometheus.HistogramVec
	HTTPRequestsInFlight prometheus.Gauge
	DatabaseConnections  prometheus.Gauge
	RedisConnections     prometheus.Gauge
}

// NewMonitoringService creates a new monitoring service
func NewMonitoringService(config MonitoringConfig) (*MonitoringService, error) {
	// Initialize Jaeger tracing
	tracerProvider, err := InitJaeger(config.Jaeger)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Jaeger: %w", err)
	}

	// Initialize tracer
	tracer := otel.Tracer(config.Jaeger.ServiceName)

	// Initialize Prometheus registry
	registry, err := InitPrometheus(config.Prometheus)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Prometheus: %w", err)
	}

	// If Prometheus is disabled, create a dummy registry
	if registry == nil {
		registry = prometheus.NewRegistry()
	}

	// Initialize metrics
	metrics := &Metrics{
		HTTPRequestsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests",
			},
			[]string{"method", "endpoint", "status_code"},
		),
		HTTPRequestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_request_duration_seconds",
				Help:    "HTTP request duration in seconds",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"method", "endpoint", "status_code"},
		),
		HTTPRequestsInFlight: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "http_requests_in_flight",
				Help: "Number of HTTP requests currently being processed",
			},
		),
		DatabaseConnections: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "database_connections_active",
				Help: "Number of active database connections",
			},
		),
		RedisConnections: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "redis_connections_active",
				Help: "Number of active Redis connections",
			},
		),
	}

	// Register metrics with Prometheus
	if config.Prometheus.Enabled {
		registry.MustRegister(
			metrics.HTTPRequestsTotal,
			metrics.HTTPRequestDuration,
			metrics.HTTPRequestsInFlight,
			metrics.DatabaseConnections,
			metrics.RedisConnections,
		)
	}

	return &MonitoringService{
		tracerProvider: tracerProvider,
		tracer:         tracer,
		registry:       registry,
		metrics:        metrics,
		serviceName:    config.Jaeger.ServiceName,
	}, nil
}

// GetTracer returns the OpenTelemetry tracer
func (m *MonitoringService) GetTracer() trace.Tracer {
	return m.tracer
}

// GetRegistry returns the Prometheus registry
func (m *MonitoringService) GetRegistry() *prometheus.Registry {
	return m.registry
}

// GetMetrics returns the metrics collection
func (m *MonitoringService) GetMetrics() *Metrics {
	return m.metrics
}

// StartSpan starts a new tracing span
func (m *MonitoringService) StartSpan(ctx context.Context, name string) (context.Context, monitoring.Span) {
	ctx, span := m.tracer.Start(ctx, name)
	return ctx, NewSpanAdapter(span)
}

// RecordHTTPRequest records HTTP request metrics
func (m *MonitoringService) RecordHTTPRequest(method, endpoint, statusCode string, duration time.Duration) {
	m.metrics.HTTPRequestsTotal.WithLabelValues(method, endpoint, statusCode).Inc()
	m.metrics.HTTPRequestDuration.WithLabelValues(method, endpoint, statusCode).Observe(duration.Seconds())
}

// IncrementInFlightRequests increments the in-flight requests counter
func (m *MonitoringService) IncrementInFlightRequests() {
	m.metrics.HTTPRequestsInFlight.Inc()
}

// DecrementInFlightRequests decrements the in-flight requests counter
func (m *MonitoringService) DecrementInFlightRequests() {
	m.metrics.HTTPRequestsInFlight.Dec()
}

// SetDatabaseConnections sets the number of active database connections
func (m *MonitoringService) SetDatabaseConnections(count float64) {
	m.metrics.DatabaseConnections.Set(count)
}

// SetRedisConnections sets the current number of Redis connections
func (m *MonitoringService) SetRedisConnections(count float64) {
	m.metrics.RedisConnections.Set(count)
}

// PrometheusHandler returns a Gin handler for Prometheus metrics endpoint
func (m *MonitoringService) PrometheusHandler() gin.HandlerFunc {
	return PrometheusHandler(m.registry)
}

// HealthHandler returns a simple health check handler
func (m *MonitoringService) HealthHandler() gin.HandlerFunc {
	return HealthHandler(m.serviceName)
}

// GetPrometheusHandler returns the Prometheus handler as http.Handler
func (m *MonitoringService) GetPrometheusHandler() http.Handler {
	return PrometheusHTTPHandler(m.registry)
}

// GetHealthHandler returns the health handler as http.Handler
func (m *MonitoringService) GetHealthHandler() http.Handler {
	return HealthHTTPHandler(m.serviceName)
}

// NewStringAttribute creates a new string attribute for tracing
func (m *MonitoringService) NewStringAttribute(key, value string) monitoring.Attribute {
	return NewStringAttribute(key, value)
}

// NewIntAttribute creates a new int attribute for tracing
func (m *MonitoringService) NewIntAttribute(key string, value int64) monitoring.Attribute {
	return NewIntAttribute(key, int(value))
}

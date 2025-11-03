package monitoring

import (
	"context"
	"net/http"
	"time"
)

// Span represents a tracing span interface
type Span interface {
	End()
	SetAttributes(attributes ...Attribute)
}

// Attribute represents a span attribute
type Attribute interface {
	Key() string
	Value() interface{}
}

// MonitoringService defines the interface for monitoring capabilities
type MonitoringService interface {
	// Tracing methods
	StartSpan(ctx context.Context, name string) (context.Context, Span)

	// Metrics methods
	RecordHTTPRequest(method, endpoint, statusCode string, duration time.Duration)
	IncrementInFlightRequests()
	DecrementInFlightRequests()
	SetDatabaseConnections(count float64)
	SetRedisConnections(count float64)
}

// MonitoringHandlers defines handlers for monitoring endpoints
type MonitoringHandlers interface {
	GetPrometheusHandler() http.Handler
	GetHealthHandler() http.Handler
}

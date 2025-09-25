package monitoring

import (
	"context"
	"fmt"
	"log"

	"github.com/achmdndy/safa-life-api/src/domain/shared"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// JaegerTracingProvider implements the TracingProvider interface for Jaeger
type JaegerTracingProvider struct{}

// NewJaegerTracingProvider creates a new Jaeger tracing provider
func NewJaegerTracingProvider() shared.TracingProvider {
	return &JaegerTracingProvider{}
}

// Initialize sets up Jaeger tracing and returns a cleanup function
func (j *JaegerTracingProvider) Initialize(config shared.TracingConfig) (func(), error) {
	// Create OTLP HTTP exporter
	// Use the endpoint directly without adding HTTP prefix since it's already handled by Docker networking
	exporter, err := otlptracehttp.New(
		context.Background(),
		otlptracehttp.WithEndpoint(config.Endpoint),
		otlptracehttp.WithInsecure(), // Use HTTP instead of HTTPS for local development
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create OTLP exporter: %v", err)
	}

	// Create resource with service information
	res, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(config.ServiceName),
			semconv.ServiceVersionKey.String("1.0.0"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %v", err)
	}

	// Create trace provider
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
		trace.WithSampler(trace.AlwaysSample()),
	)

	// Set global trace provider
	otel.SetTracerProvider(tp)

	// Set global propagator
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	log.Printf("Jaeger tracing initialized with service name: %s", config.ServiceName)

	// Return cleanup function
	cleanup := func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}

	return cleanup, nil
}
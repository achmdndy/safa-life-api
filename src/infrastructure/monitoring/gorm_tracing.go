package monitoring

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	gormotel "gorm.io/plugin/opentelemetry/tracing"
)

// InitGormTracing initializes GORM tracing plugin
func InitGormTracing(db *gorm.DB, config DatabaseTracingConfig) error {
	if !config.Enabled {
		return nil
	}

	// Create the tracing plugin with basic configuration
	plugin := gormotel.NewPlugin()

	// Use the plugin with GORM
	if err := db.Use(plugin); err != nil {
		return fmt.Errorf("failed to register GORM tracing plugin: %w", err)
	}

	// Add custom callbacks for additional tracing attributes
	if err := addCustomTracingCallbacks(db, config); err != nil {
		return fmt.Errorf("failed to add custom tracing callbacks: %w", err)
	}

	return nil
}

// addCustomTracingCallbacks adds custom callbacks for enhanced tracing
func addCustomTracingCallbacks(db *gorm.DB, config DatabaseTracingConfig) error {
	// Add callback for query operations
	if err := db.Callback().Query().Before("gorm:query").Register("tracing:before_query", func(db *gorm.DB) {
		ctx := db.Statement.Context
		if ctx != nil {
			span := trace.SpanFromContext(ctx)
			if span.IsRecording() {
				span.SetAttributes(
					attribute.String("db.operation", "query"),
					attribute.String("db.table", db.Statement.Table),
					attribute.String("service.name", config.ServiceName),
					attribute.String("service.version", config.ServiceVersion),
					attribute.String("environment", config.Environment),
				)
			}
		}
	}); err != nil {
		return fmt.Errorf("failed to register query tracing callback: %w", err)
	}

	// Add callback for create operations
	if err := db.Callback().Create().Before("gorm:create").Register("tracing:before_create", func(db *gorm.DB) {
		ctx := db.Statement.Context
		if ctx != nil {
			span := trace.SpanFromContext(ctx)
			if span.IsRecording() {
				span.SetAttributes(
					attribute.String("db.operation", "create"),
					attribute.String("db.table", db.Statement.Table),
					attribute.String("service.name", config.ServiceName),
					attribute.String("service.version", config.ServiceVersion),
					attribute.String("environment", config.Environment),
				)
			}
		}
	}); err != nil {
		return fmt.Errorf("failed to register create tracing callback: %w", err)
	}

	// Add callback for update operations
	if err := db.Callback().Update().Before("gorm:update").Register("tracing:before_update", func(db *gorm.DB) {
		ctx := db.Statement.Context
		if ctx != nil {
			span := trace.SpanFromContext(ctx)
			if span.IsRecording() {
				span.SetAttributes(
					attribute.String("db.operation", "update"),
					attribute.String("db.table", db.Statement.Table),
					attribute.String("service.name", config.ServiceName),
					attribute.String("service.version", config.ServiceVersion),
					attribute.String("environment", config.Environment),
				)
			}
		}
	}); err != nil {
		return fmt.Errorf("failed to register update tracing callback: %w", err)
	}

	// Add callback for delete operations
	if err := db.Callback().Delete().Before("gorm:delete").Register("tracing:before_delete", func(db *gorm.DB) {
		ctx := db.Statement.Context
		if ctx != nil {
			span := trace.SpanFromContext(ctx)
			if span.IsRecording() {
				span.SetAttributes(
					attribute.String("db.operation", "delete"),
					attribute.String("db.table", db.Statement.Table),
					attribute.String("service.name", config.ServiceName),
					attribute.String("service.version", config.ServiceVersion),
					attribute.String("environment", config.Environment),
				)
			}
		}
	}); err != nil {
		return fmt.Errorf("failed to register delete tracing callback: %w", err)
	}

	return nil
}

// WithDatabaseTracing wraps a function with database tracing context
func WithDatabaseTracing(ctx context.Context, operationName string, fn func(context.Context) error) error {
	tracer := otel.Tracer("safalife-database")
	ctx, span := tracer.Start(ctx, operationName)
	defer span.End()

	// Add common database attributes
	span.SetAttributes(
		attribute.String("db.system", "postgresql"),
		attribute.String("db.name", "safalife"),
		attribute.String("operation.name", operationName),
	)

	err := fn(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetAttributes(
			attribute.String("error.type", "database_error"),
			attribute.String("error.message", err.Error()),
		)
	}

	return err
}

// WithDatabaseOperation creates a span for a specific database operation
func WithDatabaseOperation(ctx context.Context, operation, table string, fn func(context.Context) error) error {
	tracer := otel.Tracer("safalife-database")
	spanName := fmt.Sprintf("db.%s.%s", operation, table)
	ctx, span := tracer.Start(ctx, spanName)
	defer span.End()

	// Add operation-specific attributes
	span.SetAttributes(
		attribute.String("db.system", "postgresql"),
		attribute.String("db.name", "safalife"),
		attribute.String("db.operation", operation),
		attribute.String("db.table", table),
	)

	err := fn(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetAttributes(
			attribute.String("error.type", "database_error"),
			attribute.String("error.message", err.Error()),
		)
	}

	return err
}

// AddDatabaseMetrics adds database-related attributes to the current span
func AddDatabaseMetrics(ctx context.Context, rowsAffected int64, queryDuration time.Duration) {
	span := trace.SpanFromContext(ctx)
	if span.IsRecording() {
		span.SetAttributes(
			attribute.Int64("db.rows_affected", rowsAffected),
			attribute.Float64("db.query_duration_ms", float64(queryDuration.Nanoseconds())/1e6),
		)
	}
}

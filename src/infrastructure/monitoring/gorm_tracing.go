package monitoring

import (
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

const (
	gormSpanKey        = "gorm:span"
	callBackBeforeName = "otel:before"
	callBackAfterName  = "otel:after"
)

// GormTracingPlugin implements GORM plugin interface for OpenTelemetry tracing
type GormTracingPlugin struct {
	tracer trace.Tracer
}

// NewGormTracingPlugin creates a new GORM tracing plugin
func NewGormTracingPlugin() *GormTracingPlugin {
	return &GormTracingPlugin{
		tracer: otel.Tracer("gorm"),
	}
}

// Name returns the plugin name
func (p *GormTracingPlugin) Name() string {
	return "otel-gorm-tracing"
}

// Initialize initializes the plugin
func (p *GormTracingPlugin) Initialize(db *gorm.DB) error {
	// Register callbacks for all operations
	if err := p.registerCallbacks(db); err != nil {
		return fmt.Errorf("failed to register GORM tracing callbacks: %w", err)
	}
	return nil
}

// registerCallbacks registers before and after callbacks for all GORM operations
func (p *GormTracingPlugin) registerCallbacks(db *gorm.DB) error {
	// Query operations
	if err := db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, p.before); err != nil {
		return err
	}
	if err := db.Callback().Query().After("gorm:query").Register(callBackAfterName, p.after); err != nil {
		return err
	}

	// Create operations
	if err := db.Callback().Create().Before("gorm:create").Register(callBackBeforeName, p.before); err != nil {
		return err
	}
	if err := db.Callback().Create().After("gorm:create").Register(callBackAfterName, p.after); err != nil {
		return err
	}

	// Update operations
	if err := db.Callback().Update().Before("gorm:update").Register(callBackBeforeName, p.before); err != nil {
		return err
	}
	if err := db.Callback().Update().After("gorm:update").Register(callBackAfterName, p.after); err != nil {
		return err
	}

	// Delete operations
	if err := db.Callback().Delete().Before("gorm:delete").Register(callBackBeforeName, p.before); err != nil {
		return err
	}
	if err := db.Callback().Delete().After("gorm:delete").Register(callBackAfterName, p.after); err != nil {
		return err
	}

	// Row operations
	if err := db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, p.before); err != nil {
		return err
	}
	if err := db.Callback().Row().After("gorm:row").Register(callBackAfterName, p.after); err != nil {
		return err
	}

	// Raw operations
	if err := db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, p.before); err != nil {
		return err
	}
	if err := db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, p.after); err != nil {
		return err
	}

	return nil
}

// before callback starts a new span for the database operation
func (p *GormTracingPlugin) before(db *gorm.DB) {
	// Skip if context is nil or if we're already in a span to avoid recursion
	if db.Statement == nil || db.Statement.Context == nil {
		return
	}

	ctx := db.Statement.Context
	
	// Check if we already have a span to avoid double-wrapping
	if span := trace.SpanFromContext(ctx); span.SpanContext().IsValid() {
		// If we already have an active span, just store it
		db.Set(gormSpanKey, span)
		return
	}

	// Create span name based on operation
	spanName := p.getSpanName(db)
	
	// Start new span
	ctx, span := p.tracer.Start(ctx, spanName, trace.WithSpanKind(trace.SpanKindClient))
	
	// Set basic span attributes
	p.setSpanAttributes(span, db)
	
	// Update context and store span
	db.Statement.Context = ctx
	db.Set(gormSpanKey, span)
}

// after callback finishes the span and records any errors
func (p *GormTracingPlugin) after(db *gorm.DB) {
	spanValue, exists := db.Get(gormSpanKey)
	if !exists {
		return
	}

	span, ok := spanValue.(trace.Span)
	if !ok {
		return
	}

	defer span.End()

	// Record error if any
	if db.Error != nil {
		span.RecordError(db.Error)
		span.SetStatus(codes.Error, db.Error.Error())
	} else {
		span.SetStatus(codes.Ok, "")
	}

	// Add additional attributes after execution only if we have valid data
	if db.Statement != nil && db.Statement.SQL.String() != "" {
		span.SetAttributes(
			attribute.String("db.statement", db.Statement.SQL.String()),
			attribute.Int64("db.rows_affected", db.RowsAffected),
		)
	}
}

// getSpanName generates appropriate span name based on the operation
func (p *GormTracingPlugin) getSpanName(db *gorm.DB) string {
	if db.Statement.Schema != nil && db.Statement.Schema.Table != "" {
		return fmt.Sprintf("gorm:%s %s", p.getOperationType(db), db.Statement.Schema.Table)
	}
	return fmt.Sprintf("gorm:%s", p.getOperationType(db))
}

// getOperationType determines the type of database operation
func (p *GormTracingPlugin) getOperationType(db *gorm.DB) string {
	switch {
	case db.Statement.SQL.String() != "":
		// Raw SQL - try to determine from SQL content
		sql := db.Statement.SQL.String()
		switch {
		case len(sql) >= 6 && sql[:6] == "SELECT":
			return "select"
		case len(sql) >= 6 && sql[:6] == "INSERT":
			return "insert"
		case len(sql) >= 6 && sql[:6] == "UPDATE":
			return "update"
		case len(sql) >= 6 && sql[:6] == "DELETE":
			return "delete"
		default:
			return "raw"
		}
	case db.Statement.Dest != nil:
		// Query operation (SELECT)
		return "select"
	default:
		// Determine by GORM context
		if db.Statement.Schema != nil {
			// Check if this is a write operation by looking at the context
			if db.Statement.ReflectValue.IsValid() {
				// Check the kind of reflection value to determine operation type
				kind := db.Statement.ReflectValue.Kind()
				if kind.String() != "" {
					// This is likely a create/update operation
					if db.Statement.Schema.PrioritizedPrimaryField != nil {
						// Has primary key, likely update
						return "update"
					}
					return "create"
				}
			}
		}
		return "query"
	}
}

// setSpanAttributes sets common span attributes for database operations
func (p *GormTracingPlugin) setSpanAttributes(span trace.Span, db *gorm.DB) {
	// Basic safety check
	if db == nil || db.Statement == nil {
		return
	}

	// Database system
	span.SetAttributes(
		attribute.String("db.system", "postgresql"),
		attribute.String("db.operation", p.getOperationType(db)),
	)

	// Table name if available
	if db.Statement.Schema != nil && db.Statement.Schema.Table != "" {
		span.SetAttributes(
			attribute.String("db.sql.table", db.Statement.Schema.Table),
		)
	}

	// Database name - use a simple approach to avoid potential crashes
	if db.Statement.ConnPool != nil {
		span.SetAttributes(
			attribute.String("db.name", "safa_life"), // Use configured database name
		)
	}

	// Add timing
	span.SetAttributes(
		attribute.String("db.started_at", time.Now().Format(time.RFC3339)),
	)
}
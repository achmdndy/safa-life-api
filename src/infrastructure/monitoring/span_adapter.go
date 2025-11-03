package monitoring

import (
	"fmt"

	"github.com/safalife/core-api/src/domain/monitoring"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// SpanAdapter adapts OpenTelemetry span to domain interface
type SpanAdapter struct {
	span trace.Span
}

// NewSpanAdapter creates a new span adapter
func NewSpanAdapter(span trace.Span) monitoring.Span {
	return &SpanAdapter{span: span}
}

// End ends the span
func (s *SpanAdapter) End() {
	s.span.End()
}

// SetAttributes sets attributes on the span
func (s *SpanAdapter) SetAttributes(attributes ...monitoring.Attribute) {
	attrs := make([]attribute.KeyValue, len(attributes))
	for i, attr := range attributes {
		var value attribute.Value
		switch v := attr.Value().(type) {
		case string:
			value = attribute.StringValue(v)
		case int:
			value = attribute.IntValue(v)
		case int64:
			value = attribute.Int64Value(v)
		case float64:
			value = attribute.Float64Value(v)
		case bool:
			value = attribute.BoolValue(v)
		default:
			value = attribute.StringValue(fmt.Sprintf("%v", v))
		}
		attrs[i] = attribute.KeyValue{
			Key:   attribute.Key(attr.Key()),
			Value: value,
		}
	}
	s.span.SetAttributes(attrs...)
}

// AttributeAdapter adapts attribute to domain interface
type AttributeAdapter struct {
	key   string
	value interface{}
}

// NewStringAttribute creates a new string attribute
func NewStringAttribute(key, value string) monitoring.Attribute {
	return &AttributeAdapter{key: key, value: value}
}

// NewIntAttribute creates a new int attribute
func NewIntAttribute(key string, value int) monitoring.Attribute {
	return &AttributeAdapter{key: key, value: value}
}

// Key returns the attribute key
func (a *AttributeAdapter) Key() string {
	return a.key
}

// Value returns the attribute value
func (a *AttributeAdapter) Value() interface{} {
	return a.value
}

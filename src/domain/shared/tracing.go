package shared

// TracingConfig represents the configuration needed for tracing
type TracingConfig struct {
	ServiceName string
	Endpoint    string
}

// TracingProvider defines the interface for distributed tracing providers
// This interface is in domain layer to avoid circular dependencies
type TracingProvider interface {
	// Initialize sets up tracing and returns a cleanup function
	Initialize(config TracingConfig) (func(), error)
}

// TracingFactory defines the interface for creating tracing providers
// This allows the application layer to create infrastructure tracing providers
// without directly importing infrastructure packages
type TracingFactory interface {
	CreateTracingProvider() TracingProvider
}
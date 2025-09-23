package shared

import (
	"github.com/achmdndy/safa-life-api/src/domain/health"
)

// Container defines the interface for dependency injection
// This interface is in domain layer to avoid circular dependencies
type Container interface {
	// Repository access
	GetHealthRepository() health.CheckerRepository
	
	// Service access  
	GetHealthService() *health.Service
	
	// Lifecycle management
	Close() error
}

// QueryHandlerInterface defines the interface for query handlers
// This is in domain layer to avoid importing application layer
type QueryHandlerInterface interface {
	Handle(query interface{}) (interface{}, error)
}

// ContainerFactory defines the interface for creating containers
// This allows the application layer to create infrastructure containers
// without directly importing infrastructure packages
type ContainerFactory interface {
	CreateContainer(dbConfig DatabaseConfig, redisConfig RedisConfig) Container
}

// DatabaseConfig represents database configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}

// RedisConfig represents Redis configuration
type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}
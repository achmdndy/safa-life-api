package container

import (
	"github.com/achmdndy/safa-life-api/src/application/health/interfaces"
	"github.com/achmdndy/safa-life-api/src/application/health/queries"
	"github.com/achmdndy/safa-life-api/src/domain/health"
	"github.com/achmdndy/safa-life-api/src/domain/shared"
	"github.com/achmdndy/safa-life-api/src/infrastructure/container"
)

// DatabaseConfig represents database configuration for application layer
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}

// RedisConfig represents Redis configuration for application layer
type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

// ApplicationContainer defines application-specific dependencies
// This interface is in application layer for presentation to use
type ApplicationContainer interface {
	shared.Container
	
	// Query Handler access - using application interface
	GetHealthQueryHandler() interfaces.QueryHandler
}

// AppContainer implements the ApplicationContainer interface
// This is in application layer and uses dependency injection through domain interfaces
type AppContainer struct {
	// Infrastructure container injected through domain interface
	infraContainer shared.Container

	// Query Handlers
	getHealthQueryHandler *queries.GetHealthQueryHandler
}

// NewAppContainer creates a new application container with configuration
func NewAppContainer(dbConfig DatabaseConfig, redisConfig RedisConfig) ApplicationContainer {
	// Create infrastructure container with real implementation
	infraContainer := container.NewInfraContainer(
		container.DatabaseConfig{
			Host:     dbConfig.Host,
			Port:     dbConfig.Port,
			User:     dbConfig.User,
			Password: dbConfig.Password,
			DBName:   dbConfig.DBName,
			SSLMode:  dbConfig.SSLMode,
			TimeZone: dbConfig.TimeZone,
		},
		container.RedisConfig{
			Host:     redisConfig.Host,
			Port:     redisConfig.Port,
			Password: redisConfig.Password,
			DB:       redisConfig.DB,
		},
	)
	
	appContainer := &AppContainer{
		infraContainer: infraContainer,
	}
	appContainer.initializeQueryHandlers()
	return appContainer
}

// GetHealthRepository returns the health repository from infrastructure container
func (c *AppContainer) GetHealthRepository() health.CheckerRepository {
	return c.infraContainer.GetHealthRepository()
}

// GetHealthService returns the health service from infrastructure container
func (c *AppContainer) GetHealthService() *health.Service {
	return c.infraContainer.GetHealthService()
}

// GetHealthQueryHandler returns the health query handler as application interface
func (c *AppContainer) GetHealthQueryHandler() interfaces.QueryHandler {
	return c.getHealthQueryHandler
}

// Close gracefully shuts down all connections through infrastructure container
func (c *AppContainer) Close() error {
	return c.infraContainer.Close()
}

// initializeQueryHandlers creates query handler instances using injected dependencies
func (c *AppContainer) initializeQueryHandlers() {
	// Get health service from infrastructure container
	healthService := c.infraContainer.GetHealthService()
	
	// Create query handler with the service
	c.getHealthQueryHandler = queries.NewGetHealthQueryHandler(healthService)
}
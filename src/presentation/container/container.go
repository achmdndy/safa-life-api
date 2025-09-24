package container

import (
	"github.com/achmdndy/safa-life-api/src/application/container"
	healthHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/health"
	quranHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/quran"
	"github.com/achmdndy/safa-life-api/src/presentation/routes"
)

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

// ServerConfig represents server configuration
type ServerConfig struct {
	Host string
	Port string
	Env  string
}

// PresentationContainer holds presentation layer dependencies
// This only imports from application layer, following Clean Architecture
type PresentationContainer struct {
	// Application container (dependency injection from lower layers)
	appContainer container.ApplicationContainer

	// HTTP Handlers (presentation layer specific)
	HealthHandler *healthHandler.Handler
	QuranHandler  *quranHandler.Handler

	// Routes (presentation layer specific)
	RouteConfig routes.RouteConfig
}

// NewPresentationContainer creates a new presentation container with configuration
func NewPresentationContainer(dbConfig DatabaseConfig, redisConfig RedisConfig) *PresentationContainer {
	// Create application container with configuration
	appContainer := container.NewAppContainer(
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
	
	presentationContainer := &PresentationContainer{
		appContainer: appContainer,
	}
	
	presentationContainer.initializeHTTPHandlers()
	presentationContainer.initializeRoutes()
	
	return presentationContainer
}

// GetApplicationContainer returns the application container
func (c *PresentationContainer) GetApplicationContainer() container.ApplicationContainer {
	return c.appContainer
}

// Close gracefully shuts down all connections
func (c *PresentationContainer) Close() error {
	return c.appContainer.Close()
}

// initializeHTTPHandlers creates HTTP handler instances
func (c *PresentationContainer) initializeHTTPHandlers() {
	queryHandler := c.appContainer.GetHealthQueryHandler()
	c.HealthHandler = healthHandler.NewHandler(queryHandler)

	quranQueryHandler := c.appContainer.GetQuranQueryHandler()
	quranCommandHandler := c.appContainer.GetQuranCommandHandler()
	c.QuranHandler = quranHandler.NewHandler(quranQueryHandler, quranCommandHandler)
}

// initializeRoutes sets up route configuration
func (c *PresentationContainer) initializeRoutes() {
	c.RouteConfig = routes.RouteConfig{
		HealthHandler: c.HealthHandler,
		QuranHandler:  c.QuranHandler,
	}
}
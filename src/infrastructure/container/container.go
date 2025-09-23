package container

import (
	"log"

	"github.com/achmdndy/safa-life-api/src/domain/health"
	"github.com/achmdndy/safa-life-api/src/domain/shared"
	"github.com/achmdndy/safa-life-api/src/infrastructure/configs"
	healthInfra "github.com/achmdndy/safa-life-api/src/infrastructure/health"
)

// InfraContainerFactory implements the ContainerFactory interface
type InfraContainerFactory struct{}

// NewInfraContainerFactory creates a new infrastructure container factory
func NewInfraContainerFactory() shared.ContainerFactory {
	return &InfraContainerFactory{}
}

// CreateContainer creates a new infrastructure container with configuration
func (f *InfraContainerFactory) CreateContainer(dbConfig shared.DatabaseConfig, redisConfig shared.RedisConfig) shared.Container {
	return NewInfraContainer(
		DatabaseConfig{
			Host:     dbConfig.Host,
			Port:     dbConfig.Port,
			User:     dbConfig.User,
			Password: dbConfig.Password,
			DBName:   dbConfig.DBName,
			SSLMode:  dbConfig.SSLMode,
			TimeZone: dbConfig.TimeZone,
		},
		RedisConfig{
			Host:     redisConfig.Host,
			Port:     redisConfig.Port,
			Password: redisConfig.Password,
			DB:       redisConfig.DB,
		},
	)
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

// InfraContainer implements the domain container interfaces
// This is in infrastructure layer and handles all infrastructure concerns
type InfraContainer struct {
	// Infrastructure
	database *configs.Database
	redis    *configs.Redis

	// Repositories
	healthRepository health.CheckerRepository

	// Configuration
	dbConfig    DatabaseConfig
	redisConfig RedisConfig
}

// NewInfraContainer creates a new infrastructure container with injected configuration
func NewInfraContainer(dbConfig DatabaseConfig, redisConfig RedisConfig) shared.Container {
	container := &InfraContainer{
		dbConfig:    dbConfig,
		redisConfig: redisConfig,
	}
	container.initializeInfrastructure()
	container.initializeRepositories()
	return container
}

// GetHealthRepository returns the health repository
func (c *InfraContainer) GetHealthRepository() health.CheckerRepository {
	return c.healthRepository
}

// GetHealthService returns the health service
func (c *InfraContainer) GetHealthService() *health.Service {
	// Create service with repository from infrastructure
	return health.NewService(c.healthRepository)
}

// Close gracefully shuts down all connections
func (c *InfraContainer) Close() error {
	if c.database != nil {
		if err := c.database.Close(); err != nil {
			return err
		}
	}
	if c.redis != nil {
		if err := c.redis.Close(); err != nil {
			return err
		}
	}
	return nil
}

// initializeInfrastructure sets up database and Redis connections using injected configuration
func (c *InfraContainer) initializeInfrastructure() {
	// Initialize database with injected config
	db, err := configs.NewDatabase(configs.DatabaseConfig{
		Host:     c.dbConfig.Host,
		Port:     c.dbConfig.Port,
		User:     c.dbConfig.User,
		Password: c.dbConfig.Password,
		DBName:   c.dbConfig.DBName,
		SSLMode:  c.dbConfig.SSLMode,
		TimeZone: c.dbConfig.TimeZone,
	})
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	c.database = db

	// Initialize Redis with injected config
	redis, err := configs.NewRedis(configs.RedisConfig{
		Host:     c.redisConfig.Host,
		Port:     c.redisConfig.Port,
		Password: c.redisConfig.Password,
		DB:       c.redisConfig.DB,
	})
	if err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}
	c.redis = redis
}

// initializeRepositories creates repository instances
func (c *InfraContainer) initializeRepositories() {
	c.healthRepository = healthInfra.NewCheckerRepository(c.database, c.redis)
}
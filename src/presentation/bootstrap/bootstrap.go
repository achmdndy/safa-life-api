package bootstrap

import (
	"fmt"
	"log"

	"github.com/achmdndy/safa-life-api/src/presentation/container"
	"github.com/achmdndy/safa-life-api/src/presentation/monitoring"
	"github.com/achmdndy/safa-life-api/src/presentation/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Config represents the configuration needed to bootstrap the application
type Config struct {
	DB struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
		SSLMode  string
		TimeZone string
	}
	Redis struct {
		Host     string
		Port     int
		Password string
		DB       int
	}
	Server struct {
		Host string
		Port int
	}
	Tracing struct {
		ServiceName string
		Endpoint    string
	}
}

// Bootstrap initializes and starts the application with proper dependency injection
// This function now expects a properly injected containerFactory, tracingProvider, and monitoringMiddleware
func Bootstrap(config Config, containerFactory *container.PresentationContainerFactory, tracingProvider monitoring.TracingProvider) error {
	fmt.Println("🚀 Starting application...")

	// Initialize tracing if provider is provided
	var tracingCleanup func()
	if tracingProvider != nil && config.Tracing.ServiceName != "" {
		fmt.Println("🔍 Initializing distributed tracing...")
		cleanup, err := tracingProvider.Initialize(monitoring.TracingConfig{
			ServiceName: config.Tracing.ServiceName,
			Endpoint:    config.Tracing.Endpoint,
		})
		if err != nil {
			log.Printf("Warning: Failed to initialize tracing: %v", err)
		} else {
			tracingCleanup = cleanup
		}
	}

	// Create presentation container with injected factory
	presentationContainer := container.NewPresentationContainer(
		containerFactory,
		container.DatabaseConfig{
			Host:     config.DB.Host,
			Port:     config.DB.Port,
			User:     config.DB.User,
			Password: config.DB.Password,
			DBName:   config.DB.Name,
			SSLMode:  config.DB.SSLMode,
			TimeZone: config.DB.TimeZone,
		},
		container.RedisConfig{
			Host:     config.Redis.Host,
			Port:     config.Redis.Port,
			Password: config.Redis.Password,
			DB:       config.Redis.DB,
		},
	)
	defer func() {
		if err := presentationContainer.Close(); err != nil {
			log.Printf("Error closing container: %v", err)
		}
		// Cleanup tracing
		if tracingCleanup != nil {
			tracingCleanup()
		}
	}()

	// Setup routes with integrated monitoring and tracing
	fmt.Println("🛣️  Setting up routes with integrated monitoring and tracing...")
	router := routes.SetupRoutes(presentationContainer.RouteConfig)

	// Setup Swagger
	fmt.Println("📚 Setting up Swagger documentation...")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	serverAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	fmt.Printf("🌙 Safa Life API is running on %s\n", serverAddr)
	fmt.Printf("📊 Prometheus metrics available at http://%s/metrics\n", serverAddr)
	fmt.Printf("📚 Swagger documentation available at http://%s/swagger/index.html\n", serverAddr)

	if err := router.Run(serverAddr); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	return nil
}

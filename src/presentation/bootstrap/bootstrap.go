package bootstrap

import (
	"fmt"
	"log"

	"github.com/achmdndy/safa-life-api/src/presentation/container"
	"github.com/achmdndy/safa-life-api/src/presentation/middlewares"
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
}

// Bootstrap initializes and starts the application with proper dependency injection
// This function now expects a properly injected containerFactory
func Bootstrap(config Config, containerFactory *container.PresentationContainerFactory) error {
	fmt.Println("🚀 Starting application...")

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
	}()

	// Setup routes with complete router (includes middleware)
	fmt.Println("🛣️  Setting up routes with middleware...")
	router := routes.SetupRoutes(presentationContainer.RouteConfig)

	// Setup monitoring middleware
	fmt.Println("📊 Setting up monitoring...")
	monitoringMiddleware := middlewares.NewMonitoringMiddleware()
	monitoringMiddleware.Setup(router, "safa-life-api")

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

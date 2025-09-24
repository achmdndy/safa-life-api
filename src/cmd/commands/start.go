package commands

import (
	"fmt"
	"log"

	"github.com/achmdndy/safa-life-api/src/cmd/core"
	"github.com/achmdndy/safa-life-api/src/infrastructure/monitoring"
	"github.com/achmdndy/safa-life-api/src/presentation/container"
	"github.com/achmdndy/safa-life-api/src/presentation/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	
	// Swagger imports
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "github.com/achmdndy/safa-life-api/docs" // This will be generated
)

// @title Safa Life API
// @version 1.0
// @description Al-Quran API with comprehensive Islamic content
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		Start()
	},
}

func Start() {
	fmt.Println("📦 Initializing dependencies...")

	// Initialize Jaeger tracing
	fmt.Println("🔍 Initializing Jaeger tracing...")
	jaegerCleanup := monitoring.InitJaeger(monitoring.JaegerConfig{
		ServiceName: core.Config.Monitoring.Jaeger.ServiceName,
		Endpoint:    core.Config.Monitoring.Jaeger.Endpoint,
	})
	defer jaegerCleanup()

	cont := container.NewPresentationContainer(
		container.DatabaseConfig{
			Host:     core.Config.DB.Host,
			Port:     core.Config.DB.Port,
			User:     core.Config.DB.User,
			Password: core.Config.DB.Password,
			DBName:   core.Config.DB.Name,
			SSLMode:  core.Config.DB.SSLMode,
			TimeZone: core.Config.DB.TimeZone,
		},
		container.RedisConfig{
			Host:     core.Config.Redis.Host,
			Port:     core.Config.Redis.Port,
			Password: core.Config.Redis.Password,
			DB:       core.Config.Redis.DB,
		},
	)
	defer func() {
		if err := cont.Close(); err != nil {
			log.Printf("Error closing container: %v", err)
		}
	}()

	// Create router first
	fmt.Println("🛣️  Creating router...")
	router := gin.New()

	// Setup monitoring middleware FIRST
	fmt.Println("📊 Setting up monitoring...")
	monitoringMiddleware := monitoring.NewMonitoringMiddleware()
	monitoringMiddleware.Setup(router, "safa-life-api")

	// Setup Swagger
	fmt.Println("📚 Setting up Swagger documentation...")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup routes AFTER monitoring middleware
	fmt.Println("🛣️  Setting up routes...")
	routes.SetupRoutesWithRouter(router, cont.RouteConfig)

	// Start server
	serverAddr := fmt.Sprintf("%s:%d", core.Config.Server.Host, core.Config.Port)
	fmt.Printf("🌙 Safa Life API is running on %s\n", serverAddr)
	fmt.Printf("📊 Prometheus metrics available at http://%s/metrics\n", serverAddr)
	fmt.Printf("📚 Swagger documentation available at http://%s/swagger/index.html\n", serverAddr)

	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

package commands

// @title Safalife API
// @version 1.0.0
// @description Safalife API is a comprehensive backend service for Safalife's competition management system. This API provides endpoints for health monitoring, user management, competition tracking, and administrative functions. Built with Go using Clean Architecture principles, it offers high performance and reliability for managing large-scale competitions.
// @termsOfService http://swagger.io/terms/

// @contact.name Safalife API Support Team
// @contact.url http://www.safalife.com/support
// @contact.email api-support@safalife.com

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /v1
// @schemes http https

// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Enter the token with the `Bearer: ` prefix, e.g. "Bearer abcde12345"

// @tag.name Surahs
// @tag.description Surah management operations including creation, retrieval, updates, and deletion of Quran chapters

// @tag.name Ayahs
// @tag.description Ayah (verse) management operations including CRUD operations and retrieval by Surah or Juz

// @tag.name Juz
// @tag.description Juz (Para) management operations including creation, retrieval, updates, and deletion of Quran sections

// @externalDocs.description Safalife API Documentation
// @externalDocs.url https://docs.safalife.com/api

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"github.com/safalife/core-api/cli/core"
	_ "github.com/safalife/core-api/docs" // Import docs for swagger
	appContainer "github.com/safalife/core-api/src/application/container"
	domainContainer "github.com/safalife/core-api/src/domain/container"
	"github.com/safalife/core-api/src/infrastructure/configs"
	infraContainer "github.com/safalife/core-api/src/infrastructure/container"
	"github.com/safalife/core-api/src/infrastructure/monitoring"
	infraStorage "github.com/safalife/core-api/src/infrastructure/storage"
	presentationContainer "github.com/safalife/core-api/src/presentation/container"
	"github.com/safalife/core-api/src/presentation/routes"
)

func getEnvironment() string {
	env := os.Getenv("SAFALIFE_APP_ENV")
	if env == "" {
		env = "development"
	}
	return env
}

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Safalife API",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize configuration with the parsed config flag
		core.InitConfig(core.ConfigFlag)
		RunServer()
	},
}

func RunServer() {
	fmt.Println("\n┌─ Safalife API Server Initialization ────────────────────────────────────────┐")
	fmt.Println("│                                                                             │")
	fmt.Println("│  Starting Safalife API Server...                                            │")
	fmt.Println("│                                                                             │")
	fmt.Println("└─────────────────────────────────────────────────────────────────────────────┘")

	// Initialize database connection
	fmt.Print("\n[1/5] Initializing database connection... ")
	dbConfig := configs.DatabaseConfig{
		Host:     core.Config.DB.Host,
		User:     core.Config.DB.User,
		Password: core.Config.DB.Password,
		Name:     core.Config.DB.Name,
		Port:     core.Config.DB.Port,
		SSLMode:  core.Config.DB.SSLMode,
		TimeZone: core.Config.DB.TimeZone,
	}
	if err := configs.InitDatabase(dbConfig); err != nil {
		fmt.Println("FAILED")
		log.Fatalf("Failed to initialize database: %v", err)
	}
	fmt.Println("SUCCESS")

	// Initialize Redis connection
	fmt.Print("[2/5] Initializing Redis connection... ")
	redisConfig := configs.RedisConfig{
		Host:     core.Config.Redis.Host,
		Port:     core.Config.Redis.Port,
		Password: core.Config.Redis.Password,
		DB:       core.Config.Redis.DB,
	}
	if err := configs.InitRedis(redisConfig); err != nil {
		fmt.Println("FAILED")
		log.Fatalf("Failed to initialize Redis: %v", err)
	}
	fmt.Println("SUCCESS")

	// Get underlying sql.DB from GORM
	fmt.Print("[3/6] Setting up database adapter... ")
	gormDB := configs.GetDB()
	sqlDB, err := gormDB.DB()
	if err != nil {
		fmt.Println("FAILED")
		log.Fatalf("Failed to get underlying sql.DB: %v", err)
	}
	fmt.Println("SUCCESS")

	// Initialize monitoring service
	fmt.Print("[4/6] Initializing monitoring service... ")
	monitoringConfig := monitoring.MonitoringConfig{
		Jaeger: monitoring.JaegerConfig{
			Endpoint:       core.Config.Monitoring.Jaeger.Endpoint,
			ServiceName:    core.Config.Monitoring.Jaeger.ServiceName,
			ServiceVersion: core.Config.Monitoring.Jaeger.ServiceVersion,
			Environment:    core.Config.Monitoring.Jaeger.Environment,
		},
		Prometheus: monitoring.PrometheusConfig{
			Enabled: core.Config.Monitoring.Prometheus.Enabled,
		},
	}
	monitoringService, err := monitoring.NewMonitoringService(monitoringConfig)
	if err != nil {
		fmt.Println("FAILED")
		log.Fatalf("Failed to initialize monitoring service: %v", err)
	}
	fmt.Println("SUCCESS")

	// Setup dependency injection containers
	fmt.Print("[5/6] Initializing dependency containers... ")
	// Build S3 config from loaded app config
	s3Cfg := infraStorage.S3Config{
		Enabled:         core.Config.Storage.S3.Enabled,
		Bucket:          core.Config.Storage.S3.Bucket,
		Region:          core.Config.Storage.S3.Region,
		AccessKeyID:     core.Config.Storage.S3.AccessKeyID,
		SecretAccessKey: core.Config.Storage.S3.SecretAccessKey,
		Endpoint:        core.Config.Storage.S3.Endpoint,
		UsePathStyle:    core.Config.Storage.S3.UsePathStyle,
		PublicURLBase:   core.Config.Storage.S3.PublicURLBase,
	}

	infraCont := infraContainer.NewInfrastructureContainer(
		context.Background(),
		sqlDB,
		gormDB,
		configs.GetRedis(),
		monitoringService,
		core.Config.JWT.Secret,
		core.Config.JWT.AccessTokenTTL,
		core.Config.JWT.RefreshTokenTTL,
		core.Config.JWT.Issuer,
		s3Cfg,
	)
	domainCont := domainContainer.NewDomainContainer(
		infraCont.MonitoringService,
		infraCont.TransactionManager,
		infraCont.UUIDGenerator,
		infraCont.StorageService,
		infraCont.SurahRepository,
		infraCont.AyahRepository,
		infraCont.JuzRepository,
		infraCont.TranslationEditionRepository,
		infraCont.AyahTranslationRepository,
		infraCont.ReciterRepository,
		infraCont.AyahAudioFileRepository,
	)
	appCont := appContainer.NewApplicationContainer(domainCont)
	presentationCont := presentationContainer.NewPresentationContainer(appCont)
	fmt.Println("SUCCESS")

	// Setup Gin router
	fmt.Print("[6/6] Configuring HTTP router and routes... ")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, presentationCont)
	fmt.Println("SUCCESS")

	// Start HTTP server
	port := fmt.Sprintf("%d", core.Config.App.Port)
	fmt.Printf("\n┌─ Safalife API Server ────────────────────────────────────────────────────────┐\n")
	fmt.Printf("│                                                                              │\n")
	fmt.Printf("│  Server Address: http://localhost:%s                                       │\n", port)
	fmt.Printf("│  Environment: %s                                                    │\n", getEnvironment())
	fmt.Printf("│  Health Check: http://localhost:%s/v1/health                           │\n", port)
	fmt.Printf("│  API Documentation: http://localhost:%s/swagger/index.html                 │\n", port)
	fmt.Printf("│                                                                              │\n")
	fmt.Printf("│  Status: RUNNING                                                             │\n")
	fmt.Printf("│                                                                              │\n")
	fmt.Printf("└──────────────────────────────────────────────────────────────────────────────┘\n")
	go func() {
		if err := router.Run(":" + port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Setup graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("\n┌─ Graceful Shutdown ──────────────────────────────────────────────────────────┐")
		fmt.Println("│                                                                              │")
		fmt.Println("│  Initiating graceful shutdown...                                             │")
		fmt.Println("│                                                                              │")
		fmt.Println("└──────────────────────────────────────────────────────────────────────────────┘")

		// Close database connection
		fmt.Print("\n[1/2] Closing database connection... ")
		if err := configs.CloseDatabase(); err != nil {
			fmt.Println("ERROR")
			log.Printf("Error closing database: %v", err)
		} else {
			fmt.Println("SUCCESS")
		}

		// Close Redis connection
		fmt.Print("[2/2] Closing Redis connection... ")
		if err := configs.CloseRedis(); err != nil {
			fmt.Println("ERROR")
			log.Printf("Error closing Redis: %v", err)
		} else {
			fmt.Println("SUCCESS")
		}

		fmt.Println("\n┌─ Shutdown Complete ──────────────────────────────────────────────────────────┐")
		fmt.Println("│                                                                              │")
		fmt.Println("│  Safalife API Server has been stopped successfully.                          │")
		fmt.Println("│  Thank you for using Safalife API!                                           │")
		fmt.Println("│                                                                              │")
		fmt.Println("└──────────────────────────────────────────────────────────────────────────────┘")
		os.Exit(0)
	}()

	fmt.Println("\nPress Ctrl+C to stop the server")

	// Keep the server running
	select {}
}

package commands

import (
	"log"

	appContainer "github.com/achmdndy/safa-life-api/src/application/container"
	"github.com/achmdndy/safa-life-api/src/cmd/core"
	"github.com/achmdndy/safa-life-api/src/domain/shared"
	infraContainer "github.com/achmdndy/safa-life-api/src/infrastructure/container"
	"github.com/achmdndy/safa-life-api/src/infrastructure/monitoring"
	"github.com/achmdndy/safa-life-api/src/presentation/bootstrap"
	"github.com/achmdndy/safa-life-api/src/presentation/di"
	presentationMonitoring "github.com/achmdndy/safa-life-api/src/presentation/monitoring"
	"github.com/spf13/cobra"

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

// AppContainerFactoryWrapper wraps infrastructure factory to implement application ContainerFactory
type AppContainerFactoryWrapper struct {
	infraFactory shared.ContainerFactory
}

// CreateContainer implements the application ContainerFactory interface
func (w *AppContainerFactoryWrapper) CreateContainer(dbConfig shared.DatabaseConfig, redisConfig shared.RedisConfig) shared.Container {
	return w.infraFactory.CreateContainer(dbConfig, redisConfig)
}

// Ensure AppContainerFactoryWrapper implements appContainer.ContainerFactory
var _ appContainer.ContainerFactory = (*AppContainerFactoryWrapper)(nil)

func Start() {
	// Convert core.Config to bootstrap.Config
	config := bootstrap.Config{
		DB: struct {
			Host     string
			Port     int
			User     string
			Password string
			Name     string
			SSLMode  string
			TimeZone string
		}{
			Host:     core.Config.DB.Host,
			Port:     core.Config.DB.Port,
			User:     core.Config.DB.User,
			Password: core.Config.DB.Password,
			Name:     core.Config.DB.Name,
			SSLMode:  core.Config.DB.SSLMode,
			TimeZone: core.Config.DB.TimeZone,
		},
		Redis: struct {
			Host     string
			Port     int
			Password string
			DB       int
		}{
			Host:     core.Config.Redis.Host,
			Port:     core.Config.Redis.Port,
			Password: core.Config.Redis.Password,
			DB:       core.Config.Redis.DB,
		},
		Server: struct {
			Host string
			Port int
		}{
			Host: core.Config.Server.Host,
			Port: core.Config.Port,
		},
		Tracing: struct {
			ServiceName string
			Endpoint    string
		}{
			ServiceName: core.Config.Monitoring.Jaeger.ServiceName,
			Endpoint:    core.Config.Monitoring.Jaeger.Endpoint,
		},
	}

	// Proper Dependency Injection Chain: Infrastructure -> Application -> Presentation

	// Step 1: Create Infrastructure Factory (lowest layer)
	infraFactory := infraContainer.NewInfraContainerFactory()

	// Step 2: Wrap Infrastructure Factory to implement Application ContainerFactory interface
	appFactory := &AppContainerFactoryWrapper{
		infraFactory: infraFactory,
	}

	// Step 3: Create Jaeger tracing provider (infrastructure layer)
	jaegerProvider := monitoring.NewJaegerTracingProvider()

	// Step 3.6: Create presentation layer implementations with injected functions
	// This avoids direct imports from presentation to infrastructure/domain
	tracingProviderImpl := presentationMonitoring.NewConcreteTracingProvider(func(config presentationMonitoring.TracingConfig) (func(), error) {
		// Convert presentation config to domain config
		domainConfig := shared.TracingConfig{
			ServiceName: config.ServiceName,
			Endpoint:    config.Endpoint,
		}
		return jaegerProvider.Initialize(domainConfig)
	})

	// Step 4: Create Presentation DI with Application Factory (Clean Architecture compliant)
	presentationDI := di.NewPresentationDI(appFactory)

	// Step 5: Bootstrap with proper dependency injection chain including tracing
	if err := presentationDI.Bootstrap(config, tracingProviderImpl); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
}

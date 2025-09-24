package container

import (
	"github.com/achmdndy/safa-life-api/src/application/container"
	audioHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/audio"
	healthHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/health"
	quranHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/quran"
	reciterHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/reciter"
	resourceHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/resource"
	storyHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/story"
	tafsirHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/tafsir"
	tajweedHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/tajweed"
	topicHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/topic"
	translationHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/translation"
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
	HealthHandler      *healthHandler.Handler
	QuranHandler       *quranHandler.Handler
	TajweedHandler     *tajweedHandler.Handler
	TranslationHandler *translationHandler.Handler
	ReciterHandler     *reciterHandler.Handler
	AudioHandler       *audioHandler.Handler
	ResourceHandler    *resourceHandler.Handler
	StoryHandler       *storyHandler.Handler
	TafsirHandler      *tafsirHandler.Handler
	TopicHandler       *topicHandler.Handler

	// Routes (presentation layer specific)
	RouteConfig routes.RouteConfig
}

// NewPresentationContainer creates a new presentation container with configuration
func NewPresentationContainer(presentationFactory *PresentationContainerFactory, dbConfig DatabaseConfig, redisConfig RedisConfig) *PresentationContainer {
	// Get the application container factory from presentation factory
	appContainerFactory := presentationFactory.GetApplicationContainerFactory()
	
	// Create application container with configuration using factory
	appContainer := container.NewAppContainer(
		appContainerFactory,
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

	tajweedQueryHandler := c.appContainer.GetTajweedQueryHandler()
	tajweedCommandHandler := c.appContainer.GetTajweedCommandHandler()
	c.TajweedHandler = tajweedHandler.NewHandler(tajweedQueryHandler, tajweedCommandHandler)

	translationQueryHandler := c.appContainer.GetTranslationQueryHandler()
	translationCommandHandler := c.appContainer.GetTranslationCommandHandler()
	c.TranslationHandler = translationHandler.NewHandler(translationQueryHandler, translationCommandHandler)

	reciterQueryHandler := c.appContainer.GetReciterQueryHandler()
	reciterCommandHandler := c.appContainer.GetReciterCommandHandler()
	c.ReciterHandler = reciterHandler.NewHandler(reciterQueryHandler, reciterCommandHandler)

	audioQueryHandler := c.appContainer.GetAudioQueryHandler()
	audioCommandHandler := c.appContainer.GetAudioCommandHandler()
	c.AudioHandler = audioHandler.NewHandler(audioQueryHandler, audioCommandHandler)

	resourceQueryHandler := c.appContainer.GetResourceQueryHandler()
	resourceCommandHandler := c.appContainer.GetResourceCommandHandler()
	c.ResourceHandler = resourceHandler.NewHandler(resourceQueryHandler, resourceCommandHandler)

	storyQueryHandler := c.appContainer.GetStoryQueryHandler()
	storyCommandHandler := c.appContainer.GetStoryCommandHandler()
	c.StoryHandler = storyHandler.NewHandler(storyQueryHandler, storyCommandHandler)

	tafsirQueryHandler := c.appContainer.GetTafsirQueryHandler()
	tafsirCommandHandler := c.appContainer.GetTafsirCommandHandler()
	c.TafsirHandler = tafsirHandler.NewHandler(tafsirQueryHandler, tafsirCommandHandler)

	topicQueryHandler := c.appContainer.GetTopicQueryHandler()
	topicCommandHandler := c.appContainer.GetTopicCommandHandler()
	c.TopicHandler = topicHandler.NewHandler(topicQueryHandler, topicCommandHandler)
}

// initializeRoutes sets up route configuration
func (c *PresentationContainer) initializeRoutes() {
	c.RouteConfig = routes.RouteConfig{
		HealthHandler:      c.HealthHandler,
		QuranHandler:       c.QuranHandler,
		TajweedHandler:     c.TajweedHandler,
		TranslationHandler: c.TranslationHandler,
		ReciterHandler:     c.ReciterHandler,
		AudioHandler:       c.AudioHandler,
		ResourceHandler:    c.ResourceHandler,
		StoryHandler:       c.StoryHandler,
		TafsirHandler:      c.TafsirHandler,
		TopicHandler:       c.TopicHandler,
	}
}
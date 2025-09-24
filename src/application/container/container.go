package container

import (
	audioCommand "github.com/achmdndy/safa-life-api/src/application/audio/command"
	audioQuery "github.com/achmdndy/safa-life-api/src/application/audio/query"
	"github.com/achmdndy/safa-life-api/src/application/health/interfaces"
	healthQuery "github.com/achmdndy/safa-life-api/src/application/health/query"
	quranCommand "github.com/achmdndy/safa-life-api/src/application/quran/command"
	quranQuery "github.com/achmdndy/safa-life-api/src/application/quran/query"
	reciterCommand "github.com/achmdndy/safa-life-api/src/application/reciter/command"
	reciterQuery "github.com/achmdndy/safa-life-api/src/application/reciter/query"
	resourceCommand "github.com/achmdndy/safa-life-api/src/application/resource/command"
	resourceQuery "github.com/achmdndy/safa-life-api/src/application/resource/query"
	storyCommand "github.com/achmdndy/safa-life-api/src/application/story/command"
	storyQuery "github.com/achmdndy/safa-life-api/src/application/story/query"
	tafsirCommand "github.com/achmdndy/safa-life-api/src/application/tafsir/command"
	tafsirQuery "github.com/achmdndy/safa-life-api/src/application/tafsir/query"
	tajweedCommand "github.com/achmdndy/safa-life-api/src/application/tajweed/command"
	tajweedQuery "github.com/achmdndy/safa-life-api/src/application/tajweed/query"
	topicCommand "github.com/achmdndy/safa-life-api/src/application/topic/command"
	topicQuery "github.com/achmdndy/safa-life-api/src/application/topic/query"
	translationCommand "github.com/achmdndy/safa-life-api/src/application/translation/command"
	translationQuery "github.com/achmdndy/safa-life-api/src/application/translation/query"
	"github.com/achmdndy/safa-life-api/src/domain/audio"
	"github.com/achmdndy/safa-life-api/src/domain/health"
	"github.com/achmdndy/safa-life-api/src/domain/quran"
	"github.com/achmdndy/safa-life-api/src/domain/reciter"
	"github.com/achmdndy/safa-life-api/src/domain/resource"
	"github.com/achmdndy/safa-life-api/src/domain/shared"
	"github.com/achmdndy/safa-life-api/src/domain/story"
	"github.com/achmdndy/safa-life-api/src/domain/tafsir"
	"github.com/achmdndy/safa-life-api/src/domain/tajweed"
	"github.com/achmdndy/safa-life-api/src/domain/topic"
	"github.com/achmdndy/safa-life-api/src/domain/translation"
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

// ContainerFactory defines the interface for creating containers in application layer
// This interface is defined in application layer to avoid dependency on domain layer
type ContainerFactory interface {
	CreateContainer(dbConfig shared.DatabaseConfig, redisConfig shared.RedisConfig) shared.Container
}

// ApplicationContainer defines application-specific dependencies
// This interface is in application layer for presentation to use
type ApplicationContainer interface {
	shared.Container

	// Query Handler access - using application interface
	GetHealthQueryHandler() interfaces.QueryHandler
	GetAudioQueryHandler() *audioQuery.QueryHandler
	GetAudioCommandHandler() *audioCommand.CommandHandler
	GetQuranQueryHandler() *quranQuery.QueryHandler
	GetQuranCommandHandler() *quranCommand.CommandHandler
	GetReciterQueryHandler() *reciterQuery.QueryHandler
	GetReciterCommandHandler() *reciterCommand.CommandHandler
	GetResourceQueryHandler() *resourceQuery.QueryHandler
	GetResourceCommandHandler() *resourceCommand.CommandHandler
	GetStoryQueryHandler() *storyQuery.QueryHandler
	GetStoryCommandHandler() *storyCommand.CommandHandler
	GetTafsirQueryHandler() *tafsirQuery.QueryHandler
	GetTafsirCommandHandler() *tafsirCommand.CommandHandler
	GetTajweedQueryHandler() *tajweedQuery.QueryHandler
	GetTajweedCommandHandler() *tajweedCommand.CommandHandler
	GetTopicQueryHandler() *topicQuery.QueryHandler
	GetTopicCommandHandler() *topicCommand.CommandHandler
	GetTranslationQueryHandler() *translationQuery.QueryHandler
	GetTranslationCommandHandler() *translationCommand.CommandHandler
}

// AppContainer implements ApplicationContainer interface
// This is in application layer and uses dependency injection through domain interfaces
type AppContainer struct {
	// Domain container interface - no direct infrastructure dependency
	infraContainer shared.Container

	// Application-specific handlers
	getHealthQueryHandler     *healthQuery.GetHealthQueryHandler
	audioQueryHandler         *audioQuery.QueryHandler
	audioCommandHandler       *audioCommand.CommandHandler
	quranQueryHandler         *quranQuery.QueryHandler
	quranCommandHandler       *quranCommand.CommandHandler
	reciterQueryHandler       *reciterQuery.QueryHandler
	reciterCommandHandler     *reciterCommand.CommandHandler
	resourceQueryHandler      *resourceQuery.QueryHandler
	resourceCommandHandler    *resourceCommand.CommandHandler
	storyQueryHandler         *storyQuery.QueryHandler
	storyCommandHandler       *storyCommand.CommandHandler
	tafsirQueryHandler        *tafsirQuery.QueryHandler
	tafsirCommandHandler      *tafsirCommand.CommandHandler
	tajweedQueryHandler       *tajweedQuery.QueryHandler
	tajweedCommandHandler     *tajweedCommand.CommandHandler
	topicQueryHandler         *topicQuery.QueryHandler
	topicCommandHandler       *topicCommand.CommandHandler
	translationQueryHandler   *translationQuery.QueryHandler
	translationCommandHandler *translationCommand.CommandHandler
}

// NewAppContainer creates a new application container using ContainerFactory
// This follows Clean Architecture by using dependency injection through application interfaces
func NewAppContainer(containerFactory ContainerFactory, dbConfig DatabaseConfig, redisConfig RedisConfig) ApplicationContainer {
	// Create infrastructure container through application factory interface
	infraContainer := containerFactory.CreateContainer(
		shared.DatabaseConfig{
			Host:     dbConfig.Host,
			Port:     dbConfig.Port,
			User:     dbConfig.User,
			Password: dbConfig.Password,
			DBName:   dbConfig.DBName,
			SSLMode:  dbConfig.SSLMode,
			TimeZone: dbConfig.TimeZone,
		},
		shared.RedisConfig{
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

// GetQuranRepository returns the quran repository from infrastructure container
func (c *AppContainer) GetQuranRepository() quran.QuranRepository {
	return c.infraContainer.GetQuranRepository()
}

// GetReciterRepository returns the reciter repository from infrastructure container
func (c *AppContainer) GetReciterRepository() reciter.ReciterRepository {
	return c.infraContainer.GetReciterRepository()
}

// GetTajweedRepository returns the tajweed repository from infrastructure container
func (c *AppContainer) GetTajweedRepository() tajweed.TajweedRepository {
	return c.infraContainer.GetTajweedRepository()
}

// GetTranslationRepository returns the translation repository from infrastructure container
func (c *AppContainer) GetTranslationRepository() translation.FullTranslationRepository {
	return c.infraContainer.GetTranslationRepository()
}

// GetAudioRepository returns the audio repository from infrastructure container
func (c *AppContainer) GetAudioRepository() audio.AyahAudioRepository {
	return c.infraContainer.GetAudioRepository()
}

// GetResourceRepository returns the resource repository from infrastructure container
func (c *AppContainer) GetResourceRepository() resource.ResourceRepository {
	return c.infraContainer.GetResourceRepository()
}

// GetStoryRepository returns the story repository from infrastructure container
func (c *AppContainer) GetStoryRepository() story.FullStoryRepository {
	return c.infraContainer.GetStoryRepository()
}

// GetTafsirRepository returns the tafsir repository from infrastructure container
func (c *AppContainer) GetTafsirRepository() tafsir.FullTafsirRepository {
	return c.infraContainer.GetTafsirRepository()
}

// GetTopicRepository returns the topic repository from infrastructure container
func (c *AppContainer) GetTopicRepository() topic.FullTopicRepository {
	return c.infraContainer.GetTopicRepository()
}

// GetQuranService returns the quran service from infrastructure container
func (c *AppContainer) GetQuranService() quran.QuranService {
	return c.infraContainer.GetQuranService()
}

// GetReciterService returns the reciter service from infrastructure container
func (c *AppContainer) GetReciterService() reciter.ReciterService {
	return c.infraContainer.GetReciterService()
}

// GetTajweedService returns the tajweed service from infrastructure container
func (c *AppContainer) GetTajweedService() tajweed.TajweedService {
	return c.infraContainer.GetTajweedService()
}

// GetTranslationService returns the translation service from infrastructure container
func (c *AppContainer) GetTranslationService() translation.TranslationService {
	return c.infraContainer.GetTranslationService()
}

// GetAudioService returns the audio service from infrastructure container
func (c *AppContainer) GetAudioService() audio.AudioService {
	return c.infraContainer.GetAudioService()
}

// GetResourceService returns the resource service from infrastructure container
func (c *AppContainer) GetResourceService() resource.ResourceService {
	return c.infraContainer.GetResourceService()
}

// GetStoryService returns the story service from infrastructure container
func (c *AppContainer) GetStoryService() story.StoryService {
	return c.infraContainer.GetStoryService()
}

// GetTafsirService returns the tafsir service from infrastructure container
func (c *AppContainer) GetTafsirService() tafsir.TafsirService {
	return c.infraContainer.GetTafsirService()
}

// GetTopicService returns the topic service from infrastructure container
func (c *AppContainer) GetTopicService() topic.TopicService {
	return c.infraContainer.GetTopicService()
}

// GetQuranQueryHandler returns the quran query handler
func (c *AppContainer) GetQuranQueryHandler() *quranQuery.QueryHandler {
	return c.quranQueryHandler
}

// GetQuranCommandHandler returns the quran command handler
func (c *AppContainer) GetQuranCommandHandler() *quranCommand.CommandHandler {
	return c.quranCommandHandler
}

// GetReciterQueryHandler returns the reciter query handler
func (c *AppContainer) GetReciterQueryHandler() *reciterQuery.QueryHandler {
	return c.reciterQueryHandler
}

// GetReciterCommandHandler returns the reciter command handler
func (c *AppContainer) GetReciterCommandHandler() *reciterCommand.CommandHandler {
	return c.reciterCommandHandler
}

// GetTajweedQueryHandler returns the tajweed query handler
func (c *AppContainer) GetTajweedQueryHandler() *tajweedQuery.QueryHandler {
	return c.tajweedQueryHandler
}

// GetTajweedCommandHandler returns the tajweed command handler
func (c *AppContainer) GetTajweedCommandHandler() *tajweedCommand.CommandHandler {
	return c.tajweedCommandHandler
}

// GetTranslationQueryHandler returns the translation query handler
func (c *AppContainer) GetTranslationQueryHandler() *translationQuery.QueryHandler {
	return c.translationQueryHandler
}

// GetTranslationCommandHandler returns the translation command handler
func (c *AppContainer) GetTranslationCommandHandler() *translationCommand.CommandHandler {
	return c.translationCommandHandler
}

// GetAudioQueryHandler returns the audio query handler
func (c *AppContainer) GetAudioQueryHandler() *audioQuery.QueryHandler {
	return c.audioQueryHandler
}

// GetAudioCommandHandler returns the audio command handler
func (c *AppContainer) GetAudioCommandHandler() *audioCommand.CommandHandler {
	return c.audioCommandHandler
}

// GetResourceQueryHandler returns the resource query handler
func (c *AppContainer) GetResourceQueryHandler() *resourceQuery.QueryHandler {
	return c.resourceQueryHandler
}

// GetResourceCommandHandler returns the resource command handler
func (c *AppContainer) GetResourceCommandHandler() *resourceCommand.CommandHandler {
	return c.resourceCommandHandler
}

// GetStoryQueryHandler returns the story query handler
func (c *AppContainer) GetStoryQueryHandler() *storyQuery.QueryHandler {
	return c.storyQueryHandler
}

// GetStoryCommandHandler returns the story command handler
func (c *AppContainer) GetStoryCommandHandler() *storyCommand.CommandHandler {
	return c.storyCommandHandler
}

// GetTafsirQueryHandler returns the tafsir query handler
func (c *AppContainer) GetTafsirQueryHandler() *tafsirQuery.QueryHandler {
	return c.tafsirQueryHandler
}

// GetTafsirCommandHandler returns the tafsir command handler
func (c *AppContainer) GetTafsirCommandHandler() *tafsirCommand.CommandHandler {
	return c.tafsirCommandHandler
}

// GetTopicQueryHandler returns the topic query handler
func (c *AppContainer) GetTopicQueryHandler() *topicQuery.QueryHandler {
	return c.topicQueryHandler
}

// GetTopicCommandHandler returns the topic command handler
func (c *AppContainer) GetTopicCommandHandler() *topicCommand.CommandHandler {
	return c.topicCommandHandler
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
	c.getHealthQueryHandler = healthQuery.NewGetHealthQueryHandler(healthService)

	// Get quran service from infrastructure container
	quranService := c.infraContainer.GetQuranService()
	quranRepository := c.infraContainer.GetQuranRepository()

	// Create quran handlers
	c.quranQueryHandler = quranQuery.NewQueryHandler(quranService)
	c.quranCommandHandler = quranCommand.NewCommandHandler(quranService, quranRepository)

	// Get reciter service from infrastructure container
	reciterService := c.infraContainer.GetReciterService()
	reciterRepository := c.infraContainer.GetReciterRepository()

	// Create reciter handlers
	c.reciterQueryHandler = reciterQuery.NewQueryHandler(reciterService)
	c.reciterCommandHandler = reciterCommand.NewCommandHandler(reciterService, reciterRepository)

	// Get tajweed service from infrastructure container
	tajweedService := c.infraContainer.GetTajweedService()
	tajweedRepository := c.infraContainer.GetTajweedRepository()

	// Create tajweed handlers
	c.tajweedQueryHandler = tajweedQuery.NewQueryHandler(tajweedService)
	c.tajweedCommandHandler = tajweedCommand.NewCommandHandler(tajweedService, tajweedRepository)

	// Get translation service from infrastructure container
	translationService := c.infraContainer.GetTranslationService()
	translationRepository := c.infraContainer.GetTranslationRepository()

	// Create translation handlers
	c.translationQueryHandler = translationQuery.NewQueryHandler(translationService)
	c.translationCommandHandler = translationCommand.NewCommandHandler(translationService, translationRepository)

	// Get audio service from infrastructure container
	audioService := c.infraContainer.GetAudioService()
	audioRepository := c.infraContainer.GetAudioRepository()

	// Create audio handlers
	c.audioQueryHandler = audioQuery.NewQueryHandler(audioService)
	c.audioCommandHandler = audioCommand.NewCommandHandler(audioService, audioRepository)

	// Get resource service from infrastructure container
	resourceService := c.infraContainer.GetResourceService()
	resourceRepository := c.infraContainer.GetResourceRepository()

	// Create resource handlers
	c.resourceQueryHandler = resourceQuery.NewQueryHandler(resourceService)
	c.resourceCommandHandler = resourceCommand.NewCommandHandler(resourceService, resourceRepository)

	// Get story service from infrastructure container
	storyService := c.infraContainer.GetStoryService()
	storyRepository := c.infraContainer.GetStoryRepository()

	// Create story handlers
	c.storyQueryHandler = storyQuery.NewQueryHandler(storyService)
	c.storyCommandHandler = storyCommand.NewCommandHandler(storyService, storyRepository)

	// Get tafsir service from infrastructure container
	tafsirService := c.infraContainer.GetTafsirService()
	tafsirRepository := c.infraContainer.GetTafsirRepository()

	// Create tafsir handlers
	c.tafsirQueryHandler = tafsirQuery.NewQueryHandler(tafsirService)
	c.tafsirCommandHandler = tafsirCommand.NewCommandHandler(tafsirService, tafsirRepository)

	// Get topic service from infrastructure container
	topicService := c.infraContainer.GetTopicService()
	topicRepository := c.infraContainer.GetTopicRepository()

	// Create topic handlers
	c.topicQueryHandler = topicQuery.NewQueryHandler(topicService)
	c.topicCommandHandler = topicCommand.NewCommandHandler(topicService, topicRepository)
}

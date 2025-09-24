package container

import (
	"log"

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
	"github.com/achmdndy/safa-life-api/src/infrastructure/configs"
	audioInfra "github.com/achmdndy/safa-life-api/src/infrastructure/audio"
	healthInfra "github.com/achmdndy/safa-life-api/src/infrastructure/health"
	quranInfra "github.com/achmdndy/safa-life-api/src/infrastructure/quran"
	reciterInfra "github.com/achmdndy/safa-life-api/src/infrastructure/reciter"
	resourceInfra "github.com/achmdndy/safa-life-api/src/infrastructure/resource"
	storyInfra "github.com/achmdndy/safa-life-api/src/infrastructure/story"
	tafsirInfra "github.com/achmdndy/safa-life-api/src/infrastructure/tafsir"
	tajweedInfra "github.com/achmdndy/safa-life-api/src/infrastructure/tajweed"
	topicInfra "github.com/achmdndy/safa-life-api/src/infrastructure/topic"
	translationInfra "github.com/achmdndy/safa-life-api/src/infrastructure/translation"
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
	// Infrastructure components
	database *configs.Database
	redis    *configs.Redis

	// Repositories
	audioRepository       audio.AyahAudioRepository
	healthRepository      health.CheckerRepository
	quranRepository       quran.QuranRepository
	reciterRepository     reciter.ReciterRepository
	resourceRepository    resource.ResourceRepository
	storyRepository       story.FullStoryRepository
	tafsirRepository      tafsir.FullTafsirRepository
	tajweedRepository     tajweed.TajweedRepository
	topicRepository       topic.FullTopicRepository
	translationRepository translation.FullTranslationRepository

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

// GetAudioRepository returns the audio repository
func (c *InfraContainer) GetAudioRepository() audio.AyahAudioRepository {
	return c.audioRepository
}

// GetHealthRepository returns the health repository
func (c *InfraContainer) GetHealthRepository() health.CheckerRepository {
	return c.healthRepository
}

// GetQuranRepository returns the quran repository
func (c *InfraContainer) GetQuranRepository() quran.QuranRepository {
	return c.quranRepository
}

// GetReciterRepository returns the reciter repository
func (c *InfraContainer) GetReciterRepository() reciter.ReciterRepository {
	return c.reciterRepository
}

// GetResourceRepository returns the resource repository
func (c *InfraContainer) GetResourceRepository() resource.ResourceRepository {
	return c.resourceRepository
}

// GetStoryRepository returns the story repository
func (c *InfraContainer) GetStoryRepository() story.FullStoryRepository {
	return c.storyRepository
}

// GetTafsirRepository returns the tafsir repository
func (c *InfraContainer) GetTafsirRepository() tafsir.FullTafsirRepository {
	return c.tafsirRepository
}

// GetTajweedRepository returns the tajweed repository
func (c *InfraContainer) GetTajweedRepository() tajweed.TajweedRepository {
	return c.tajweedRepository
}

// GetTopicRepository returns the topic repository
func (c *InfraContainer) GetTopicRepository() topic.FullTopicRepository {
	return c.topicRepository
}

// GetTranslationRepository returns the translation repository
func (c *InfraContainer) GetTranslationRepository() translation.FullTranslationRepository {
	return c.translationRepository
}

// GetAudioService returns the audio service
func (c *InfraContainer) GetAudioService() audio.AudioService {
	// Create service with repository from infrastructure
	return audio.NewAudioService(c.audioRepository)
}

// GetHealthService returns the health service
func (c *InfraContainer) GetHealthService() *health.Service {
	// Create service with repository from infrastructure
	return health.NewService(c.healthRepository)
}

// GetQuranService returns the quran service
func (c *InfraContainer) GetQuranService() quran.QuranService {
	// Create service with repository from infrastructure
	return quran.NewQuranService(c.quranRepository)
}

// GetReciterService returns the reciter service
func (c *InfraContainer) GetReciterService() reciter.ReciterService {
	// Create service with repository from infrastructure
	return reciter.NewReciterService(c.reciterRepository)
}

// GetResourceService returns the resource service
func (c *InfraContainer) GetResourceService() resource.ResourceService {
	// Create service with repository from infrastructure
	return resource.NewResourceService(c.resourceRepository)
}

// GetStoryService returns the story service
func (c *InfraContainer) GetStoryService() story.StoryService {
	// Create service with repository from infrastructure
	return story.NewStoryService(c.storyRepository)
}

// GetTafsirService returns the tafsir service
func (c *InfraContainer) GetTafsirService() tafsir.TafsirService {
	// Create service with repository from infrastructure
	return tafsir.NewTafsirService(c.tafsirRepository)
}

// GetTajweedService returns the tajweed service
func (c *InfraContainer) GetTajweedService() tajweed.TajweedService {
	// Create service with repository from infrastructure
	return tajweed.NewTajweedService(c.tajweedRepository)
}

// GetTopicService returns the topic service
func (c *InfraContainer) GetTopicService() topic.TopicService {
	// Create service with repository from infrastructure
	return topic.NewTopicService(c.topicRepository)
}

// GetTranslationService returns the translation service
func (c *InfraContainer) GetTranslationService() translation.TranslationService {
	// Create service with repository from infrastructure
	return translation.NewTranslationService(c.translationRepository)
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
	c.audioRepository = audioInfra.NewAyahAudioRepository(c.database.DB)
	c.healthRepository = healthInfra.NewCheckerRepository(c.database, c.redis)
	c.quranRepository = quranInfra.NewQuranRepository(c.database.DB)
	c.reciterRepository = reciterInfra.NewReciterRepository(c.database.DB)
	c.resourceRepository = resourceInfra.NewResourceRepository(c.database.DB)
	c.storyRepository = storyInfra.NewStoryRepository(c.database.DB)
	c.tafsirRepository = tafsirInfra.NewTafsirRepository(c.database.DB)
	c.tajweedRepository = tajweedInfra.NewTajweedRepository(c.database.DB)
	c.topicRepository = topicInfra.NewTopicRepository(c.database.DB)
	c.translationRepository = translationInfra.NewTranslationRepository(c.database.DB)
}
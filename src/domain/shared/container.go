package shared

import (
	"github.com/achmdndy/safa-life-api/src/domain/audio"
	"github.com/achmdndy/safa-life-api/src/domain/health"
	"github.com/achmdndy/safa-life-api/src/domain/quran"
	"github.com/achmdndy/safa-life-api/src/domain/reciter"
	"github.com/achmdndy/safa-life-api/src/domain/resource"
	"github.com/achmdndy/safa-life-api/src/domain/story"
	"github.com/achmdndy/safa-life-api/src/domain/tafsir"
	"github.com/achmdndy/safa-life-api/src/domain/tajweed"
	"github.com/achmdndy/safa-life-api/src/domain/topic"
	"github.com/achmdndy/safa-life-api/src/domain/translation"
)

// Container defines the interface for dependency injection
// This interface is in domain layer to avoid circular dependencies
type Container interface {
	// Repository access
	GetHealthRepository() health.CheckerRepository
	GetQuranRepository() quran.QuranRepository
	GetReciterRepository() reciter.ReciterRepository
	GetTajweedRepository() tajweed.TajweedRepository
	GetTranslationRepository() translation.FullTranslationRepository
	GetAudioRepository() audio.AyahAudioRepository
	GetResourceRepository() resource.ResourceRepository
	GetStoryRepository() story.FullStoryRepository
	GetTafsirRepository() tafsir.FullTafsirRepository
	GetTopicRepository() topic.FullTopicRepository
	
	// Service access  
	GetHealthService() *health.Service
	GetQuranService() quran.QuranService
	GetReciterService() reciter.ReciterService
	GetTajweedService() tajweed.TajweedService
	GetTranslationService() translation.TranslationService
	GetAudioService() audio.AudioService
	GetResourceService() resource.ResourceService
	GetStoryService() story.StoryService
	GetTafsirService() tafsir.TafsirService
	GetTopicService() topic.TopicService
	
	// Lifecycle management
	Close() error
}

// QueryHandlerInterface defines the interface for query handlers
// This is in domain layer to avoid importing application layer
type QueryHandlerInterface interface {
	Handle(query interface{}) (interface{}, error)
}

// ContainerFactory defines the interface for creating containers
// This allows the application layer to create infrastructure containers
// without directly importing infrastructure packages
type ContainerFactory interface {
	CreateContainer(dbConfig DatabaseConfig, redisConfig RedisConfig) Container
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
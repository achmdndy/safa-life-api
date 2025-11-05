package container

import (
	"context"
	"database/sql"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/safalife/core-api/src/domain/core"
	"github.com/safalife/core-api/src/domain/quran"
	domStorage "github.com/safalife/core-api/src/domain/storage"
	infraCore "github.com/safalife/core-api/src/infrastructure/core"
	"github.com/safalife/core-api/src/infrastructure/monitoring"
	infraQuran "github.com/safalife/core-api/src/infrastructure/quran"
	infraStorage "github.com/safalife/core-api/src/infrastructure/storage"
)

// InfrastructureContainer holds infrastructure dependencies
type InfrastructureContainer struct {
	MonitoringService *monitoring.MonitoringService

	// Core infrastructure
	TransactionManager core.ContextTransactionManager
	UUIDGenerator      core.UUIDGenerator

	// Storage service
	StorageService domStorage.StorageServiceInterface

	// Quran repositories
	SurahRepository              quran.SurahRepositoryInterface
	AyahRepository               quran.AyahRepositoryInterface
	JuzRepository                quran.JuzRepositoryInterface
	TranslationEditionRepository quran.TranslationEditionRepositoryInterface
	AyahTranslationRepository    quran.AyahTranslationRepositoryInterface
	ReciterRepository            quran.ReciterRepositoryInterface
	AyahAudioFileRepository      quran.AyahAudioFileRepositoryInterface
	QuranRepository              quran.QuranRepositoryInterface
}

// NewInfrastructureContainer creates a new infrastructure container
func NewInfrastructureContainer(ctx context.Context, db *sql.DB, gormDB *gorm.DB, redisClient *redis.Client, monitoringService *monitoring.MonitoringService, jwtSecret string, accessTokenTTL int, refreshTokenTTL int, issuer string, s3Cfg infraStorage.S3Config) *InfrastructureContainer {
	// Initialize core infrastructure
	transactionManager := infraCore.NewGormContextTransactionManager()
	uuidGenerator := infraCore.NewUUIDGenerator()

	// Initialize Storage service (S3)
	var storageService domStorage.StorageServiceInterface
	if s3Cfg.Enabled {
		svc, err := infraStorage.NewS3StorageService(ctx, s3Cfg)
		if err == nil {
			storageService = svc
		}
	}

	// Initialize Quran infrastructure
	quranMapper := infraQuran.NewMapper()
	surahRepo := infraQuran.NewSurahRepository(gormDB, quranMapper)
	ayahRepo := infraQuran.NewAyahRepository(gormDB, quranMapper)
	juzRepo := infraQuran.NewJuzRepository(gormDB, quranMapper)
	translationEditionRepo := infraQuran.NewTranslationEditionRepository(gormDB, quranMapper)
	ayahTranslationRepo := infraQuran.NewAyahTranslationRepository(gormDB, quranMapper)
	reciterRepo := infraQuran.NewReciterRepository(gormDB, quranMapper)
	ayahAudioFileRepo := infraQuran.NewAyahAudioFileRepository(gormDB, quranMapper)
	quranRepo := quran.NewQuranRepository(surahRepo, ayahRepo, juzRepo, transactionManager)

	return &InfrastructureContainer{
		MonitoringService: monitoringService,

		// Core infrastructure
		TransactionManager: transactionManager,
		UUIDGenerator:      uuidGenerator,

		// Storage service
		StorageService: storageService,

		// Quran repositories
		SurahRepository:              surahRepo,
		AyahRepository:               ayahRepo,
		JuzRepository:                juzRepo,
		TranslationEditionRepository: translationEditionRepo,
		AyahTranslationRepository:    ayahTranslationRepo,
		ReciterRepository:            reciterRepo,
		AyahAudioFileRepository:      ayahAudioFileRepo,
		QuranRepository:              quranRepo,
	}
}

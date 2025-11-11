package container

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/safalife/core-api/src/domain/core"
	domPrayer "github.com/safalife/core-api/src/domain/prayertimes"
	"github.com/safalife/core-api/src/domain/quran"
	domStorage "github.com/safalife/core-api/src/domain/storage"
	infraConfigs "github.com/safalife/core-api/src/infrastructure/configs"
	infraCore "github.com/safalife/core-api/src/infrastructure/core"
	"github.com/safalife/core-api/src/infrastructure/monitoring"
	infraPrayer "github.com/safalife/core-api/src/infrastructure/prayertimes"
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

	// Prayer times repository
	PrayerTimesRepository domPrayer.PrayerTimesRepositoryInterface

	// Quran repositories
	SurahRepository              quran.SurahRepositoryInterface
	AyahRepository               quran.AyahRepositoryInterface
	JuzRepository                quran.JuzRepositoryInterface
	TranslationEditionRepository quran.TranslationEditionRepositoryInterface
	AyahTranslationRepository    quran.AyahTranslationRepositoryInterface
	ReciterRepository            quran.ReciterRepositoryInterface
	AyahAudioFileRepository      quran.AyahAudioFileRepositoryInterface
	BookmarkAyahRepository       quran.BookmarkAyahRepositoryInterface
	LastReadRepository           quran.LastReadRepositoryInterface
	ProgressHatamRepository      quran.ProgressHatamRepositoryInterface
	QuranRepository              quran.QuranRepositoryInterface
}

// NewInfrastructureContainer creates a new infrastructure container
func NewInfrastructureContainer(ctx context.Context, db *sql.DB, gormDB *gorm.DB, redisClient *redis.Client, monitoringService *monitoring.MonitoringService, jwtSecret string, accessTokenTTL int, refreshTokenTTL int, issuer string, s3Cfg infraStorage.S3Config, ptCfg infraConfigs.PrayerTimesProvidersConfig) *InfrastructureContainer {
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
	bookmarkAyahRepo := infraQuran.NewBookmarkAyahRepository(gormDB, quranMapper)
	lastReadRepo := infraQuran.NewLastReadRepository(gormDB, quranMapper)
	progressHatamRepo := infraQuran.NewProgressHatamRepository(gormDB, quranMapper)
	quranRepo := quran.NewQuranRepository(surahRepo, ayahRepo, juzRepo, transactionManager)

	// Initialize PrayerTimes providers and repository
	var aladhanHTTP *http.Client
	if ptCfg.AladhanTimeoutSeconds > 0 {
		aladhanHTTP = &http.Client{Timeout: time.Duration(ptCfg.AladhanTimeoutSeconds) * time.Second}
	} else {
		aladhanHTTP = &http.Client{Timeout: 15 * time.Second}
	}
	hablullahProvider := infraPrayer.NewHablullahProvider()
	aladhanProvider := infraPrayer.NewAladhanProvider(aladhanHTTP, ptCfg.AladhanBaseURL)
	prayerRepo := infraPrayer.NewCompositePrayerTimesRepository(hablullahProvider, aladhanProvider)

	return &InfrastructureContainer{
		MonitoringService: monitoringService,

		// Core infrastructure
		TransactionManager: transactionManager,
		UUIDGenerator:      uuidGenerator,

		// Storage service
		StorageService: storageService,

		// Prayer times repository
		PrayerTimesRepository: prayerRepo,

		// Quran repositories
		SurahRepository:              surahRepo,
		AyahRepository:               ayahRepo,
		JuzRepository:                juzRepo,
		TranslationEditionRepository: translationEditionRepo,
		AyahTranslationRepository:    ayahTranslationRepo,
		ReciterRepository:            reciterRepo,
		AyahAudioFileRepository:      ayahAudioFileRepo,
		BookmarkAyahRepository:       bookmarkAyahRepo,
		LastReadRepository:           lastReadRepo,
		ProgressHatamRepository:      progressHatamRepo,
		QuranRepository:              quranRepo,
	}
}

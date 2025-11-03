package container

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/safalife/core-api/src/domain/core"
	"github.com/safalife/core-api/src/domain/quran"
	infraCore "github.com/safalife/core-api/src/infrastructure/core"
	"github.com/safalife/core-api/src/infrastructure/monitoring"
	infraQuran "github.com/safalife/core-api/src/infrastructure/quran"
)

// InfrastructureContainer holds infrastructure dependencies
type InfrastructureContainer struct {
	MonitoringService *monitoring.MonitoringService

	// Core infrastructure
	TransactionManager core.ContextTransactionManager
	UUIDGenerator      core.UUIDGenerator

	// Quran repositories
	SurahRepository quran.SurahRepositoryInterface
	AyahRepository  quran.AyahRepositoryInterface
	JuzRepository   quran.JuzRepositoryInterface
	QuranRepository quran.QuranRepositoryInterface
}

// NewInfrastructureContainer creates a new infrastructure container
func NewInfrastructureContainer(db *sql.DB, gormDB *gorm.DB, redisClient *redis.Client, monitoringService *monitoring.MonitoringService, jwtSecret string, accessTokenTTL int, refreshTokenTTL int, issuer string) *InfrastructureContainer {
	// Initialize core infrastructure
	transactionManager := infraCore.NewGormContextTransactionManager()
	uuidGenerator := infraCore.NewUUIDGenerator()

	// Initialize Quran infrastructure
	quranMapper := infraQuran.NewMapper()
	surahRepo := infraQuran.NewSurahRepository(gormDB, quranMapper)
	ayahRepo := infraQuran.NewAyahRepository(gormDB, quranMapper)
	juzRepo := infraQuran.NewJuzRepository(gormDB, quranMapper)
	quranRepo := quran.NewQuranRepository(surahRepo, ayahRepo, juzRepo, transactionManager)

	return &InfrastructureContainer{
		MonitoringService: monitoringService,

		// Core infrastructure
		TransactionManager: transactionManager,
		UUIDGenerator:      uuidGenerator,

		// Quran repositories
		SurahRepository: surahRepo,
		AyahRepository:  ayahRepo,
		JuzRepository:   juzRepo,
		QuranRepository: quranRepo,
	}
}

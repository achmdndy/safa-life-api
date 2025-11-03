package configs

import (
	"fmt"
	"log"
	"time"

	"github.com/safalife/core-api/src/infrastructure/monitoring"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DatabaseConfig represents database configuration
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Port     int    `mapstructure:"port"`
	SSLMode  string `mapstructure:"ssl_mode"`
	TimeZone string `mapstructure:"timezone"`
}

var DB *gorm.DB

// CustomNamingStrategy preserves PascalCase for both table and column names
type CustomNamingStrategy struct {
	schema.NamingStrategy
}

// ColumnName preserves the original column name (PascalCase) with quotes
func (ns CustomNamingStrategy) ColumnName(table, column string) string {
	return `"` + column + `"`
}

// TableName preserves the original table name (PascalCase)
func (ns CustomNamingStrategy) TableName(table string) string {
	return table
}

// InitDatabase initializes the database connection using GORM
func InitDatabase(config DatabaseConfig) error {
	return InitDatabaseWithTracing(config, monitoring.DatabaseTracingConfig{
		ServiceName:    "safalife-api",
		ServiceVersion: "1.0.0",
		Environment:    "development",
		Enabled:        true,
		IncludeParams:  false,
	})
}

// InitDatabaseWithTracing initializes the database connection with tracing support
func InitDatabaseWithTracing(config DatabaseConfig, tracingConfig monitoring.DatabaseTracingConfig) error {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.Host,
		config.User,
		config.Password,
		config.Name,
		config.Port,
		config.SSLMode,
		config.TimeZone,
	)

	// Open database connection with custom naming strategy
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: CustomNamingStrategy{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	})

	// Enable debug mode to see SQL queries
	db = db.Debug()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Initialize GORM tracing
	if tracingErr := monitoring.InitGormTracing(db, tracingConfig); tracingErr != nil {
		log.Printf("Warning: Failed to initialize GORM tracing: %v", err)
		// Continue without tracing rather than failing completely
	} else {
		log.Println("GORM tracing initialized successfully")
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	DB = db
	log.Println("Database connected successfully")
	return nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}

// WithTransaction executes a function within a database transaction
func WithTransaction(fn func(tx *gorm.DB) error) error {
	return DB.Transaction(fn)
}

// CloseDatabase closes the database connection
func CloseDatabase() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return fmt.Errorf("failed to get underlying sql.DB: %w", err)
		}
		if err := sqlDB.Close(); err != nil {
			return fmt.Errorf("failed to close database: %w", err)
		}
		log.Println("Database connection closed")
	}
	return nil
}

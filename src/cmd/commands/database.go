package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/achmdndy/safa-life-api/src/cmd/core"
	"github.com/achmdndy/safa-life-api/src/infrastructure/configs"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// getDatabase creates a database connection using the current configuration
func getDatabase() (*gorm.DB, error) {
	dbConfig := configs.DatabaseConfig{
		Host:     core.Config.DB.Host,
		Port:     core.Config.DB.Port,
		User:     core.Config.DB.User,
		Password: core.Config.DB.Password,
		DBName:   core.Config.DB.Name,
		SSLMode:  core.Config.DB.SSLMode,
		TimeZone: core.Config.DB.TimeZone,
	}

	database, err := configs.NewDatabase(dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return database.GetDB(), nil
}

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔄 Migrating database...")

		db, err := getDatabase()
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		for _, model := range core.RegisterModels() {
			err := db.Debug().AutoMigrate(model.Model)
			if err != nil {
				log.Fatalf("Failed to migrate model %T: %v", model.Model, err)
			}
		}

		fmt.Println("✅ Database migrated successfully.")
	},
}

var MigrateFreshCmd = &cobra.Command{
	Use:   "migrate:fresh",
	Short: "Drop all tables and migrate the database fresh",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔄 Migrating database fresh (dropping all tables)...")

		db, err := getDatabase()
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		// Drop all tables first
		models := core.RegisterModels()
		for i := len(models) - 1; i >= 0; i-- {
			err := db.Migrator().DropTable(models[i].Model)
			if err != nil {
				log.Printf("Warning: Failed to drop table for model %T: %v", models[i].Model, err)
			}
		}

		// Recreate all tables
		for _, model := range models {
			err := db.Debug().AutoMigrate(model.Model)
			if err != nil {
				log.Fatalf("Failed to migrate model %T: %v", model.Model, err)
			}
		}

		fmt.Println("✅ Database migrated fresh successfully.")
	},
}

var DBSeedCmd = &cobra.Command{
	Use:   "db:seed",
	Short: "Seed the database with initial data",
	Long: `Seed the database with initial data.

Examples:
  # Seed all models
  safa-life-api db:seed

  # Seed specific model
  safa-life-api db:seed --model quran

  # Force seed (override existing data)
  safa-life-api db:seed --force

  # Seed specific model with force
  safa-life-api db:seed --model quran --force`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🌱 Seeding database...")

		db, err := getDatabase()
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		// Get flags
		model, _ := cmd.Flags().GetString("model")
		force, _ := cmd.Flags().GetBool("force")

		// Initialize seeder registry
		registry := core.NewSeederRegistry()

		// Check if data already exists (unless force is used)
		if !force {
			var count int64
			db.Table("surahs").Count(&count)
			if count > 0 {
				fmt.Println("⚠️  Database already contains data. Use --force to override.")
				return
			}
		}

		// If force is used, wipe existing data first
		if force {
			fmt.Println("🗑️  Force flag detected, wiping existing data...")
			models := core.RegisterModels()
			for _, modelInfo := range models {
				if err := db.Unscoped().Where("1 = 1").Delete(modelInfo.Model).Error; err != nil {
					log.Printf("Warning: Failed to delete data from %T: %v", modelInfo.Model, err)
				}
			}
			fmt.Println("✅ Existing data wiped.")
		}

		// Run seeders
		if model != "" {
			// Seed specific model
			if err := registry.SeedSpecific(db, model); err != nil {
				log.Fatalf("Failed to seed %s: %v", model, err)
			}
		} else {
			// Seed all models
			if err := registry.SeedAll(db); err != nil {
				log.Fatalf("Failed to seed database: %v", err)
			}
		}

		fmt.Println("✅ Database seeded successfully.")
	},
}

var DBWipeCmd = &cobra.Command{
	Use:   "db:wipe",
	Short: "Wipe all data from the database (truncate tables)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🧹 Wiping database...")

		db, err := getDatabase()
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		// Disable foreign key checks temporarily
		db.Exec("SET FOREIGN_KEY_CHECKS = 0")

		// Truncate all tables
		models := core.RegisterModels()
		for _, model := range models {
			tableName := db.NamingStrategy.TableName(fmt.Sprintf("%T", model.Model))
			if tableName != "" {
				err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", tableName)).Error
				if err != nil {
					// Try DELETE if TRUNCATE fails (some databases don't support TRUNCATE)
					err = db.Exec(fmt.Sprintf("DELETE FROM %s", tableName)).Error
					if err != nil {
						log.Printf("Warning: Failed to wipe table %s: %v", tableName, err)
					}
				}
			}
		}

		// Re-enable foreign key checks
		db.Exec("SET FOREIGN_KEY_CHECKS = 1")

		fmt.Println("✅ Database wiped successfully.")
	},
}

var SchemaDumpCmd = &cobra.Command{
	Use:   "schema:dump",
	Short: "Dump the database schema to a SQL file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📤 Dumping database schema...")

		// Create dumps directory if it doesn't exist
		dumpsDir := "dumps"
		if err := os.MkdirAll(dumpsDir, 0755); err != nil {
			log.Fatalf("Failed to create dumps directory: %v", err)
		}

		// Generate filename with timestamp
		timestamp := time.Now().Format("20060102_150405")
		filename := filepath.Join(dumpsDir, fmt.Sprintf("schema_%s.sql", timestamp))

		// Use pg_dump to dump schema only
		dumpCmd := exec.Command("pg_dump",
			"-h", core.Config.DB.Host,
			"-p", fmt.Sprintf("%d", core.Config.DB.Port),
			"-U", core.Config.DB.User,
			"-d", core.Config.DB.Name,
			"--schema-only",
			"--no-owner",
			"--no-privileges",
			"-f", filename,
		)

		// Set password via environment variable
		dumpCmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", core.Config.DB.Password))

		if err := dumpCmd.Run(); err != nil {
			log.Fatalf("Failed to dump schema: %v", err)
		}

		fmt.Printf("✅ Schema dumped to: %s\n", filename)
	},
}

var SchemaRestoreCmd = &cobra.Command{
	Use:   "schema:restore [file]",
	Short: "Restore the database schema from a SQL file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		fmt.Printf("📥 Restoring database schema from: %s\n", filename)

		// Check if file exists
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			log.Fatalf("Schema file not found: %s", filename)
		}

		// Use psql to restore schema
		restoreCmd := exec.Command("psql",
			"-h", core.Config.DB.Host,
			"-p", fmt.Sprintf("%d", core.Config.DB.Port),
			"-U", core.Config.DB.User,
			"-d", core.Config.DB.Name,
			"-f", filename,
		)

		// Set password via environment variable
		restoreCmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", core.Config.DB.Password))

		if err := restoreCmd.Run(); err != nil {
			log.Fatalf("Failed to restore schema: %v", err)
		}

		fmt.Println("✅ Schema restored successfully.")
	},
}

var DBBackupCmd = &cobra.Command{
	Use:   "db:backup",
	Short: "Create a full backup of the database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("💾 Creating database backup...")

		// Create backups directory if it doesn't exist
		backupsDir := "backups"
		if err := os.MkdirAll(backupsDir, 0755); err != nil {
			log.Fatalf("Failed to create backups directory: %v", err)
		}

		// Generate filename with timestamp
		timestamp := time.Now().Format("20060102_150405")
		filename := filepath.Join(backupsDir, fmt.Sprintf("backup_%s.sql", timestamp))

		// Use pg_dump to create full backup
		backupCmd := exec.Command("pg_dump",
			"-h", core.Config.DB.Host,
			"-p", fmt.Sprintf("%d", core.Config.DB.Port),
			"-U", core.Config.DB.User,
			"-d", core.Config.DB.Name,
			"--no-owner",
			"--no-privileges",
			"-f", filename,
		)

		// Set password via environment variable
		backupCmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", core.Config.DB.Password))

		if err := backupCmd.Run(); err != nil {
			log.Fatalf("Failed to create backup: %v", err)
		}

		fmt.Printf("✅ Database backup created: %s\n", filename)
	},
}

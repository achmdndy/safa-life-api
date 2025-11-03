package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/safalife/core-api/cli/core"
	"github.com/safalife/core-api/cli/seeders"
	"github.com/safalife/core-api/src/infrastructure/configs"
	"github.com/safalife/core-api/src/infrastructure/quran"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// Seeder interface for database seeders
type Seeder interface {
	GetName() string
	Seed(db *gorm.DB, force bool) error
}

// registerModels returns all models that need to be registered for migration
func registerModels() []interface{} {
	return []interface{}{
		// Quran models
		&quran.SurahModel{},
		&quran.AyahModel{},
		&quran.JuzModel{},
	}
}

// SurahSeeder implements the Seeder interface for Surah data
type SurahSeeder struct{}

func (s *SurahSeeder) GetName() string {
	return "quran"
}

func (s *SurahSeeder) Seed(db *gorm.DB, force bool) error {
	seeder := &seeders.SurahSeeder{}
	return seeder.SeedWithForce(db, force)
}

// JuzSeeder implements the Seeder interface for Juz data
type JuzSeeder struct{}

func (j *JuzSeeder) GetName() string {
	return "juz"
}

func (j *JuzSeeder) Seed(db *gorm.DB, force bool) error {
	seeder := &seeders.JuzSeeder{}
	return seeder.SeedWithForce(db, force)
}

func getDatabase() (*gorm.DB, error) {
	// Check if database is already initialized
	if db := configs.GetDB(); db != nil {
		return db, nil
	}

	// Initialize database connection
	dbConfig := configs.DatabaseConfig{
		Host:     core.Config.DB.Host,
		User:     core.Config.DB.User,
		Password: core.Config.DB.Password,
		Name:     core.Config.DB.Name,
		Port:     core.Config.DB.Port,
		SSLMode:  core.Config.DB.SSLMode,
		TimeZone: core.Config.DB.TimeZone,
	}

	if err := configs.InitDatabase(dbConfig); err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	// Get the initialized database
	db := configs.GetDB()
	if db == nil {
		return nil, fmt.Errorf("database connection is nil after initialization")
	}

	return db, nil
}

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize configuration with the parsed config flag
		core.InitConfig(core.ConfigFlag)
		fmt.Println("Migrating database...")

		// Get database connection
		db, err := getDatabase()
		if err != nil {
			fmt.Printf("Error connecting to database: %v\n", err)
			os.Exit(1)
		}

		// Get all models to migrate
		models := registerModels()

		// Run AutoMigrate for all models
		if err := db.AutoMigrate(models...); err != nil {
			fmt.Printf("Error migrating database: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Database migrated successfully. %d models processed.\n", len(models))
	},
}

var MigrateFreshCmd = &cobra.Command{
	Use:   "migrate:fresh",
	Short: "Drop all tables and migrate the database fresh",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize configuration with the parsed config flag
		core.InitConfig(core.ConfigFlag)
		fmt.Println("Migrating database fresh (dropping all tables)...")

		// Get database connection
		db, err := getDatabase()
		if err != nil {
			fmt.Printf("Error connecting to database: %v\n", err)
			os.Exit(1)
		}

		// Get all models to migrate
		models := registerModels()

		// Drop all tables first
		fmt.Println("Dropping existing tables...")
		for _, model := range models {
			if err := db.Migrator().DropTable(model); err != nil {
				fmt.Printf("Warning: Could not drop table for model %T: %v\n", model, err)
			}
		}

		// Run AutoMigrate for all models
		fmt.Println("Creating fresh tables...")
		if err := db.AutoMigrate(models...); err != nil {
			fmt.Printf("Error migrating database: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Database migrated fresh successfully. %d models processed.\n", len(models))
	},
}

var DBSeedCmd = &cobra.Command{
	Use:   "db:seed",
	Short: "Seed the database with sample data",
	Long: `Seed the database with sample data.

Examples:
  safalife db:seed                    # Seed all models
  safalife db:seed -m quran           # Seed only Quran data
  safalife db:seed --force            # Force seed (wipe existing data)`,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize configuration with the parsed config flag
		core.InitConfig(core.ConfigFlag)
		fmt.Println("Seeding database...")

		db, err := getDatabase()
		if err != nil {
			fmt.Printf("Error connecting to database: %v\n", err)
			os.Exit(1)
		}

		// Get flags
		model, _ := cmd.Flags().GetString("model")
		force, _ := cmd.Flags().GetBool("force")

		// Initialize available seeders
		seeders := map[string]Seeder{
			"quran": &SurahSeeder{},
			"juz":   &JuzSeeder{},
		}

		// Seed based on model flag
		if model == "" {
			// Seed all models
			fmt.Println("Seeding all available models...")
			for name, seeder := range seeders {
				fmt.Printf("Seeding %s...\n", name)
				if err := seeder.Seed(db, force); err != nil {
					if force {
						fmt.Printf("Warning: Error seeding %s (continuing due to --force): %v\n", name, err)
					} else {
						fmt.Printf("Error seeding %s: %v\n", name, err)
						os.Exit(1)
					}
				} else {
					fmt.Printf("✓ %s seeded successfully\n", name)
				}
			}
		} else {
			// Seed specific model
			seeder, exists := seeders[model]
			if !exists {
				fmt.Printf("Unknown model: %s. Available models: ", model)
				for name := range seeders {
					fmt.Printf("%s ", name)
				}
				fmt.Println()
				os.Exit(1)
			}

			fmt.Printf("Seeding %s...\n", model)
			if err := seeder.Seed(db, force); err != nil {
				fmt.Printf("Error seeding %s: %v\n", model, err)
				os.Exit(1)
			}
			fmt.Printf("✓ %s seeded successfully\n", model)
		}

		fmt.Println("Database seeded successfully.")
	},
}

var DBWipeCmd = &cobra.Command{
	Use:   "db:wipe",
	Short: "Wipe all data from the database (truncate tables)",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize configuration with the parsed config flag
		core.InitConfig(core.ConfigFlag)
		fmt.Println("🧹 Wiping database...")

		db, err := getDatabase()
		if err != nil {
			fmt.Printf("Failed to connect to database: %v\n", err)
			return
		}

		// Disable foreign key checks temporarily
		db.Exec("SET FOREIGN_KEY_CHECKS = 0")

		// Truncate all tables
		models := registerModels()
		for _, model := range models {
			tableName := db.NamingStrategy.TableName(fmt.Sprintf("%T", model))
			if tableName != "" {
				err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", tableName)).Error
				if err != nil {
					// Try DELETE if TRUNCATE fails (some databases don't support TRUNCATE)
					err = db.Exec(fmt.Sprintf("DELETE FROM %s", tableName)).Error
					if err != nil {
						fmt.Printf("Warning: Failed to wipe table %s: %v\n", tableName, err)
					}
				}
			}
		}

		// Re-enable foreign key checks
		db.Exec("SET FOREIGN_KEY_CHECKS = 1")

		fmt.Println("Database wiped successfully.")
	},
}

var SchemaDumpCmd = &cobra.Command{
	Use:   "schema:dump",
	Short: "Dump the database schema to a SQL file",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize configuration with the parsed config flag
		core.InitConfig(core.ConfigFlag)
		fmt.Println("Dumping database schema...")

		// Create dumps directory if it doesn't exist
		dumpsDir := "dumps"
		if err := os.MkdirAll(dumpsDir, 0755); err != nil {
			fmt.Printf("Failed to create dumps directory: %v\n", err)
			return
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
			fmt.Printf("Failed to dump schema: %v\n", err)
			return
		}

		fmt.Printf("Schema dumped to: %s\n", filename)
	},
}

var SchemaRestoreCmd = &cobra.Command{
	Use:   "schema:restore [file]",
	Short: "Restore the database schema from a SQL file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize configuration with the parsed config flag
		core.InitConfig(core.ConfigFlag)
		filename := args[0]
		fmt.Printf("Restoring database schema from: %s\n", filename)

		// Check if file exists
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			fmt.Printf("Schema file not found: %s\n", filename)
			return
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
			fmt.Printf("Failed to restore schema: %v\n", err)
			return
		}

		fmt.Println("Schema restored successfully.")
	},
}

var DBBackupCmd = &cobra.Command{
	Use:   "db:backup",
	Short: "Create a full backup of the database",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize configuration with the parsed config flag
		core.InitConfig(core.ConfigFlag)
		fmt.Println("Creating database backup...")

		// Create backups directory if it doesn't exist
		backupsDir := "backups"
		if err := os.MkdirAll(backupsDir, 0755); err != nil {
			fmt.Printf("Failed to create backups directory: %v\n", err)
			return
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
			fmt.Printf("Failed to create backup: %v\n", err)
			return
		}

		fmt.Printf("Database backup created: %s\n", filename)
	},
}

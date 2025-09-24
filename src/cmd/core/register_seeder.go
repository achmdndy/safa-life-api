package core

import (
	"fmt"
	"gorm.io/gorm"
	quran "github.com/achmdndy/safa-life-api/src/infrastructure/quran"
)

// Seeder interface defines the contract for all seeders
type Seeder interface {
	Seed(db *gorm.DB) error
	GetName() string
}

// QuranSeeder implements the Seeder interface for Quran data
type QuranSeeder struct {
	name string
}

// NewQuranSeeder creates a new QuranSeeder instance
func NewQuranSeeder() *QuranSeeder {
	return &QuranSeeder{
		name: "quran",
	}
}

// GetName returns the seeder name
func (qs *QuranSeeder) GetName() string {
	return qs.name
}

// Seed implements the seeding logic for Quran data
func (qs *QuranSeeder) Seed(db *gorm.DB) error {
	fmt.Println("🌱 Seeding Quran data...")
	
	// Initialize faker
	faker := quran.NewQuranFaker()
	
	// Step 1: Generate and seed Surahs first (they have no dependencies)
	fmt.Printf("   📖 Generating and seeding Surahs...\n")
	surahs := faker.SurahFaker.GenerateBatch(20)
	if err := db.Create(&surahs).Error; err != nil {
		return fmt.Errorf("failed to seed surahs: %w", err)
	}
	fmt.Printf("   ✅ Seeded %d Surahs\n", len(surahs))
	
	// Step 2: Generate and seed Juzs with valid Surah references
	fmt.Printf("   📚 Generating and seeding Juzs...\n")
	juzs := faker.JuzFaker.GenerateBatchWithValidSurahs(30, surahs)
	if err := db.Create(&juzs).Error; err != nil {
		return fmt.Errorf("failed to seed juzs: %w", err)
	}
	fmt.Printf("   ✅ Seeded %d Juzs\n", len(juzs))
	
	// Step 3: Generate and seed Ayahs with valid Surah and Juz references
	fmt.Printf("   📝 Generating and seeding Ayahs...\n")
	ayahs := faker.AyahFaker.GenerateForAllSurahsWithValidJuz(surahs, juzs)
	
	// Seed Ayahs in batches to avoid memory issues
	batchSize := 1000
	for i := 0; i < len(ayahs); i += batchSize {
		end := i + batchSize
		if end > len(ayahs) {
			end = len(ayahs)
		}
		
		batch := ayahs[i:end]
		if err := db.Create(&batch).Error; err != nil {
			return fmt.Errorf("failed to seed ayahs batch %d-%d: %w", i, end, err)
		}
		fmt.Printf("   ✅ Seeded Ayahs batch %d-%d\n", i+1, end)
	}
	
	fmt.Println("✅ Quran data seeded successfully!")
	return nil
}

// SeederRegistry manages all available seeders
type SeederRegistry struct {
	seeders map[string]Seeder
}

// NewSeederRegistry creates a new SeederRegistry instance
func NewSeederRegistry() *SeederRegistry {
	registry := &SeederRegistry{
		seeders: make(map[string]Seeder),
	}
	
	// Register all seeders
	registry.Register(NewQuranSeeder())
	
	return registry
}

// Register adds a seeder to the registry
func (sr *SeederRegistry) Register(seeder Seeder) {
	sr.seeders[seeder.GetName()] = seeder
}

// GetSeeder returns a seeder by name
func (sr *SeederRegistry) GetSeeder(name string) (Seeder, bool) {
	seeder, exists := sr.seeders[name]
	return seeder, exists
}

// GetAllSeeders returns all registered seeders
func (sr *SeederRegistry) GetAllSeeders() map[string]Seeder {
	return sr.seeders
}

// GetSeederNames returns all available seeder names
func (sr *SeederRegistry) GetSeederNames() []string {
	names := make([]string, 0, len(sr.seeders))
	for name := range sr.seeders {
		names = append(names, name)
	}
	return names
}

// SeedAll runs all registered seeders
func (sr *SeederRegistry) SeedAll(db *gorm.DB) error {
	fmt.Println("🌱 Running all seeders...")
	
	for name, seeder := range sr.seeders {
		fmt.Printf("🔄 Running seeder: %s\n", name)
		if err := seeder.Seed(db); err != nil {
			return fmt.Errorf("seeder %s failed: %w", name, err)
		}
	}
	
	fmt.Println("✅ All seeders completed successfully!")
	return nil
}

// SeedSpecific runs a specific seeder by name
func (sr *SeederRegistry) SeedSpecific(db *gorm.DB, name string) error {
	seeder, exists := sr.GetSeeder(name)
	if !exists {
		return fmt.Errorf("seeder '%s' not found. Available seeders: %v", name, sr.GetSeederNames())
	}
	
	fmt.Printf("🌱 Running specific seeder: %s\n", name)
	if err := seeder.Seed(db); err != nil {
		return fmt.Errorf("seeder %s failed: %w", name, err)
	}
	
	fmt.Printf("✅ Seeder %s completed successfully!\n", name)
	return nil
}
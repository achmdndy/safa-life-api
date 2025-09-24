package main

import (
	"log"

	"github.com/achmdndy/safa-life-api/src/cmd/commands"
	"github.com/achmdndy/safa-life-api/src/cmd/core"
)

func main() {
	core.RegisterCommand(commands.StartCmd)
	core.RegisterCommand(commands.MigrateCmd)
	core.RegisterCommand(commands.MigrateFreshCmd)
	
	// Add flags to DBSeedCmd
	commands.DBSeedCmd.Flags().StringP("model", "m", "", "Specify which model to seed (e.g., 'quran')")
	commands.DBSeedCmd.Flags().BoolP("force", "f", false, "Force seed by wiping existing data first")
	core.RegisterCommand(commands.DBSeedCmd)
	
	core.RegisterCommand(commands.DBWipeCmd)
	core.RegisterCommand(commands.SchemaDumpCmd)
	core.RegisterCommand(commands.SchemaRestoreCmd)
	core.RegisterCommand(commands.DBBackupCmd)
	core.InitConfig(core.ConfigFlag)

	if err := core.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}

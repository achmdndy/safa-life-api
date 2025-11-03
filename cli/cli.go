package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/safalife/core-api/cli/commands"
	"github.com/safalife/core-api/cli/core"
)

func printBanner() {
	banner := `
╔══════════════════════════════════════════════════════════════════╗
║                                                                  ║
║   ███████╗ █████╗ ███████╗ █████╗ ██╗     ██╗███████╗███████╗    ║
║   ██╔════╝██╔══██╗██╔════╝██╔══██╗██║     ██║██╔════╝██╔════╝    ║
║   ███████╗███████║█████╗  ███████║██║     ██║█████╗  █████╗      ║
║   ╚════██║██╔══██║██╔══╝  ██╔══██║██║     ██║██╔══╝  ██╔══╝      ║
║   ███████║██║  ██║██║     ██║  ██║███████╗██║██║     ███████╗    ║
║   ╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝  ╚═╝╚══════╝╚═╝╚═╝     ╚══════╝    ║
║                                                                  ║
║                       Safalife API                               ║
║            High-Performance Go Backend Service                   ║
║                                                                  ║
╚══════════════════════════════════════════════════════════════════╝
`
	fmt.Print(banner)
}

func printSystemInfo() {
	fmt.Println("\n┌─ System Information ─────────────────────────────────────────────────────────┐")
	fmt.Printf("│ Version: v1.0.0                                                              │\n")
	fmt.Printf("│ Environment: %s                                                           │\n", getEnvironment())
	fmt.Printf("│ Go Version: %s                                                            │\n", getGoVersion())
	fmt.Println("└──────────────────────────────────────────────────────────────────────────────┘")
}

func getEnvironment() string {
	env := os.Getenv("SAFALIFE_APP_ENV")
	if env == "" {
		env = "development"
	}
	return env
}

func getGoVersion() string {
	// This would typically come from runtime.Version() but keeping it simple
	return "1.21+"
}

func printAvailableCommands() {
	fmt.Println("\n┌─ Available Commands ─────────────────────────────────────────────────────────┐")
	fmt.Println("│                                                                              │")
	fmt.Println("│  start     Start the Safalife API server                                    │")
	fmt.Println("│  help      Show help information                                            │")
	fmt.Println("│                                                                              │")
	fmt.Println("│  Examples:                                                                   │")
	fmt.Println("│    ./safalife start                    # Start with default config          │")
	fmt.Println("│    ./safalife start --config prod      # Start with production config       │")
	fmt.Println("│    ./safalife help                     # Show detailed help                 │")
	fmt.Println("│                                                                              │")
	fmt.Println("└──────────────────────────────────────────────────────────────────────────────┘")
}

func RunCmd() {
	// Show banner and info only if no arguments provided
	if len(os.Args) == 1 {
		printBanner()
		printSystemInfo()
		printAvailableCommands()
		fmt.Println("\nFor more information, run: ./safalife help")
		return
	}

	// Register commands
	core.RegisterCmd(commands.StartCmd)
	core.RegisterCmd(commands.MigrateCmd)
	core.RegisterCmd(commands.MigrateFreshCmd)
	core.RegisterCmd(commands.DBSeedCmd)
	core.RegisterCmd(commands.DBWipeCmd)
	core.RegisterCmd(commands.SchemaDumpCmd)
	core.RegisterCmd(commands.SchemaRestoreCmd)
	core.RegisterCmd(commands.DBBackupCmd)

	// Initialize flags for commands
	commands.DBSeedCmd.Flags().StringP("model", "m", "", "Specific model to seed (e.g., users, articles)")
	commands.DBSeedCmd.Flags().BoolP("force", "f", false, "Force seed (override existing data)")
	commands.DBSeedCmd.Flags().String("file", "", "Seed from CSV file (location, area, company) or specify path (type:path)")

	// Execute command (config will be initialized in each command handler)
	if err := core.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}

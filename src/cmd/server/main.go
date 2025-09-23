package main

import (
	"log"

	"github.com/achmdndy/safa-life-api/src/cmd/commands"
	"github.com/achmdndy/safa-life-api/src/cmd/core"
)

func main() {
	core.RegisterCommand(commands.StartCmd)
	core.InitConfig(core.ConfigFlag)

	if err := core.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}

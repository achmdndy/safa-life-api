package core

import (
	"github.com/spf13/cobra"
)

var ConfigFlag string

var rootCmd = &cobra.Command{
	Use:   "safa-life-api",
	Short: "Safa Life API",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&ConfigFlag, "config", "", "set config file (config.yaml, config.development.yaml, config.production.yaml)")
}

func RegisterCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

package cmd

import (
	"fmt"

	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/operations"
	"github.com/spf13/cobra"
)

var prepfilePath string

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&prepfilePath, "file", "f", "Prepfile.toml", "Path to the configuration file")
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run operations using Prepfile",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	prepfile, err := config.LoadConfig(prepfilePath)
	if err != nil {
		fmt.Printf("Failed to load config file '%s': %s\n", prepfilePath, err)
		return
	}

	operations.DispatchOperations(prepfile)
}

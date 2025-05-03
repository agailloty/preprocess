package cmd

import (
	"fmt"

	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/operations"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run operations using Prefile",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	prepfile, err := config.LoadConfig("Prepfile.toml")
	if err != nil {
		fmt.Printf("Prepfile.toml not found %s", err)
	}

	operations.DispatchOperations(prepfile)
}

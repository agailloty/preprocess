package cmd

import (
	"github.com/agailloty/preprocess/dataset"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(printCmd)
}

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Display all or part of the dataset on the console",
	Run:   print,
}

func print(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		guessedTypes := dataset.ReadDatasetColumns(args[0], ",")
		for _, dt := range guessedTypes {
			dataset.DisplayColumn(dt, 5)
			println()
		}
	}
}

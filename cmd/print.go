package cmd

import (
	"fmt"

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
		data := dataset.ReadAllLines(args[0], ",")
		guessedTypes := dataset.ReadDatasetColumns(data)

		for _, dt := range guessedTypes {
			fmt.Printf("%T \n", dt)
		}
	}
}

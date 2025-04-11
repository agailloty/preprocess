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
			displayColumn(dt, 10)
			println()
		}
	}
}

func displayColumn(column dataset.DataSetColumn, n int) {
	switch v := column.(type) {
	case dataset.Float:
		fmt.Printf("%s (Float) \n", v.Name)
		for i := range n {
			fmt.Printf("%.2f ", v.Data[i])
		}
	case dataset.String:
		fmt.Printf("%s (Float)", v.Name)
		for i := range n {
			fmt.Printf("%s", v.Data[i])
		}
	case dataset.Integer:
		fmt.Printf("%s (Float)", v.Name)
		for i := range n {
			fmt.Printf("%d", v.Data[i])
		}
	}
}

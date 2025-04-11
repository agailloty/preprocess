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
			displayColumn(dt, 5)
		}
	}
}

func displayColumn(column dataset.DataSetColumn, n int) {
	switch v := column.(type) {
	case dataset.Float:
		fmt.Printf("%s (Float) \n", v.Name)
		for i := range n {
			fmt.Println(v.Data[i])
		}
	case dataset.String:
		fmt.Printf("%s (String) \n", v.Name)
		for i := range n {
			fmt.Println(v.Data[i])
		}
	case dataset.Integer:
		fmt.Printf("%s (Integer) \n", v.Name)
		for i := range n {
			fmt.Println(v.Data[i])
		}
	}
}

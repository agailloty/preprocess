package cmd

import (
	"github.com/agailloty/preprocess/dataset"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Display all or part of the dataset on the console",
	Run:   print,
}

func print(cmd *cobra.Command, args []string) {
	dataframe := dataset.ReadDataFrame(datafilename, separator, decimalSeparator, encoding)
	for _, dt := range dataframe.Columns {
		dataset.DisplayColumn(dt, 5)
		println()
	}
}

func init() {
	rootCmd.AddCommand(printCmd)
	printCmd.Run = print
	printCmd.Flags().StringVarP(&datafilename, "data", "d", "", "Path to the dataset")
	printCmd.Flags().StringVarP(&separator, "sep", "s", ",", "Separator for csv file")
	printCmd.Flags().StringVarP(&decimalSeparator, "dsep", "m", ",", "Decimal separator")
	printCmd.Flags().StringVarP(&encoding, "encoding", "e", "utf-8", "Character encoding")
}

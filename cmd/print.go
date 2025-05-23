package cmd

import (
	"github.com/agailloty/preprocess/common"
	"github.com/agailloty/preprocess/dataset"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Display all or part of the dataset on the console",
	Run:   print,
}

func print(cmd *cobra.Command, args []string) {
	dfSpecs := common.DataSpecs{
		Filename:         datafilename,
		CsvSeparator:     csvseparator,
		DecimalSeparator: decimalSeparator,
		Encoding:         encoding,
	}
	dataframe := dataset.ReadDataFrame(dfSpecs)
	for _, dt := range dataframe.Columns {
		dataset.DisplayColumn(dt, 5)
		println()
	}
}

func init() {
	rootCmd.AddCommand(printCmd)
	printCmd.Run = print
	setDataSpecFlags(printCmd)
}

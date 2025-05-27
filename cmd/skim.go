package cmd

import (
	"github.com/agailloty/preprocess/common"
	"github.com/agailloty/preprocess/skim"
	"github.com/spf13/cobra"
)

var skimCmd = &cobra.Command{
	Use:   "skim",
	Short: "Print a subset of the dataset in the console",
	Long:  "Print a subset of the dataset in the console",
	Run:   skimDf,
}

func init() {
	rootCmd.AddCommand(skimCmd)
	skimCmd.Run = skimDf
	setDataSpecFlags(skimCmd)
}

func skimDf(cmd *cobra.Command, args []string) {
	dfSpec := common.DataSpecs{
		Filename:         datafilename,
		CsvSeparator:     csvseparator,
		DecimalSeparator: decimalSeparator,
		Encoding:         encoding,
	}

	skim.SkimDf(dfSpec)
}

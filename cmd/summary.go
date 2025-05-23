package cmd

import (
	"fmt"

	"github.com/agailloty/preprocess/common"
	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/summary"
	"github.com/spf13/cobra"
)

var summaryOutput string

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Display dataset summary statistics",
	Run:   summarizeDataset,
}

func summarizeDataset(cmd *cobra.Command, args []string) {
	prepfile, err := config.LoadConfigFromPrepfile(prepfilePath)
	if err != nil {
		fmt.Printf("Failed to load config file '%s': %s\n", prepfilePath, err)
		return
	}

	if prepfile != nil {
		dataframe := dataset.ReadDataFrame(prepfile.Data)

		summary.Summarize(dataframe, summaryOutput)
	} else {
		if decimalSeparator == "" {
			decimalSeparator = ","
		}
		dataSpec := common.DataSpecs{
			Filename:         datafilename,
			CsvSeparator:     csvseparator,
			DecimalSeparator: decimalSeparator,
			Encoding:         encoding,
		}
		dataframe := dataset.ReadDataFrame(dataSpec)
		if summaryOutput == "" {
			summaryOutput = "Summaryfile.toml"
		}
		summary.Summarize(dataframe, summaryOutput)
	}
}

func init() {
	rootCmd.AddCommand(summaryCmd)
	summaryCmd.Run = summarizeDataset
	setDataSpecFlags(summaryCmd)
	summaryCmd.Flags().StringVarP(&summaryOutput, "output", "o", "Summaryfile.toml", "Output name for Summaryfile")
}

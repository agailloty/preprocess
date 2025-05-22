package cmd

import (
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/summary"
	"github.com/spf13/cobra"
)

var datasetSummary string
var separatorSummary string
var outputSummary string
var decimalSeparatorSummary string

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Display dataset summary statistics",
	Run:   summarizeDataset,
}

func summarizeDataset(cmd *cobra.Command, args []string) {
	if decimalSeparatorSummary == "" {
		decimalSeparatorSummary = ","
	}
	dataframe := dataset.ReadDataFrame(datasetSummary, separatorSummary, decimalSeparatorSummary, encoding)
	if outputSummary == "" {
		outputSummary = "Summaryfile.toml"
	}
	summary.Summarize(dataframe, outputSummary)
}

func init() {
	rootCmd.AddCommand(summaryCmd)
	summaryCmd.Run = summarizeDataset
	summaryCmd.Flags().StringVarP(&datasetSummary, "data", "d", "", "Path to the dataset")
	summaryCmd.Flags().StringVarP(&separatorSummary, "sep", "s", ",", "Separator for csv file")
	summaryCmd.Flags().StringVarP(&decimalSeparatorSummary, "dsep", "m", ",", "Decimal separator")
	summaryCmd.Flags().StringVarP(&encoding, "enc", "e", "utf-8", "Character encoding")
	summaryCmd.Flags().StringVarP(&outputSummary, "output", "o", "Summaryfile.toml", "Output name for Summaryfile")
}

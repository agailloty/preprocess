package cmd

import (
	"fmt"

	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/summary"
	"github.com/spf13/cobra"
)

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
		dataframe := dataset.ReadDataFrame(prepfile.Data.File,
			prepfile.Data.Separator,
			prepfile.Data.DecimalSeparator,
			prepfile.Data.Encoding)

		summary.Summarize(dataframe, output)
	} else {
		if decimalSeparator == "" {
			decimalSeparator = ","
		}
		dataframe := dataset.ReadDataFrame(datasetPath, separator, decimalSeparator, encoding)
		if output == "" {
			output = "Summaryfile.toml"
		}
		summary.Summarize(dataframe, output)
	}
}

func init() {
	rootCmd.AddCommand(summaryCmd)
	summaryCmd.Run = summarizeDataset
	summaryCmd.Flags().StringVarP(&datasetPath, "data", "d", "", "Path to the dataset")
	summaryCmd.Flags().StringVarP(&separator, "sep", "s", ",", "Separator for csv file")
	summaryCmd.Flags().StringVarP(&decimalSeparator, "dsep", "m", ",", "Decimal separator")
	summaryCmd.Flags().StringVarP(&encoding, "enc", "e", "utf-8", "Character encoding")
	summaryCmd.Flags().StringVarP(&output, "output", "o", "Summaryfile.toml", "Output name for Summaryfile")
}

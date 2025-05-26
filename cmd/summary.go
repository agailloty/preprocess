package cmd

import (
	"fmt"
	"time"

	"github.com/agailloty/preprocess/common"
	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/summary"
	"github.com/agailloty/preprocess/utils"
	"github.com/spf13/cobra"
)

var summaryOutput string
var makeHtml bool

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Display dataset summary statistics",
	Run:   summarizeDataset,
}

func summarizeDataset(cmd *cobra.Command, args []string) {
	start := time.Now()
	prepfile, err := config.LoadConfigFromPrepfile(prepfilePath)
	if err != nil {
		fmt.Printf("Failed to load config file '%s': %s\n", prepfilePath, err)
		return
	}

	var summaryFile summary.SummaryFile

	if prepfile != nil {
		dataframe := dataset.ReadDataFrame(prepfile.Data)

		summaryFile = summary.GetSummaryFile(dataframe)
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
		summaryFile = summary.GetSummaryFile(dataframe)
	}

	utils.SerializeStruct(summaryFile, summaryOutput)
	if makeHtml {
		summary.SummaryHtml(summaryFile, "report.html")
	}

	elapsed := time.Since(start)
	fmt.Printf("Finished in : %s\n", elapsed)
}

func init() {
	rootCmd.AddCommand(summaryCmd)
	summaryCmd.Run = summarizeDataset
	setDataSpecFlags(summaryCmd)
	setPrepfileFlag(summaryCmd)
	summaryCmd.Flags().StringVarP(&summaryOutput, "output", "o", "Summaryfile.toml", "Output name for Summaryfile")
	summaryCmd.Flags().BoolVarP(&makeHtml, "html", "t", false, "Generate HTML file")
}

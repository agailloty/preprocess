package cmd

import (
	"log"
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
var excludedColumns []string

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: SUMMARY_SHORT_DESCRIPTION,
	Long:  SUMMARY_LONG_DESCRIPTION,
	Run:   summarizeDataset,
}

func summarizeDataset(cmd *cobra.Command, args []string) {
	start := time.Now()
	var (
		summaryFile summary.SummaryFile
		dataframe   dataset.DataFrame
		err         error
	)

	getOutputFile := func() string {
		if summaryOutput == "" {
			return "Summaryfile.toml"
		}
		return summaryOutput
	}

	switch {
	case datafilename != "":
		dfSpec := common.DataSpecs{
			Filename:          datafilename,
			CsvSeparator:      csvseparator,
			DecimalSeparator:  decimalSeparator,
			Encoding:          encoding,
			MissingIdentifier: "",
		}
		dataframe = dataset.ReadDataFrame(dfSpec)
	case prepfilePath != "":
		var prepfile *config.Prepfile
		prepfile, err = config.LoadConfigFromPrepfile(prepfilePath)
		if err != nil {
			log.Printf("Failed to load config file '%s': %s\n", prepfilePath, err)
			return
		}
		if prepfile == nil {
			log.Println("Prepfile is nil.")
			return
		}
		dataframe = dataset.ReadDataFrame(prepfile.Data)
	default:
		log.Println("No data source specified. Please provide a data file or a prepfile.")
		return
	}

	summaryFile = summary.GetSummaryFile(dataframe, excludedColumns)

	outputFile := getOutputFile()
	utils.SerializeStruct(summaryFile, outputFile)

	if makeHtml {
		summary.SummaryHtml(summaryFile, "report.html")
	}

	log.Printf("Finished in : %s\n", time.Since(start))
}

func init() {
	rootCmd.AddCommand(summaryCmd)
	summaryCmd.Run = summarizeDataset
	setDataSpecFlags(summaryCmd)
	setPrepfileFlag(summaryCmd)
	summaryCmd.Flags().StringVarP(&summaryOutput, "output", "o", "Summaryfile.toml", "Output name for Summaryfile")
	summaryCmd.Flags().BoolVarP(&makeHtml, "html", "t", false, "Generate HTML file")
	summaryCmd.Flags().StringArrayVar(&excludedColumns, "exclude", excludedColumns, "Exclude columns from summary")
}

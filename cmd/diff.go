package cmd

import (
	"log"
	"time"

	"github.com/agailloty/preprocess/common"
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/summary"
	"github.com/spf13/cobra"
)

var sourcedata string
var targetdata string

func init() {
	rootCmd.AddCommand(diffCmd)
	diffCmd.Run = computeDiff
	diffCmd.Flags().StringVar(&sourcedata, "source", "", "Path to the source dataset")
	diffCmd.Flags().StringVar(&targetdata, "target", "", "Path to the target dataset")
	diffCmd.Flags().BoolVarP(&makeHtml, "html", "t", false, "Generate HTML file")
	diffCmd.Flags().BoolVarP(&nobrowser, "no-browser", "b", false, "Do not launch the browser")
	diffCmd.Flags().StringVarP(&csvseparator, "sep", "s", ",", "Separator for csv file")
	diffCmd.Flags().StringVarP(&decimalSeparator, "dsep", "m", ".", "Decimal separator")
}

var diffCmd = &cobra.Command{
	Use: "diff",
}

func computeDiff(cmd *cobra.Command, args []string) {
	start := time.Now()
	source := dataset.ReadDataFrame(common.DataSpecs{Filename: sourcedata,
		CsvSeparator:     csvseparator,
		DecimalSeparator: decimalSeparator})

	target := dataset.ReadDataFrame(common.DataSpecs{Filename: targetdata,
		CsvSeparator:     csvseparator,
		DecimalSeparator: decimalSeparator})
	diff := summary.GenerateDiffSummary(&source, &target)
	summary.DiffHtml(diff, "htmldiff.html")
	elapsed := time.Since(start)
	log.Printf("Finished preprocessing in : %s\n", elapsed)
}

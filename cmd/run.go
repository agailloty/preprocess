package cmd

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/agailloty/preprocess/common"
	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/operations"
	"github.com/agailloty/preprocess/summary"
	"github.com/agailloty/preprocess/utils"
	"github.com/spf13/cobra"
)

var columnList []string
var operationList []string
var numerics bool
var showDiff bool

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Run = run
	setDataSpecFlags(runCmd)
	setPrepfileFlag(runCmd)
	runCmd.Flags().StringArrayVar(&columnList, "column", []string{}, "Target column(s) for preprocessing")
	runCmd.Flags().StringArrayVar(&operationList, "op", []string{}, "Preprocessing operation(s) (e.g., fillna:method=mean)")
	runCmd.Flags().BoolVar(&numerics, "numerics", false, "Apply operations only on numeric columns")
	runCmd.Flags().BoolVar(&showDiff, "show-diff", false, "Show a summary of the differences in a HTML page")
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: RUN_SHORT_DESCRIPTION,
	Long:  RUN_LONG_DESCRIPTION,
}

func run(cmd *cobra.Command, args []string) {
	start := time.Now()
	validateFlags()
	elapsed := time.Since(start)
	log.Printf("Finished preprocessing in : %s\n", elapsed)
}

func validateFlags() {
	isFileProvided := runCmd.Flags().Changed("file")
	isDataProvided := runCmd.Flags().Changed("data")
	isColumnListProvided := runCmd.Flags().Changed("column")
	isOperationListProvided := runCmd.Flags().Changed("op")

	var alteredDf dataset.DataFrame

	providedNFlags := btoi(isFileProvided) +
		btoi(isDataProvided) +
		btoi(isColumnListProvided) +
		btoi(isOperationListProvided)

	// if no flag is provided or only --file, then run using Prepfile
	if providedNFlags == 0 || isFileProvided {
		prepfile, err := config.LoadConfigFromPrepfile(prepfilePath)
		if err != nil {
			log.Fatalf("Failed to load config file '%s': %s\n", prepfilePath, err)
		}

		alteredDf = operations.DispatchOperations(prepfile)
	}

	if isDataProvided && isColumnListProvided && isOperationListProvided {
		dfSpecs := common.DataSpecs{Filename: datafilename, CsvSeparator: csvseparator, DecimalSeparator: decimalSeparator, Encoding: encoding}
		prepfile := config.MakeConfigFromCommandsArgs(dfSpecs, columnList, operationList)
		alteredDf = operations.DispatchOperations(prepfile)
	}

	if showDiff {
		oldFileName := alteredDf.DataSpecs.Filename
		if _, err := os.Stat(oldFileName); errors.Is(err, os.ErrNotExist) {
			log.Printf("Cannot compute diff because %s no longer exists", oldFileName)
			return
		}
		baseDf := dataset.ReadDataFrame(alteredDf.DataSpecs)
		diffName := baseDf.Name + "_diff.html"
		diffSum := summary.GenerateDiffSummary(&baseDf, &alteredDf)
		summary.DiffHtml(diffSum, diffName)
		utils.OpenBrowser(diffName)
	}

}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

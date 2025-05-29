package cmd

import (
	"log"
	"time"

	"github.com/agailloty/preprocess/common"
	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/operations"
	"github.com/spf13/cobra"
)

var columnList []string
var operationList []string
var numerics bool

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Run = run
	setDataSpecFlags(runCmd)
	setPrepfileFlag(runCmd)
	runCmd.Flags().StringArrayVar(&columnList, "column", []string{}, "Target column(s) for preprocessing")
	runCmd.Flags().StringArrayVar(&operationList, "op", []string{}, "Preprocessing operation(s) (e.g., fillna:method=mean)")
	runCmd.Flags().BoolVar(&numerics, "numerics", false, "Apply operations only on numeric columns")
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

		operations.DispatchOperations(prepfile)
	}

	if isDataProvided && isColumnListProvided && isOperationListProvided {
		dfSpecs := common.DataSpecs{Filename: datafilename, CsvSeparator: csvseparator, DecimalSeparator: decimalSeparator, Encoding: encoding}
		prepfile := config.MakeConfigFromCommandsArgs(dfSpecs, columnList, operationList)
		operations.DispatchOperations(prepfile)
	}

}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

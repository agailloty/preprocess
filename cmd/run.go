package cmd

import (
	"fmt"

	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/operations"
	"github.com/spf13/cobra"
)

var prepfilePath string
var datasetPath string
var columnList []string
var operationList []string
var numerics bool

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Run = run
	runCmd.Flags().StringVarP(&prepfilePath, "file", "f", "Prepfile.toml", "Path to the configuration file")
	runCmd.Flags().StringVarP(&datasetPath, "data", "d", "", "Path to the dataset file")
	runCmd.Flags().StringArrayVar(&columnList, "column", []string{}, "Target column(s) for preprocessing")
	runCmd.Flags().StringArrayVar(&operationList, "op", []string{}, "Preprocessing operation(s) (e.g., fillna:method=mean)")
	runCmd.Flags().BoolVar(&numerics, "numerics", false, "Apply operations only on numeric columns")
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run operations using Prepfile",
}

func run(cmd *cobra.Command, args []string) {
	validateFlags()
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
			fmt.Printf("Failed to load config file '%s': %s\n", prepfilePath, err)
			return
		}

		operations.DispatchOperations(prepfile)
	}

}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

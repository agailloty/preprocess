package operations

import (
	"fmt"

	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
)

func DispatchOperations(prepfile *config.Config) {
	df := dataset.ReadDataFrame(prepfile.Data.File, prepfile.Data.Separator)
	fmt.Printf("Successfully read dataset %s \n", prepfile.Data.File)

	for _, col := range df.Columns {
		if prepfile.Data.NumericOperations != nil {
			applyOperationsOnNumericColumns(&df, prepfile.Data.NumericOperations.Preprocess)
		}
		// If there are column specific operation
		found, columnConfig := findColumnConfig(prepfile.Data.Columns, col.GetName())
		if found {
			preprocessOps := columnConfig.Preprocess
			if col.GetType() == "int" || col.GetType() == "float" {
				applyNumericOperationsOnColumn(preprocessOps, col, &df)
			} else if col.GetType() == "string" {
				applyTextOperationsOnColumn(preprocessOps, col, &df)
			}
		}
		RenameColumn(col, prepfile.Data.Columns)
	}

	if prepfile.PostProcess.SortDataset != nil {
		SortDatasetColumns(df, prepfile.PostProcess.SortDataset.Descending)
	}

	ExportCsv(df, prepfile.PostProcess.FileName)
}

func findColumnConfig(columns []config.ColumnConfig, name string) (found bool, result config.ColumnConfig) {
	for _, value := range columns {
		if value.Name == name {
			found = true
			result = value
			break
		}
	}
	return found, result
}

func applyNumericOperationsOnColumn(preprocessOps *[]config.PreprocessOp, col dataset.DataSetColumn, df *dataset.DataFrame) {
	if preprocessOps != nil {
		for _, prep := range *preprocessOps {
			if prep.Op == "fillna" && prep.Method == "" && prep.Value != "" {
				replaceMissingValues(col, prep.Value)
			} else if prep.Op == "fillna" && prep.Method != "" {
				replaceMissingWithStats(col, prep.Method)
			}

			// Transform operation come after filling missing values
			if prep.Op == "normalize" {
				if prep.Method == "zscore" {
					applyZScoreToEveryElement(col, df)
				} else if prep.Method == "minmax" {
					applyMinMaxScoreToEveryElement(col)
				}
			}

			if prep.Op == "discretize" && prep.Method == "binning" && prep.Bins != nil {
				makeBinsFromNumericColumns(col, prep.Bins, df, true)
			}
		}
	}
}

func applyTextOperationsOnColumn(preprocessOps *[]config.PreprocessOp, col dataset.DataSetColumn, df *dataset.DataFrame) {
	if preprocessOps != nil {
		for _, prep := range *preprocessOps {
			if prep.Op == "fillna" && prep.Method == "" && prep.Value != "" {
				replaceMissingValues(col, prep.Value)
			}
		}
	}
}

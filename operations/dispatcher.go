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

		found, columnConfig := findColumnConfig(prepfile.Data.Columns, col.GetName())
		if found {
			preprocessOps := columnConfig.Preprocess
			if preprocessOps != nil {
				for _, prep := range *preprocessOps {
					if prep.Op == "fillna" && prep.Method == "" && prep.Value != "" {
						replaceMissingValues(col, prep.Value)
					} else if prep.Op == "fillna" && prep.Method != "" {
						replaceMissingWithStats(col, prep.Method)
					}

					// Transform operation come after filling missing values
					if prep.Op == "normalize" && prep.Method == "zscore" {
						applyZScoreToEveryElement(col)
					}

				}

			}
		}
		RenameColumn(col, prepfile.Data.Columns)
	}

	if prepfile.PostProcess.SortDataset != nil {
		SortDatasetColumns(df, prepfile.PostProcess.SortDataset.Descending)
	}

	ExportCsv(df, prepfile.PostProcess.Export)
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

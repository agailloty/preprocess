package operations

import (
	"fmt"

	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
)

func DispatchOperations(prepfile *config.Prepfile) {
	df := dataset.ReadDataFrame(prepfile.Data.DataSpecs)
	fmt.Printf("Successfully read dataset %s \n", prepfile.Data.DataSpecs.Filename)

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
				applyTextOperationsOnColumn(preprocessOps, col)
			}
		}
		RenameColumn(col, prepfile.Data.Columns)
	}

	if prepfile.PostProcess.DropColumns != nil {
		for _, columnToDelete := range *prepfile.PostProcess.DropColumns {
			df.DeleteColumnByName(columnToDelete)
		}
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

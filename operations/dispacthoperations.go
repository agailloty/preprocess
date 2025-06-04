package operations

import (
	"log"

	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
)

func DispatchOperations(prepfile *config.Prepfile) {
	df := dataset.ReadDataFrame(prepfile.Data)
	log.Printf("Successfully read dataset %s \n", prepfile.Data.Filename)

	operationStats := summarizeOperations(prepfile, &df)

	if prepfile.Preprocess.NumericOperations != nil {
		excluded := prepfile.Preprocess.NumericOperations.ExcludeCols
		dispatchDatasetNumericOperations(&df, prepfile.Preprocess.NumericOperations.Operations, excluded)
	}

	if prepfile.Preprocess.TextOperations != nil {
		applyOperationsOnTextColumns(&df, prepfile.Preprocess.TextOperations.Operations)
	}

	for _, col := range df.Columns {
		// If there are column specific operation
		found, columnConfig := findColumnConfig(prepfile.Preprocess.Columns, col.GetName())
		if found {
			preprocessOps := columnConfig.Operations
			if col.GetType() == "int" || col.GetType() == "float" {
				dispatchColumnNumericOperations(&df, col, preprocessOps, []string{})
			} else if col.GetType() == "string" {
				applyTextOperationsOnColumn(&df, preprocessOps, col)
			}
		}
		RenameColumn(col, prepfile.Preprocess.Columns)
	}

	if prepfile.PostProcess.DropColumns != nil {
		for _, columnToDelete := range *prepfile.PostProcess.DropColumns {
			df.DeleteColumnByName(columnToDelete)
		}
	}

	if prepfile.PostProcess.SortDataset != nil {
		SortDatasetColumns(df, prepfile.PostProcess.SortDataset.Descending)
	}

	if prepfile.PostProcess.DataSetSplit != nil {
		operateSplit(&df, prepfile.PostProcess.DataSetSplit)
	}
	ExportCsv(df, prepfile.PostProcess.FileName)

	operationStats.logOperationStats()
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

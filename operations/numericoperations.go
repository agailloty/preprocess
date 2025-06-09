package operations

import (
	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/statistics"
	"github.com/agailloty/preprocess/utils"
)

func applySingleNumericOperation(df *dataset.DataFrame, operation config.PreprocessOp, col dataset.DataSetColumn) {
	if operation.Op == OP_FILLNA && operation.Method == "" && operation.Value != "" {
		replaceMissingValues(col, operation.Value)
	} else if operation.Op == OP_FILLNA && operation.Method != "" {
		fillMissingWithStats(col, operation.Method)
	}

	// Transform operation come after filling missing values
	if operation.Op == OP_SCALE {
		if operation.Method == METHOD_SCALE_ZSCORE {
			statistics.ScaleWithZscore(col, df)
		} else if operation.Method == METHOD_SCALE_MINMAX {
			statistics.ScaleWithMinMax(col, df)
		}
	}
}

func dispatchDatasetNumericOperations(df *dataset.DataFrame, operations *[]config.PreprocessOp, excluded []string) {
	if operations == nil {
		return
	}
	for _, column := range df.Columns {
		// Do not process excluded columns
		if utils.Contains(excluded, column.GetName()) {
			continue
		}

		if column.GetType() != "int" && column.GetType() != "float" {
			continue
		}

		for _, op := range *operations {
			applySingleNumericOperation(df, op, column)
		}
	}
}

func dispatchColumnNumericOperations(df *dataset.DataFrame, column dataset.DataSetColumn, operations *[]config.PreprocessOp, excluded []string) {
	if operations == nil {
		return
	}
	// Do not process excluded columns
	if utils.Contains(excluded, column.GetName()) {
		return
	}

	// Process only numeric columns
	if column.GetType() != "int" && column.GetType() != "float" {
		return
	}

	for _, op := range *operations {
		applySingleNumericOperation(df, op, column)
	}
}

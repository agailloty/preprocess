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

	if operation.Op == OP_DISCRETIZE && operation.Method == METHOD_DISCRETIZE_BINNING && operation.Bins != nil {
		makeBinsFromNumericColumns(col, operation.Bins, df, true)
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

func makeBinsFromNumericColumns(column dataset.DataSetColumn, bins []config.BinningOperation, df *dataset.DataFrame, overrideColumn bool) {
	binData := make([]dataset.Nullable[string], column.Length())

	switch v := column.(type) {
	case *dataset.Float:
		for i, val := range v.Data {
			binFound := false
			for _, bin := range bins {
				if val.Value >= bin.Lower && val.Value <= bin.Upper {
					binData[i] = dataset.Nullable[string]{IsValid: true, Value: bin.Label}
					binFound = true
					break
				}
				if !binFound {
					binData[i] = dataset.Nullable[string]{IsValid: true, Value: column.ValueAt(i)}
				}
			}
		}
	case *dataset.Integer:
		for i, val := range v.Data {
			binFound := false
			for _, bin := range bins {
				if float64(val.Value) >= bin.Lower && float64(val.Value) <= bin.Upper {
					binData[i] = dataset.Nullable[string]{IsValid: true, Value: bin.Label}
					binFound = true
					break
				}
				if !binFound {
					binData[i] = dataset.Nullable[string]{IsValid: true, Value: column.ValueAt(i)}
				}
			}
		}
	}

	columnName := column.GetName()
	if !overrideColumn {
		columnName = columnName + "_C"
	}
	binnedColumn := dataset.String{
		Name: columnName,
		Data: binData,
	}

	if overrideColumn {
		for i := range df.Columns {
			if df.Columns[i].GetName() == columnName {
				df.Columns[i] = &binnedColumn
			}
		}
	} else {
		df.Columns = append(df.Columns, &binnedColumn)
	}
}

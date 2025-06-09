package operations

import (
	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/utils"
)

func parseOperations(ops []config.PreprocessOp, df *dataset.DataFrame, col dataset.DataSetColumn) []operationRunner {
	operationRunners := []operationRunner{}
	for _, op := range ops {
		if op.Op == OP_DUMMY {
			strCol, ok := col.(*dataset.String)
			if !ok {
				continue
			}
			operationRunners = append(operationRunners, parseDummy(op, df, strCol))
		}

		if op.Op == OP_FILLNA {
			switch v := col.(type) {
			case *dataset.Float:
				operationRunners = append(operationRunners, parseNumericFillna(op, df, v))
			case *dataset.Integer:
				operationRunners = append(operationRunners, parseNumericFillna(op, df, v))
			case *dataset.String:
				operationRunners = append(operationRunners, parseStringFillna(op, df, v))
			}
		}
	}

	return operationRunners
}

func parseDummy(op config.PreprocessOp, df *dataset.DataFrame, col *dataset.String) dummyOperation {
	parsedOp := dummyOperation{
		df:                  df,
		col:                 col,
		isExcluded:          utils.Contains(op.ExcludeCols, col.Name),
		dropLast:            op.DummyDropLast,
		prefixColName:       op.DummyPrefixColName,
		continueWithTooMany: op.DummyContinueWithTooManyValues,
	}

	return parsedOp
}

func parseNumericFillna(op config.PreprocessOp, df *dataset.DataFrame, col dataset.DataSetColumn) fillnaNumericOperation {
	value, ok := op.Value.(float64)

	parsedOp := fillnaNumericOperation{
		df:             df,
		col:            col,
		method:         op.Method,
		value:          value,
		isValueNumeric: ok,
	}

	return parsedOp
}

func parseStringFillna(op config.PreprocessOp, df *dataset.DataFrame, col *dataset.String) fillnaStringOperation {
	value, _ := op.Value.(string)

	parsedOp := fillnaStringOperation{
		df:     df,
		col:    col,
		method: op.Method,
		value:  value,
	}

	return parsedOp
}

func parseRenameColumn(op config.ColumnConfig, df *dataset.DataFrame, col *dataset.DataSetColumn) renameOperation {
	parsedOp := renameOperation{
		df:      df,
		col:     col,
		newName: op.NewName,
	}

	return parsedOp
}

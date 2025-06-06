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

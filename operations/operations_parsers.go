package operations

import (
	"strings"

	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/statistics"
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

		if op.Op == OP_DISCRETIZE && op.Method == METHOD_DISCRETIZE_BINNING && op.Bins != nil {
			operationRunners = append(operationRunners, parseBins(op, df, &col))
		}

		if op.Op == OP_CLEAN && op.Method != "" {
			strCol, ok := col.(*dataset.String)
			if !ok {
				continue
			}
			operationRunners = append(operationRunners, parseClean(op, df, strCol))
		}

		if op.Op == OP_GROUP && len(op.Options) > 0 {
			strCol, ok := col.(*dataset.String)
			if !ok {
				continue
			}
			operationRunners = append(operationRunners, parseGroup(op, df, strCol))
		}

		if op.Op == OP_SCALE {
			if col.GetType() == "int" || col.GetType() == "float" {
				operationRunners = append(operationRunners, parseScale(op, df, col))
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

func parseBins(op config.PreprocessOp, df *dataset.DataFrame, col *dataset.DataSetColumn) binOperation {
	parsedOps := binOperation{
		df:             df,
		col:            col,
		bins:           op.Bins,
		overrideColumn: true,
	}
	return parsedOps
}

func parseClean(op config.PreprocessOp, df *dataset.DataFrame, col *dataset.String) cleanOperation {
	var cleanFunc stringFuction

	if op.Method == METHOD_CLEAN_TRIMWS {
		cleanFunc = trimWhitespace
	}
	if op.Method == METHOD_CLEAN_LOWER {
		cleanFunc = strings.ToLower
	}

	if op.Method == METHOD_CLEAN_UPPER {
		cleanFunc = strings.ToUpper
	}

	if op.Method == METHOD_CLEAN_TITLE {
		cleanFunc = strings.ToTitle
	}

	parsedOp := cleanOperation{
		df:        df,
		col:       col,
		cleanFunc: cleanFunc,
	}

	return parsedOp
}

func parseGroup(op config.PreprocessOp, df *dataset.DataFrame, col *dataset.String) groupOperation {
	parsedOps := groupOperation{
		df:      df,
		col:     col,
		options: op.Options,
	}

	return parsedOps
}

func parseScale(op config.PreprocessOp, df *dataset.DataFrame, col dataset.DataSetColumn) scaleOperation {
	var scaleFunc statistics.ScaleFunc
	if op.Method == METHOD_SCALE_MINMAX {
		scaleFunc = statistics.ComputeMinMaxScore
	} else {
		scaleFunc = statistics.ComputeZScore
	}
	parsedOp := scaleOperation{
		df:        df,
		col:       col,
		scaleFunc: scaleFunc,
	}

	return parsedOp
}

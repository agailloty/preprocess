package operations

import (
	"strings"

	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/utils"
)

func applyOperationsOnTextColumns(df *dataset.DataFrame, operations *[]config.PreprocessOp) {
	for _, column := range df.Columns {
		if column.GetType() == "string" {
			applyTextOperationsOnColumn(df, operations, column)
		}
	}
}

func applyTextOperationsOnColumn(df *dataset.DataFrame, preprocessOps *[]config.PreprocessOp, col dataset.DataSetColumn) {
	if preprocessOps == nil {
		return
	}

	for _, prep := range *preprocessOps {
		if prep.Op == OP_FILLNA && prep.Method == "" && prep.Value != "" {
			replaceMissingValues(col, prep.Value)
		}
		var stringOps []stringFuction
		if prep.Op == OP_CLEAN {
			if prep.Method == METHOD_CLEAN_TRIMWS {
				stringOps = append(stringOps, trimWhitespace)
			}
			if prep.Method == METHOD_CLEAN_UPPER {
				stringOps = append(stringOps, strings.ToUpper)
			}
			if prep.Method == METHOD_CLEAN_LOWER {
				stringOps = append(stringOps, strings.ToLower)
			}
		}
		applyStringOperation(col, stringOps)

		if prep.Op == OP_DUMMY {
			if len(prep.ExcludeCols) > 0 && utils.Contains(prep.ExcludeCols, col.GetName()) {
				continue
			}
			addDummyToDataframe(df, col, prep.DummyDropLast, prep.DummyPrefixColName, prep.DummyContinueWithTooManyValues)
		}
	}
}

type stringFuction func(value string) string

func applyStringOperation(column dataset.DataSetColumn, operations []stringFuction) {
	switch v := column.(type) {
	case *dataset.String:
		for i := range v.Data {
			for _, stringFunc := range operations {
				v.Data[i].Value = stringFunc(v.Data[i].Value)
			}
		}
	}
}

func trimWhitespace(value string) string {
	return strings.Trim(value, " ")
}

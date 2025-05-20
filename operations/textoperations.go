package operations

import (
	"strings"

	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
)

func applyTextOperationsOnColumn(preprocessOps *[]config.PreprocessOp, col dataset.DataSetColumn) {
	if preprocessOps == nil {
		return
	}

	for _, prep := range *preprocessOps {
		if prep.Op == "fillna" && prep.Method == "" && prep.Value != "" {
			replaceMissingValues(col, prep.Value)
		}
		var stringOps []stringFuction
		if prep.Op == "clean" {
			if prep.Method == "trimws" {
				stringOps = append(stringOps, trimWhitespace)
			}
			if prep.Method == "upper" {
				stringOps = append(stringOps, strings.ToUpper)
			}
			if prep.Method == "lower" {
				stringOps = append(stringOps, strings.ToLower)
			}
		}
		applyStringOperation(col, stringOps)
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

package operations

import (
	"strings"

	"github.com/agailloty/preprocess/dataset"
)

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

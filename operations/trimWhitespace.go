package operations

import (
	"strings"

	"github.com/agailloty/preprocess/dataset"
)

type stringFuction func(value string) string

func applyStringOperation(column *dataset.String, operations []stringFuction) {
	for i := range column.Data {
		for _, stringFunc := range operations {
			column.Data[i].Value = stringFunc(column.Data[i].Value)
		}
	}
}

func trimWhitespace(value string) string {
	return strings.Trim(value, " ")
}

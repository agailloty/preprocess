package operations

import (
	"errors"
	"strings"

	"github.com/agailloty/preprocess/dataset"
)

type cleanOperation struct {
	df        *dataset.DataFrame
	col       *dataset.String
	cleanFunc stringFuction
}

func (o cleanOperation) run() []operationResult {
	results := []operationResult{}
	results = append(results, op_clean(o))
	return results
}

func op_clean(specs cleanOperation) operationResult {

	if specs.cleanFunc == nil {
		return operationResult{isSuccess: false, opErrors: []error{errors.New("UNKNOWN_CLEAN_OPERATION")}}
	}

	for i := range specs.col.Data {
		specs.col.Data[i].Value = specs.cleanFunc(specs.col.Data[i].Value)
	}
	return operationResult{isSuccess: true}
}

type stringFuction func(value string) string

func trimWhitespace(value string) string {
	return strings.Trim(value, " ")
}

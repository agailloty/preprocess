package operations

import (
	"errors"

	"github.com/agailloty/preprocess/dataset"
)

type fillnaStringOperation struct {
	df     *dataset.DataFrame
	col    *dataset.String
	method string
	value  string
}

func op_fillnaString(specs fillnaStringOperation) operationResult {
	if specs.col == nil {
		return operationResult{isSuccess: false, opErrors: []error{errors.New("NIL_COLUMN")}}
	}

	if specs.method == "" && specs.value == "" {
		return operationResult{isSuccess: false, opErrors: []error{errors.New("EMPTY_METHOD_AND_VALUE")}}
	}

	if specs.method == "" {
		fillMissingStringWithValue(specs.col, specs.value)
	}

	return operationResult{isSuccess: true}
}

func (o fillnaStringOperation) run() []operationResult {
	results := []operationResult{}
	results = append(results, op_fillnaString(o))
	return results
}

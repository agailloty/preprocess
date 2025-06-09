package operations

import (
	"errors"

	"github.com/agailloty/preprocess/dataset"
)

type renameOperation struct {
	df      *dataset.DataFrame
	col     *dataset.DataSetColumn
	newName string
}

func (o renameOperation) run() []operationResult {
	results := []operationResult{}
	results = append(results, op_rename(o))
	return results
}

func (o renameOperation) isIndependant() bool {
	return true
}

func op_rename(specs renameOperation) operationResult {
	if specs.newName == "" {
		return operationResult{isSuccess: false, opErrors: []error{errors.New("NEW_NAME_EMPTY")}}
	}

	if specs.col == nil {
		return operationResult{isSuccess: false, opErrors: []error{errors.New("COLUMN_IS_NIL")}}
	}

	if specs.newName == (*specs.col).GetName() {
		return operationResult{isSuccess: false, opErrors: []error{errors.New("NEW_NAME_SAME_AS_OLD")}}
	}

	(*specs.col).SetName(specs.newName)

	return operationResult{isSuccess: true}
}

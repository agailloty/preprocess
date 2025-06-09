package operations

import (
	"github.com/agailloty/preprocess/dataset"
)

type dummyOperation struct {
	df                  *dataset.DataFrame
	col                 *dataset.String
	isExcluded          bool
	dropLast            bool
	prefixColName       bool
	continueWithTooMany bool
}

func op_dummy(specs dummyOperation) operationResult {
	dummyCols, err := makeDummy(specs)
	if err != nil {
		return operationResult{isSuccess: false, opErrors: []error{err}}
	}
	for _, dummy := range dummyCols {
		specs.df.Columns = append(specs.df.Columns, &dummy)
	}
	// Delete initial column in all case to avoid adding dummy multiple times
	specs.df.DeleteColumn(specs.col)

	return operationResult{isSuccess: true}
}

func (o dummyOperation) run() []operationResult {
	results := []operationResult{}
	results = append(results, op_dummy(o))
	return results
}

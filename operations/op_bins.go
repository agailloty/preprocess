package operations

import (
	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
)

type binOperation struct {
	df             *dataset.DataFrame
	col            *dataset.DataSetColumn
	overrideColumn bool
	bins           []config.BinningOperation
}

func (o binOperation) run() []operationResult {
	results := []operationResult{}
	results = append(results, op_bins(o))
	return results
}

func op_bins(specs binOperation) operationResult {
	err := makeBinsFromNumericColumns(specs)
	if err != nil {
		return operationResult{isSuccess: false, opErrors: []error{err}}
	}
	return operationResult{isSuccess: true}
}

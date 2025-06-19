package operations

import (
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/statistics"
)

type scaleOperation struct {
	df        *dataset.DataFrame
	col       dataset.DataSetColumn
	scaleFunc statistics.ScaleFunc
}

func (o scaleOperation) run() []operationResult {
	results := []operationResult{}
	results = append(results, op_scale(o))
	return results
}

func op_scale(specs scaleOperation) operationResult {
	statistics.ScaleNumericColumn(specs.col, specs.df, specs.scaleFunc)
	return operationResult{isSuccess: true}
}

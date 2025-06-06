package operations

import (
	"github.com/agailloty/preprocess/dataset"
)

type operationResult struct {
	isSuccess bool
	opErrors  []error
}

type singleOperationFunc func(df dataset.DataFrame,
	column dataset.DataSetColumn, operationParams ...any) operationResult

package operations

import (
	"errors"

	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/statistics"
)

type fillnaNumericOperation struct {
	df             *dataset.DataFrame
	col            dataset.DataSetColumn
	method         string
	value          float64
	isValueNumeric bool
}

func (o fillnaNumericOperation) run() []operationResult {
	results := []operationResult{}
	results = append(results, op_fillnaNumeric(o))
	return results
}

func (o fillnaNumericOperation) isIndependant() bool {
	return true
}

func op_fillnaNumeric(specs fillnaNumericOperation) operationResult {

	if specs.method == "" {
		if !specs.isValueNumeric {
			return operationResult{isSuccess: false, opErrors: []error{errors.New("FILLING_VALUE_NOT_NUMERIC")}}
		} else {
			fillMissingNumericWithValue(specs.col, specs.value)
		}
	}

	if specs.method != "" {
		fillMissingWithStats(specs.col, specs.method)
	}

	return operationResult{isSuccess: true}
}

func fillMissingNumericWithValue(column dataset.DataSetColumn, value float64) {
	switch v := column.(type) {
	case *dataset.Integer:
		fillMissingIntegerWithValue(v, int(value))
	case *dataset.Float:
		fillMissingFloatWithValue(v, value)
	}
}

func fillMissingWithStats(column dataset.DataSetColumn, method string) {
	switch v := column.(type) {
	case *dataset.Integer:
		var statFunc func(numbers []int) float64
		switch method {
		case "mean":
			statFunc = statistics.Mean
		case "median":
			statFunc = statistics.Median
		}
		fillMissingIntegerColumnWithStatsFunc(v, statFunc)
	case *dataset.Float:
		var statFunc func(numbers []float64) float64
		switch method {
		case "mean":
			statFunc = statistics.Mean
		case "median":
			statFunc = statistics.Median
		}
		fillMissingFloatColumnWithStatsFunc(v, statFunc)
	}
}

func fillMissingIntegerColumnWithStatsFunc(column *dataset.Integer, f func(numbers []int) float64) {
	var validData []int
	for _, data := range column.Data {
		if data.IsValid {
			validData = append(validData, data.Value)
		}
	}
	replaceValue := int(f(validData))
	fillMissingIntegerWithValue(column, replaceValue)
}

func fillMissingFloatColumnWithStatsFunc[T statistics.Number](column dataset.DataSetColumn, f func(numbers []T) float64) {
	floatColumn, ok := column.(*dataset.Float)
	if ok {
		var validData []T
		for _, data := range floatColumn.Data {
			if data.IsValid {
				validData = append(validData, T(data.Value))
			}
		}
		replaceValue := f(validData)
		replaceMissingValues(column, replaceValue)
	}
}

package operations

import (
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/statistics"
)

func replaceMissingWithStats(column dataset.DataSetColumn, method string) {
	switch v := column.(type) {
	case *dataset.Integer:
		var statFunc func(numbers []int) float64
		switch method {
		case "mean":
			statFunc = statistics.Mean
		case "median":
			statFunc = statistics.Median
		}
		replaceIntegerColumnWithStatsFunc(v, statFunc)
	case *dataset.Float:
		var statFunc func(numbers []float32) float64
		switch method {
		case "mean":
			statFunc = statistics.Mean
		case "median":
			statFunc = statistics.Median
		}
		replaceFloatColumnWithStatsFunc(v, statFunc)
	}
}

func replaceIntegerColumnWithStatsFunc(column *dataset.Integer, f func(numbers []int) float64) {
	var validData []int
	for _, data := range column.Data {
		if data.IsValid {
			validData = append(validData, data.Value)
		}
	}
	replaceValue := int(f(validData))
	fillMissingIntegerWithValue(column, replaceValue)
}

func replaceFloatColumnWithStatsFunc[T statistics.Number](column dataset.DataSetColumn, f func(numbers []T) float64) {
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

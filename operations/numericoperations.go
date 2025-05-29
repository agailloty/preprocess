package operations

import (
	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/statistics"
)

func applyOperationsOnNumericColumns(df *dataset.DataFrame, operations *[]config.PreprocessOp) {
	for _, column := range df.Columns {
		if column.GetType() == "int" || column.GetType() == "float" {
			applyNumericOperationsOnColumn(operations, column, df)
		}
	}
}

func applyNumericOperationsOnColumn(preprocessOps *[]config.PreprocessOp, col dataset.DataSetColumn, df *dataset.DataFrame) {
	if preprocessOps != nil {
		for _, prep := range *preprocessOps {
			if prep.Op == "fillna" && prep.Method == "" && prep.Value != "" {
				replaceMissingValues(col, prep.Value)
			} else if prep.Op == "fillna" && prep.Method != "" {
				replaceMissingWithStats(col, prep.Method)
			}

			// Transform operation come after filling missing values
			if prep.Op == "scale" {
				if prep.Method == "zscore" {
					statistics.ScaleWithZscore(col, df)
				} else if prep.Method == "minmax" {
					statistics.ScaleWithMinMax(col, df)
				}
			}

			if prep.Op == "discretize" && prep.Method == "binning" && prep.Bins != nil {
				makeBinsFromNumericColumns(col, prep.Bins, df, true)
			}
		}
	}
}

func makeBinsFromNumericColumns(column dataset.DataSetColumn, bins []config.BinningOperation, df *dataset.DataFrame, overrideColumn bool) {
	binData := make([]dataset.Nullable[string], column.Length())

	switch v := column.(type) {
	case *dataset.Float:
		for i, val := range v.Data {
			binFound := false
			for _, bin := range bins {
				if val.Value >= bin.Lower && val.Value <= bin.Upper {
					binData[i] = dataset.Nullable[string]{IsValid: true, Value: bin.Label}
					binFound = true
					break
				}
				if !binFound {
					binData[i] = dataset.Nullable[string]{IsValid: true, Value: column.ValueAt(i)}
				}
			}
		}
	case *dataset.Integer:
		for i, val := range v.Data {
			binFound := false
			for _, bin := range bins {
				if float64(val.Value) >= bin.Lower && float64(val.Value) <= bin.Upper {
					binData[i] = dataset.Nullable[string]{IsValid: true, Value: bin.Label}
					binFound = true
					break
				}
				if !binFound {
					binData[i] = dataset.Nullable[string]{IsValid: true, Value: column.ValueAt(i)}
				}
			}
		}
	}

	columnName := column.GetName()
	if !overrideColumn {
		columnName = columnName + "_C"
	}
	binnedColumn := dataset.String{
		Name: columnName,
		Data: binData,
	}

	if overrideColumn {
		for i := range df.Columns {
			if df.Columns[i].GetName() == columnName {
				df.Columns[i] = &binnedColumn
			}
		}
	} else {
		df.Columns = append(df.Columns, &binnedColumn)
	}
}

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
		var statFunc func(numbers []float64) float64
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

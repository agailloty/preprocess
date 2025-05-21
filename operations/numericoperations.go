package operations

import (
	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/statistics"
	"github.com/agailloty/preprocess/utils"
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
			if prep.Op == "normalize" {
				if prep.Method == "zscore" {
					applyZScoreToEveryElement(col, df)
				} else if prep.Method == "minmax" {
					applyMinMaxScoreToEveryElement(col)
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
				if float32(val.Value) >= bin.Lower && float32(val.Value) <= bin.Upper {
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

func applyZScoreToEveryElement(column dataset.DataSetColumn, df *dataset.DataFrame) {
	switch v := column.(type) {
	case *dataset.Integer:
		// Convert into to float32
		validData := extractNonNullInts(v.Data)
		mu := statistics.Mean(validData)
		sigma := statistics.StdDev(validData)
		zScores := make([]dataset.Nullable[float32], column.Length())
		for i := range v.Data {
			zScore := statistics.ComputeZScore(float32(v.Data[i].Value), float32(mu), float32(sigma))
			zScores[i] = dataset.Nullable[float32]{IsValid: v.Data[i].IsValid, Value: zScore}
		}
		newColumn := dataset.Float{Name: column.GetName(), Data: zScores}
		utils.OverrideDataFrameColumn(df, column.GetName(), &newColumn)

	case *dataset.Float:
		validData := extractNonNullFloats(v.Data)
		mu := statistics.Mean(validData)
		sigma := statistics.StdDev(validData)
		for i := range v.Data {
			zScore := statistics.ComputeZScore(float32(v.Data[i].Value), float32(mu), float32(sigma))
			v.Data[i] = dataset.Nullable[float32]{IsValid: v.Data[i].IsValid, Value: zScore}
		}
	}
}

func applyMinMaxScoreToEveryElement(column dataset.DataSetColumn) {
	switch v := column.(type) {
	case *dataset.Integer:
		validData := extractNonNullInts(v.Data)
		min, max := statistics.MinMax(validData)
		for i := range v.Data {
			zScore := statistics.ComputeMinMaxScore(float32(v.Data[i].Value), float32(min), float32(max))
			v.Data[i] = dataset.Nullable[int]{IsValid: v.Data[i].IsValid, Value: int(zScore)}
		}

	case *dataset.Float:
		validData := extractNonNullFloats(v.Data)
		mu := statistics.Mean(validData)
		sigma := statistics.StdDev(validData)
		for i := range v.Data {
			zScore := statistics.ComputeZScore(float32(v.Data[i].Value), float32(mu), float32(sigma))
			v.Data[i] = dataset.Nullable[float32]{IsValid: v.Data[i].IsValid, Value: zScore}
		}
	}
}

func extractNonNullInts(data []dataset.Nullable[int]) []int {
	result := make([]int, len(data))
	for i, item := range data {
		if item.IsValid {
			result[i] = item.Value
		}
	}
	return result
}

func extractNonNullFloats(data []dataset.Nullable[float32]) []float32 {
	result := make([]float32, len(data))
	for i, item := range data {
		if item.IsValid {
			result[i] = item.Value
		}
	}
	return result
}

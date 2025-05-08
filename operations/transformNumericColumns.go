package operations

import (
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/statistics"
)

func applyZScoreToEveryElement(column dataset.DataSetColumn) {
	switch v := column.(type) {
	case *dataset.Integer:
		validData := extractIntsWithDefault(v.Data, 0)
		mu := statistics.Mean(validData)
		sigma := statistics.StdDev(validData)
		for i := range v.Data {
			zScore := statistics.ComputeZScore(float32(v.Data[i].Value), float32(mu), float32(sigma))
			v.Data[i] = dataset.Nullable[int]{IsValid: true, Value: int(zScore)}
		}

	case *dataset.Float:
		validData := extractFloatsWithDefault(v.Data, 0)
		mu := statistics.Mean(validData)
		sigma := statistics.StdDev(validData)
		for i := range v.Data {
			zScore := statistics.ComputeZScore(float32(v.Data[i].Value), float32(mu), float32(sigma))
			v.Data[i] = dataset.Nullable[float32]{IsValid: true, Value: zScore}
		}
	}
}

func applyMinMaxScoreToEveryElement(column dataset.DataSetColumn) {
	switch v := column.(type) {
	case *dataset.Integer:
		validData := extractIntsWithDefault(v.Data, 0)
		min, max := statistics.MinMax(validData)
		for i := range v.Data {
			zScore := statistics.ComputeMinMaxScore(float32(v.Data[i].Value), float32(min), float32(max))
			v.Data[i] = dataset.Nullable[int]{IsValid: true, Value: int(zScore)}
		}

	case *dataset.Float:
		validData := extractFloatsWithDefault(v.Data, 0)
		mu := statistics.Mean(validData)
		sigma := statistics.StdDev(validData)
		for i := range v.Data {
			zScore := statistics.ComputeZScore(float32(v.Data[i].Value), float32(mu), float32(sigma))
			v.Data[i] = dataset.Nullable[float32]{IsValid: true, Value: zScore}
		}
	}
}

func extractIntsWithDefault(data []dataset.Nullable[int], defaultValue int) []int {
	result := make([]int, len(data))
	for i, item := range data {
		if item.IsValid {
			result[i] = item.Value
		} else {
			result[i] = defaultValue
		}
	}
	return result
}

func extractFloatsWithDefault(data []dataset.Nullable[float32], defaultValue float32) []float32 {
	result := make([]float32, len(data))
	for i, item := range data {
		if item.IsValid {
			result[i] = item.Value
		} else {
			result[i] = defaultValue
		}
	}
	return result
}

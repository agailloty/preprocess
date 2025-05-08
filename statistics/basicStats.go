package statistics

import (
	"math"
	"sort"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Mean[T Number](numbers []T) float64 {
	if len(numbers) == 0 {
		return 0
	}
	var sum float64
	for _, v := range numbers {
		sum += float64(v)
	}
	return sum / float64(len(numbers))
}

func Variance[T Number](values []T) float64 {
	mean := Mean(values)
	var sumSq float64
	for _, v := range values {
		diff := float64(v) - mean
		sumSq += diff * diff
	}
	return sumSq / float64(len(values))
}

func StdDev[T Number](values []T) float64 {
	return math.Sqrt(Variance(values))
}

func Median[T Number](values []T) float64 {
	if len(values) == 0 {
		return 0
	}
	sorted := make([]float64, len(values))
	for i, v := range values {
		sorted[i] = float64(v)
	}
	sort.Float64s(sorted)

	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2
	}
	return sorted[mid]
}

func MinMax[T Number](values []T) (min, max T) {
	min, max = values[0], values[0]
	for _, v := range values[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

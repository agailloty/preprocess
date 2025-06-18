package statistics

type ScaleFunc func(x float64, mu float64, sigma float64) float64

func ComputeZScore(x float64, mu float64, sigma float64) float64 {
	if sigma == 0 {
		return 0
	}

	return (x - mu) / sigma
}

func ComputeMinMaxScore(x float64, min float64, max float64) float64 {
	return (x - min) / (max - min)
}

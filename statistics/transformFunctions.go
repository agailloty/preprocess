package statistics

func ComputeZScore(x float32, mu float32, sigma float32) float32 {
	if sigma == 0 {
		return 0
	}

	return (x - mu) / sigma
}

func ComputeMinMaxScore(x float32, min float32, max float32) float32 {
	return (x - min) / (max - min)
}

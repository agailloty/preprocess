package statistics

func ComputeZScore(x float32, mu float32, sigma float32) float32 {
	if sigma == 0 {
		return 0
	}

	return (x - mu) / sigma
}

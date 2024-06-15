package common

func Clamp(v, min, max float64) float64 {
	return ClampMax(ClampMin(v, min), max)
}

func ClampMax(v, max float64) float64 {
	if v > max {
		return max
	}

	return v
}

func ClampMin(v, min float64) float64 {
	if v < min {
		return min
	}

	return v
}

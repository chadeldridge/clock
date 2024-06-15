package incrementers

func Clamp(v, min, max int) int {
	return ClampMax(ClampMin(v, min), max)
}

func ClampMax(v, max int) int {
	if v > max {
		return max
	}

	return v
}

func ClampMin(v, min int) int {
	if v < min {
		return min
	}

	return v
}

package incrementers

func IsClamped(v, min, max int) bool {
	return v >= min && v <= max
}

func Clamp(v, min, max int) int {
	return ClampMax(ClampMin(v, min), max)
}

func ClampMin(v, min int) int {
	if v < min {
		return min
	}

	return v
}

func ClampMax(v, max int) int {
	if v > max {
		return max
	}

	return v
}

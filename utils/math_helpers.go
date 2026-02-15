package utils

import "math"

// Mean calculates the average of a slice of values
func Mean(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}

	var sum float64
	for _, v := range values {
		sum += v
	}

	return sum / float64(len(values))
}

// StandardDeviation calculates the standard deviation of a slice of values
func StandardDeviation(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}

	mean := Mean(values)
	var variance float64

	for _, v := range values {
		variance += math.Pow(v-mean, 2)
	}

	return math.Sqrt(variance / float64(len(values)))
}

// PercentageChange calculates the percentage change between two values
func PercentageChange(old, new float64) float64 {
	if old == 0 {
		return 0
	}

	return ((new - old) / old) * 100
}

// BasisPoints calculates basis points difference between value and reference
func BasisPoints(value, reference float64) float64 {
	if reference == 0 {
		return 0
	}

	return ((value - reference) / reference) * 10000
}

// AbsoluteChange calculates the absolute change between two values
func AbsoluteChange(old, new float64) float64 {
	return new - old
}

// Max returns the maximum of two float64 values
func Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// Min returns the minimum of two float64 values
func Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

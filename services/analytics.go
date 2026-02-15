package services

import "math"

func CalculateSpread(bid, ask float64) float64 {
	return ask - bid
}

func LiquidityScore(volume, spread float64) float64 {
	return volume / (spread + 1)
}

func Volatility(prices []float64) float64 {
	if len(prices) == 0 {
		return 0
	}

	var mean float64
	for _, p := range prices {
		mean += p
	}
	mean /= float64(len(prices))

	var variance float64
	for _, p := range prices {
		variance += math.Pow(p-mean, 2)
	}

	return math.Sqrt(variance / float64(len(prices)))
}

func TrendingScore(volume, volatility float64) float64 {
	return volume * volatility
}

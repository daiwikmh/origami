package models

import "time"

// Trade represents a single trade execution
type Trade struct {
	MarketID  string
	Price     float64
	Quantity  float64
	Timestamp time.Time
	IsBuy     bool
}

// PriceHistory maintains a rolling window of prices for a market
type PriceHistory struct {
	MarketID string
	Prices   []float64
	Times    []time.Time
	MaxSize  int
}

// NewPriceHistory creates a new price history with default window size
func NewPriceHistory(marketID string) *PriceHistory {
	return &PriceHistory{
		MarketID: marketID,
		Prices:   make([]float64, 0, 100),
		Times:    make([]time.Time, 0, 100),
		MaxSize:  100,
	}
}

// AddPrice appends a price and maintains the rolling window
func (ph *PriceHistory) AddPrice(price float64, timestamp time.Time) {
	ph.Prices = append(ph.Prices, price)
	ph.Times = append(ph.Times, timestamp)

	// Keep only the last MaxSize entries
	if len(ph.Prices) > ph.MaxSize {
		ph.Prices = ph.Prices[len(ph.Prices)-ph.MaxSize:]
		ph.Times = ph.Times[len(ph.Times)-ph.MaxSize:]
	}
}

// GetPrices returns the current price window
func (ph *PriceHistory) GetPrices() []float64 {
	return ph.Prices
}

// GetLatestPrice returns the most recent price
func (ph *PriceHistory) GetLatestPrice() (float64, bool) {
	if len(ph.Prices) == 0 {
		return 0, false
	}
	return ph.Prices[len(ph.Prices)-1], true
}

package models

type MarketSummary struct {
	MarketID       string  `json:"market_id"`
	BaseDenom      string  `json:"base_denom"`
	QuoteDenom     string  `json:"quote_denom"`
	Volume         float64 `json:"volume"`
	Price          float64 `json:"price"`
	Volatility     float64 `json:"volatility"`
	LiquidityScore float64 `json:"liquidity_score"`
}

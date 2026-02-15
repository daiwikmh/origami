package models

import "time"

// OrderbookDepth provides multi-level orderbook metrics
type OrderbookDepth struct {
	BidDepth5   float64 `json:"bid_depth_5"`
	AskDepth5   float64 `json:"ask_depth_5"`
	BidDepth10  float64 `json:"bid_depth_10"`
	AskDepth10  float64 `json:"ask_depth_10"`
	TotalBids   int     `json:"total_bids"`
	TotalAsks   int     `json:"total_asks"`
	Spread      float64 `json:"spread"`
	SpreadBps   float64 `json:"spread_bps"`
	MidPrice    float64 `json:"mid_price"`
}

// MarketAnalytics contains comprehensive market metrics
type MarketAnalytics struct {
	MarketID         string          `json:"market_id"`
	BaseDenom        string          `json:"base_denom"`
	QuoteDenom       string          `json:"quote_denom"`
	CurrentPrice     float64         `json:"current_price"`
	Volume24h        float64         `json:"volume_24h"`
	PriceChange24h   float64         `json:"price_change_24h"`
	PriceChange24hPct float64        `json:"price_change_24h_pct"`
	Volatility       float64         `json:"volatility"`
	LiquidityScore   float64         `json:"liquidity_score"`
	TrendingScore    float64         `json:"trending_score"`
	OrderbookDepth   *OrderbookDepth `json:"orderbook_depth,omitempty"`
	Timestamp        time.Time       `json:"timestamp"`
}

// TrendingMarket represents a market in trending rankings
type TrendingMarket struct {
	MarketID    string  `json:"market_id"`
	Symbol      string  `json:"symbol"`
	Score       float64 `json:"score"`
	Volume24h   float64 `json:"volume_24h"`
	Volatility  float64 `json:"volatility"`
	PriceChange float64 `json:"price_change_pct"`
}

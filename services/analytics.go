package services

import (
	"math"
	"sort"
	"time"

	"github.com/daiwikmh/origami/cache"
	"github.com/daiwikmh/origami/models"
	"github.com/daiwikmh/origami/utils"
)

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

// CalculateRealLiquidity computes liquidity from orderbook depth
func CalculateRealLiquidity(orderbookData interface{}) float64 {
	obMap, ok := orderbookData.(map[string]interface{})
	if !ok {
		return 0
	}

	orderbook, ok := obMap["orderbook"].(map[string]interface{})
	if !ok {
		return 0
	}

	buys, _ := utils.ExtractOrderbookLevels(orderbook, "buys")
	sells, _ := utils.ExtractOrderbookLevels(orderbook, "sells")

	// Calculate total liquidity in top 10 levels
	var bidLiquidity, askLiquidity float64

	limit := 10
	for i := 0; i < limit && i < len(buys); i++ {
		bidLiquidity += buys[i].Price * buys[i].Quantity
	}

	for i := 0; i < limit && i < len(sells); i++ {
		askLiquidity += sells[i].Price * sells[i].Quantity
	}

	return bidLiquidity + askLiquidity
}

// CalculateOrderbookDepth computes multi-level depth metrics
func CalculateOrderbookDepth(orderbookData interface{}) *models.OrderbookDepth {
	obMap, ok := orderbookData.(map[string]interface{})
	if !ok {
		return nil
	}

	orderbook, ok := obMap["orderbook"].(map[string]interface{})
	if !ok {
		return nil
	}

	buys, _ := utils.ExtractOrderbookLevels(orderbook, "buys")
	sells, _ := utils.ExtractOrderbookLevels(orderbook, "sells")

	if len(buys) == 0 || len(sells) == 0 {
		return nil
	}

	depth := &models.OrderbookDepth{
		TotalBids: len(buys),
		TotalAsks: len(sells),
	}

	// Calculate depth at different levels
	for i := 0; i < 5 && i < len(buys); i++ {
		depth.BidDepth5 += buys[i].Quantity * buys[i].Price
	}

	for i := 0; i < 5 && i < len(sells); i++ {
		depth.AskDepth5 += sells[i].Quantity * sells[i].Price
	}

	for i := 0; i < 10 && i < len(buys); i++ {
		depth.BidDepth10 += buys[i].Quantity * buys[i].Price
	}

	for i := 0; i < 10 && i < len(sells); i++ {
		depth.AskDepth10 += sells[i].Quantity * sells[i].Price
	}

	// Calculate spread
	bestBid := buys[0].Price
	bestAsk := sells[0].Price
	depth.Spread = bestAsk - bestBid
	depth.MidPrice = (bestBid + bestAsk) / 2

	if depth.MidPrice > 0 {
		depth.SpreadBps = (depth.Spread / depth.MidPrice) * 10000
	}

	return depth
}

// CalculateMarketVolatility calculates volatility using cached price history
func CalculateMarketVolatility(marketID string, dataCache *cache.DataCache) float64 {
	history, exists := dataCache.GetPriceHistory(marketID)
	if !exists || len(history.Prices) < 2 {
		return 0
	}

	return Volatility(history.Prices)
}

// ComputeMarketAnalytics calculates comprehensive analytics for a market
func ComputeMarketAnalytics(marketID string, dataCache *cache.DataCache) *models.MarketAnalytics {
	// Get orderbook
	orderbookData, found := dataCache.GetOrderbook(marketID)
	if !found {
		return nil
	}

	// Get price history
	priceHistory, _ := dataCache.GetPriceHistory(marketID)

	// Extract market info from orderbook
	obMap, ok := orderbookData.(map[string]interface{})
	if !ok {
		return nil
	}

	orderbook, ok := obMap["orderbook"].(map[string]interface{})
	if !ok {
		return nil
	}

	buys, _ := utils.ExtractOrderbookLevels(orderbook, "buys")
	sells, _ := utils.ExtractOrderbookLevels(orderbook, "sells")

	if len(buys) == 0 || len(sells) == 0 {
		return nil
	}

	// Calculate current price (mid price)
	currentPrice := (buys[0].Price + sells[0].Price) / 2

	// Calculate orderbook depth
	depth := CalculateOrderbookDepth(orderbookData)

	// Calculate volatility
	volatility := 0.0
	if priceHistory != nil && len(priceHistory.Prices) > 0 {
		volatility = Volatility(priceHistory.Prices)
	}

	// Calculate 24h price change
	priceChange24h := 0.0
	priceChange24hPct := 0.0
	if priceHistory != nil && len(priceHistory.Prices) > 0 {
		oldPrice := priceHistory.Prices[0]
		priceChange24h = currentPrice - oldPrice
		priceChange24hPct = utils.PercentageChange(oldPrice, currentPrice)
	}

	// Calculate volume (simplified - from trades if available)
	volume24h := 0.0
	trades, found := dataCache.GetTrades(marketID)
	if found {
		for _, trade := range trades {
			volume24h += trade.Price * trade.Quantity
		}
	}

	// Calculate liquidity score
	spread := depth.Spread
	liquidityScore := LiquidityScore(volume24h, spread)

	// Calculate trending score
	trendingScore := CalculateTrendingScoreEnhanced(volume24h, volatility, math.Abs(priceChange24hPct))

	return &models.MarketAnalytics{
		MarketID:         marketID,
		BaseDenom:        "",  // Would extract from market data
		QuoteDenom:       "",  // Would extract from market data
		CurrentPrice:     currentPrice,
		Volume24h:        volume24h,
		PriceChange24h:   priceChange24h,
		PriceChange24hPct: priceChange24hPct,
		Volatility:       volatility,
		LiquidityScore:   liquidityScore,
		TrendingScore:    trendingScore,
		OrderbookDepth:   depth,
		Timestamp:        time.Now(),
	}
}

// CalculateTrendingScoreEnhanced uses weighted formula
// Volume: 40%, Volatility: 30%, Price Change: 30%
func CalculateTrendingScoreEnhanced(volume, volatility, priceChange float64) float64 {
	// Normalize components (simple scaling)
	volumeScore := math.Log10(volume + 1) * 0.4
	volatilityScore := volatility * 0.3
	priceChangeScore := priceChange * 0.3

	return volumeScore + volatilityScore + priceChangeScore
}

// GetTopTrendingMarkets returns top N markets by trending score
func GetTopTrendingMarkets(limit int, dataCache *cache.DataCache) []*models.TrendingMarket {
	allAnalytics := dataCache.GetAllAnalytics()

	// Convert to trending markets
	trendingMarkets := make([]*models.TrendingMarket, 0, len(allAnalytics))
	for _, analytics := range allAnalytics {
		trendingMarkets = append(trendingMarkets, &models.TrendingMarket{
			MarketID:    analytics.MarketID,
			Symbol:      analytics.BaseDenom + "/" + analytics.QuoteDenom,
			Score:       analytics.TrendingScore,
			Volume24h:   analytics.Volume24h,
			Volatility:  analytics.Volatility,
			PriceChange: analytics.PriceChange24hPct,
		})
	}

	// Sort by score descending
	sort.Slice(trendingMarkets, func(i, j int) bool {
		return trendingMarkets[i].Score > trendingMarkets[j].Score
	})

	// Return top N
	if len(trendingMarkets) > limit {
		trendingMarkets = trendingMarkets[:limit]
	}

	return trendingMarkets
}

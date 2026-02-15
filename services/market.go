package services

import (
	"sort"
	"time"

	"github.com/daiwikmh/origami/cache"
	"github.com/daiwikmh/origami/clients"
	"github.com/daiwikmh/origami/models"
)

var dataCache *cache.DataCache

// InitMarketService initializes the market service with cache
func InitMarketService(cache *cache.DataCache) {
	dataCache = cache
}

func GetMarkets() (map[string]interface{}, error) {
	// Try cache first
	if dataCache != nil {
		if cached, found := dataCache.GetMarkets(); found {
			return cached.(map[string]interface{}), nil
		}
	}

	// Fallback to API
	data, err := clients.FetchMarkets()
	if err == nil && dataCache != nil {
		dataCache.SetMarkets(data, 10*time.Second)
	}

	return data, err
}

func GetOrderbook(marketID string) (map[string]interface{}, error) {
	// Try cache first
	if dataCache != nil {
		if cached, found := dataCache.GetOrderbook(marketID); found {
			return cached.(map[string]interface{}), nil
		}
	}

	// Fallback to API
	data, err := clients.FetchOrderbook(marketID)
	if err == nil && dataCache != nil {
		dataCache.SetOrderbook(marketID, data, 5*time.Second)
	}

	return data, err
}

// GetMarketAnalytics retrieves analytics for a specific market
func GetMarketAnalytics(marketID string) *models.MarketAnalytics {
	if dataCache == nil {
		return nil
	}

	// Try to get from cache
	analytics, found := dataCache.GetAnalytics(marketID)
	if found {
		return analytics
	}

	// Compute on-demand if not cached
	return ComputeMarketAnalytics(marketID, dataCache)
}

// GetTopMarkets returns top markets sorted by specified criteria
func GetTopMarkets(sortBy string, limit int) []*models.MarketAnalytics {
	if dataCache == nil {
		return nil
	}

	allAnalytics := dataCache.GetAllAnalytics()

	// Sort based on criteria
	switch sortBy {
	case "volume":
		sort.Slice(allAnalytics, func(i, j int) bool {
			return allAnalytics[i].Volume24h > allAnalytics[j].Volume24h
		})
	case "volatility":
		sort.Slice(allAnalytics, func(i, j int) bool {
			return allAnalytics[i].Volatility > allAnalytics[j].Volatility
		})
	case "trending":
		sort.Slice(allAnalytics, func(i, j int) bool {
			return allAnalytics[i].TrendingScore > allAnalytics[j].TrendingScore
		})
	case "price_change":
		sort.Slice(allAnalytics, func(i, j int) bool {
			return allAnalytics[i].PriceChange24hPct > allAnalytics[j].PriceChange24hPct
		})
	default:
		// Default to trending score
		sort.Slice(allAnalytics, func(i, j int) bool {
			return allAnalytics[i].TrendingScore > allAnalytics[j].TrendingScore
		})
	}

	// Return top N
	if len(allAnalytics) > limit {
		allAnalytics = allAnalytics[:limit]
	}

	return allAnalytics
}

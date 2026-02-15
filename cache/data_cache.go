package cache

import (
	"sync"
	"time"

	"github.com/daiwikmh/origami/models"
)

// CacheEntry holds cached data with expiration
type CacheEntry struct {
	Data      interface{}
	ExpiresAt time.Time
}

// DataCache provides thread-safe in-memory caching with TTL
type DataCache struct {
	markets      *CacheEntry
	orderbooks   map[string]*CacheEntry
	trades       map[string][]*models.Trade
	priceHistory map[string]*models.PriceHistory
	analytics    map[string]*models.MarketAnalytics
	mu           sync.RWMutex
}

// NewDataCache initializes an empty cache
func NewDataCache() *DataCache {
	cache := &DataCache{
		orderbooks:   make(map[string]*CacheEntry),
		trades:       make(map[string][]*models.Trade),
		priceHistory: make(map[string]*models.PriceHistory),
		analytics:    make(map[string]*models.MarketAnalytics),
	}

	// Start cleanup goroutine
	go cache.cleanupLoop()

	return cache
}

// SetMarkets stores market list with TTL
func (c *DataCache) SetMarkets(data interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.markets = &CacheEntry{
		Data:      data,
		ExpiresAt: time.Now().Add(ttl),
	}
}

// GetMarkets retrieves cached markets if not expired
func (c *DataCache) GetMarkets() (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.markets == nil || c.isExpired(c.markets) {
		return nil, false
	}

	return c.markets.Data, true
}

// SetOrderbook stores orderbook for a market with TTL
func (c *DataCache) SetOrderbook(marketID string, data interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.orderbooks[marketID] = &CacheEntry{
		Data:      data,
		ExpiresAt: time.Now().Add(ttl),
	}
}

// GetOrderbook retrieves cached orderbook if not expired
func (c *DataCache) GetOrderbook(marketID string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, exists := c.orderbooks[marketID]
	if !exists || c.isExpired(entry) {
		return nil, false
	}

	return entry.Data, true
}

// SetTrades replaces trade history for a market
func (c *DataCache) SetTrades(marketID string, trades []*models.Trade) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.trades[marketID] = trades
}

// GetTrades retrieves trade history for a market
func (c *DataCache) GetTrades(marketID string) ([]*models.Trade, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	trades, exists := c.trades[marketID]
	return trades, exists
}

// AppendTrade adds a new trade and maintains rolling window (max 1000)
func (c *DataCache) AppendTrade(marketID string, trade *models.Trade) {
	c.mu.Lock()
	defer c.mu.Unlock()

	trades := c.trades[marketID]
	trades = append(trades, trade)

	// Keep last 1000 trades
	if len(trades) > 1000 {
		trades = trades[len(trades)-1000:]
	}

	c.trades[marketID] = trades
}

// SetPriceHistory stores price history for a market
func (c *DataCache) SetPriceHistory(marketID string, history *models.PriceHistory) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.priceHistory[marketID] = history
}

// GetPriceHistory retrieves price history for a market
func (c *DataCache) GetPriceHistory(marketID string) (*models.PriceHistory, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	history, exists := c.priceHistory[marketID]
	return history, exists
}

// SetAnalytics stores computed analytics for a market
func (c *DataCache) SetAnalytics(marketID string, analytics *models.MarketAnalytics) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.analytics[marketID] = analytics
}

// GetAnalytics retrieves cached analytics for a market
func (c *DataCache) GetAnalytics(marketID string) (*models.MarketAnalytics, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	analytics, exists := c.analytics[marketID]
	if !exists {
		return nil, false
	}

	// Check if analytics are fresh (within 15 seconds)
	if time.Since(analytics.Timestamp) > 15*time.Second {
		return nil, false
	}

	return analytics, true
}

// GetAllAnalytics retrieves all cached analytics (for trending calculations)
func (c *DataCache) GetAllAnalytics() []*models.MarketAnalytics {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := make([]*models.MarketAnalytics, 0, len(c.analytics))
	for _, analytics := range c.analytics {
		// Only include fresh analytics (within 15 seconds)
		if time.Since(analytics.Timestamp) <= 15*time.Second {
			result = append(result, analytics)
		}
	}

	return result
}

// isExpired checks if a cache entry has expired (not thread-safe, caller must lock)
func (c *DataCache) isExpired(entry *CacheEntry) bool {
	return time.Now().After(entry.ExpiresAt)
}

// CleanExpired removes expired entries
func (c *DataCache) CleanExpired() {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Clean markets
	if c.markets != nil && c.isExpired(c.markets) {
		c.markets = nil
	}

	// Clean orderbooks
	for marketID, entry := range c.orderbooks {
		if c.isExpired(entry) {
			delete(c.orderbooks, marketID)
		}
	}

	// Clean old analytics (older than 1 minute)
	for marketID, analytics := range c.analytics {
		if time.Since(analytics.Timestamp) > 60*time.Second {
			delete(c.analytics, marketID)
		}
	}
}

// cleanupLoop runs periodic cleanup
func (c *DataCache) cleanupLoop() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		c.CleanExpired()
	}
}

package workers

import (
	"log"
	"sync"
	"time"

	"github.com/daiwikmh/origami/cache"
	"github.com/daiwikmh/origami/clients"
	"github.com/daiwikmh/origami/models"
	"github.com/daiwikmh/origami/services"
	"github.com/daiwikmh/origami/utils"
)

// DataCollector manages background data collection workers
type DataCollector struct {
	cache    *cache.DataCache
	stopChan chan bool
	wg       sync.WaitGroup
}

// NewDataCollector creates a new data collector
func NewDataCollector(dataCache *cache.DataCache) *DataCollector {
	return &DataCollector{
		cache:    dataCache,
		stopChan: make(chan bool),
	}
}

// Start begins all background workers
func (dc *DataCollector) Start() {
	log.Println("Starting background workers...")

	dc.wg.Add(5)

	go dc.collectMarkets()
	go dc.collectOrderbooks()
	go dc.collectTrades()
	go dc.updatePriceHistory()
	go dc.computeAnalytics()

	log.Println("Background workers started")
}

// Stop gracefully stops all workers
func (dc *DataCollector) Stop() {
	log.Println("Stopping background workers...")
	close(dc.stopChan)
	dc.wg.Wait()
	log.Println("Background workers stopped")
}

// collectMarkets fetches market list every 10 seconds
func (dc *DataCollector) collectMarkets() {
	defer dc.wg.Done()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	// Initial fetch
	dc.fetchAndCacheMarkets()

	for {
		select {
		case <-dc.stopChan:
			return
		case <-ticker.C:
			dc.fetchAndCacheMarkets()
		}
	}
}

func (dc *DataCollector) fetchAndCacheMarkets() {
	data, err := clients.FetchMarkets()
	if err != nil {
		log.Printf("Error fetching markets: %v", err)
		return
	}

	dc.cache.SetMarkets(data, 10*time.Second)
	log.Println("Markets updated")
}

// collectOrderbooks fetches orderbooks for active markets every 5 seconds
func (dc *DataCollector) collectOrderbooks() {
	defer dc.wg.Done()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-dc.stopChan:
			return
		case <-ticker.C:
			dc.fetchAndCacheOrderbooks()
		}
	}
}

func (dc *DataCollector) fetchAndCacheOrderbooks() {
	// Get market list from cache
	marketsData, found := dc.cache.GetMarkets()
	if !found {
		return
	}

	marketsMap, ok := marketsData.(map[string]interface{})
	if !ok {
		return
	}

	marketsList, ok := marketsMap["markets"].([]interface{})
	if !ok {
		return
	}

	// Fetch orderbooks for top 50 markets (to avoid overwhelming the API)
	limit := 50
	if len(marketsList) < limit {
		limit = len(marketsList)
	}

	for i := 0; i < limit; i++ {
		market, ok := marketsList[i].(map[string]interface{})
		if !ok {
			continue
		}

		marketID := utils.ParseString(market["marketId"])
		if marketID == "" {
			continue
		}

		// Fetch orderbook
		orderbook, err := clients.FetchOrderbook(marketID)
		if err != nil {
			log.Printf("Error fetching orderbook for %s: %v", marketID, err)
			continue
		}

		dc.cache.SetOrderbook(marketID, orderbook, 5*time.Second)
	}

	log.Printf("Orderbooks updated for %d markets", limit)
}

// collectTrades fetches recent trades every 10 seconds
func (dc *DataCollector) collectTrades() {
	defer dc.wg.Done()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-dc.stopChan:
			return
		case <-ticker.C:
			dc.fetchAndCacheTrades()
		}
	}
}

func (dc *DataCollector) fetchAndCacheTrades() {
	// Get market list from cache
	marketsData, found := dc.cache.GetMarkets()
	if !found {
		return
	}

	marketsMap, ok := marketsData.(map[string]interface{})
	if !ok {
		return
	}

	marketsList, ok := marketsMap["markets"].([]interface{})
	if !ok {
		return
	}

	// Fetch trades for top 30 markets
	limit := 30
	if len(marketsList) < limit {
		limit = len(marketsList)
	}

	for i := 0; i < limit; i++ {
		market, ok := marketsList[i].(map[string]interface{})
		if !ok {
			continue
		}

		marketID := utils.ParseString(market["marketId"])
		if marketID == "" {
			continue
		}

		// Fetch recent trades
		tradesData, err := clients.FetchTrades(marketID, 100)
		if err != nil {
			log.Printf("Error fetching trades for %s: %v", marketID, err)
			continue
		}

		// Parse trades
		dc.parseTrades(marketID, tradesData)
	}

	log.Printf("Trades updated for %d markets", limit)
}

func (dc *DataCollector) parseTrades(marketID string, data map[string]interface{}) {
	tradesArray, ok := data["trades"].([]interface{})
	if !ok {
		return
	}

	trades := make([]*models.Trade, 0, len(tradesArray))

	for _, t := range tradesArray {
		tradeMap, ok := t.(map[string]interface{})
		if !ok {
			continue
		}

		price, err := utils.ParseFloat(tradeMap["price"])
		if err != nil {
			continue
		}

		quantity, err := utils.ParseFloat(tradeMap["quantity"])
		if err != nil {
			continue
		}

		timestamp := time.Now() // Would parse from API in production
		isBuy := utils.ParseString(tradeMap["tradeDirection"]) == "buy"

		trades = append(trades, &models.Trade{
			MarketID:  marketID,
			Price:     price,
			Quantity:  quantity,
			Timestamp: timestamp,
			IsBuy:     isBuy,
		})
	}

	dc.cache.SetTrades(marketID, trades)
}

// updatePriceHistory updates price history from recent data every 60 seconds
func (dc *DataCollector) updatePriceHistory() {
	defer dc.wg.Done()

	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-dc.stopChan:
			return
		case <-ticker.C:
			dc.extractPriceHistory()
		}
	}
}

func (dc *DataCollector) extractPriceHistory() {
	// Get market list from cache
	marketsData, found := dc.cache.GetMarkets()
	if !found {
		return
	}

	marketsMap, ok := marketsData.(map[string]interface{})
	if !ok {
		return
	}

	marketsList, ok := marketsMap["markets"].([]interface{})
	if !ok {
		return
	}

	for _, m := range marketsList {
		market, ok := m.(map[string]interface{})
		if !ok {
			continue
		}

		marketID := utils.ParseString(market["marketId"])
		if marketID == "" {
			continue
		}

		// Get or create price history
		history, exists := dc.cache.GetPriceHistory(marketID)
		if !exists {
			history = models.NewPriceHistory(marketID)
		}

		// Try to get price from orderbook
		orderbookData, found := dc.cache.GetOrderbook(marketID)
		if found {
			if price := dc.extractMidPrice(orderbookData); price > 0 {
				history.AddPrice(price, time.Now())
				dc.cache.SetPriceHistory(marketID, history)
			}
		}
	}

	log.Println("Price history updated")
}

func (dc *DataCollector) extractMidPrice(orderbookData interface{}) float64 {
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

	if len(buys) == 0 || len(sells) == 0 {
		return 0
	}

	return (buys[0].Price + sells[0].Price) / 2
}

// computeAnalytics calculates analytics for all markets every 15 seconds
func (dc *DataCollector) computeAnalytics() {
	defer dc.wg.Done()

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	// Initial computation
	time.Sleep(5 * time.Second) // Wait for initial data
	dc.calculateAllAnalytics()

	for {
		select {
		case <-dc.stopChan:
			return
		case <-ticker.C:
			dc.calculateAllAnalytics()
		}
	}
}

func (dc *DataCollector) calculateAllAnalytics() {
	// Get market list from cache
	marketsData, found := dc.cache.GetMarkets()
	if !found {
		return
	}

	marketsMap, ok := marketsData.(map[string]interface{})
	if !ok {
		return
	}

	marketsList, ok := marketsMap["markets"].([]interface{})
	if !ok {
		return
	}

	count := 0
	for _, m := range marketsList {
		market, ok := m.(map[string]interface{})
		if !ok {
			continue
		}

		marketID := utils.ParseString(market["marketId"])
		if marketID == "" {
			continue
		}

		// Compute analytics for this market
		analytics := services.ComputeMarketAnalytics(marketID, dc.cache)
		if analytics != nil {
			dc.cache.SetAnalytics(marketID, analytics)
			count++
		}
	}

	log.Printf("Analytics computed for %d markets", count)
}

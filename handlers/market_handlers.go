package handlers

import (
	"strconv"

	"github.com/daiwikmh/origami/services"
	"github.com/gin-gonic/gin"
)

// Raw markets
func GetMarkets(c *gin.Context) {
	data, err := services.GetMarkets()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, data)
}

// Market Summary (simplified + enriched)
func GetMarketSummary(c *gin.Context) {
	data, err := services.GetMarkets()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	markets := data["markets"].([]interface{})

	result := []gin.H{}

	for _, m := range markets {
		market := m.(map[string]interface{})

		result = append(result, gin.H{
			"marketId": market["marketId"],
			"base":     market["baseDenom"],
			"quote":    market["quoteDenom"],
		})
	}

	c.JSON(200, result)
}

// Liquidity endpoint
func GetLiquidity(c *gin.Context) {
	id := c.Param("id")

	ob, err := services.GetOrderbook(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Calculate real liquidity from orderbook depth
	liquidity := services.CalculateRealLiquidity(ob)
	depth := services.CalculateOrderbookDepth(ob)

	response := gin.H{
		"market_id":       id,
		"liquidity_score": liquidity,
	}

	if depth != nil {
		response["orderbook_depth"] = depth
	}

	c.JSON(200, response)
}

// Trending markets (using real analytics)
func GetTrending(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}
	if limit > 50 {
		limit = 50
	}

	markets := services.GetTopMarkets("trending", limit)

	c.JSON(200, gin.H{
		"markets": markets,
		"count":   len(markets),
	})
}

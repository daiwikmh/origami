package handlers

import (
	// "net/http"
	"sort"

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

	orderbook := ob["orderbook"].(map[string]interface{})

	buys := orderbook["buys"].([]interface{})
	sells := orderbook["sells"].([]interface{})

	// dummy example
	liquidity := len(buys) + len(sells)

	c.JSON(200, gin.H{
		"marketId": id,
		"liquidity": liquidity,
	})
}

// Trending markets (dummy for now)
func GetTrending(c *gin.Context) {
	data := []struct {
		Market string  `json:"market"`
		Score  float64 `json:"score"`
	}{
		{"BTC/USDT", 120},
		{"ETH/USDT", 90},
		{"INJ/USDT", 150},
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Score > data[j].Score
	})

	c.JSON(200, data)
}

package handlers

import (
	"github.com/daiwikmh/origami/services"
	"github.com/gin-gonic/gin"
)

// GetMarketAnalytics returns comprehensive analytics for a market
func GetMarketAnalytics(c *gin.Context) {
	marketID := c.Param("id")

	analytics := services.GetMarketAnalytics(marketID)
	if analytics == nil {
		c.JSON(404, gin.H{"error": "Market not found or analytics unavailable"})
		return
	}

	c.JSON(200, analytics)
}

// GetVolatility returns volatility metric for a market
func GetVolatility(c *gin.Context) {
	marketID := c.Param("id")

	analytics := services.GetMarketAnalytics(marketID)
	if analytics == nil {
		c.JSON(404, gin.H{"error": "Market not found"})
		return
	}

	c.JSON(200, gin.H{
		"market_id":  marketID,
		"volatility": analytics.Volatility,
		"timestamp":  analytics.Timestamp,
	})
}

// GetOrderbookDepth returns detailed orderbook depth metrics
func GetOrderbookDepth(c *gin.Context) {
	marketID := c.Param("id")

	analytics := services.GetMarketAnalytics(marketID)
	if analytics == nil || analytics.OrderbookDepth == nil {
		c.JSON(404, gin.H{"error": "Market not found or orderbook unavailable"})
		return
	}

	c.JSON(200, analytics.OrderbookDepth)
}

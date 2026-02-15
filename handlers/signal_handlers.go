package handlers

import (
	"strconv"

	"github.com/daiwikmh/origami/services"
	"github.com/gin-gonic/gin"
)

// GetHotMarkets returns trending markets with high scores
func GetHotMarkets(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	if limit > 50 {
		limit = 50 // Cap at 50
	}

	markets := services.GetTopMarkets("trending", limit)

	c.JSON(200, gin.H{
		"markets": markets,
		"count":   len(markets),
	})
}

// GetVolatilityRanking returns markets sorted by volatility
func GetVolatilityRanking(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	if limit > 50 {
		limit = 50
	}

	markets := services.GetTopMarkets("volatility", limit)

	c.JSON(200, gin.H{
		"markets": markets,
		"count":   len(markets),
	})
}

// GetVolumeLeaders returns markets sorted by 24h volume
func GetVolumeLeaders(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	if limit > 50 {
		limit = 50
	}

	markets := services.GetTopMarkets("volume", limit)

	c.JSON(200, gin.H{
		"markets": markets,
		"count":   len(markets),
	})
}

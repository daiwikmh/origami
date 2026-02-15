package handlers

import (
	"github.com/daiwikmh/origami/auth"
	"github.com/daiwikmh/origami/models"
	"github.com/gin-gonic/gin"
)

var keyStore *auth.KeyStore

// InitAdminHandlers initializes admin handlers with key store
func InitAdminHandlers(ks *auth.KeyStore) {
	keyStore = ks
}

// GenerateAPIKey creates a new API key
func GenerateAPIKey(c *gin.Context) {
	var req struct {
		Name      string `json:"name" binding:"required"`
		RateLimit int    `json:"rate_limit"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	// Default rate limit: 100 requests per minute
	if req.RateLimit == 0 {
		req.RateLimit = 100
	}

	apiKey := keyStore.GenerateKey(req.Name, req.RateLimit)

	c.JSON(201, gin.H{
		"api_key":    apiKey.Key,
		"name":       apiKey.Name,
		"rate_limit": apiKey.RateLimit,
		"created_at": apiKey.CreatedAt,
		"message":    "API key created successfully. Store it securely - it won't be shown again.",
	})
}

// ListAPIKeys returns all API keys (without showing the actual key)
func ListAPIKeys(c *gin.Context) {
	keys := keyStore.ListKeys()

	result := make([]gin.H, 0, len(keys))
	for _, key := range keys {
		result = append(result, gin.H{
			"key_preview":   maskKey(key.Key),
			"name":          key.Name,
			"created_at":    key.CreatedAt,
			"last_used_at":  key.LastUsedAt,
			"request_count": key.RequestCount,
			"rate_limit":    key.RateLimit,
			"is_active":     key.IsActive,
		})
	}

	c.JSON(200, gin.H{
		"keys":  result,
		"count": len(result),
	})
}

// RevokeAPIKey deactivates an API key
func RevokeAPIKey(c *gin.Context) {
	var req struct {
		Key string `json:"key" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	success := keyStore.RevokeKey(req.Key)
	if !success {
		c.JSON(404, gin.H{"error": "api key not found"})
		return
	}

	c.JSON(200, gin.H{"message": "API key revoked successfully"})
}

// GetUsageStats returns usage statistics
func GetUsageStats(c *gin.Context) {
	stats := keyStore.GetUsageStats()
	c.JSON(200, stats)
}

// GetKeyUsage returns usage for a specific key
func GetKeyUsage(c *gin.Context) {
	// Get API key from context
	apiKey, exists := c.Get("api_key")
	if !exists {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	keyStr := apiKey.(string)

	// Get usage stats
	stats := keyStore.GetUsageStats()
	keyStats, exists := stats.KeyStats[keyStr]
	if !exists {
		c.JSON(404, gin.H{"error": "usage data not found"})
		return
	}

	c.JSON(200, keyStats)
}

// GetRateLimitInfo returns rate limit info for current key
func GetRateLimitInfo(c *gin.Context) {
	apiKeyObj, exists := c.Get("api_key_obj")
	if !exists {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	key := apiKeyObj.(*models.APIKey)

	c.JSON(200, gin.H{
		"rate_limit":    key.RateLimit,
		"window":        "1 minute",
		"request_count": key.RequestCount,
	})
}

// maskKey masks an API key for display
func maskKey(key string) string {
	if len(key) < 12 {
		return "***"
	}
	return key[:6] + "..." + key[len(key)-6:]
}

// AdminDashboard serves the admin dashboard HTML
func AdminDashboard(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, dashboardHTML)
}

// TestAPIEndpoint serves the API testing interface
func TestAPIEndpoint(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, testingHTML)
}

// GetSystemInfo returns system information and available endpoints
func GetSystemInfo(c *gin.Context) {
	endpoints := []gin.H{
		{
			"path":        "/origami/markets",
			"method":      "GET",
			"description": "Get all spot markets from Injective",
		},
		{
			"path":        "/origami/markets/summary",
			"method":      "GET",
			"description": "Get simplified market summary",
		},
		{
			"path":        "/origami/markets/:id/liquidity",
			"method":      "GET",
			"description": "Get liquidity metrics for a market",
		},
		{
			"path":        "/origami/markets/:id/analytics",
			"method":      "GET",
			"description": "Get comprehensive analytics for a market",
		},
		{
			"path":        "/origami/markets/:id/volatility",
			"method":      "GET",
			"description": "Get volatility indicator for a market",
		},
		{
			"path":        "/origami/markets/:id/depth",
			"method":      "GET",
			"description": "Get orderbook depth for a market",
		},
		{
			"path":        "/origami/signals/trending",
			"method":      "GET",
			"description": "Get trending markets",
			"params":      "?limit=10",
		},
		{
			"path":        "/origami/signals/hot",
			"method":      "GET",
			"description": "Get hot markets with highest scores",
			"params":      "?limit=10",
		},
		{
			"path":        "/origami/signals/volatile",
			"method":      "GET",
			"description": "Get most volatile markets",
			"params":      "?limit=10",
		},
		{
			"path":        "/origami/signals/volume",
			"method":      "GET",
			"description": "Get volume leaders",
			"params":      "?limit=10",
		},
		{
			"path":        "/origami/nft/verify/:address",
			"method":      "GET",
			"description": "Verify NFT ownership for an address",
		},
		{
			"path":        "/origami/nft/verify/batch",
			"method":      "POST",
			"description": "Batch verify NFT ownership for multiple addresses",
		},
	}

	c.JSON(200, gin.H{
		"name":        "Origami API",
		"version":     "1.0.0",
		"description": "Production-ready Injective intelligence API with authentication and rate limiting",
		"endpoints":   endpoints,
		"docs":        "/dashboard",
		"test":        "/test",
	})
}

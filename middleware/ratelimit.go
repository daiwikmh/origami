package middleware

import (
	"github.com/daiwikmh/origami/auth"
	"github.com/daiwikmh/origami/models"
	"github.com/gin-gonic/gin"
)

// RateLimiter enforces per-key rate limits
func RateLimiter(keyStore *auth.KeyStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get API key from context (set by auth middleware)
		apiKey, exists := c.Get("api_key")
		if !exists {
			c.JSON(500, gin.H{"error": "internal server error"})
			c.Abort()
			return
		}

		keyStr := apiKey.(string)

		// Get key object to check rate limit
		keyObj, exists := c.Get("api_key_obj")
		if !exists {
			c.JSON(500, gin.H{"error": "internal server error"})
			c.Abort()
			return
		}

		key := keyObj.(*models.APIKey)

		// Check rate limit
		if !keyStore.CheckRateLimit(keyStr, key.RateLimit) {
			c.JSON(429, gin.H{
				"error":       "rate limit exceeded",
				"rate_limit":  key.RateLimit,
				"window":      "1 minute",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

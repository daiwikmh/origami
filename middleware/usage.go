package middleware

import (
	"github.com/daiwikmh/origami/auth"
	"github.com/gin-gonic/gin"
)

// UsageTracker tracks API usage per key
func UsageTracker(keyStore *auth.KeyStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get API key from context (set by auth middleware)
		apiKey, exists := c.Get("api_key")
		if !exists {
			c.Next()
			return
		}

		keyStr := apiKey.(string)
		endpoint := c.Request.Method + " " + c.FullPath()

		// Update last used timestamp
		keyStore.UpdateLastUsed(keyStr)

		// Track endpoint usage
		keyStore.TrackEndpointUsage(keyStr, endpoint)

		c.Next()
	}
}

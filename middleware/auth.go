package middleware

import (
	"strings"

	"github.com/daiwikmh/origami/auth"
	"github.com/gin-gonic/gin"
)

// APIKeyAuth validates API key from Authorization header
func APIKeyAuth(keyStore *auth.KeyStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// Check if Authorization header is present
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "missing api key"})
			c.Abort()
			return
		}

		// Extract Bearer token
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "invalid api key"})
			c.Abort()
			return
		}

		apiKey := parts[1]

		// Validate key
		key, valid := keyStore.ValidateKey(apiKey)
		if !valid {
			c.JSON(401, gin.H{"error": "invalid api key"})
			c.Abort()
			return
		}

		// Store key in context for later use
		c.Set("api_key", apiKey)
		c.Set("api_key_obj", key)

		c.Next()
	}
}

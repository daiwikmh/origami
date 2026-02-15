package main

import (
	"github.com/daiwikmh/origami/auth"
	"github.com/daiwikmh/origami/handlers"
	"github.com/daiwikmh/origami/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(keyStore *auth.KeyStore) *gin.Engine {
	r := gin.Default()

	// Public endpoints (no auth required)
	r.GET("/", handlers.GetSystemInfo)
	r.GET("/dashboard", handlers.AdminDashboard)
	r.GET("/test", handlers.TestAPIEndpoint)
	r.GET("/docs", handlers.ServeDocs)

	// Admin endpoints (no auth for demo purposes - add auth in production)
	admin := r.Group("/admin")
	{
		admin.POST("/keys/generate", handlers.GenerateAPIKey)
		admin.GET("/keys", handlers.ListAPIKeys)
		admin.POST("/keys/revoke", handlers.RevokeAPIKey)
		admin.GET("/usage", handlers.GetUsageStats)
	}

	// Protected API routes under /origami namespace
	origami := r.Group("/origami")
	origami.Use(middleware.APIKeyAuth(keyStore))
	origami.Use(middleware.RateLimiter(keyStore))
	origami.Use(middleware.UsageTracker(keyStore))
	{
		// Market endpoints
		origami.GET("/markets", handlers.GetMarkets)
		origami.GET("/markets/summary", handlers.GetMarketSummary)
		origami.GET("/markets/:id/liquidity", handlers.GetLiquidity)

		// Analytics endpoints
		origami.GET("/markets/:id/analytics", handlers.GetMarketAnalytics)
		origami.GET("/markets/:id/volatility", handlers.GetVolatility)
		origami.GET("/markets/:id/depth", handlers.GetOrderbookDepth)

		// Signal endpoints
		origami.GET("/signals/trending", handlers.GetTrending)
		origami.GET("/signals/hot", handlers.GetHotMarkets)
		origami.GET("/signals/volatile", handlers.GetVolatilityRanking)
		origami.GET("/signals/volume", handlers.GetVolumeLeaders)

		// User endpoints
		origami.GET("/me", handlers.GetKeyUsage)
		origami.GET("/me/limits", handlers.GetRateLimitInfo)

		// NFT verification endpoints
		origami.GET("/nft/verify/:address", handlers.VerifyNFTOwnership)
		origami.POST("/nft/verify/batch", handlers.BatchVerifyNFTOwnership)
	}

	return r
}

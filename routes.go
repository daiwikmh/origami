package main

import (
	"github.com/daiwikmh/origami/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/markets", handlers.GetMarkets)
	r.GET("/markets/summary", handlers.GetMarketSummary)
	r.GET("/markets/:id/liquidity", handlers.GetLiquidity)
	r.GET("/signals/trending", handlers.GetTrending)

	return r
}

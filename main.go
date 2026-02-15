package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/daiwikmh/origami/auth"
	"github.com/daiwikmh/origami/cache"
	"github.com/daiwikmh/origami/handlers"
	"github.com/daiwikmh/origami/services"
	"github.com/daiwikmh/origami/workers"
)

func main() {
	log.Println("Initializing Origami API Platform...")

	// Initialize API key store
	keyStore := auth.NewKeyStore()
	log.Println("API key store initialized")

	// Print default API key for testing
	keys := keyStore.ListKeys()
	if len(keys) > 0 {
		fmt.Println("\n" + strings.Repeat("=", 70))
		fmt.Println("  DEFAULT API KEY FOR TESTING")
		fmt.Println(strings.Repeat("=", 70))
		fmt.Printf("  Name: %s\n", keys[0].Name)
		fmt.Printf("  Key:  %s\n", keys[0].Key)
		fmt.Printf("  Rate Limit: %d requests/minute\n", keys[0].RateLimit)
		fmt.Println(strings.Repeat("=", 70))
		fmt.Println()
	}

	// Initialize cache
	dataCache := cache.NewDataCache()
	log.Println("Cache initialized")

	// Initialize services with cache
	services.InitMarketService(dataCache)
	log.Println("Services initialized")

	// Initialize handlers
	handlers.InitAdminHandlers(keyStore)
	log.Println("Handlers initialized")

	// Start background workers
	collector := workers.NewDataCollector(dataCache)
	collector.Start()

	// Setup HTTP server
	r := SetupRouter(keyStore)
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in goroutine
	go func() {
		log.Println("Starting server on :8080...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	log.Println("Server started successfully on :8080")
	log.Println("API is ready to accept requests")
	log.Println("\nAccess the dashboard at: http://localhost:8080/dashboard")
	log.Println("Test API endpoints at: http://localhost:8080/test")

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Stop background workers
	collector.Stop()

	// Shutdown HTTP server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited gracefully")
}

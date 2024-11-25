package main

import (
	"context"
	"go-server/internal/config"
	"go-server/internal/db"
	"go-server/internal/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to MongoDB
	db.Connect(cfg)

	// Initialize router
	router := routes.Setup()

	// Create HTTP server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Run server in a goroutine
	go func() {
		log.Println("Server running on port 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %s", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown error: %s", err)
	}

	db.Disconnect()
	log.Println("Server gracefully stopped")
}

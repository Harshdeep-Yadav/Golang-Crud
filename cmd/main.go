package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/harsh/book_crud/config"
	"github.com/harsh/book_crud/middleware"
	"github.com/harsh/book_crud/routes"
	"net/http"
)

func main() {
	// Initialize logger
	logger := log.New(os.Stdout, "[BOOK-CRUD] ", log.LstdFlags)

	// Initialize DB
	if err := config.ConnectDB(); err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}
	logger.Println("Successfully connected to database")

	// Create router
	router := mux.NewRouter()

	// Add middleware
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.RecoveryMiddleware)
	router.Use(middleware.CORSMiddleware)

	// Register routes
	routes.RegisterBookRoutes(router)
	logger.Println("Routes registered successfully")

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	serverAddr := ":" + port

	// Create server with timeouts
	server := &http.Server{
		Addr:         serverAddr,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Channel to listen for errors coming from the server
	serverErrors := make(chan error, 1)

	// Start server in a goroutine
	go func() {
		logger.Printf("Server is running on port %s", port)
		serverErrors <- server.ListenAndServe()
	}()

	// Channel to listen for an interrupt or terminate signal from the OS
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Blocking select waiting for either a server error or a shutdown signal
	select {
	case err := <-serverErrors:
		logger.Fatalf("Error starting server: %v", err)
	case sig := <-shutdown:
		logger.Printf("Server is shutting down due to %v signal", sig)
		// Give outstanding requests 5 seconds to complete
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not stop server gracefully: %v", err)
		}
	}
}

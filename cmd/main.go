package main

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/harsh/book_crud/config"

	"github.com/harsh/book_crud/routes"

	"net/http"
)

func main() {

	// Initialize DB
	config.ConnectDB()

	// Create router
	router := mux.NewRouter()

	// // Register routes
	routes.RegisterBookRoutes(router)
	fmt.Println("Routes registered")

	// Start server
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

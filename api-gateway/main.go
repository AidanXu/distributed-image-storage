package main

import (
	"log"
	"net/http"

	"api-gateway/internal/handlers"
)

func main() {
    http.HandleFunc("/", handlers.BaseHandler)

    log.Println("Starting API gateway on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

package main

import (
	"api-gateway/internal/handlers"
	"log"
	"net/http"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    // Explicitly handle OPTIONS method for preflight requests
    r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Just return no content - CORS headers are added at the middleware
        w.WriteHeader(http.StatusNoContent)
    })

    r.HandleFunc("/login", handlers.LoginHandler).Methods("POST", "OPTIONS")
    r.HandleFunc("/storage", handlers.StorageHandler).Methods("GET", "POST", "OPTIONS")

    // Setup CORS to allow everything for demonstration; adjust as needed
    corsHandler := gorillaHandlers.CORS(
        gorillaHandlers.AllowedOrigins([]string{"*"}), // This should be adjusted in production
        gorillaHandlers.AllowedMethods([]string{"POST", "GET", "OPTIONS"}),
        gorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
        gorillaHandlers.AllowCredentials(),
    )(r)

    log.Println("Starting API gateway on port 8080...")
    if err := http.ListenAndServe(":8080", corsHandler); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
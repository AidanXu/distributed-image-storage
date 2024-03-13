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

    r.PathPrefix("/login").HandlerFunc(handlers.LoginHandler).Methods("POST", "OPTIONS")
    r.PathPrefix("/storage").HandlerFunc(handlers.StorageHandler).Methods("GET", "POST", "OPTIONS")

    corsHandler := gorillaHandlers.CORS(
        gorillaHandlers.AllowedOrigins([]string{"*"}), // Allow all origins
        gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // Allow common methods
        gorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), // Allow common headers
        gorillaHandlers.AllowCredentials(),
    )(r)


    log.Println("Starting API gateway on port 8080...")
    if err := http.ListenAndServe(":8080", corsHandler); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
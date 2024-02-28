package main

import (
	handlers "api-gateway/internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.PathPrefix("/login").HandlerFunc(handlers.LoginHandler)

    log.Println("Starting API gateway on port 8080...")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

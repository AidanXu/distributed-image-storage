package main

import (
	"log"
	"net/http"

	auth "login-service/internal/auth"
	handlers "login-service/internal/handlers"
)

func main() {
	auth.InitDB("host=db user=testuser password=12345 dbname=testdb port=5432 sslmode=disable")
	auth.DB.AutoMigrate(&auth.User{})


    http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/login/register", handlers.RegisterHandler)

    log.Println("Login service listening on port 9090...")
    if err := http.ListenAndServe(":9090", nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

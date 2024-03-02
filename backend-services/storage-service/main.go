package main

import (
	"log"
	"net/http"

	handlers "storage-service/internal/handlers"
)

func main() {

    http.HandleFunc("/storage/upload", handlers.UploadHandler)
    http.HandleFunc("/storage", handlers.GetPhotos)

    log.Println("Storage service listening on port 9091...")
    if err := http.ListenAndServe(":9091", nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

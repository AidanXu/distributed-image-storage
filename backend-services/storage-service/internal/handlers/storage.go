package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	auth "storage-service/internal/auth"

	jwt "github.com/golang-jwt/jwt/v4"
	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func init() {
    var err error
    // Initialize the MinIO client
    minioClient, err = minio.New("minio:9000", &minio.Options{
        Creds:  credentials.NewStaticV4("minioadmin", "minioadmin", ""),
        Secure: false, 
    })
    if err != nil {
        log.Fatalf("Could not initialize MinIO client: %v", err)
    }
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
    token, err := auth.ValidateToken(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        http.Error(w, "Cannot extract claims", http.StatusInternalServerError)
        return
    }

    username, ok := claims["sub"].(string)
    if !ok {
        http.Error(w, "Cannot extract username from token", http.StatusInternalServerError)
        return
    }

    // Parse multipart form with a max memory of 50MB.
    if err := r.ParseMultipartForm(50 << 20); err != nil {
        http.Error(w, "Error parsing multipart form: "+err.Error(), http.StatusBadRequest)
        return
    }

    files := r.MultipartForm.File["image"]
    var uploadedFiles []string

    for _, fileHeader := range files {
        file, err := fileHeader.Open()
        if err != nil {
            http.Error(w, "Invalid file upload", http.StatusBadRequest)
            return
        }
        defer file.Close()

        fileName := fileHeader.Filename
        fileType := fileHeader.Header.Get("Content-Type")

        if fileType != "image/jpeg" && fileType != "image/png" {
            http.Error(w, "Unsupported file type", http.StatusBadRequest)
            return
        }

        // Construct the object name using the username for organization
        objectName := fmt.Sprintf("%s/%s", username, fileName)

        // Upload the file to MinIO
        _, err = minioClient.PutObject(r.Context(), "photobucket", objectName, file, fileHeader.Size, minio.PutObjectOptions{ContentType: fileType})
        if err != nil {
            log.Printf("Failed to upload: %v", err)
            http.Error(w, "Failed to upload file", http.StatusInternalServerError)
            return
        }

        uploadedFiles = append(uploadedFiles, fileName)
    }

    response := map[string]interface{}{
        "message": "Files uploaded successfully",
        "uploadedFiles": uploadedFiles,
    }

    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Error encoding response", http.StatusInternalServerError)
        return
    }

    log.Printf("Uploaded files: %v", uploadedFiles)
}
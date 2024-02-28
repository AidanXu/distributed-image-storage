package handlers

import (
	"encoding/json"
	"net/http"

	dto "login-service/internal/DTO"
	auth "login-service/internal/auth"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        var loginReq dto.LoginRequest
        err := json.NewDecoder(r.Body).Decode(&loginReq)
        if err != nil {
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }

        var user auth.User
        result := auth.DB.Where("username = ?", loginReq.Username).First(&user)
        if result.Error != nil {
            http.Error(w, "Invalid username or password", http.StatusUnauthorized)
            return
        }

        err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
        if err != nil {
            http.Error(w, "Invalid username or password", http.StatusUnauthorized)
            return
        }

        //TODO: add jwt here probably
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
    } else {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var regReq dto.RegisterRequest
    if err := json.NewDecoder(r.Body).Decode(&regReq); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if err := auth.CreateUser(regReq.Username, regReq.Password); err != nil {
        http.Error(w, "Failed to register user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Registration successful"})
}
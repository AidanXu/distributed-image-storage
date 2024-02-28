package handlers

import (
	"encoding/json"
	"net/http"

	dto "login-service/internal/DTO"
	auth "login-service/internal/auth"

	jwt "github.com/dgrijalva/jwt-go"
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

        // Generate JWT
        tokenString, err := GenerateJWT(user.Username)
        if err != nil {
            http.Error(w, "Failed to generate token", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{
            "token":   tokenString,
            "message": "Login successful",
        })
    } else {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var regReq dto.LoginRequest
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

var jwtKey = []byte("xdd123") // temp secret for now

func GenerateJWT(username string) (string, error) {
	claims := &jwt.StandardClaims{
		Subject:   username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
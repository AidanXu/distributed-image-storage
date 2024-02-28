package handlers

import (
	"encoding/json"
	"net/http"
)

func StorageRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("reached storage endpoint")
}

// var jwtKey = []byte("xdd123") // temp secret for now

// func ValidateToken(w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
// 	authHeader := r.Header.Get("Authorization")
// 	if authHeader == "" {
// 		return nil, errors.New("authorization header is required")
// 	}

// 	// Expecting "Bearer <token>"
// 	bearerToken := strings.Split(authHeader, " ")
// 	if len(bearerToken) != 2 {
// 		return nil, errors.New("invalid token format")
// 	}

// 	tokenString := bearerToken[1]
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Validate the alg is what you expect:
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("unexpected signing method")
// 		}

// 		return jwtKey, nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	if !token.Valid {
// 		return nil, errors.New("invalid token")
// 	}

// 	return token, nil
// }
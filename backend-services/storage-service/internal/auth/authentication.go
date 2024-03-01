package auth

import (
	"errors"
	"net/http"
	"strings"

	jwt "github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("xdd123") // temp secret for now

func ValidateToken(r *http.Request) (*jwt.Token, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("Authorization header is required")
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		return nil, errors.New("Invalid token format")
	}

	tokenString := bearerToken[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return token, nil
}
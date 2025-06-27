package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"os"
)

var AccessTokenSecret = []byte(os.Getenv("ACCESS_SECRET"))
var RefreshTokenSecret = []byte(os.Getenv("REFRESH_SECRET"))

func GenerateTokens(userID, email string) (string, string, error) {
	accessToken, err := generateToken(userID, email, AccessTokenSecret, 15*time.Minute)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := generateToken(userID, email, RefreshTokenSecret, 7*24*time.Hour)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func generateToken(userID, email string, secret []byte, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func VerifyToken(tokenString string, secret []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"os"
)

var AccessTokenSecret = []byte(os.Getenv("ACCESS_SECRET"))
var RefreshTokenSecret = []byte(os.Getenv("REFRESH_SECRET"))

func GenerateTokens(userID, email, role string) (string, string, error) {
	accessToken, err := generateToken(userID, email, role, AccessTokenSecret, 15*time.Minute)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := generateToken(userID, email, role, RefreshTokenSecret, 7*24*time.Hour)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func generateToken(userID, email, role string, secret []byte, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"email":  email,
		"role":   role,
		"exp":    time.Now().Add(duration).Unix(),
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

type UserClaims struct {
	UserID string
	Email  string
	Role   string
}

func ParseUserClaims(tokenStr string, isAccessToken bool) (*UserClaims, error) {
	secret := AccessTokenSecret
	if !isAccessToken {
		secret = RefreshTokenSecret
	}

	claimsMap, err := VerifyToken(tokenStr, secret)
	if err != nil {
		return nil, err
	}

	userID, _ := claimsMap["userId"].(string)
	email, _ := claimsMap["email"].(string)
	role, _ := claimsMap["role"].(string)

	if userID == "" || email == "" || role == "" {
		return nil, errors.New("invalid claims data")
	}

	return &UserClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
	}, nil
}

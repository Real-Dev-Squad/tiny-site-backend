package utils

import (
	"errors"
	"time"

	"github.com/Real-Dev-Squad/tiny-site-backend/config"
	"github.com/Real-Dev-Squad/tiny-site-backend/models"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrInvalidToken            = errors.New("invalid token")
	ErrTokenExpired            = errors.New("token has expired")
)

func GenerateToken(user *models.User) (string, error) {
	key := []byte(config.JwtSecret)
	expiryTime := time.Now().Add(time.Duration(config.JwtValidity) * time.Hour).UTC()

	claims := jwt.MapClaims{
		"iss":    config.JwtIssuer,
		"exp":    expiryTime.Unix(),
		"iat":    time.Now().UTC().Unix(),
		"email":  user.Email,
		"userID": user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(key)
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	key := []byte(config.JwtSecret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigningMethod
		}
		return key, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
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

	// Parsing the token

	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {

		//validatint the algo
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigningMethod
		}
		return key, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, ErrTokenExpired
		}
		return nil, err
	}

	// Validating  the token and casting the claims :P
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

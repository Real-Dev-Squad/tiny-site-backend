package utils

import (
	"errors"
	"time"

	"github.com/Real-Dev-Squad/tiny-site-backend/config"
	"github.com/Real-Dev-Squad/tiny-site-backend/models"
	"github.com/golang-jwt/jwt/v5"
)

/*
 * GenerateToken generates a JWT token for the user
 */
func GenerateToken(user *models.User) (string, error) {
	issuer := config.JwtIssuer
	key := []byte(config.JwtSecret)
	tokenValidityInHours := config.JwtValidity
	tokenExpiryTime := time.Now().Add(time.Duration(tokenValidityInHours) * time.Hour).UTC().Unix()
	claims := jwt.MapClaims{
		"iss":    issuer,
		"exp":    tokenExpiryTime,
		"email":  user.Email,
		"userID": user.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(key)
	return tokenString, err
}

/*
 * VerifyToken verifies the token and returns the email of the user
 */
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, errors.New("token has expired")
		}
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/Real-Dev-Squad/tiny-site-backend/models"
	"github.com/golang-jwt/jwt/v5"
)

/*
 * GenerateToken generates a JWT token for the user
 */
func GenerateToken(user *models.User) (string, error) {
	issuer := os.Getenv("JWT_ISSUER")
	key := []byte(os.Getenv("JWT_SECRET"))

	tokenValidityInHours, err := strconv.ParseInt(os.Getenv("JWT_VALIDITY_IN_HOURS"), 10, 8)

	if err != nil {
		return "", err
	}

	tokenExpiryTime := time.Now().Add(time.Second * time.Duration(tokenValidityInHours)).UTC().Format(time.RFC3339)

	t := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"iss":   issuer,
		"exp":   tokenExpiryTime,
		"email": user.Email,
	})

	token, error := t.SignedString(key)

	return token, error
}

/*
 * VerifyToken verifies the token and returns the email of the user
 */
func VerifyToken(tokenString string) (string, error) {
	var claims jwt.MapClaims = nil

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if token.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if c, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return "", err
	} else {
		claims = c
	}

	expiryTime, err := time.Parse(time.RFC3339, claims["exp"].(string))

	if err != nil {
		return "", err
	}

	if time.Now().UTC().After(expiryTime) {
		return "", errors.New("token has expired")
	}

	return claims["email"].(string), nil
}

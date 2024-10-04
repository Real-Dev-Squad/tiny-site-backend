package unit

import (
	"os"
	"testing"
	"time"

	"github.com/Real-Dev-Squad/tiny-site-backend/config"
	"github.com/Real-Dev-Squad/tiny-site-backend/models"
	"github.com/Real-Dev-Squad/tiny-site-backend/utils"
	"github.com/golang-jwt/jwt/v5"
)

func TestMain(m *testing.M) {
	utils.LoadEnv("../../environments/dev.env")

	code := m.Run()

	os.Exit(code)
}

func TestGenerateJWT(t *testing.T) {
	dummyUser := &models.User{
		Email: "test@gmail.com",
		ID:    123,
	}

	token, err := utils.GenerateToken(dummyUser)

	if err != nil {
		t.Fatalf("Expected nil but got %v", err)
	}

	if len(token) == 0 {
		t.Fatalf("Empty token of length")
	}
}

func TestVerifyJWT(t *testing.T) {
	t.Run("ValidToken", func(t *testing.T) {
		dummyUser := &models.User{
			Email: "test@gmail.com",
			ID:    123,
		}

		validToken, generateTokenError := utils.GenerateToken(dummyUser)
		if generateTokenError != nil {
			t.Fatalf("Error: %v", generateTokenError)
		}

		claims, validTokenError := utils.VerifyToken(validToken)
		if validTokenError != nil {
			t.Fatalf("Error: %v", validTokenError)
		}

		if claims["email"] != dummyUser.Email {
			t.Fatalf("Expected email %v but got %v", dummyUser.Email, claims["email"])
		}

		if claims["userID"] != float64(dummyUser.ID) {
			t.Fatalf("Expected userID %v but got %v", dummyUser.ID, claims["userID"])
		}
	})

	t.Run("ExpiredToken", func(t *testing.T) {
		expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
			"iss":    config.JwtIssuer,
			"exp":    time.Now().Add(-time.Hour).Unix(),
			"email":  "test@gmail.com",
			"userID": 123,
		})

		key := []byte(config.JwtSecret)
		expiredTokenString, _ := expiredToken.SignedString(key)

		_, expiredTokenError := utils.VerifyToken(expiredTokenString)
		if expiredTokenError != utils.ErrTokenExpired {
			t.Fatalf("Expected error %v but got %v", utils.ErrTokenExpired, expiredTokenError)
		}
	})

	t.Run("InvalidToken", func(t *testing.T) {
		invalidToken := "invalid.token.here"

		_, invalidTokenError := utils.VerifyToken(invalidToken)
		if invalidTokenError == nil {
			t.Fatalf("Expected an error but got nil")
		}
	})
}

func TestVerifyJWTForOneYear(t *testing.T) {
	os.Setenv("JWT_VALIDITY_IN_HOURS", "8760")
	defer os.Setenv("JWT_VALIDITY_IN_HOURS", "24")

	dummyUser := &models.User{
		Email: "test@gmail.com",
		ID:    123,
	}

	validToken, generateTokenError := utils.GenerateToken(dummyUser)
	if generateTokenError != nil {
		t.Fatalf("Error: %v", generateTokenError)
	}

	claims, validTokenError := utils.VerifyToken(validToken)
	if validTokenError != nil {
		t.Fatalf("Error: %v", validTokenError)
	}

	if claims["email"] != dummyUser.Email {
		t.Fatalf("Expected email %v but got %v", dummyUser.Email, claims["email"])
	}

	if claims["userID"] != float64(dummyUser.ID) {
		t.Fatalf("Expected userID %v but got %v", dummyUser.ID, claims["userID"])
	}
}

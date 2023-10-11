package unit

import (
	"os"
	"testing"

	"github.com/Real-Dev-Squad/tiny-site-backend/models"
	"github.com/Real-Dev-Squad/tiny-site-backend/utils"
)

func TestMain(m *testing.M) {
	utils.LoadEnv("../../.env")

	code := m.Run()

	os.Exit(code)
}

func TestGenerateJWT(t *testing.T) {
	dummyUser := &models.User{
		Email: "test@gmail.com",
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
	dummyUser := &models.User{
		Email: "test@gmail.com",
	}

	validToken, generateTokenError := utils.GenerateToken(dummyUser)
	expiredToken := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZ21haWwuY29tIiwiZXhwIjoiMjAyMy0xMC0wMVQxOTo1Njo0OS4zOTc5NzEyWiIsImlzcyI6Indpc2VlLWJhY2tlbmQifQ.h11JtaPg-ITKR8UXTyz_Q7pJU_3gYyXwIkqX7lI1UK2nVkvxQvkyN23-u3wj8fV5mNIvp-ePTOp-7odsPcGC_g"

	if generateTokenError != nil {
		t.Fatalf("Error: %v", generateTokenError)
	}

	_, expiredTokenError := utils.VerifyToken(expiredToken)

	if expiredTokenError == nil {
		t.Fatalf("Expected error but got nil")
	}

	email, validTokenError := utils.VerifyToken(validToken)

	if validTokenError != nil {
		t.Fatalf("Error: %v", validTokenError)
	}

	if email != dummyUser.Email {
		t.Fatalf("Expected %v but got %v", dummyUser.Email, email)
	}
}

package tests

import (
	"os"
	"testing"

	"github.com/Real-Dev-Squad/tiny-site-backend/models"
	"github.com/Real-Dev-Squad/tiny-site-backend/utils"
	"github.com/uptrace/bun"
)

var db *bun.DB

func TestMain(m *testing.M) {
	utils.LoadEnv("../../.env")
	dsn := os.Getenv("TEST_DB_URL")
	db = utils.SetupDBConnection(dsn)

	defer db.Close()

	code := m.Run()

	os.Exit(code)
}

func generateValidAuthToken() string {
	user := &models.User{
		UserName: "testuser",
		Email:    "test@example.com",
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		panic(err)
	}

	return token
}

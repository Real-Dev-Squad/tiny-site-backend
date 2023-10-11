package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Real-Dev-Squad/tiny-site-backend/routes"
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

func TestGetUsers(t *testing.T) {
	router := routes.SetupV1Routes(db)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/v1/users", nil)

	router.ServeHTTP(w, req)

	if err != nil {
		t.Fatal(err)
	}

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}
}

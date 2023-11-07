package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Real-Dev-Squad/tiny-site-backend/routes"
)

func TestLogout(t *testing.T) {
	router := routes.SetupV1Routes(db)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/v1/auth/logout", nil)

	router.ServeHTTP(w, req)

	if err != nil {
		t.Fatal(err)
	}

	if w.Code != http.StatusFound {
		t.Errorf("Expected status code %d but got %d", http.StatusFound, w.Code)
	}
}

func TestLogin(t *testing.T) {
	router := routes.SetupV1Routes(db)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/v1/auth/google/login", nil)

	router.ServeHTTP(w, req)

	if err != nil {
		t.Fatal(err)
	}

	if w.Code != http.StatusFound {
		t.Errorf("Expected status code %d but got %d", http.StatusFound, w.Code)
	}
}

func TestCallback(t *testing.T) {
	router := routes.SetupV1Routes(db)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/v1/auth/google/callback", nil)

	router.ServeHTTP(w, req)

	if err != nil {
		t.Fatal(err)
	}

	if w.Code != http.StatusFound {
		t.Errorf("Expected status code %d but got %d", http.StatusFound, w.Code)
	}
}
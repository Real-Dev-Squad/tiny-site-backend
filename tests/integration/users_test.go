package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"time"

	controller "github.com/Real-Dev-Squad/tiny-site-backend/controllers"
	"github.com/Real-Dev-Squad/tiny-site-backend/dtos"
	"github.com/Real-Dev-Squad/tiny-site-backend/routes"
	"github.com/gin-gonic/gin"
)

func TestGetUsers(t *testing.T) {
	router := gin.Default()

	router.GET("v1/users", func(ctx *gin.Context) {
		controller.GetUserList(ctx, db)
	})
	// routes.UserRoutes(router.Group("/v1"), db)

	w := httptest.NewRecorder()

	token := generateValidAuthToken()

	if token == "" {
		t.Fatal("Failed to generate valid auth token")
	}

	req, err := http.NewRequest("GET", "/v1/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.AddCookie(&http.Cookie{Name: "token", Value: token})

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}
}

func TestGetUsersUnauthorized(t *testing.T) {

	router := gin.Default()

	router.GET("v1/users", func(ctx *gin.Context) {
		controller.GetUserList(ctx, db)
	})
	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/v1/users", nil)

	req.Header.Set("Authorization", "Bearer invalid_token")
	router.ServeHTTP(w, req)

	if err != nil {
		t.Fatal(err)
	}

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d but got %d", http.StatusUnauthorized, w.Code)
	}
}

func TestGetSelfUnauthorized(t *testing.T) {
	router := routes.SetupV1Routes(db)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/v1/users/self", nil)

	req.Header.Set("Authorization", "Bearer invalid_token")
	router.ServeHTTP(w, req)

	if err != nil {
		t.Fatal(err)
	}

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d but got %d", http.StatusUnauthorized, w.Code)
	}
}

func TestGetUserById(t *testing.T) {
	router := gin.Default()
	routes.UserRoutes(router.Group("/v1"), db)

	token := generateValidAuthToken()

	req, err := http.NewRequest("GET", "/v1/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.AddCookie(&http.Cookie{Name: "token", Value: token})

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	var response dtos.UserResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Message != "user fetched successfully" {
		t.Errorf("Expected message to be 'user fetched successfully' but got '%s'", response.Message)
	}

	if response.Data.ID != 1 {
		t.Errorf("Expected user ID to be 1 but got %d", response.Data.ID)
	}
}

func TestGetUserByIdUnauthorized(t *testing.T) {
	router := routes.SetupV1Routes(db)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/v1/users/1", nil)

	req.Header.Set("Authorization", "Bearer invalid_token")
	router.ServeHTTP(w, req)

	if err != nil {
		t.Fatal(err)
	}

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d but got %d", http.StatusUnauthorized, w.Code)
	}
}

func TestGetUrlsByUserIdUnauthorized(t *testing.T) {
	router := routes.SetupV1Routes(db)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/v1/user/1/urls", nil)

	req.Header.Set("Authorization", "Bearer invalid_token")
	router.ServeHTTP(w, req)

	if err != nil {
		t.Fatal(err)
	}

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d but got %d", http.StatusUnauthorized, w.Code)
	}
}

func TestURLCreationResponse(t *testing.T) {
	now := time.Now()
	response := dtos.URLCreationResponse{
		Message:   "URL created successfully",
		ShortURL:  "http://short.url",
		CreatedAt: now,
	}

	if response.Message != "URL created successfully" {
		t.Errorf("Expected message to be 'URL created successfully' but got '%s'", response.Message)
	}

	if response.ShortURL != "http://short.url" {
		t.Errorf("Expected short URL to be 'http://short.url' but got '%s'", response.ShortURL)
	}
	if !response.CreatedAt.Equal(now) {
		t.Errorf("Expected CreatedAt to be '%v' but got '%v'", now, response.CreatedAt)
	}
}

package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	controller "github.com/Real-Dev-Squad/tiny-site-backend/controllers"
	"github.com/Real-Dev-Squad/tiny-site-backend/routes"
	"github.com/gin-gonic/gin"
)

func TestCreateTinyURL(t *testing.T) {
	router := setupTestRouter()
	requestBody := map[string]interface{}{
		"OrgUrl":    "https://example.com",
		"CreatedBy": "testuser",
	}
	requestJSON, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/v1/create-tinyurl", bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Error("Failed to parse JSON response:", err)
	}

	expectedMessage := "Tiny URL created successfully"
	if msg, ok := response["message"].(string); !ok || msg != expectedMessage {
		t.Errorf("Expected response message %q, got %q", expectedMessage, msg)
	}
}

func setupTestRouter() *gin.Engine {
	router := routes.SetupV1Routes(db)
	router.POST("/v1/create-tinyurl", func(ctx *gin.Context) {
		controller.CreateTinyURL(ctx, db)
	})
	return router
}

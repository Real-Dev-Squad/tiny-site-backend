package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	controller "github.com/Real-Dev-Squad/tiny-site-backend/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func (suite *AppTestSuite) TestCreateTinyURL() {
	router := gin.Default()
	// Change the route to the one defined in your routes file.
	router.POST("/v1/tinyurl", func(ctx *gin.Context) {
		controller.CreateTinyURL(ctx, suite.db) // Use the db from the suite.
	})

	// Simplified request body for debugging.
	requestBody := map[string]interface{}{
		"originalUrl": "https://example.com",
		"createdBy":   "testuser",
		"shortUrl":    "testShort", // Include all required fields
		"userId":      1,           // Use a valid user ID from your test data
		"expiredAt":   "2023-12-31T23:59:59Z",
	}
	requestJSON, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/v1/tinyurl", bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log response for debugging.
	suite.T().Logf("Response Code: %d, Body: %s", w.Code, w.Body.String())

	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200")
	// Add more assertions as needed.
}

func (suite *AppTestSuite) TestRedirectShortURL() {
	router := gin.Default()
	router.GET("/v1/tinyurl/:shortUrl", func(ctx *gin.Context) {
		controller.RedirectShortURL(ctx, suite.db) // Use the db from the suite.
	})

	// Assuming '37fff02c' is a valid short URL in your test database
	req, _ := http.NewRequest("GET", "/v1/tinyurl/37fff02c", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusMovedPermanently, w.Code, "Expected status code to be 301")
	// Add more assertions to check the response headers and other conditions as needed.
}

func (suite *AppTestSuite) TestGetAllURLs() {
	router := gin.Default()
	router.GET("/v1/user/:id/urls", func(ctx *gin.Context) {
		controller.GetAllURLs(ctx, suite.db) // Use the db from the suite.
	})

	// Assuming '1' is a valid user ID in your test database
	req, _ := http.NewRequest("GET", "/v1/user/1/urls", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200")
	// Add more assertions to check the response body and other conditions as needed.
}


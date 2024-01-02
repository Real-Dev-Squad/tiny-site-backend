package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	controller "github.com/Real-Dev-Squad/tiny-site-backend/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestCreateTinyURL_Success tests the creation of a tiny URL with valid data and expects a successful response.
func (suite *AppTestSuite) TestCreateTinyURL_Success() {
	// Setup the router and route for creating a tiny URL
	router := gin.Default()
	router.POST("/v1/tinyurl", func(ctx *gin.Context) {
		controller.CreateTinyURL(ctx, suite.db)
	})

	requestBody := map[string]interface{}{
		"OriginalUrl": "https://example.com",
		"UserId":      1, 
	}
	requestJSON, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/v1/tinyurl", bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	suite.T().Logf("Response Code: %d, Body: %s", w.Code, w.Body.String())
	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200 for successful tiny URL creation")
}

// TestCreateTinyURL_InvalidJSON tests the creation of a tiny URL with invalid JSON and expects a bad request response.
func (suite *AppTestSuite) TestCreateTinyURL_InvalidJSON() {
	// Setup the router and route for creating a tiny URL
	router := gin.Default()
	router.POST("/v1/tinyurl", func(ctx *gin.Context) {
		controller.CreateTinyURL(ctx, suite.db)
	})

	// Create a request with invalid JSON and a recorder to test the endpoint
	req, _ := http.NewRequest("POST", "/v1/tinyurl", bytes.NewBuffer([]byte("{invalid json")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code, "Expected status code to be 400 for invalid JSON")
}

// TestCreateTinyURL_EmptyOriginalURL tests the creation of a tiny URL with an empty original URL and expects a bad request response.
func (suite *AppTestSuite) TestCreateTinyURL_EmptyOriginalURL() {
	// Setup the router and route for creating a tiny URL
	router := gin.Default()
	router.POST("/v1/tinyurl", func(ctx *gin.Context) {
		controller.CreateTinyURL(ctx, suite.db)
	})

	requestBody := map[string]interface{}{
		"OriginalUrl": "",
		"UserId":      1,
	}
	requestJSON, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/v1/tinyurl", bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code, "Expected status code to be 400 for empty original URL")
}

// TestRedirectShortURL_Success tests the redirection of a short URL to the original URL and expects a moved permanently response.
func (suite *AppTestSuite) TestRedirectShortURL_Success() {
	router := gin.Default()
	router.GET("/v1/tinyurl/:shortURL", func(ctx *gin.Context) {
		controller.RedirectShortURL(ctx, suite.db)
	})

	req, _ := http.NewRequest("GET", "/v1/tinyurl/37fff02c", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusMovedPermanently, w.Code, "Expected status code to be 301 for successful redirection")
}

// TestRedirectShortURL_NotFound tests the redirection of a non-existent short URL and expects a not found response.
func (suite *AppTestSuite) TestRedirectShortURL_NotFound() {
	router := gin.Default()
	router.GET("/v1/tinyurl/:shortURL", func(ctx *gin.Context) {
		controller.RedirectShortURL(ctx, suite.db)
	})

	req, _ := http.NewRequest("GET", "/v1/tinyurl/nonexistent", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusNotFound, w.Code, "Expected status code to be 404 for non-existent short URL")
}

// TestGetAllURLs_Success tests the retrieval of all URLs for a user and expects a successful response.
func (suite *AppTestSuite) TestGetAllURLs_Success() {
	router := gin.Default()
	router.GET("/v1/user/:id/urls", func(ctx *gin.Context) {
		controller.GetAllURLs(ctx, suite.db)
	})

	req, _ := http.NewRequest("GET", "/v1/user/1/urls", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200 for successful retrieval of all URLs")
	responseBody := w.Body.String()
	suite.T().Logf("Response Body: %s", responseBody)
	fmt.Println("response", responseBody);
}

// TestGetURLDetails_Success tests the retrieval of details for a specific short URL and expects a successful response.
func (suite *AppTestSuite) TestGetURLDetails_Success() {
	router := gin.Default()
	router.GET("/v1/urls/:shortURL", func(ctx *gin.Context) {
		controller.GetURLDetails(ctx, suite.db)
	})

	req, _ := http.NewRequest("GET", `/v1/urls/37fff02c`, nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200 for successful retrieval of URL details")
}

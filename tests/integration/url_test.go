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

// TestCreateTinyURLSuccess tests the successful creation of a tiny URL with valid data.
func (suite *AppTestSuite) TestCreateTinyURLSuccess() {
	// Setup the router and route for creating a tiny URL
	router := gin.Default()
	router.POST("/v1/tinyurl", func(ctx *gin.Context) {
		controller.CreateTinyURL(ctx, suite.db)
	})

	// Prepare a request with valid data and a recorder to test the endpoint
	requestBody := map[string]interface{}{
		"OriginalUrl": "https://example.com",
		"UserId":      1,
	}
	requestJSON, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/v1/tinyurl", bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the status code is 200 for successful tiny URL creation
	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200 for successful tiny URL creation")
}

// TestCreateTinyURLInvalidJSON tests the creation of a tiny URL with invalid JSON and expects a bad request response.
func (suite *AppTestSuite) TestCreateTinyURLInvalidJSON() {
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

	// Assert the status code is 400 for invalid JSON
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code, "Expected status code to be 400 for invalid JSON")
}

// TestCreateTinyURLEmptyOriginalURL tests the creation of a tiny URL with an empty original URL and expects a bad request response.
func (suite *AppTestSuite) TestCreateTinyURLEmptyOriginalURL() {
	// Setup the router and route for creating a tiny URL
	router := gin.Default()
	router.POST("/v1/tinyurl", func(ctx *gin.Context) {
		controller.CreateTinyURL(ctx, suite.db)
	})

	// Prepare a request with an empty original URL and a recorder to test the endpoint
	requestBody := map[string]interface{}{
		"OriginalUrl": "",
		"UserId":      1,
	}
	requestJSON, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/v1/tinyurl", bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the status code is 400 for empty original URL
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code, "Expected status code to be 400 for empty original URL")
}

// TestRedirectShortURLSuccess tests the successful redirection of a short URL to the original URL.
func (suite *AppTestSuite) TestRedirectShortURLSuccess() {
	router := gin.Default()
	router.GET("/v1/tinyurl/:shortURL", func(ctx *gin.Context) {
		controller.RedirectShortURL(ctx, suite.db)
	})

	req, _ := http.NewRequest("GET", "/v1/tinyurl/37fff02c", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the status code is 301 for successful redirection
	assert.Equal(suite.T(), http.StatusMovedPermanently, w.Code, "Expected status code to be 301 for successful redirection")
}

// TestRedirectShortURLNotFound tests the redirection of a non-existent short URL and expects a not found response.
func (suite *AppTestSuite) TestRedirectShortURLNotFound() {
	router := gin.Default()
	router.GET("/v1/tinyurl/:shortURL", func(ctx *gin.Context) {
		controller.RedirectShortURL(ctx, suite.db)
	})

	req, _ := http.NewRequest("GET", "/v1/tinyurl/nonexistent", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the status code is 404 for non-existent short URL
	assert.Equal(suite.T(), http.StatusNotFound, w.Code, "Expected status code to be 404 for non-existent short URL")
}

// TestGetAllURLsSuccess tests the successful retrieval of all URLs for a user.
func (suite *AppTestSuite) TestGetAllURLsSuccess() {
	router := gin.Default()
	router.GET("/v1/user/:id/urls", func(ctx *gin.Context) {
		controller.GetAllURLs(ctx, suite.db)
	})

	req, _ := http.NewRequest("GET", "/v1/user/1/urls", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the status code is 200 for successful retrieval of all URLs
	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200 for successful retrieval of all URLs")
	responseBody := w.Body.String()
	suite.T().Logf("Response Body: %s", responseBody)
	fmt.Println("response", responseBody)
}

// TestGetURLDetailsSuccess tests the successful retrieval of details for a specific short URL.
func (suite *AppTestSuite) TestGetURLDetailsSuccess() {
	router := gin.Default()
	router.GET("/v1/urls/:shortURL", func(ctx *gin.Context) {
		controller.GetURLDetails(ctx, suite.db)
	})

	req, _ := http.NewRequest("GET", `/v1/urls/37fff02c`, nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the status code is 200 for successful retrieval of URL details
	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200 for successful retrieval of URL details")
}

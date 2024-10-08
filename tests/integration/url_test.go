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

// TestCreateTinyURLSuccess tests the successful creation of a tiny URL with valid data.
func (suite *AppTestSuite) TestCreateTinyURLSuccess() {
	// Setup the router and route for creating a tiny URL
	router := gin.Default()

	router.Use(func(ctx *gin.Context) {
		ctx.Set("userID", int64(1))
		ctx.Next()
	})

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

// TestCreateTinyURLCustomShortURL tests the creation of a tiny URL with a custom short URL and expects a successful response.
func (suite *AppTestSuite) TestCreateTinyURLCustomShortURL() {
	router := gin.Default()

	router.Use(func(ctx *gin.Context) {
		ctx.Set("userID", int64(1))
		ctx.Next()
	})

	router.POST("/v1/tinyurl", func(ctx *gin.Context) {
		controller.CreateTinyURL(ctx, suite.db)
	})

	requestBody := map[string]interface{}{
		"OriginalUrl": "https://example.com",
		"ShortUrl":    "short",
	}
	requestJSON, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/v1/tinyurl", bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200 for successful tiny URL creation")
}

func (suite *AppTestSuite) TestCreateTinyURLCustomShortURLExists() {
	router := gin.Default()

	router.Use(func(ctx *gin.Context) {
		ctx.Set("userID", int64(1))
		ctx.Next()
	})

	router.POST("/v1/tinyurl", func(ctx *gin.Context) {
		controller.CreateTinyURL(ctx, suite.db)
	})

	requestBody := map[string]interface{}{
		"OriginalUrl": "https://rds.com",
		"ShortUrl":    "37fff",
	}
	requestJSON, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/v1/tinyurl", bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code, "Expected status code to be 400 for existing short URL")
}

func (suite *AppTestSuite) TestCreateTinyURLExistingOriginalURL() {
    router := gin.Default()

	router.Use(func(ctx *gin.Context) {
		ctx.Set("userID", int64(1))
		ctx.Next()
	})

    router.POST("/v1/tinyurl", func(ctx *gin.Context) {
        controller.CreateTinyURL(ctx, suite.db)
    })

    existingOriginalURL := "https://www.example.com/1"

    requestBody := map[string]interface{}{
        "OriginalUrl": existingOriginalURL,
    }
    requestJSON, _ := json.Marshal(requestBody)
    req, _ := http.NewRequest("POST", "/v1/tinyurl", bytes.NewBuffer(requestJSON))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200 for existing original URL")

    var responseBody map[string]interface{}
    err := json.Unmarshal(w.Body.Bytes(), &responseBody)
    assert.NoError(suite.T(), err)

    if responseBody["urlCount"] != nil {
        responseBody["urlCount"] = int(responseBody["urlCount"].(float64))
    }

    expectedResponse := map[string]interface{}{
        "message":   "Shortened URL already exists",
        "shortUrl":  "37fff",
        "urlCount":  0,
        "createdAt": responseBody["createdAt"], 
    }

    assert.Equal(suite.T(), expectedResponse, responseBody, "Response body does not match expected JSON")
}

// TestRedirectShortURLSuccess tests the successful redirection of a short URL to the original URL.
func (suite *AppTestSuite) TestRedirectShortURLSuccess() {
	router := gin.Default()
	router.GET("/v1/redirect/:shortURL", func(ctx *gin.Context) {
		controller.RedirectShortURL(ctx, suite.db)
	})

	req, _ := http.NewRequest("GET", "/v1/redirect/37fff", nil)

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
	userEmail := "john.doe@example.com"

	router.GET("/v1/urls/self", func(ctx *gin.Context) {
		ctx.Set("user", userEmail)
		controller.GetAllURLs(ctx, suite.db)
	})

	req, _ := http.NewRequest("GET", "/v1/urls/self", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200 for successful retrieval of all URLs")
}

// TestGetURLDetailsSuccess tests the successful retrieval of details for a specific short URL.
func (suite *AppTestSuite) TestGetURLDetailsSuccess() {
	router := gin.Default()
	router.GET("/v1/urls/:shortURL", func(ctx *gin.Context) {
		controller.GetURLDetails(ctx, suite.db)
	})

	req, _ := http.NewRequest("GET", `/v1/urls/37fff`, nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the status code is 200 for successful retrieval of URL details
	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200 for successful retrieval of URL details")
}

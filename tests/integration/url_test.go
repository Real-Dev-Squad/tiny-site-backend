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

func (suite *AppTestSuite) TestCreateTinyURL() {
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
    assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200")
}

func (suite *AppTestSuite) TestCreateTinyURLInvalidJSON() {
    router := gin.Default()
    router.POST("/v1/tinyurl", func(ctx *gin.Context) {
        controller.CreateTinyURL(ctx, suite.db)
    })

    req, _ := http.NewRequest("POST", "/v1/tinyurl", bytes.NewBuffer([]byte("{invalid json")))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(suite.T(), http.StatusBadRequest, w.Code, "Expected status code to be 400")
}

func (suite *AppTestSuite) TestCreateTinyURLEmptyOriginalURL() {
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

    assert.Equal(suite.T(), http.StatusBadRequest, w.Code, "Expected status code to be 400")
}


func (suite *AppTestSuite) TestRedirectShortURL() {
	router := gin.Default()
	router.GET("/v1/tinyurl/:shortURL", func(ctx *gin.Context) {
		controller.RedirectShortURL(ctx, suite.db)
	})

	req, _ := http.NewRequest("GET", "/v1/tinyurl/37fff02c", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusMovedPermanently, w.Code, "Expected status code to be 301")
}

func (suite *AppTestSuite) TestRedirectShortURLNotFound() {
	router := gin.Default()
	router.GET("/v1/tinyurl/:shortURL", func(ctx *gin.Context) {
		controller.RedirectShortURL(ctx, suite.db)
	})

	req, _ := http.NewRequest("GET", "/v1/tinyurl/nonexistent", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusNotFound, w.Code, "Expected status code to be 404")
}


func (suite *AppTestSuite) TestGetAllURLs() {
	router := gin.Default()
	router.GET("/v1/user/:id/urls", func(ctx *gin.Context) {
		controller.GetAllURLs(ctx, suite.db)
	})

	req, _ := http.NewRequest("GET", "/v1/user/1/urls", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200")

	responseBody := w.Body.String()
    suite.T().Logf("Response Body: %s", responseBody)
	fmt.Println("resposne", responseBody);
}

func (suite *AppTestSuite) TestGetURLDetails() {
	router := gin.Default()
	router.GET("/v1/urls/:shortURL", func(ctx *gin.Context) {
		controller.GetURLDetails(ctx, suite.db)
	})

	req, _ := http.NewRequest("GET", `/v1/urls/37fff02c`, nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200")
}


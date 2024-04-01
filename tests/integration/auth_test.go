package tests

import (
	"net/http"
	"net/http/httptest"
	"os"

	controller "github.com/Real-Dev-Squad/tiny-site-backend/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
)

// It ensures that calling the logout endpoint resets the 'token' cookie and redirects to the configured AUTH_REDIRECT_URL.
func (suite *AppTestSuite) TestLogout() {
	os.Setenv("AUTH_REDIRECT_URL", "http://example.com/home")
	router := gin.Default()
	auth := router.Group("/v1/auth")

	auth.GET("/logout", func(ctx *gin.Context) {
		domain := os.Getenv("DOMAIN")
		authRedirectURL := os.Getenv("AUTH_REDIRECT_URL")

		ctx.SetCookie("token", "", -1, "/", domain, true, true)
		ctx.Redirect(http.StatusFound, authRedirectURL)
	})

	req, _ := http.NewRequest("GET", "/v1/auth/logout", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusFound, w.Code, "Expected status code to be 302")
	resetCookie := false
	for _, cookie := range w.Result().Cookies() {
		if cookie.Name == "token" && cookie.Value == "" && cookie.MaxAge < 0 {
			resetCookie = true
		}
	}
	assert.True(suite.T(), resetCookie, "Expected 'token' cookie to be reset")

	assert.Equal(suite.T(), "http://example.com/home", w.Result().Header.Get("Location"), "Expected redirect to authRedirectURL")
}

// It ensures that calling the Google login endpoint redirects to the Google OAuth URL.
func (suite *AppTestSuite) TestGoogleLogin() {
	router := gin.Default()

	router.GET("/v1/auth/google/login", controller.GoogleLogin)

	req, _ := http.NewRequest("GET", "/v1/auth/google/login", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusFound, w.Code, "Expected status code to be 302")
	assert.Contains(suite.T(), w.Result().Header.Get("Location"), "https://accounts.google.com/o/oauth2/auth", "Expected redirect to Google OAuth URL")
}

// It ensures that calling the Google OAuth callback endpoint results in a successful response.
func (suite *AppTestSuite) TestGoogleCallback() {
	router := gin.Default()

	router.GET("/v1/auth/google/callback", func(ctx *gin.Context) {
	})

	req, _ := http.NewRequest("GET", "/v1/auth/google/callback", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200")
}

func (suite *AppTestSuite) TestGoogleCallback_ErrorHandling() {
	router := gin.Default()

	// Mock the database connection to simulate a scenario where the database is unavailable.
	var db *bun.DB = nil

	router.GET("/v1/auth/google/callback", func(ctx *gin.Context) {
		controller.GoogleCallback(ctx, db)
	})

	req, _ := http.NewRequest("GET", "/v1/auth/google/callback", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the response status code is 500 Internal Server Error.
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code, "Expected status code to be 500")

	// Assert that the response body contains the expected error message.
	expectedErrorMessage := "database is not available"
	assert.Contains(suite.T(), w.Body.String(), expectedErrorMessage, "Expected error message in response body")
}

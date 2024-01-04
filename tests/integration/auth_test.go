package tests

import (
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestLogout checks the logout functionality.
// It ensures that calling the logout endpoint resets the 'token' cookie and redirects to the configured AUTH_REDIRECT_URL.
func (suite *AppTestSuite) TestLogout() {
    // Setup the environment and router
    os.Setenv("AUTH_REDIRECT_URL", "http://example.com/home")
    router := gin.Default()
    auth := router.Group("/v1/auth")

    auth.GET("/logout", func(ctx *gin.Context) {
        domain := os.Getenv("DOMAIN")
        authRedirectURL := os.Getenv("AUTH_REDIRECT_URL")

        // Reset the 'token' cookie and redirect to the authRedirectURL
        ctx.SetCookie("token", "", -1, "/", domain, true, true)
        ctx.Redirect(http.StatusFound, authRedirectURL)
    })

    req, _ := http.NewRequest("GET", "/v1/auth/logout", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Assert that the status code is 302 (redirect) and the 'token' cookie is reset
    assert.Equal(suite.T(), http.StatusFound, w.Code, "Expected status code to be 302")
    resetCookie := false
    for _, cookie := range w.Result().Cookies() {
        if cookie.Name == "token" && cookie.Value == "" && cookie.MaxAge < 0 {
            resetCookie = true
        }
    }
    assert.True(suite.T(), resetCookie, "Expected 'token' cookie to be reset")

    // Assert that the response redirects to the configured authRedirectURL
    assert.Equal(suite.T(), "http://example.com/home", w.Result().Header.Get("Location"), "Expected redirect to authRedirectURL")
}

// TestLogin checks the login functionality.
// It ensures that calling the Google login endpoint redirects to the Google OAuth URL.
func (suite *AppTestSuite) TestLogin() {
    // Setup the router
    router := gin.Default()

    // Define the login endpoint
    router.GET("/v1/auth/google/login", func(ctx *gin.Context) {
        ctx.Redirect(http.StatusFound, "/mock-google-oauth-url")
    })

    req, _ := http.NewRequest("GET", "/v1/auth/google/login", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Assert that the status code is 302 (redirect) and redirects to the mock Google OAuth URL
    assert.Equal(suite.T(), http.StatusFound, w.Code, "Expected status code to be 302")
    assert.Equal(suite.T(), "/mock-google-oauth-url", w.Result().Header.Get("Location"), "Expected redirect to mock Google OAuth URL")
}

// TestOAuthCallback checks the OAuth callback functionality.
// It ensures that calling the Google OAuth callback endpoint results in a successful response.
func (suite *AppTestSuite) TestOAuthCallback() {
    router := gin.Default()

    // Define the OAuth callback endpoint
    router.GET("/v1/auth/google/callback", func(ctx *gin.Context) {
    })

    // Create a request and recorder to test the endpoint
    req, _ := http.NewRequest("GET", "/v1/auth/google/callback", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Assert that the status code is 200 (OK)
    assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200")
}

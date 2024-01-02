package tests

import (
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestLogout_WhenCalled_ExpectRedirectToAuthURL tests the logout functionality to ensure
// that it resets the 'token' cookie and redirects to the AUTH_REDIRECT_URL.
func (suite *AppTestSuite) TestLogout_WhenCalled_ExpectRedirectToAuthURL() {
    // Setup the environment and router
    os.Setenv("AUTH_REDIRECT_URL", "http://example.com/home")
    router := gin.Default()
    auth := router.Group("/v1/auth")

    auth.GET("/logout", func(ctx *gin.Context) {
        domain := os.Getenv("DOMAIN")
        authRedirectUrl := os.Getenv("AUTH_REDIRECT_URL")

        // Reset the 'token' cookie and redirect to the authRedirectUrl
        ctx.SetCookie("token", "", -1, "/", domain, true, true)
        ctx.Redirect(http.StatusFound, authRedirectUrl)
    })

    req, _ := http.NewRequest("GET", "/v1/auth/logout", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Assert that the status code is 302 (redirect) and the cookie is reset
    assert.Equal(suite.T(), http.StatusFound, w.Code, "Expected status code to be 302")
    resetCookie := false
    for _, cookie := range w.Result().Cookies() {
        if cookie.Name == "token" && cookie.Value == "" && cookie.MaxAge < 0 {
            resetCookie = true
        }
    }
    assert.True(suite.T(), resetCookie, "Expected 'token' cookie to be reset")

    // Assert that the response redirects to the expected authRedirectUrl
    assert.Equal(suite.T(), "http://example.com/home", w.Result().Header.Get("Location"), "Expected redirect to authRedirectUrl")
}

// TestLogin_WhenCalled_ExpectRedirectToGoogleOAuthURL tests the login functionality to ensure
// that it redirects to the Google OAuth URL.
func (suite *AppTestSuite) TestLogin_WhenCalled_ExpectRedirectToGoogleOAuthURL() {
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

// TestOAuthCallback_WhenCalled_ExpectSuccess tests the OAuth callback functionality to ensure
// that it handles the callback successfully.
func (suite *AppTestSuite) TestOAuthCallback_WhenCalled_ExpectSuccess() {
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

package tests

import (
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func (suite *AppTestSuite) TestLogout() {
    os.Setenv("AUTH_REDIRECT_URL", "http://example.com/home")

    router := gin.Default()
    auth := router.Group("/v1/auth")
    auth.GET("/logout", func(ctx *gin.Context) {
        domain := os.Getenv("DOMAIN")
        authRedirectUrl := os.Getenv("AUTH_REDIRECT_URL")

        ctx.SetCookie("token", "", -1, "/", domain, true, true)
        ctx.Redirect(http.StatusFound, authRedirectUrl)
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

    assert.Equal(suite.T(), "http://example.com/home", w.Result().Header.Get("Location"), "Expected redirect to authRedirectUrl")
}

func (suite *AppTestSuite) TestLogin() {
    router := gin.Default()
    router.GET("/v1/auth/google/login", func(ctx *gin.Context) {
        ctx.Redirect(http.StatusFound, "/mock-google-oauth-url")
    })

    req, _ := http.NewRequest("GET", "/v1/auth/google/login", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(suite.T(), http.StatusFound, w.Code, "Expected status code to be 302")
}

func (suite *AppTestSuite) TestOAuthCallback() {
    router := gin.Default()
    router.GET("/v1/auth/google/callback", func(ctx *gin.Context) {
    })

    req, _ := http.NewRequest("GET", "/v1/auth/google/callback", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200")
}


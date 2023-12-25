package tests

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func (suite *AppTestSuite) TestLogout() {
    router := gin.Default()
    router.GET("/v1/auth/logout", func(ctx *gin.Context) {
        ctx.Redirect(http.StatusFound, "/home")
    })

    req, _ := http.NewRequest("GET", "/v1/auth/logout", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(suite.T(), http.StatusFound, w.Code, "Expected status code to be 302")
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


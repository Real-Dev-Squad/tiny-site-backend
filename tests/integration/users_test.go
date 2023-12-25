package tests

import (
	"net/http"
	"net/http/httptest"

	controller "github.com/Real-Dev-Squad/tiny-site-backend/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func (suite *AppTestSuite) TestGetUsers() {
	router := gin.Default()

	router.GET("/v1/users", func(ctx *gin.Context) {
		controller.GetUserList(ctx, suite.db) 
	})

	req, _ := http.NewRequest("GET", "/v1/users", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200")
}

func (suite *AppTestSuite) TestGetUserByID() {
    router := gin.Default()
    userID := "1" 
    
    router.GET("/v1/users/:id", func(ctx *gin.Context) {
        controller.GetUserByID(ctx, suite.db)
    })
    
    req, _ := http.NewRequest("GET", "/v1/users/"+userID, nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    
    assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200 for existing user")
    
    reqNotFound, _ := http.NewRequest("GET", "/v1/users/nonexistent", nil)
    wNotFound := httptest.NewRecorder()
    router.ServeHTTP(wNotFound, reqNotFound)
    
    assert.Equal(suite.T(), http.StatusNotFound, wNotFound.Code, "Expected status code to be 404 for non-existing user")
}

func (suite *AppTestSuite) TestGetSelfUser() {
    router := gin.Default()
    userEmail := "john.doe@example.com"

    router.GET("/v1/users/self", func(ctx *gin.Context) {
        ctx.Set("user", userEmail)
        controller.GetSelfUser(ctx, suite.db)
    })
    
	req, _ := http.NewRequest("GET", "/v1/users/self", nil)
    req.Header.Set("user", "john.doe@example.com")
    
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    
    assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200 for existing user")

}

func (suite *AppTestSuite) TestFailedGetSelfUser() {
    router := gin.Default()
    userEmail := "nonexisting@example.com" 

    router.GET("/v1/users/self", func(ctx *gin.Context) {
        ctx.Set("user", userEmail)
        controller.GetSelfUser(ctx, suite.db)
    })
    
    reqNotFound, _ := http.NewRequest("GET", "/v1/users/self", nil)
    reqNotFound.Header.Set("user", "nonexistent@example.com") 
    
    wNotFound := httptest.NewRecorder()
    router.ServeHTTP(wNotFound, reqNotFound)
    
    assert.Equal(suite.T(), http.StatusNotFound, wNotFound.Code, "Expected status code to be 404 for non-existing user")
}


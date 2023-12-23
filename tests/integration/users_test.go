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

	// Setup the route.
	router.GET("/v1/users", func(ctx *gin.Context) {
		controller.GetUserList(ctx, suite.db) // Use the db from the suite.
	})

	// Create a new HTTP request to the route.
	req, _ := http.NewRequest("GET", "/v1/users", nil)
	w := httptest.NewRecorder()

	// Serve the HTTP request.
	router.ServeHTTP(w, req)

	// Assert the results.
	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200")
	// Add more assertions as needed.
}

func (suite *AppTestSuite) TestGetUserByID() {
    router := gin.Default()
    userID := "1"  // Assume this user exists in your test database.
    
    router.GET("/v1/users/:id", func(ctx *gin.Context) {
        controller.GetUserByID(ctx, suite.db)
    })
    
    // Test for successful retrieval
    req, _ := http.NewRequest("GET", "/v1/users/"+userID, nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    
    assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200 for existing user")
    
    // You can add more assertions here to check the response body
    
    // Test for user not found
    reqNotFound, _ := http.NewRequest("GET", "/v1/users/nonexistent", nil)
    wNotFound := httptest.NewRecorder()
    router.ServeHTTP(wNotFound, reqNotFound)
    
    assert.Equal(suite.T(), http.StatusNotFound, wNotFound.Code, "Expected status code to be 404 for non-existing user")
}

func (suite *AppTestSuite) TestGetSelfUser() {
    router := gin.Default()
    userEmail := "john.doe@example.com" // Assume this user exists in your test database.

    router.GET("/v1/users/self", func(ctx *gin.Context) {
        ctx.Set("user", userEmail)  // Mocking user's email in the context
        controller.GetSelfUser(ctx, suite.db)
    })
    
	req, _ := http.NewRequest("GET", "/v1/users/self", nil)
    req.Header.Set("user", "john.doe@example.com")
    
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    
    assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200 for existing user")
    
    // // Test for user not found - the same setup but with a non-existent user.
    // reqNotFound, _ := http.NewRequest("GET", "/v1/users/self", nil)
    // reqNotFound.Header.Set("user", "nonexistent@example.com")  // Mocking non-existent user's email.
    
    // wNotFound := httptest.NewRecorder()
    // router.ServeHTTP(wNotFound, reqNotFound)
    
    // assert.Equal(suite.T(), http.StatusNotFound, wNotFound.Code, "Expected status code to be 404 for non-existing user")
}



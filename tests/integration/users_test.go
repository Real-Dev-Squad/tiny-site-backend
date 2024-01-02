package tests

import (
	"net/http"
	"net/http/httptest"

	controller "github.com/Real-Dev-Squad/tiny-site-backend/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestGetUsers_Success tests the retrieval of a list of users and expects a successful response.
func (suite *AppTestSuite) TestGetUsers_Success() {
	// Setup the router and route
	router := gin.Default()
	router.GET("/v1/users", func(ctx *gin.Context) {
		controller.GetUserList(ctx, suite.db) 
	})

	// Create a request and recorder to test the endpoint
	req, _ := http.NewRequest("GET", "/v1/users", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200 for successful user retrieval")
}

// TestGetUserByID_ExistingUser tests the retrieval of a user by ID for an existing user and expects a successful response.
func (suite *AppTestSuite) TestGetUserByID_ExistingUser() {
	router := gin.Default()
	userID := "1" 

	router.GET("/v1/users/:id", func(ctx *gin.Context) {
		controller.GetUserByID(ctx, suite.db)
	})
	
	req, _ := http.NewRequest("GET", "/v1/users/"+userID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200 for existing user")
}

// TestGetUserByID_NonExistent tests the retrieval of a user by ID for a non-existing user and expects a not found response.
func (suite *AppTestSuite) TestGetUserByID_NonExistent() {
	router := gin.Default()
	router.GET("/v1/users/:id", func(ctx *gin.Context) {
		controller.GetUserByID(ctx, suite.db)
	})
	
	req, _ := http.NewRequest("GET", "/v1/users/999", nil) 
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusNotFound, w.Code, "Expected status code to be 404 for non-existing user ID")
}

// TestGetSelfUser_ExistingUser tests the retrieval of the user's own profile with valid credentials and expects a successful response.
func (suite *AppTestSuite) TestGetSelfUser_ExistingUser() {
	router := gin.Default()
	userEmail := "john.doe@example.com"

	router.GET("/v1/users/self", func(ctx *gin.Context) {
		ctx.Set("user", userEmail)
		controller.GetSelfUser(ctx, suite.db)
	})
	
	req, _ := http.NewRequest("GET", "/v1/users/self", nil)
	req.Header.Set("user", userEmail)
    
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected status code to be 200 for existing user")
}

// TestGetSelfUser_NonExistingUser tests the retrieval of the user's own profile with invalid credentials and expects a not found response.
func (suite *AppTestSuite) TestGetSelfUser_NonExistingUser() {
	router := gin.Default()
	userEmail := "nonexisting@example.com"

	router.GET("/v1/users/self", func(ctx *gin.Context) {
		ctx.Set("user", userEmail)
		controller.GetSelfUser(ctx, suite.db)
	})
	
	req, _ := http.NewRequest("GET", "/v1/users/self", nil)
	req.Header.Set("user", userEmail)
    
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusNotFound, w.Code, "Expected status code to be 404 for non-existing user")
}

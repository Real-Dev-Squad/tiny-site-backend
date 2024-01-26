package tests

import (
	"net/http"
	"net/http/httptest"

	controller "github.com/Real-Dev-Squad/tiny-site-backend/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestGetUsersSuccess tests the successful retrieval of a list of users.
func (suite *AppTestSuite) TestGetUsersSuccess() {
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

// TestGetUserByIDExistingUser tests the retrieval of a user by ID for an existing user and expects a successful response.
func (suite *AppTestSuite) TestGetUserByIDExistingUser() {
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

// TestGetUserByIDNonExistent tests the retrieval of a user by ID for a non-existing user and expects a not found response.
func (suite *AppTestSuite) TestGetUserByIDNonExistent() {
	router := gin.Default()
	router.GET("/v1/users/:id", func(ctx *gin.Context) {
		controller.GetUserByID(ctx, suite.db)
	})

	req, _ := http.NewRequest("GET", "/v1/users/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusNotFound, w.Code, "Expected status code to be 404 for non-existing user ID")
}

// TestGetSelfUserExistingUser tests the retrieval of the user's own profile with valid credentials and expects a successful response.
func (suite *AppTestSuite) TestGetSelfUserExistingUser() {
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

// TestGetSelfUserNonExistingUser tests the retrieval of the user's own profile with invalid credentials and expects a not found response.
func (suite *AppTestSuite) TestGetSelfUserNonExistingUser() {
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

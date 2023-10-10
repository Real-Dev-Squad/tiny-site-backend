package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"tiny-site-backend/initializers"
	"tiny-site-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(c *gin.Context) {
	payload := c.MustGet("validatedPayload").(*models.SignUpInput)

	if payload.Password != payload.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Password and password confirmation do not match"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	newUser := models.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Username:  payload.Username,
		Email:     strings.ToLower(payload.Email),
		Password:  string(hashedPassword),
		Photo:     &payload.Photo,
	}

	result := initializers.DB.Create(&newUser)

	if result.Error != nil {
		handleSignupError(c, result.Error)
		return
	}

	c.JSON(201, gin.H{"status": "success", "data": gin.H{"user": models.FilterUserRecord(&newUser)}})
}

func SignInUser(c *gin.Context) {
	validatedPayload := c.MustGet("validatedPayload").(*models.SignInInput)
	payload := *validatedPayload

	user, err := getUserByUsername(payload.Username)
	if err != nil {
		handleError(c, err)
		return
	}
	fmt.Println("Hashed Password from DB:", user.Password)
	fmt.Println("Hashed Password from Login:", payload.Password)

	if err := validatePassword(user.Password, payload.Password); err != nil {
		fmt.Println("Password comparison failed:", err)
		handleError(c, err)
		return
	}

	tokenString, err := generateAuthToken(user)
	if err != nil {
		handleError(c, err)
		return
	}

	setAuthTokenCookie(c, tokenString)

	c.JSON(http.StatusOK, gin.H{"status": "success", "token": tokenString})
}

func LogoutUser(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func handleSignupError(c *gin.Context, err error) {
	errMsg := err.Error()
	if strings.Contains(errMsg, "duplicate key value violates unique constraint") {
		c.JSON(http.StatusConflict, gin.H{"status": "error", "message": "User with that email already exists"})
	} else if strings.Contains(errMsg, "record not found") {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Record not found"})
	} else {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": errMsg})
	}
}

func generateAuthToken(user models.User) (string, error) {
	config, _ := initializers.LoadConfig(".")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["exp"] = time.Now().Add(config.JwtExpiresIn).Unix()
	claims["iat"] = time.Now().Unix()
	claims["nbf"] = time.Now().Unix()

	tokenString, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return "", fmt.Errorf("generating JWT Token failed: %v", err)
	}

	return tokenString, nil
}

func setAuthTokenCookie(c *gin.Context, tokenString string) {
	config, _ := initializers.LoadConfig(".")
	Origin := os.Getenv("DOMAIN")
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", tokenString, int(config.JwtMaxAge), "/", Origin, false, true)
}

func handleError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
}

func getUserByUsername(username string) (models.User, error) {
	var user models.User
	result := initializers.DB.First(&user, "username = ?", username)
	if result.Error != nil {
		fmt.Println("Error querying user:", result.Error)
		return models.User{}, result.Error
	}
	return user, nil
}

func validatePassword(savedPassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(savedPassword), []byte(inputPassword))
}

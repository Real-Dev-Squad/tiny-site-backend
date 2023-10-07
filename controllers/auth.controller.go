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
		errMsg := result.Error.Error()
		if strings.Contains(errMsg, "duplicate key value violates unique constraint") {
			c.JSON(409, gin.H{"status": "error", "message": "User with that email already exists"})
		} else if strings.Contains(errMsg, "record not found") {
			c.JSON(404, gin.H{"status": "error", "message": "Record not found"})
		} else {
			c.JSON(502, gin.H{"status": "error", "message": errMsg})
		}
		return
	}

	c.JSON(201, gin.H{"status": "success", "data": gin.H{"user": models.FilterUserRecord(&newUser)}})
}

func SignInUser(c *gin.Context) {
	validatedPayload := c.MustGet("validatedPayload").(*models.SignInInput)
	payload := *validatedPayload

	var user models.User
	result := initializers.DB.First(&user, "username = ?", payload.Username)
	if result.Error != nil {
		c.JSON(400, gin.H{"status": "error", "message": "Invalid username or Password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		c.JSON(400, gin.H{"status": "error", "message": "Invalid username or Password"})
		return
	}

	config, _ := initializers.LoadConfig(".")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["exp"] = time.Now().Add(config.JwtExpiresIn).Unix()
	claims["iat"] = time.Now().Unix()
	claims["nbf"] = time.Now().Unix()

	tokenString, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		c.JSON(502, gin.H{"status": "error", "message": fmt.Sprintf("generating JWT Token failed: %v", err)})
		return
	}

	Origin := os.Getenv("DOMAIN")

	c.SetSameSite(http.SameSiteNoneMode)

	c.SetCookie("token", tokenString, int(config.JwtMaxAge), "/", Origin, false, true)

	c.JSON(200, gin.H{"status": "success", "token": tokenString})
}

func LogoutUser(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(200, gin.H{"status": "success"})
}

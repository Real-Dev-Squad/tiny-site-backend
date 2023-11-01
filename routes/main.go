package routes

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func SetupV1Routes(db *bun.DB) *gin.Engine {
	var router = gin.Default()

	v1 := router.Group("v1/")
	UserRoutes(v1, db)
	AuthRoutes(v1, db)
	TinyURLRoutes(v1, db)

	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")

	config := cors.DefaultConfig()
	config.AllowOrigins = allowedOrigins
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	return router
}

func Listen(listenAddress string, db *bun.DB) {
	router := SetupV1Routes(db)
	router.Run(listenAddress)
}

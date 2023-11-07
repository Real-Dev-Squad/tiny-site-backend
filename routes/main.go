package routes

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func SetupV1Routes(db *bun.DB) *gin.Engine {
	router := gin.Default()

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")

	corsConfig := cors.Config{
		AllowOrigins:     []string{allowedOrigins},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}

	router.Use(cors.New(corsConfig))

	v1 := router.Group("v1/")
	UserRoutes(v1, db)
	AuthRoutes(v1, db)
	TinyURLRoutes(v1, db)

	return router
}

func Listen(listenAddress string, db *bun.DB) {
	router := SetupV1Routes(db)
	router.Run(listenAddress)
}

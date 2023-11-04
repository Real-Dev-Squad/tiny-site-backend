package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Real-Dev-Squad/tiny-site-backend/models"
	"github.com/Real-Dev-Squad/tiny-site-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func CreateTinyURL(ctx *gin.Context, db *bun.DB) {
	var body models.Tinyurl

	err := ctx.BindJSON(&body)

	if err != nil {
		fmt.Println("JSON Error:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "JSON Error: " + err.Error(),
		})
		return
	}

	if body.OriginalUrl == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "OrgUrl is required",
		})
		return
	}

	body.ShortUrl = utils.GenerateMD5Hash(body.OriginalUrl)

	body.CreatedAt = time.Now()

	_, err = db.NewInsert().Model(&body).Exec(ctx)

	if err != nil {
		fmt.Println("Database Error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database Error: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Tiny URL created successfully",
		"short_url": body.ShortUrl,
	})
}

func RedirectShortURL(ctx *gin.Context, db *bun.DB) {
	shortURL := ctx.Param("shortURL")

	var tinyURL models.Tinyurl
	err := db.NewSelect().
		Model(&tinyURL).
		Where("short_url = ?", shortURL).
		Scan(ctx, &tinyURL) // Use Scan to bind data to the struct.
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Short URL not found",
		})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, tinyURL.OriginalUrl)
}

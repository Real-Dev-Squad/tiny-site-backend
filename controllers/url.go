package controller

import (
	"net/http"
	"time"

	"github.com/Real-Dev-Squad/tiny-site-backend/dtos"
	"github.com/Real-Dev-Squad/tiny-site-backend/models"
	"github.com/Real-Dev-Squad/tiny-site-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func CreateTinyURL(ctx *gin.Context, db *bun.DB) {
	var body models.Tinyurl

	if err := ctx.BindJSON(&body); err != nil {
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

	if err := db.NewSelect().Model(&body).Where("original_url = ?", body.OriginalUrl).Scan(ctx, &body); err == nil {
		ctx.JSON(http.StatusOK, dtos.URLCreationResponse{
			Message:  "Tiny URL already exists",
			ShortURL: body.ShortUrl,
		})
		return
	}

	body.ShortUrl = utils.GenerateMD5Hash(body.OriginalUrl)
	body.CreatedAt = time.Now()

	if _, err := db.NewInsert().Model(&body).Exec(ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database Error: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dtos.URLCreationResponse{
		Message:  "Tiny URL created successfully",
		ShortURL: body.ShortUrl,
	})
}

func RedirectShortURL(ctx *gin.Context, db *bun.DB) {
	shortURL := ctx.Param("shortURL")

	var tinyURL models.Tinyurl
	err := db.NewSelect().
		Model(&tinyURL).
		Where("short_url = ?", shortURL).
		Scan(ctx, &tinyURL)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Short URL not found",
		})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, tinyURL.OriginalUrl)
}

func GetAllURLs(ctx *gin.Context, db *bun.DB) {
	userID := ctx.Param("id")
	var tinyURL []models.Tinyurl

	err := db.NewSelect().
		Model(&tinyURL).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Scan(ctx, &tinyURL)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "No URLs found for the user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "All URLs fetched successfully",
		"urls":    tinyURL,
	})
}

func GetURLDetails(ctx *gin.Context, db *bun.DB) {
	shortURL := ctx.Param("shortURL")
	var tinyURL models.Tinyurl

	err := db.NewSelect().
		Model(&tinyURL).
		Where("short_url = ?", shortURL).
		Scan(ctx, &tinyURL)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "No URLs found for the user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "URL fetched successfully",
		"url":     tinyURL,
	})
}

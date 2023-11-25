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
		ctx.JSON(http.StatusBadRequest, dtos.URLCreationResponse{
			Message: "Invalid JSON format: " + err.Error(),
		})
		return
	}

	if body.OriginalUrl == "" {
		ctx.JSON(http.StatusBadRequest, dtos.URLCreationResponse{
			Message: "original url is required",
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
		ctx.JSON(http.StatusInternalServerError, dtos.URLCreationResponse{
			Message: "Failed to insert into database: " + err.Error(),
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
		ctx.JSON(http.StatusNotFound, dtos.URLDetailsResponse{
			Message: "Short URL not found",
		})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, tinyURL.OriginalUrl)
}

func GetAllURLs(ctx *gin.Context, db *bun.DB) {
	userID := ctx.Param("id")
	var tinyURLs []models.Tinyurl

	err := db.NewSelect().
		Model(&tinyURLs).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Scan(ctx, &tinyURLs)

	if err != nil {
		ctx.JSON(http.StatusNotFound, dtos.UserURLsResponse{
			Message: "No URLs found for the user",
		})
		return
	}
	var urlDetails []dtos.URLDetails
	for _, tinyURL := range tinyURLs {
		urlDetails = append(urlDetails, dtos.URLDetails{
			OriginalURL: tinyURL.OriginalUrl,
			ShortURL:    tinyURL.ShortUrl,
			CreatedAt:   tinyURL.CreatedAt,
		})
	}

	ctx.JSON(http.StatusOK, dtos.UserURLsResponse{
		Message: "All URLs fetched successfully",
		URLs:    urlDetails,
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
		ctx.JSON(http.StatusNotFound, dtos.URLDetailsResponse{
			Message: "No URLs found for the user",
		})
		return
	}
	urlDetails := dtos.URLDetails{
		ID:          tinyURL.ID,
		OriginalURL: tinyURL.OriginalUrl,
		ShortURL:    tinyURL.ShortUrl,
		Comment:     tinyURL.Comment,
		UserID:      tinyURL.UserID,
		CreatedBy:   tinyURL.CreatedBy,
		ExpiredAt:   tinyURL.ExpiredAt,
		CreatedAt:   tinyURL.CreatedAt,
	}

	ctx.JSON(http.StatusOK, dtos.URLDetailsResponse{
		Message: "URL fetched successfully",
		URL:     urlDetails,
	})
}

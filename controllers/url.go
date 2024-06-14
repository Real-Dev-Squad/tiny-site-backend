package controller

import (
	"net/http"
	"strings"
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
			Message: "Invalid Request.",
		})
		return
	}

	if body.OriginalUrl == "" {
		ctx.JSON(http.StatusBadRequest, dtos.URLCreationResponse{
			Message: "URL is required",
		})
		return
	}

	var existingOriginalURL models.Tinyurl
	if err := db.NewSelect().Model(&existingOriginalURL).Where("original_url = ?", body.OriginalUrl).Limit(1).Scan(ctx); err == nil {
		ctx.JSON(http.StatusOK, dtos.URLCreationResponse{
			Message:  "Shortened URL already exists",
			ShortURL: existingOriginalURL.ShortUrl,
		})
		return
	}

	if body.ShortUrl != "" {
		if len(body.ShortUrl) < 5 {
			ctx.JSON(http.StatusBadRequest, dtos.URLCreationResponse{
				Message: "Custom short URL must be at least 5 characters long",
			})
			return
		}

		var existingURL models.Tinyurl
		if err := db.NewSelect().Model(&existingURL).Where("short_url = ?", body.ShortUrl).Limit(1).Scan(ctx); err == nil {
			ctx.JSON(http.StatusBadRequest, dtos.URLCreationResponse{
				Message: "Custom short URL already exists",
			})
			return
		}
	} else {
		generatedShortURL := utils.GenerateMD5Hash(body.OriginalUrl)
		var existingURL models.Tinyurl
		if err := db.NewSelect().Model(&existingURL).Where("short_url = ?", generatedShortURL).Limit(1).Scan(ctx); err != nil {
			body.ShortUrl = generatedShortURL
		}
	}
	count, _ := db.NewSelect().Model(models.Tinyurl{}).Where("user_id = ?", body.UserID).Count(ctx)
	body.CreatedAt = time.Now().UTC()
	if count >= 50 {
		ctx.JSON(http.StatusInternalServerError, dtos.URLCreationResponse{
			Message: "Url Limit Reached, Please Delete to Create New !",
		})
		return
	}
	if _, err := db.NewInsert().Model(&body).Exec(ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.URLCreationResponse{
			Message: "OOPS!!, Unable to process your request at this moment, Please try after sometime. ",
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

	if !strings.HasPrefix(tinyURL.OriginalUrl, "http://") && !strings.HasPrefix(tinyURL.OriginalUrl, "https://") {
		tinyURL.OriginalUrl = "http://" + tinyURL.OriginalUrl
	}

	tinyURL.AccessCount++
	tinyURL.LastAccessedAt = time.Now().UTC()

	_, err = db.NewUpdate().
		Model(&tinyURL).
		Column("access_count", "last_accessed_at").
		WherePK().
		Exec(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update access count and timestamp",
		})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, tinyURL.OriginalUrl)
}

func GetAllURLs(ctx *gin.Context, db *bun.DB) {
	userEmail, _ := ctx.Get("user")
	var user models.User
	var tinyURLs []models.Tinyurl

	userModelError := db.NewSelect().
		Model(&user).
		Where("email = ?", userEmail).
		Scan(ctx, &user)

	if userModelError != nil {
		ctx.JSON(http.StatusNotFound, dtos.UserURLsResponse{
			Message: "User not found",
		})
		return
	}
	userID := user.ID

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
func DeleteURL(ctx *gin.Context, db *bun.DB) {
	id, _ := ctx.Params.Get("id")
	_, err := db.NewUpdate().Model(&models.Tinyurl{}).Set("is_deleted=?", true).Set("deleted_at=?", time.Now().UTC()).Where("id = ?", id).Exec(ctx)

	if err != nil {
		ctx.JSON(http.StatusNotFound, dtos.UserURLsResponse{
			Message: "No URLs found",
		})
		return
	}

	ctx.JSON(http.StatusOK, dtos.UserURLsResponse{
		Message: "Url Set to deleted",
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
		ctx.JSON(http.StatusNotFound, dtos.URLNotFoundResponse{
			Message: "URL not found",
		})
		return
	}
	urlDetails := dtos.URLDetails{
		ID:             tinyURL.ID,
		OriginalURL:    tinyURL.OriginalUrl,
		ShortURL:       tinyURL.ShortUrl,
		Comment:        tinyURL.Comment,
		UserID:         tinyURL.UserID,
		CreatedBy:      tinyURL.CreatedBy,
		ExpiredAt:      tinyURL.ExpiredAt,
		CreatedAt:      tinyURL.CreatedAt,
		AccessCount:    tinyURL.AccessCount,
		LastAccessedAt: tinyURL.LastAccessedAt,
	}

	ctx.JSON(http.StatusOK, dtos.URLDetailsResponse{
		Message: "URL fetched successfully",
		URL:     urlDetails,
	})
}

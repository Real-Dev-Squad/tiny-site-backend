package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/Real-Dev-Squad/tiny-site-backend/models"
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

	if body.OrgUrl == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "OrgUrl is required",
		})
		return
	}

	body.ShortUrl = generateMD5Hash(body.OrgUrl)

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
		"message": "Tiny URL created successfully",
	})
}

func generateMD5Hash(url string) string {
	url = url + time.Nanosecond.String()
	hash := md5.New()
	hash.Write([]byte(url))
	return hex.EncodeToString(hash.Sum(nil))[:8]
}

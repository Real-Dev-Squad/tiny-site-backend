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

func GetTinyURLs(ctx *gin.Context, db *bun.DB) {
	var tinyURLs []models.Tinyurl
	err := db.NewSelect().Model(&tinyURLs).OrderExpr("id ASC").Limit(10).Scan(ctx)

	if err != nil {
		fmt.Println("Error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": tinyURLs,
	})
}

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

	body.ShortUrl = generateMD5Hash(body.OrgUrl)
	body.CreatedAt = time.Now()
	// body.CreatedBy = body.CreatedBy
	// body.ExpiredAt = time.Now().AddDate(0, 0, 7)

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
		"data":    "",
	})
}

func generateMD5Hash(url string) string {
	url = url + time.Nanosecond.String()
	hash := md5.New()
	hash.Write([]byte(url))
	return hex.EncodeToString(hash.Sum(nil))[:8]
}

// func generateShortURL() string {
// 	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
// 	shortURL := make([]byte, 6)
// 	for i := range shortURL {
// 		shortURL[i] = characters[rand.Intn(len(characters))]
// 	}
// 	return string(shortURL)
// }

// func decryptShortURL(shortURL string) string {
// 	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
// 	decryptedURL := ""
// 	for _, char := range shortURL {
// 		decryptedURL += string(characters[len(characters)-strings.IndexRune(characters, rune(char))-1])
// 	}
// 	return decryptedURL
// }

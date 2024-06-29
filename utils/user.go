package utils

import (
	"github.com/Real-Dev-Squad/tiny-site-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func IncrementURLCount(userID int64, db *bun.DB, ctx *gin.Context) error {
	_, err := db.NewUpdate().
		Model((*models.User)(nil)).
		Set("url_count = url_count + 1").
		Where("id = ?", userID).
		Exec(ctx)
	return err
}

func DecrementURLCount(userID int64, db *bun.DB, ctx *gin.Context) error {
	_, err := db.NewUpdate().
		Model((*models.User)(nil)).
		Set("url_count = url_count - 1").
		Where("id = ?", userID).
		Exec(ctx)
	return err
}

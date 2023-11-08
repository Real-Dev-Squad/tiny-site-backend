package controller

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/uptrace/bun"
    "github.com/Real-Dev-Squad/tiny-site-backend/models"
)

func GetUserList(ctx *gin.Context, db *bun.DB) {
    var users []models.User
    err := db.NewSelect().Model(&users).OrderExpr("id ASC").Limit(10).Scan(ctx)

    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "message": "error",
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "Users fetched successfully",
        "data":    users,
    })
}

func GetUserByID(ctx *gin.Context, db *bun.DB) {
    id := ctx.Param("id")

    var user models.User
    err := db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx)

    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "message": "error",
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "User fetched successfully",
        "data":    user,
    })
}

func GetSelfUser(ctx *gin.Context, db *bun.DB) {
    userEmail, _ := ctx.Get("user")

    var user models.User
    err := db.NewSelect().Model(&user).Where("email = ?", userEmail).Scan(ctx)

    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "message": "error",
        })
        return
    }

    ctx.JSON(http.StatusOK, user)
}

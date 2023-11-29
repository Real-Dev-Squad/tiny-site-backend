package controller

import (
	"net/http"

	"github.com/Real-Dev-Squad/tiny-site-backend/dtos"
	"github.com/Real-Dev-Squad/tiny-site-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func GetUserList(ctx *gin.Context, db *bun.DB) {
	var users []models.User
	err := db.NewSelect().Model(&users).OrderExpr("id ASC").Limit(10).Scan(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.UserListResponse{
			Message: "Failed to fetch users: " + err.Error(),
		})
		return
	}

	if len(users) == 0 {
		ctx.JSON(http.StatusNotFound, dtos.UserListResponse{
			Message: "No users found",
		})
		return
	}

	var dtoUsers []dtos.User
	for _, user := range users {
		dtoUsers = append(dtoUsers, dtos.User{
			ID:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		})
	}

	ctx.JSON(http.StatusOK, dtos.UserListResponse{
		Message: "users fetched successfully",
		Data:    dtoUsers,
	})
}

func GetUserByID(ctx *gin.Context, db *bun.DB) {
	id := ctx.Param("id")

	var user models.User
	err := db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx)

	if err != nil {
		ctx.JSON(http.StatusNotFound, dtos.UserResponse{
			Message: "User not found: " + err.Error(),
		})
		return
	}

	dtoUser := dtos.User{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, dtos.UserResponse{
		Message: "user fetched successfully",
		Data:    dtoUser,
	})
}

func GetSelfUser(ctx *gin.Context, db *bun.DB) {
	userEmail, _ := ctx.Get("user")

	var user models.User
	err := db.NewSelect().Model(&user).Where("email = ?", userEmail).Scan(ctx)

	if err != nil {
		ctx.JSON(http.StatusNotFound, dtos.UserResponse{
			Message: "User not found: " + err.Error(),
		})
		return
	}

	dtoUser := dtos.User{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, dtos.UserResponse{
		Message: "user fetched successfully",
		Data:    dtoUser,
	})
}

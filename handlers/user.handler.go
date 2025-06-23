package handlers

import (
	"net/http"

	"github.com/Kaa-dan/skill-management/database"
	"github.com/Kaa-dan/skill-management/managers"
	"github.com/Kaa-dan/skill-management/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	groupName   string
	userManager *managers.UserManager
}

func NewUserHandleFrom(userManager *managers.UserManager) *UserHandler {
	return &UserHandler{"api/users", userManager}
}

func (userHandler *UserHandler) RegisterUserApis(r *gin.Engine) {
	userGroup := r.Group(userHandler.groupName)
	userGroup.GET("", userHandler.Create)

}

func (userHandler *UserHandler) Create(ctx *gin.Context) {
	//creating user
	database.DB.Create(&models.UserModel{FullName: "nithin raj", Email: "nithraj.dev"})

	ctx.JSON(http.StatusOK, gin.H{"msg": "success"})
}

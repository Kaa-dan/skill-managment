package handlers

import (
	"fmt"
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
	userGroup.POST("", userHandler.Create)

}

func (userHandler *UserHandler) Create(ctx *gin.Context) {

	var user_data struct {
		Full_name string `json:"full_name"`
		Email     string `json:"email"`
	}

	err := ctx.BindJSON(&user_data)

	//error
	if err != nil {
		fmt.Println("failed", err)
	}

	//creating user

	database.DB.Create(&models.UserModel{FullName: user_data.Full_name, Email: user_data.Email})

	ctx.JSON(http.StatusOK, gin.H{"msg": "success"})
}

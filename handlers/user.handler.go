package handlers

import (
	"fmt"
	"net/http"

	"github.com/Kaa-dan/skill-management/common"
	"github.com/Kaa-dan/skill-management/managers"
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

	user_data := common.NewUserCreationInput()

	err := ctx.BindJSON(&user_data)

	//error
	if err != nil {
		fmt.Println("failed", err)
	}

	//creating user
	newUser, err := userHandler.userManager.Create(user_data)

	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, newUser)
}

func (userHandler *UserHandler) List(ctx *gin.Context) {

	//creating user
	allUsers, err := userHandler.userManager.List()

	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, allUsers)
}

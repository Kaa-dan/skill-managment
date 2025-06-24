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
	userGroup.GET("", userHandler.List)
	userGroup.GET(":user_id/", userHandler.Detail)
	userGroup.DELETE(":user_id/", userHandler.Delete)

}

func (userHandler *UserHandler) Create(ctx *gin.Context) {

	user_data := common.NewUserCreationInput()

	err := ctx.BindJSON(&user_data)

	//error
	if err != nil {
		common.BadResponse(ctx, "failed to bind data ")
		return
	}

	//creating user
	newUser, err := userHandler.userManager.Create(user_data)

	if err != nil {
		common.BadResponse(ctx, "failed to create user")
		return
	}
	ctx.JSON(http.StatusOK, newUser)
}

func (userHandler *UserHandler) List(ctx *gin.Context) {

	allUsers, err := userHandler.userManager.List()

	if err != nil {
		common.BadResponse(ctx, "failed to get user")
		return
	}
	ctx.JSON(http.StatusOK, allUsers)
}

func (userHandler *UserHandler) Detail(ctx *gin.Context) {
	userId, ok := ctx.Params.Get("user_id")

	if !ok {
		fmt.Println("invalid user id ")
	}

	user, err := userHandler.userManager.Detail(userId)

	if user.ID == 0 {
		common.BadResponse(ctx, "no user")
		return
	}

	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, user)
}

func (userHandler *UserHandler) Delete(ctx *gin.Context) {
	userId, ok := ctx.Params.Get("user_id")
	if !ok {
		fmt.Println("invalid user id ")
	}
	err := userHandler.userManager.Delete(userId)
	if err != nil {
		common.BadResponse(ctx, "failed to delete user")
		return
	}
	common.SuccessResponse(ctx, "Deleted user")
}

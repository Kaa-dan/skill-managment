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
	userManager managers.UserManager
}

func NewUserHandleFrom(userManager managers.UserManager) *UserHandler {
	return &UserHandler{"api/users", userManager}
}

func (handlers *UserHandler) RegisterUserApis(r *gin.Engine) {
	userGroup := r.Group(handlers.groupName)
	userGroup.POST("", handlers.Create)
	userGroup.GET("", handlers.List)
	userGroup.GET(":user_id/", handlers.Detail)
	userGroup.DELETE(":user_id/", handlers.Delete)
	userGroup.PATCH(":user_id/", handlers.Update)

}

func (handlers *UserHandler) Create(ctx *gin.Context) {

	user_data := common.NewUserCreationInput()

	err := ctx.BindJSON(&user_data)

	//error
	if err != nil {
		common.BadResponse(ctx, "failed to bind data ")
		return
	}

	//creating user
	newUser, err := handlers.userManager.Create(user_data)

	if err != nil {
		common.BadResponse(ctx, "failed to create user")
		return
	}
	ctx.JSON(http.StatusOK, newUser)
}

func (handlers *UserHandler) List(ctx *gin.Context) {

	allUsers, err := handlers.userManager.List()

	if err != nil {
		common.BadResponse(ctx, "failed to get user")
		return
	}
	ctx.JSON(http.StatusOK, allUsers)
}

func (handlers *UserHandler) Detail(ctx *gin.Context) {
	userId, ok := ctx.Params.Get("user_id")

	if !ok {
		fmt.Println("invalid user id ")
	}

	user, err := handlers.userManager.Detail(userId)

	if user.ID == 0 {
		common.BadResponse(ctx, "no user")
		return
	}

	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, user)
}

func (handlers *UserHandler) Delete(ctx *gin.Context) {
	userId, ok := ctx.Params.Get("user_id")
	if !ok {
		fmt.Println("invalid user id ")
	}
	err := handlers.userManager.Delete(userId)
	if err != nil {
		common.BadResponse(ctx, "failed to delete user")
		return
	}
	common.SuccessResponse(ctx, "Deleted user")
}

func (handlers *UserHandler) Update(ctx *gin.Context) {
	userId, ok := ctx.Params.Get("user_id")
	if !ok {
		fmt.Println("invalid user id ")
	}

	user_data := common.NewUserUpdateInput()

	err := ctx.BindJSON(&user_data)

	//error
	if err != nil {
		common.BadResponse(ctx, "failed to bind data ")
		return
	}

	//creating user
	updatedUser, err := handlers.userManager.Update(user_data, userId)

	if err != nil {
		common.BadResponse(ctx, "failed to update user")
		return
	}
	ctx.JSON(http.StatusOK, updatedUser)
}

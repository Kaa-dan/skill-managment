package handlers

import (
	"github.com/Kaa-dan/skill-management/managers"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	groupName   string
	userManager *managers.UserManager
}

func NewUserHandleFrom(userManager *managers.UserManager) *UserHandler {
	return &UserHandler{"user", userManager}
}

func (userHandler *UserHandler) RegisterUserApis(r *gin.Engine) {
	userGroup := r.Group(userHandler.groupName)
}

func (userHandler *UserHandler) Create() {

}

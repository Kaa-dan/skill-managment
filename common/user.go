package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// create user
type UserCreationInput struct {
	Fullname string `json:"full_name"`
	Email    string `json:"email"`
}

func NewUserCreationInput() *UserCreationInput {
	return &UserCreationInput{}
}

// update user
type UserUpdateInput struct {
	Fullname string `json:"full_name"`
	Email    string `json:"email"`
}

func NewUserUpdateInput() *UserCreationInput {
	return &UserCreationInput{}
}

// response
type requestResponse struct {
	Message string `json:"message"`
	Status  uint   `json:"status"`
}

func SuccessResponse(ctx *gin.Context, msg string) {
	response := requestResponse{
		Message: msg,
		Status:  http.StatusOK,
	}
	ctx.JSON(http.StatusOK, response)

}

func BadResponse(ctx *gin.Context, msg string) {
	response := requestResponse{
		Message: msg,
		Status:  http.StatusBadRequest,
	}
	ctx.JSON(http.StatusBadRequest, response)

}

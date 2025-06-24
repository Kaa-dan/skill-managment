package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserCreationInput struct {
	Fullname string `json:"full_name"`
	Email    string `json:"email"`
}

func NewUserCreationInput() *UserCreationInput {
	return &UserCreationInput{}
}

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

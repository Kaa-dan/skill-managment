package main

import (
	"github.com/Kaa-dan/skill-management/handlers"
	"github.com/Kaa-dan/skill-management/managers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	userManager := managers.NewUserManager()
	userHandler := handlers.NewUserHandleFrom(userManager)

	userHandler.RegisterUserApis(router)

	router.Run() // listen and server on 0.0.0.0:8080
}

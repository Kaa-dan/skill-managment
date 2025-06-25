package main

import (
	"github.com/Kaa-dan/skill-management/database"
	"github.com/Kaa-dan/skill-management/handlers"
	"github.com/Kaa-dan/skill-management/managers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	//db
	database.Initialize()
}

func main() {

	//router setup
	router := gin.Default()

	router.Use(cors.Default())

	userManager := managers.NewUserManager()

	userHandler := handlers.NewUserHandler(userManager)

	userHandler.RegisterUserApis(router)

	router.Run() // listen and server on 0.0.0.0:8080
}

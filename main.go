package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "welcome"})
	})

	router.Run() // listen and server on 0.0.0.0:8080
}

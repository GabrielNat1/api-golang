package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "primeiros testes",
		})
	})

	server.Run(":8000")

}

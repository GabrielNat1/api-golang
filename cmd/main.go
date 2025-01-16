package main

import (
	"api-go/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	ProductController := controller.NewProductController()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "primeiros testes",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")

}

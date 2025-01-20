package main

import (
	"api-go/controller"
	"api-go/db"
	"api-go/middleware"
	"api-go/repository"
	"api-go/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDb()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)

	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	ProductController := controller.NewProductController(*ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "primeiros testes",
		})
	})

	protected := server.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/products", ProductController.GetProducts)
		protected.POST("/product", ProductController.CreateProduct)
		protected.GET("/product/:productId", ProductController.GetProductsById)
		protected.PUT("/product", ProductController.UpdateProduct)
		protected.DELETE("/product/:productId", ProductController.DeleteProduct)
	}

	server.Run(":8000")

}

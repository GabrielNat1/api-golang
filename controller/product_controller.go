package controller

import (
	"api-go/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct{}

func NewProductController() *productController {
	return &productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{ID: 1, Name: "Produto 1", Price: 100.00},
	}
	ctx.JSON(http.StatusOK, products)
}

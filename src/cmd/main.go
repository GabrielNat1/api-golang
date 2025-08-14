package main

import (
	"api-go/src/config"
	"api-go/src/controller"
	"api-go/src/database"
	"api-go/src/middleware"
	"api-go/src/pkg/logger"
	"api-go/src/repository"
	"api-go/src/usecase"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// Initialize config
	cfg := config.LoadConfig()

	// Initialize logger
	logger.Init("development")

	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}
	defer db.Close()

	// Initialize repositories and services
	productRepository := repository.NewProductRepository(db)
	productUseCase := usecase.NewProductUseCase(productRepository)
	productController := controller.NewProductController(*productUseCase)

	// Initialize Gin
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	// Health check
	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})

	// Protected routes
	protected := server.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/products", productController.GetProducts)
		protected.POST("/product", productController.CreateProduct)
		protected.GET("/product/:productId", productController.GetProductsById)
		protected.PUT("/product", productController.UpdateProduct)
		protected.DELETE("/product/:productId", productController.DeleteProduct)
	}

	// Server configuration
	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: server,
	}

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Shutdown server
	logger.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	logger.Info("Server exiting")
}

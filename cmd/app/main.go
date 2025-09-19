package main

import (
	"Webserver/internal/config"
	"Webserver/internal/database"
	"Webserver/internal/handlers"
	"Webserver/internal/repositories"
	"Webserver/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// 1. Ввести грамотную работу с переменными
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	var sellerRepo *repositories.SellerRepository

	// 2. Подключить  постргре 16
	db, err := database.Connect(cfg)
	if err != nil {
		panic(err)
	}

	// 3. Подключить сущности к орм - gorm

	sellerRepo = repositories.NewSellerRepository(db)
	sellerSrv := services.NewSellerService(sellerRepo)
	sellerHandler := handlers.NewSellerHandler(sellerSrv)

	r := gin.Default()

	// Добавляем ping endpoint
	r.GET("/ping", ping)

	// sellers routes
	sellerRoutes := r.Group("/sellers")
	{
		// CRUD
		sellerRoutes.POST("", sellerHandler.Create)
		sellerRoutes.GET("", sellerHandler.GetAll)      // Все продавцы
		sellerRoutes.GET("/:id", sellerHandler.GetByID) // Конкретный продавец
		sellerRoutes.PATCH("/:id", sellerHandler.Update)
		sellerRoutes.DELETE("/:id", sellerHandler.Delete)
	}

	// cars routes - нужно создать соответствующие handlers!
	carRoutes := r.Group("/cars")
	{
		// CRUD
		// Временные заглушки - нужно реализовать handlers для cars
		carRoutes.POST("", func(c *gin.Context) {
			c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
		})
		carRoutes.GET("", func(c *gin.Context) {
			c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
		})
		carRoutes.PATCH("/:id", func(c *gin.Context) {
			c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
		})
		carRoutes.DELETE("/:id", func(c *gin.Context) {
			c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
		})
	}

	r.Run(":" + cfg.HTTPort)
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

package main

import (
	"Webserver/internal/handlers"
	"Webserver/internal/repositories"
	"Webserver/internal/services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	sellerV := os.Getenv("SELLER_PROVIDER_VERSION")
	var sellerRepo *repositories.SellerRepository
	if sellerV == "1" {
		sellerRepo = repositories.NewSellerRepository()
	} else {
		sellerRepo = repositories.NewSellerRepository2()
	}

	sellerSrv := services.NewSellerService(sellerRepo)
	sellerHandler := handlers.NewSellerHandler(sellerSrv)

	r := gin.Default()

	{
		r.GET("/ping", ping)
	}

	// models/cars
	{
		// CRUD
		r.POST("/sellers", sellerHandler.Create)
		r.GET("/sellers", sellerHandler.GetAll)
		r.PATCH("/sellers/:id", sellerHandler.Update)
		r.DELETE("/sellers/:id/", sellerHandler.Delete)
		// CRUD

	}

	// models/cars
	{
		// CRUD

		r.POST("/cars", handlers.Create)
		r.GET("/cars", handlers.GetAll)
		r.PATCH("/cars/:id", handlers.Update)
		r.DELETE("/cars/:id/", handlers.Delete)
		// CRUD

	}

	r.Run()
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

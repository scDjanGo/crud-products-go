package main

import (
	appcors "products/cors"
	"products/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(appcors.CORS_SETTINGS()))

	v1 := router.Group("/api/v1")
	{
		v1.GET("/products", handlers.GetProducts)
		v1.POST("/products", handlers.PostProduct)
		v1.PATCH("/products/:id", handlers.PatchProduct)
		v1.DELETE("/products/:id", handlers.DeleteProduct)
	}

	router.Run(":8080")
}

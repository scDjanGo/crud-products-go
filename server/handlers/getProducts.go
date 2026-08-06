package handlers

import (
	"net/http"

	"products/db"
	"products/dto"

	"github.com/gin-gonic/gin"
)


func GetProducts(c *gin.Context) {


	 data := []dto.ListProduct{}


	for _, value := range db.PRODUCTS_DB {
		var product dto.ListProduct

		product.Id = value.Id
		product.Title = value.Title
		product.Description = value.Description
		
		data = append(data, product)
	}


	c.JSON(http.StatusOK,data)
}



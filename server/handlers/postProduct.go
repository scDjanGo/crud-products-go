package handlers

import (
	"net/http"
	"products/db"
	"products/dto"

	"github.com/gin-gonic/gin"
)

func PostProduct(c *gin.Context) {

	var newProduct dto.Product

	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	db.PRODUCTS_DB[db.LARGEST_ID] = newProduct
	db.LARGEST_ID++
	c.JSON(http.StatusCreated, newProduct)

}

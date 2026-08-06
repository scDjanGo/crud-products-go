package handlers

import (
	"net/http"
	"strconv"

	"products/db"
	"products/dto"

	"github.com/gin-gonic/gin"
)

func PatchProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "INvalid id"})
		return
	}


	var input dto.PatchProduct

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentProduct, exits := db.PRODUCTS_DB[id]

	if !exits {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}


	currentProduct.Title = *input.Title
	currentProduct.Description = *input.Description
	currentProduct.Price = *input.Price
	currentProduct.Quantity_In_Stock = *input.Quantity_In_Stock
	currentProduct.Discount = *input.Discount

	db.PRODUCTS_DB[id] = currentProduct


	c.JSON(http.StatusOK, currentProduct)



}

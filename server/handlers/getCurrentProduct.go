package handlers

import (
	"net/http"
	"products/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCurrentProduct(c *gin.Context) {

	id, ok := strconv.Atoi(c.Param("id"))

	if ok != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не валидное число"})
		return
	}

	product, exits := db.PRODUCTS_DB[id]

	if !exits {

		c.JSON(http.StatusNotFound, gin.H{"error": "Продукт с таким ID не найден"})
		return
	}

	c.JSON(http.StatusOK, product)
}

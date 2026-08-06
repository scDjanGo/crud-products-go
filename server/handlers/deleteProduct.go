package handlers

import (
	"net/http"
	"products/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не валидный ID"})
		return
	}

	 _, exits := db.PRODUCTS_DB[id]

	if !exits {
		c.JSON(http.StatusNotFound, gin.H{"error": "Продукт с таким ID не найден"})
		return
	}

	
	 delete(db.PRODUCTS_DB, id)

	 c.JSON(http.StatusNoContent, gin.H{"action": "Продукт был удалён"})



}
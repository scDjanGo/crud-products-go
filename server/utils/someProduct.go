package utils



import (
	"products/db"
)


func SomeProduct(id int) bool {

	_, exits := db.PRODUCTS_DB[id]

	return  exits
} 
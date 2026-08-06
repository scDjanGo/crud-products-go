package db 




import (
	"products/dto"
)






var PRODUCTS_DB = make(map[int]dto.Product)


var LARGEST_ID int = 1
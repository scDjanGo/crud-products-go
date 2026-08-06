package dto

type ListProduct struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Product struct {
	Id                int     `json:"id"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	Quantity_In_Stock int     `json:"quantity_in_stock,string"`
	Price             float64 `json:"price,string"`
	Discount          float64 `json:"discount,string"`
}

type PatchProduct struct {
	Title             *string  `json:"title"`
	Description       *string  `json:"description"`
	Quantity_In_Stock *int     `json:"quantity_in_stock,string"`
	Price             *float64 `json:"price,string"`
	Discount          *float64 `json:"discount,string"`
}

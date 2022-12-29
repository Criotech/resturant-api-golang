package types

type QueryProducts struct {
	MaxPrice float64 `form:"maxPrice"`
	MinPrice float64 `form:"minPrice"`
	Page     int64   `form:"page"`
	Limit    int64   `form:"limit"`
	Category string  `form:"category"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=2,max=100"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

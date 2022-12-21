package types

type QueryProducts struct {
	MaxPrice float64 `json:"maxPrice"`
	MinPrice float64 `json:"minPrice"`
	Page     int64   `json:"page"`
	Limit    int64   `json:"limit"`
	Category string  `json:"category"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=2,max=100"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

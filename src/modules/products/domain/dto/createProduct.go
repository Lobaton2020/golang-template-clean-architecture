package products

type CreateProductDto struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}
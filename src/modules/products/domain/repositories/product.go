package products

import (
	entities "golang-template-clean-architecture/src/modules/products/domain/entities"
)
type ProductRepository interface{
	FindAll(page, limit int) ([]entities.Product, error)
	Create(p entities.Product) error
}
package products

import (
	dao "golang-template-clean-architecture/src/infraestructure/db/dao"
	dto "golang-template-clean-architecture/src/modules/products/domain/dto"
	products "golang-template-clean-architecture/src/modules/products/domain/entities"
	repository "golang-template-clean-architecture/src/modules/products/domain/repositories"

	"github.com/gofiber/fiber/v2/log"
)

type CreateProductUsecase struct{
	repo repository.ProductRepository
}
func NewCreateProductUsecase(repo *dao.PostgreProductDao) *CreateProductUsecase{
	return &CreateProductUsecase{
		repo: repo,
	}
}
func (puc *CreateProductUsecase) Execute(data *dto.CreateProductDto) (err error) {
	// Where millions of validations
	err = puc.repo.Create(products.Product{
		Name: data.Name,
		Price: data.Price,
	});
	log.Info("Producto creado con exito")
	return
}
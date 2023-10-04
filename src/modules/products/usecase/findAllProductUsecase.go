package products

import (
	dto "golang-template-clean-architecture/src/common/dto"
	dao "golang-template-clean-architecture/src/infraestructure/db/dao"
	entities "golang-template-clean-architecture/src/modules/products/domain/entities"
	repository "golang-template-clean-architecture/src/modules/products/domain/repositories"
)

type FindAllProductUsecase struct{
	repo repository.ProductRepository
}
func NewFindAllProductUsecase(repo *dao.PostgreProductDao) *FindAllProductUsecase{
	return &FindAllProductUsecase{
		repo: repo,
	}
}
func (puc *FindAllProductUsecase) Execute(data *dto.PaginationDto) (result []entities.Product, err error) {
	result, err = puc.repo.FindAll(data.Page, data.Limit);
	if err != nil{
		return
	}
	return
}
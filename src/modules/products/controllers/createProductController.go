package products

import (
	"fmt"

	config "golang-template-clean-architecture/src/common/config"
	r "golang-template-clean-architecture/src/common/response"
	dto "golang-template-clean-architecture/src/modules/products/domain/dto"
	usecase "golang-template-clean-architecture/src/modules/products/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)
type CreateProductController struct{
	Usecase *usecase.CreateProductUsecase
	Result *r.Result
}
func NewCreateProductController(usecase *usecase.CreateProductUsecase, r *r.Result) *CreateProductController{
	return &CreateProductController{
		Usecase: usecase,
		Result: r,
	}
}

func (ph *CreateProductController) ValidateRequest(c *fiber.Ctx) (product dto.CreateProductDto, err error){
	if err = c.BodyParser(&product); err != nil {
		return product, fmt.Errorf(config.BAD_REQUEST, err.Error())
	}
	if err = validator.New().Struct(&product); err != nil {
		return product, fmt.Errorf(config.BAD_REQUEST, err.Error())
	}
	return
}


func (ph *CreateProductController) Run(c *fiber.Ctx) (err error) {
	data, err := ph.ValidateRequest(c)
	if err != nil{
		return ph.Result.Bad(c, "Error at validation: "+ err.Error())
	}
	err = ph.Usecase.Execute(&data)
	if err != nil {
		return ph.Result.Bad(c, "Something went wrong: "+err.Error())
	}
	ph.Result.Ok(c)
	return
}
package products

import (
	"fmt"

	config "golang-template-clean-architecture/src/common/config"
	dto "golang-template-clean-architecture/src/common/dto"
	r "golang-template-clean-architecture/src/common/response"
	usecase "golang-template-clean-architecture/src/modules/products/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)
type FindAllProductController struct{
	Usecase *usecase.FindAllProductUsecase
	Result *r.Result
}
func NewFindAllProductController(usecase *usecase.FindAllProductUsecase, r *r.Result) *FindAllProductController{
	return &FindAllProductController{
		Usecase: usecase,
		Result: r,
	}
}

func (t *FindAllProductController) ValidateRequest(c *fiber.Ctx) (pagination dto.PaginationDto, err error){
	if err = c.QueryParser(&pagination); err != nil {
		return pagination, fmt.Errorf(config.BAD_REQUEST, err.Error())
	}
	if err = validator.New().Struct(&pagination); err != nil {
		return pagination, fmt.Errorf(config.BAD_REQUEST, err.Error())
	}
	return
}


func (t *FindAllProductController) Run(c *fiber.Ctx) error {
	data, err := t.ValidateRequest(c)
	if err != nil{
		return t.Result.Bad(c, "Error at validation: "+ err.Error())
	}
	results, err := t.Usecase.Execute(&data)
	if err != nil {
		return t.Result.Bad(c, "Something went wrong: "+err.Error())
	}
	t.Result.Ok(c, results)
	return nil
}
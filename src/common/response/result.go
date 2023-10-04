package common

import (
	"github.com/gofiber/fiber/v2"
)

func NewResult() *Result{
	return &Result{}
}

type Result struct{}

func (r *Result) Ok(c *fiber.Ctx, data ...interface{}) error {
	var content interface{}
	if(len(data) > 0){
		content = data[0]
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"isError": false,
		"data": content,
	})
}

func (r *Result) Error(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"isError": true,
		"data": data,
	})
}

func (r *Result) Bad(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"isError": true,
		"data": data,
	})
}

func (r *Result) Custom(c *fiber.Ctx, data interface{}, statusCode int) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"isError": true,
		"data": data,
	})
}
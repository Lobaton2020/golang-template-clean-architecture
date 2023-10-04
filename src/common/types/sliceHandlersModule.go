package common

import "github.com/gofiber/fiber/v2"

type HandlerModule struct {
	Handler func(*fiber.Ctx) error
	Route   string
	Method  interface{}
}

type SliceHandlers struct {
	Prefix string
	Routes []HandlerModule
}

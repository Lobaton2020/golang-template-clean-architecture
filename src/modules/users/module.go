package users

import (
	types "golang-template-clean-architecture/src/common/types"

	r "golang-template-clean-architecture/src/common/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)


func configureModuleRoutes(r *r.Result, h *types.HandlersStore)  {

	handlersModuleUsers := &types.SliceHandlers{
		Prefix: "users",
		Routes: []types.HandlerModule{
			{
				Route:   "/",
				Method:  http.MethodPost,
				Handler: func(c *fiber.Ctx) error {
					return r.Ok(c, fiber.Map{"message":"POST Users, success test"})
				},
			},
			{
				Route:   "/",
				Method:  http.MethodGet,
				Handler: func(c *fiber.Ctx) error {
					return r.Ok(c, fiber.Map{"message":"GET Users, success test"})
				},
			},
			{
				Route:   "/disable",
				Method:  http.MethodGet,
				Handler: func(c *fiber.Ctx) error {
					return r.Ok(c, fiber.Map{"message":"GET Users, success JIJI"})
				},
			},
		},
	}
	h.Handlers = append(h.Handlers, *handlersModuleUsers)

}
// Toda la estructura de carpetas del modulo products debe splicarse aqui en dado caso
func ModuleProviders() []fx.Option{
	return []fx.Option{
		fx.Invoke(configureModuleRoutes),
	}
}

package infraestructure

import (
	"context"
	"errors"
	"fmt"
	config "golang-template-clean-architecture/src/common/config"
	types "golang-template-clean-architecture/src/common/types"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/fx"
)

func setRoutesByModule(app *fiber.App, h *types.HandlersStore){
	for _, handlerModule := range h.Handlers{
		route := app.Group("/"+handlerModule.Prefix)
		for _, routeItem := range handlerModule.Routes {
			log.Infof("%v %v%v",routeItem.Method, handlerModule.Prefix, routeItem.Route )
			if routeItem.Method == http.MethodGet{
				route.Get(routeItem.Route, routeItem.Handler)
			}
			if routeItem.Method == http.MethodPost{
				route.Post(routeItem.Route, routeItem.Handler)
			}
			if routeItem.Method == http.MethodPut{
				route.Put(routeItem.Route, routeItem.Handler)
			}
			if routeItem.Method == http.MethodDelete{
				route.Delete(routeItem.Route, routeItem.Handler)
			}
			if routeItem.Method == http.MethodPatch{
				route.Patch(routeItem.Route, routeItem.Handler)
			}
		}
	}
}
func errorHandler(c *fiber.Ctx, err error) error{
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	return c.Status(code).JSON(fiber.Map{
		"isError": true,
		"message": err.Error(),
	})
}
func NewHttpFiberServer(lc fx.Lifecycle, h *types.HandlersStore, cfg * config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	setRoutesByModule(app, h)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting fiber server on port "+cfg.App.Port)
			go app.Listen(":"+cfg.App.Port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}

package products

import (
	types "golang-template-clean-architecture/src/common/types"
	dao "golang-template-clean-architecture/src/infraestructure/db/dao"
	controller "golang-template-clean-architecture/src/modules/products/controllers"
	usecase "golang-template-clean-architecture/src/modules/products/usecase"
	"net/http"

	"go.uber.org/fx"
)


func configureModuleRoutes(
	ctrlCreateProduct *controller.CreateProductController,
	ctrlFindAllProduct *controller.FindAllProductController,
	h *types.HandlersStore,
	) {

	handlersModuleProducts := &types.SliceHandlers{
		Prefix: "products",
		Routes: []types.HandlerModule{
			{
				Route:   "/",
				Method:  http.MethodPost,
				Handler: ctrlCreateProduct.Run,
			},
			{
				Route:   "/",
				Method:  http.MethodGet,
				Handler: ctrlFindAllProduct.Run,
			},
		},
	}
	h.Handlers = append(h.Handlers, *handlersModuleProducts)
}

func ModuleProviders() []fx.Option{
	return []fx.Option{
		fx.Provide(dao.NewPostgreProductDao),
		fx.Provide(controller.NewCreateProductController),
		fx.Provide(controller.NewFindAllProductController),
		fx.Provide(usecase.NewCreateProductUsecase),
		fx.Provide(usecase.NewFindAllProductUsecase),
		fx.Invoke(configureModuleRoutes),
	}
}

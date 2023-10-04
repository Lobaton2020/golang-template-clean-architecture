package products

import (
	config "golang-template-clean-architecture/src/common/config"
	result "golang-template-clean-architecture/src/common/response"
	types "golang-template-clean-architecture/src/common/types"
	db "golang-template-clean-architecture/src/infraestructure/db/adapter"
	dao "golang-template-clean-architecture/src/infraestructure/db/dao"
	server "golang-template-clean-architecture/src/infraestructure/server"
	controller "golang-template-clean-architecture/src/modules/products/controllers"
	usecase "golang-template-clean-architecture/src/modules/products/usecase"

	"net/http"

	"go.uber.org/fx"
)


func configureModuleRoutes(
	ctrlCreateProduct *controller.CreateProductController,
	ctrlFindAllProduct *controller.FindAllProductController,
	) *types.SliceHandlers {
	return &types.SliceHandlers{
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
}

func init() {

	fx.New(
		fx.Provide(result.NewResult),
		fx.Provide(config.NewConfig),
		fx.Provide(db.NewDBConnection),
		fx.Provide(dao.NewPostgreProductDao),
		fx.Provide(controller.NewCreateProductController),
		fx.Provide(usecase.NewCreateProductUsecase),
		fx.Provide(controller.NewFindAllProductController),
		fx.Provide(usecase.NewFindAllPProductUsecase),
		fx.Provide(configureModuleRoutes),
		fx.Invoke(server.NewHttpFiberServer),
	  ).Run()
}

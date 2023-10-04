package main

import (
	server "golang-template-clean-architecture/src/infraestructure/server"
	productModule "golang-template-clean-architecture/src/modules/products"
	userModule "golang-template-clean-architecture/src/modules/users"
)


func main() {
	app := server.ProvidersStore{}
	app.Init()
	app.AddModule(productModule.ModuleProviders())
	app.AddModule(userModule.ModuleProviders())
	app.Up()
}
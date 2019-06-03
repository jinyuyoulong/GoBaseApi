package main

import (
	"project-api/src/bootstrap/route"
	"project-api/src/bootstrap/service"

	"github.com/kataras/iris"
	"github.com/pelletier/go-toml"
)

func main() {
	app := iris.New()

	setApplication(app)
}

func setApplication(app *iris.Application) {

	route.SetRoute(app)

	var port string

	service.BuildContainer()

	di := service.GetDi()
	container := di.Container
	container.Invoke(func(config *toml.Tree) {
		port = config.Get("app.port").(string)
	})

	app.Run(iris.Addr(port))
}

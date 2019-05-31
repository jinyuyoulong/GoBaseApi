package main

import (
	"github.com/kataras/iris"
	"project-api/src/app/bootstrap/route"
	"project-api/src/app/config"
	"project-api/src/app/bootstrap/diserver"
)

func main() {
	app := iris.New()
	diserver.NewServices(app)
	route.Configure(app)
	port := new(config.Config).New().Get("app.port").(string)
	app.Run(iris.Addr(port))
}

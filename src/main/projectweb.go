package main

import (
	"project-api/src/app/bootstrap/diserver"
	"project-api/src/app/bootstrap/route"
	"project-api/src/app/library/helper"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	diserver.NewServices(app)
	route.SetRoute(app)
	port := new(helper.Helper).NewConfig().Get("app.port").(string)
	app.Run(iris.Addr(port))
}

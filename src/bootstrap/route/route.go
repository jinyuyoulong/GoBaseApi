package route

import (
	"project-api/src/controller"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// APIRoute registers the necessary routes to the app.
func APIRoute(route *iris.Application) {
	apiC := new(controller.APIController)
	api := mvc.New(route.Party("/api"))
	api.Handle(apiC)

	// -------------------------------------------------------
	route.Get("/api/setroute", apiC.GetIndexHandler)
	route.Get("/api/user", apiC.GetUser)
}

// SetRoute 配置路由
func SetRoute(route *iris.Application) {
	APIRoute(route)
}

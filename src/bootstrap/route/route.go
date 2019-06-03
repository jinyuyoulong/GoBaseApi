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
	route.Get("/setroute", apiC.GetIndexHandler)
	route.Get("/user", apiC.GetUser)
	// a.Get("/follower/{id:long}", indexC.GetFollowerHandler)
	// b.Get("/like/{id:long}", GetLikeHandler)
}

// SetRoute 配置路由
func SetRoute(route *iris.Application) {
	APIRoute(route)
}

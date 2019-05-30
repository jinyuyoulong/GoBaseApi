package route

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"xxx.com/projectweb/src/app/controller"
)

// Configure registers the necessary routes to the app.
func Configure(a *iris.Application) {
	apiC := new(controller.APIController)
	api := mvc.New(a.Party("/api"))
	api.Handle(apiC)

	// -------------------------------------------------------
	a.Get("/setroute", apiC.GetIndexHandler)
	a.Get("/user", apiC.GetUser)
	// a.Get("/follower/{id:long}", indexC.GetFollowerHandler)
	// b.Get("/like/{id:long}", GetLikeHandler)
}

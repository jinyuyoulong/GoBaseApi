package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/pelletier/go-toml"

	"project-api/src/bootstrap/service"
	"project-api/src/models"
)

// APIController index controller
type APIController struct {
	Ctx iris.Context
}

//  使用DI 中的数据
func useContainer() {
	// 测试外部文件 获取 config 数据
	container := service.GetDi().Container
	container.Invoke(func(config *toml.Tree) {
		println("测试外部文件 获取 config 数据", config.Get("database.dirver").(string))
	})
}

// Get /
func (c *APIController) Get(ctx iris.Context) {
	user := models.CreateUser()
	datalist := user.GetAll()

	ctx.JSON(APIResult(true, datalist, ""))
}

// GetBy /?id=2
func (c *APIController) GetBy(id int) mvc.Result {
	user := models.CreateUser()
	if id < 0 {
		return mvc.Response{
			Path: "/",
		}
	}
	data := user.GetUserInfo(id)
	value := c.Ctx.GetCookie("name")
	println("cookie :", value)

	return mvc.Response{
		ContentType: "application/json",
		Text:        data.Username,
	}

}

// GetIndexHandler 单点路由测试
func (c *APIController) GetIndexHandler(ctx iris.Context) {
	ctx.Writef("Hello from method: %s and path: %s", ctx.Method(), ctx.Path())
}

// GetUser 获取user表数据
func (c *APIController) GetUser(ctx iris.Context) {
	user := models.CreateUser()
	datalist := user.GetAll()

	ctx.JSON(APIResult(true, datalist, ""))
}

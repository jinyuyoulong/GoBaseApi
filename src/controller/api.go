package controller

import (
	"github.com/kataras/iris"
	"github.com/pelletier/go-toml"

	"project-api/src/bootstrap/service"
	"project-api/src/library/session"
	"project-api/src/models"
)

// APIController index controller
type APIController struct {
	Ctx iris.Context
}

// Get /
func (c *APIController) Get() {
	user := models.CreateUser()
	datalist := user.GetAll()

	c.Ctx.JSON(APIResult(true, datalist, ""))
}

// GetBy /api?id=2
func (c *APIController) GetBy(id int) {
	user := models.CreateUser()
	data := user.GetUserInfo(id)
	c.Ctx.JSON(APIResult(true, data, ""))
}

// GetIndexHandler 单点路由测试
func (c *APIController) GetIndexHandler(ctx iris.Context) {
	// 测试外部文件 获取 config 数据
	service.GetDi().Container.Invoke(func(config *toml.Tree) {
		ctx.Writef("Hello from method: %s and path: %s", ctx.Method(), ctx.Path())
		ctx.Writef("测试外部文件 获取 config 数据 %s", config.Get("database.dirver").(string))
	})
}

// GetUser 获取user表数据
func (c *APIController) GetUser(ctx iris.Context) {
	user := models.CreateUser()
	datalist := user.GetAll()

	ctx.JSON(APIResult(true, datalist, ""))
}

// GetSet /set set session in redis
func (c *APIController) GetSet(context iris.Context) {
	session.SessionSet(context, "name", "api")
}

// GetSession /session 获取session
func (c *APIController) GetSession(context iris.Context) {
	name := session.SessionGet(context, "name")
	context.Writef("The name on the /set was: %s", name)
}

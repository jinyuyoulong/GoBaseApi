package controller

import (
	"fmt"
	"log"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"

	"project-api/src/app/bootstrap/diserver"
	"project-api/src/app/bootstrap/service"
	"project-api/src/app/library/datasource"
	"project-api/src/app/library/helper"
	"project-api/src/app/models"
)

// APIController index controller
type APIController struct {
	Ctx iris.Context
}

const (
	commonTitle string = "测试资料库"
)

//  使用DI 中的数据
func useContainer() {
	// 测试外部文件 获取 config 数据
	container := diserver.GetDI().Container
	container.Invoke(func(helper *helper.Helper) {
		println("测试外部文件 获取 config 数据", helper.NewConfig().Get("database.dirver").(string))
	})
}

// Get /
func (c *APIController) Get(ctx iris.Context) {
	Service := service.NewprojectapiService()
	datalist := Service.GetAll()

	ctx.JSON(ApiResult(true, datalist, ""))
}

// GetBy /?id=2
func (c *APIController) GetBy(id int) mvc.Result {
	Service := service.NewprojectapiService()
	if id < 0 {
		return mvc.Response{
			Path: "/",
		}
	}
	data := Service.Get(id)
	value := c.Ctx.GetCookie("name")
	println("cookie :", value)

	return mvc.Response{
		ContentType: "application/json",
		Text:        data.Avatar,
	}

}

// GetSearch /search?country=瑞士
func (c *APIController) GetSearch(ctx iris.Context) {
	Service := service.NewprojectapiService()
	country := c.Ctx.URLParam("country")
	datalist := Service.Search(country)
	fmt.Println("search:", datalist)
	ctx.JSON(ApiResult(true, datalist, ""))
}

// GetIndexHandler 单点路由测试
func (c *APIController) GetIndexHandler(ctx iris.Context) {
	ctx.Writef("Hello from method: %s and path: %s", ctx.Method(), ctx.Path())
}

// GetUser 获取user表数据
func (c *APIController) GetUser(ctx iris.Context) {
	Service := service.NewUserService()
	datalist := Service.GetAll()

	ctx.JSON(ApiResult(true, datalist, ""))
}

// GetClearcache 手动清除缓存 分布式数据同步
func (c *APIController) GetClearcache() mvc.Result {
	err := datasource.InstanceMaster().ClearCache(&models.StarInfo{})
	if err != nil {
		log.Fatal(err)
	}

	return mvc.Response{
		Text: "xorm 缓存清除成功",
	}
}

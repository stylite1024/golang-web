package router

import (
	"iris-web/internal/handler"
	"iris-web/web"

	"github.com/kataras/iris/v12"
)

func InitRouter(app *iris.Application) {
	// 静态文件路由
	staticFileRouter(app)
	// api路由
	apiRouter(app)
}

func staticFileRouter(app *iris.Application) {
	app.HandleDir("/", web.StaticFS, iris.DirOptions{
		IndexName:      "index.html", // 设置 index.html 作为默认页面
	})

	// app.Get("/", func(c iris.Context) {
	// 	c.Redirect("/index.html", iris.StatusMovedPermanently)
	// })


}

func apiRouter(app *iris.Application) {
	v1 := app.Party("/api/v1")
	{
		v1.Get("/test", handler.TestApi)
	}
}

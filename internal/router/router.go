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
	// app.HandleDir("/frontend", iris.Dir("web/frontend"))
	// app.HandleDir("/backend", iris.Dir("web/backend"))
	app.HandleDir("/frontend", web.FrontendFS)
	app.HandleDir("/backend", web.BackendFS)

	app.Get("/", func(c iris.Context) {
		c.Redirect("/frontend/index.html", iris.StatusMovedPermanently)
	})
	app.Get("/admin", func(c iris.Context) {
		c.Redirect("/backend/index.html", iris.StatusMovedPermanently)
	})
}

func apiRouter(app *iris.Application) {
	v1 := app.Party("/api/v1")
	{
		v1.Get("/test", handler.TestApi)
	}
}

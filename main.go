package main

import (
	"txwzgl/controller"
	"txwzgl/tools"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
)

func main() {

	tools.InitDB()
	app := iris.New()
	app.RegisterView(iris.HTML("./templates", ".html"))

	tools.Initclient()
	tools.Sess.UseDatabase(tools.INCLIENT)

	app.StaticWeb("/static", "./static")
	app.Use(func(ctx iris.Context) {
		ph := ctx.Path()

		if ph == "/login" {
			ctx.Next()
		} else {
			s := tools.Sess.Start(ctx).GetString("name")
			if s == "" {
				ctx.Redirect("/static/login.html")
			} else {
				ctx.Next()
			}
		}
	})
	app.Get("/gdzc_list", controller.ListGdzc)
	app.Get("/delete_gdzc", controller.DeleteGdzc)
	app.Post("/update_gdzc", controller.UpdateGdzcP)
	app.Get("/update_gdzc", controller.UpdateGdzcG)
	app.Post("/insert_gdzc", controller.InsertGdzc)
	app.Post("/login", controller.Login)
	app.Get("/logout", controller.Logout)

	app.Get("/", func(ctx iris.Context) {
		ctx.Redirect("/static/login.html", iris.StatusTemporaryRedirect)
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

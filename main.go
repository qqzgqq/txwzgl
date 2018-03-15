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

	// 固定资产列表增删改
	app.Get("/list_gdzc", controller.ListGdzc)
	app.Get("/delete_gdzc", controller.DeleteGdzc)
	app.Post("/update_gdzc", controller.UpdateGdzcP)
	app.Get("/update_gdzc", controller.UpdateGdzcG)
	app.Post("/insert_gdzc", controller.InsertGdzc)

	// 人员列表
	app.Get("/list_people", controller.ListPeople)
	app.Get("/insert_people", controller.InsertPeopleG)
	app.Post("/insert_people", controller.InsertPeopleP)
	app.Get("/update_people", controller.UpdatePeopleG)
	app.Post("/update_people", controller.UpdatePeopleP)
	app.Get("/delete_people", controller.DeletePeople)
	app.Get("/search_people", controller.SearchPeople)

	// 用户登录
	app.Post("/login", controller.Login)
	// 注销用户
	app.Get("/logout", controller.Logout)

	app.Get("/", func(ctx iris.Context) {
		ctx.Redirect("/static/login.html", iris.StatusTemporaryRedirect)
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

package controller

import (
	"fmt"
	"txwzgl/tools"

	"github.com/kataras/iris"
)

// InsertPeopleG 公司人员列表
func InsertPeopleG(ctx iris.Context) {

	adff := tools.Sess.Start(ctx).GetString("name")
	ctx.ViewData("sname", adff)
	ctx.View("people_insert.html")

}

// InsertPeopleP 公司人员列表
func InsertPeopleP(ctx iris.Context) {
	name := ctx.FormValue("name")
	position := ctx.FormValue("position")
	bumen := ctx.FormValue("bumen")
	lxfs := ctx.FormValue("lxfs")
	if name == "" {
		fmt.Println("请正确填写信息", name)
	} else {
		tools.DB.MustExec("insert into people(name,position,bumen,telephone) value (?,?,?,?)", name, position, bumen, lxfs)
	}

	ctx.Redirect("/list_people")

}

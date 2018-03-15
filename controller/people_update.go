package controller

import (
	"fmt"
	"txwzgl/model"
	"txwzgl/tools"

	"github.com/kataras/iris"
)

// UpdatePeopleG 公司人员列表
func UpdatePeopleG(ctx iris.Context) {
	adff := tools.Sess.Start(ctx).GetString("name")
	ctx.ViewData("sname", adff)
	updateid := ctx.FormValue("updateid")
	peoplesql := []model.People{}
	tools.DB.Select(&peoplesql, "select * from people where id = ?", updateid)
	fmt.Println("get:", peoplesql)
	ctx.ViewData("peoplesql", peoplesql)
	ctx.View("people_update.html")
}

// UpdatePeopleP 公司人员列表
func UpdatePeopleP(ctx iris.Context) {
	updateid := ctx.FormValue("idpeople")

	name := ctx.FormValue("name")
	position := ctx.FormValue("position")
	bumen := ctx.FormValue("bumen")
	lxfs := ctx.FormValue("lxfs")
	fmt.Println("post id ", updateid, name, position)
	tools.DB.MustExec("update people set name = ?, position = ?, bumen = ?, telephone = ? where id = ?", name, position, bumen, lxfs, updateid)
	ctx.Redirect("/list_people")
}

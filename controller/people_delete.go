package controller

import (
	"txwzgl/tools"

	"github.com/kataras/iris"
)

// DeletePeople 公司人员列表
func DeletePeople(ctx iris.Context) {
	deleteidl := ctx.FormValue("deleteid_l")
	tools.DB.MustExec("delete from people where id = ?", deleteidl)
	ctx.Redirect("/list_people")
}

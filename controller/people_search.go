package controller

import (
	"fmt"
	"txwzgl/model"
	"txwzgl/tools"

	"github.com/kataras/iris"
)

// SearchPeople 人员查找
func SearchPeople(ctx iris.Context) {

	adff := tools.Sess.Start(ctx).GetString("name")
	ctx.ViewData("sname", adff)
	xingming := ctx.FormValue("xingming")
	peoplesqll := []model.People{}

	if xingming == "" {
		tools.DB.Select(&peoplesqll, "select * from people")

	} else {
		tools.DB.Select(&peoplesqll, "select * from people where name = ?", xingming)

	}
	fmt.Println("search id is :", ctx.FormValue("idpeople"))
	ctx.ViewData("peoplesqll", peoplesqll)
	ctx.View("people_search.html")

}

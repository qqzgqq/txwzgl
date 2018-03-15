package controller

import (
	"strconv"
	"txwzgl/model"
	"txwzgl/tools"

	"github.com/kataras/iris"
)

// ListPeople 公司人员列表
func ListPeople(ctx iris.Context) {
	xsye := GetSqlxsyspeople()
	var aaa []int
	for i := 0; i < xsye; i++ {
		aaa = append(aaa, i+1)

	}

	adff := tools.Sess.Start(ctx).GetString("name")
	ctx.ViewData("sname", adff)
	ctx.ViewData("sqlnum", aaa)
	//ctx.FormValue() must don't be null
	qunayeen := ctx.URLParams()["qunaye"]

	if qunayeen == "" {
		qunayeen = "1"
	}
	//string to int
	qunayeZ, _ := strconv.Atoi(qunayeen)
	qunayeS := qunayeZ*15 - 15

	peopleList := []model.People{}

	tools.DB.Select(&peopleList, "SELECT * FROM people limit ?,?", qunayeS, 15)

	ctx.ViewData("sqll", peopleList)

	ctx.View("people_list.html")
}

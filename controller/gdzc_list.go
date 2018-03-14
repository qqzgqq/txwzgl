package controller

import (
	"strconv"
	"txwzgl/model"
	"txwzgl/tools"

	"github.com/kataras/iris"
)

// ListGdzc 获取固定资产列表信息
func ListGdzc(ctx iris.Context) {

	xsye := GetSqlxsys()
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

	sqlList := []model.Gdzc{}
	questionlist := []model.Gdzc{}
	placelist := []model.Gdzc{}
	questiontypelist := []model.Gdzc{}
	countlist := []model.Gdzc{}
	tools.DB.Select(&sqlList, "SELECT * FROM gdzc limit ?,?", qunayeS, 15)
	tools.DB.Select(&questionlist, "SELECT question_list FROM moban where question_list !='' group by question_list")
	tools.DB.Select(&placelist, "SELECT place_list FROM moban where place_list !='' group by place_list")
	tools.DB.Select(&questiontypelist, "SELECT questiontype_list FROM moban where questiontype_list !='' group by questiontype_list")
	tools.DB.Select(&countlist, "SELECT count_list FROM moban where count_list !='' group by count_list")

	ctx.ViewData("sqll", sqlList)
	ctx.ViewData("question_list", questionlist)
	ctx.ViewData("place_list", placelist)
	ctx.ViewData("questiontype_list", questiontypelist)
	ctx.ViewData("count_list", countlist)
	ctx.View("gdzc_list.html")
	// }
}

package controller

import (
	"fmt"
	"txwzgl/tools"

	"github.com/kataras/iris"
)

// DeleteGdzc delete固定资产表信息
func DeleteGdzc(ctx iris.Context) {
	deleteidgdzc := ctx.FormValue("deleteid_gdzc")
	deleteidgdzcurl := ctx.FormValue("deleteid_gdzc_url")
	tools.DB.MustExec("delete from gdzc where id=?", deleteidgdzc)
	fmt.Println(deleteidgdzc, "删除成功")
	ctx.Redirect(deleteidgdzcurl, iris.StatusTemporaryRedirect)

}

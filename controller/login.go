package controller

import (
	"txwzgl/tools"

	"github.com/kataras/iris"
)

// Login 验证用户信息
func Login(ctx iris.Context) {
	user := ctx.FormValue("lguser")
	passwd := ctx.FormValue("lgpasswd")
	var loginc int
	tools.DB.Get(&loginc, "select count(*) from user where name = ? and passwd = ?", user, passwd)
	if loginc == 1 {
		s := tools.Sess.Start(ctx)
		s.Set("name", user)
		ctx.Redirect("/list_gdzc")
	} else {
		ctx.HTML("用户不存在")
	}

}

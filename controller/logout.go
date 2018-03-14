package controller

import (
	"txwzgl/tools"

	"github.com/kataras/iris"
)

// Logout 退出登录信息并删除sesson
func Logout(ctx iris.Context) {
	s := tools.Sess.Start(ctx).GetString("name")
	if s != "" {
		tools.Sess.Start(ctx).Delete("name")
	}
	ctx.Redirect("/static/login.html")
}

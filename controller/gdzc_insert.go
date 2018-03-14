package controller

import (
	"fmt"
	"strconv"
	"time"
	"txwzgl/tools"

	"github.com/kataras/iris"
)

// InsertGdzc 插入固定资产表信息
func InsertGdzc(ctx iris.Context) {
	zcbm := ctx.FormValue("zcbm")
	place := ctx.FormValue("place")
	zclx := ctx.FormValue("zclx")
	pinpai := ctx.FormValue("pinpai")
	ggxh := ctx.FormValue("ggxh")
	pzmx := ctx.FormValue("pzmx")
	sn := ctx.FormValue("sn")
	gzsj := ctx.FormValue("gzsj")
	gmqd := ctx.FormValue("gmqd")
	cfdd := ctx.FormValue("cfdd")
	zcyz := ctx.FormValue("zcyz")
	sfsy := ctx.FormValue("sfsy")
	wxzt := ctx.FormValue("wxzt")
	zsqk := ctx.FormValue("zsqk")
	bfqx := ctx.FormValue("bfqx")
	bz := ctx.FormValue("bz")
	if gzsj == "" {
		tools.DB.MustExec("insert into gdzc(zcbm, place, zclx, pinpai, xinghao, pzmx, sn, buytime, buyqd, cfdd, zcyz, syzt, whzt, zsqk, bfqx,llbfrq,bfsjc,lubf,sfbf,bftime,bz) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, "--", gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, "--", "--", "--", "--", "--", "--", bz)
	} else {
		if bfqx == "" {
			tools.DB.MustExec("insert into gdzc(zcbm, place, zclx, pinpai, xinghao, pzmx, sn, buytime, buyqd, cfdd, zcyz, syzt, whzt, zsqk, bfqx,llbfrq,bfsjc,lubf,sfbf,bftime,bz) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, "--", "--", "--", "--", "--", "--", bz)
		} else {
			buytimeT, _ := time.Parse("2006-01-02", gzsj)
			bfqxT, _ := strconv.Atoi(bfqx)
			llbfrq := buytimeT.AddDate(bfqxT, 0, 0).Format("2006-01-02")
			llbfrqT, _ := time.Parse("2006-01-02", llbfrq)
			bfsjcT := (llbfrqT.Unix() - time.Now().Unix()) / 60 / 60 / 24
			bfqxS := bfqx + "年"
			bfsjcS := strconv.FormatInt(bfsjcT, 10) + "天"

			tools.DB.MustExec("insert into gdzc(zcbm, place, zclx, pinpai, xinghao, pzmx, sn, buytime, buyqd, cfdd, zcyz, syzt, whzt, zsqk, bfqx,llbfrq,bfsjc,lubf,sfbf,bftime,bz) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, bfqxS, llbfrq, bfsjcS, llbfrq, "--", "--", bz)

			fmt.Println("购买日期", gzsj)
			fmt.Println("理论到期时间", llbfrq)
			fmt.Println("报废时间差", bfsjcT, "天")

		}
	}

	zxsys := GetSqlxsys()
	zxsS := strconv.Itoa(zxsys)
	zxsZ := "/gdzc_list/?qunaye=" + zxsS
	fmt.Println("去哪页", zxsZ)

	ctx.Redirect(zxsZ)
}

package controller

import (
	"fmt"
	"strconv"
	"time"
	"txwzgl/model"
	"txwzgl/tools"

	"github.com/kataras/iris"
)

// UpdateGdzcG 修改固定资产信息get方法
func UpdateGdzcG(ctx iris.Context) {
	updategdzcid := ctx.FormValue("updateid_gdzc")
	updategdzcsql := []model.Gdzc{}
	tools.DB.Select(&updategdzcsql, "select * from gdzc where id = ?", updategdzcid)
	ctx.ViewData("updategdzcsql", updategdzcsql)
	ctx.View("gdzc_update.html")
}

// UpdateGdzcP 修改固定资产信息post方法
func UpdateGdzcP(ctx iris.Context) {
	id := ctx.FormValue("id")
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
	sfbf := ctx.FormValue("sfbf")
	bftime := ctx.FormValue("bftime")
	bz := ctx.FormValue("bz")
	gzsj2 := ctx.FormValue("gzsj2")
	zxsS := strconv.Itoa(GetSqlxsys())
	zxsZ := "/gdzc_list/?qunaye=" + zxsS
	if gzsj == "" {
		if gzsj2 == "" {
			tools.DB.MustExec("update gdzc set zcbm = ?, place = ?, zclx = ?, pinpai = ?, xinghao = ?,pzmx = ?, sn = ?,buytime = ?,buyqd = ?,cfdd = ?, zcyz = ?, syzt = ?, whzt = ?, zsqk = ?,bfqx = ?,llbfrq = ?,bfsjc = ?,lubf = ?,sfbf = ?,bftime = ?,bz =? where id = ?", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, bfqx, "--", "--", "--", sfbf, bftime, bz, id)
			fmt.Println("gzsj为空，gzsj2为空")
		} else if gzsj2 == "--" {
			tools.DB.MustExec("update gdzc set zcbm = ?, place = ?, zclx = ?, pinpai = ?, xinghao = ?,pzmx = ?, sn = ?,buytime = ?,buyqd = ?,cfdd = ?, zcyz = ?, syzt = ?, whzt = ?, zsqk = ?,bfqx = ?,llbfrq = ?,bfsjc = ?,lubf = ?,sfbf = ?,bftime = ?,bz =? where id = ?", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj2, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, "---", "---", "---", "---", sfbf, bftime, bz, id)
			fmt.Println("gzsj为空，gzsj2为--")
		} else {
			buytimeT, _ := time.Parse("2006-01-02", gzsj2)

			fmt.Println("报废期限为------：", bfqx)
			_, err := strconv.Atoi(bfqx)
			if err != nil {
				if bfqx == "" {
					tools.DB.MustExec("update gdzc set zcbm = ?, place = ?, zclx = ?, pinpai = ?, xinghao = ?,pzmx = ?, sn = ?,buytime = ?,buyqd = ?,cfdd = ?, zcyz = ?, syzt = ?, whzt = ?, zsqk = ?,bfqx = ?,llbfrq = ?,bfsjc = ?,lubf = ?,sfbf = ?,bftime = ?,bz =? where id = ?", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj2, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, "---", "---", "---", "---", sfbf, bftime, bz, id)
				} else {
					bfqxd := []rune(bfqx)
					bfqxs := string(bfqxd[0 : len(bfqxd)-1])
					fmt.Println("报废期限为---------------------------：", bfqxs)
					bfqxS, _ := strconv.Atoi(bfqxs)
					llbfrq := buytimeT.AddDate(bfqxS, 0, 0).Format("2006-01-02")
					llbfrqT, _ := time.Parse("2006-01-02", llbfrq)
					bfsjcT := (llbfrqT.Unix() - time.Now().Unix()) / 60 / 60 / 24
					bfsjcS := strconv.FormatInt(bfsjcT, 10) + "天"
					tools.DB.MustExec("update gdzc set zcbm = ?, place = ?, zclx = ?, pinpai = ?, xinghao = ?,pzmx = ?, sn = ?,buytime = ?,buyqd = ?,cfdd = ?, zcyz = ?, syzt = ?, whzt = ?, zsqk = ?,bfqx = ?,llbfrq = ?,bfsjc = ?,lubf = ?,sfbf = ?,bftime = ?,bz =? where id = ?", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj2, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, bfqx, llbfrq, bfsjcS, llbfrq, sfbf, bftime, bz, id)
				}
			} else {
				bfqxS, _ := strconv.Atoi(bfqx)
				llbfrq := buytimeT.AddDate(bfqxS, 0, 0).Format("2006-01-02")
				llbfrqT, _ := time.Parse("2006-01-02", llbfrq)
				bfsjcT := (llbfrqT.Unix() - time.Now().Unix()) / 60 / 60 / 24
				bfsjcS := strconv.FormatInt(bfsjcT, 10) + "天"
				tools.DB.MustExec("update gdzc set zcbm = ?, place = ?, zclx = ?, pinpai = ?, xinghao = ?,pzmx = ?, sn = ?,buytime = ?,buyqd = ?,cfdd = ?, zcyz = ?, syzt = ?, whzt = ?, zsqk = ?,bfqx = ?,llbfrq = ?,bfsjc = ?,lubf = ?,sfbf = ?,bftime = ?,bz =? where id = ?", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj2, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, bfqx, llbfrq, bfsjcS, llbfrq, sfbf, bftime, bz, id)
			}

		}
	} else {
		buytimeT, _ := time.Parse("2006-01-02", gzsj)
		_, err := strconv.Atoi(bfqx)
		if err != nil {
			if bfqx == "---" {
				tools.DB.MustExec("update gdzc set zcbm = ?, place = ?, zclx = ?, pinpai = ?, xinghao = ?,pzmx = ?, sn = ?,buytime = ?,buyqd = ?,cfdd = ?, zcyz = ?, syzt = ?, whzt = ?, zsqk = ?,bfqx = ?,llbfrq = ?,bfsjc = ?,lubf = ?,sfbf = ?,bftime = ?,bz =? where id = ?", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, "---", "---", "---", "---", sfbf, bftime, bz, id)
			} else if bfqx == "" {
				tools.DB.MustExec("update gdzc set zcbm = ?, place = ?, zclx = ?, pinpai = ?, xinghao = ?,pzmx = ?, sn = ?,buytime = ?,buyqd = ?,cfdd = ?, zcyz = ?, syzt = ?, whzt = ?, zsqk = ?,bfqx = ?,llbfrq = ?,bfsjc = ?,lubf = ?,sfbf = ?,bftime = ?,bz =? where id = ?", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, "---", "---", "---", "---", sfbf, bftime, bz, id)
			} else {
				bfqxd := []rune(bfqx)
				bfqxs := string(bfqxd[0 : len(bfqxd)-1])
				fmt.Println("gzsj不为空")
				fmt.Println("报废期限为---------------------------：", bfqx)
				bfqxS, _ := strconv.Atoi(bfqxs)
				llbfrq := buytimeT.AddDate(bfqxS, 0, 0).Format("2006-01-02")
				llbfrqT, _ := time.Parse("2006-01-02", llbfrq)
				bfsjcT := (llbfrqT.Unix() - time.Now().Unix()) / 60 / 60 / 24
				bfsjcS := strconv.FormatInt(bfsjcT, 10) + "天"
				tools.DB.MustExec("update gdzc set zcbm = ?, place = ?, zclx = ?, pinpai = ?, xinghao = ?,pzmx = ?, sn = ?,buytime = ?,buyqd = ?,cfdd = ?, zcyz = ?, syzt = ?, whzt = ?, zsqk = ?,bfqx = ?,llbfrq = ?,bfsjc = ?,lubf = ?,sfbf = ?,bftime = ?,bz =? where id = ?", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, bfqx, llbfrq, bfsjcS, llbfrq, sfbf, bftime, bz, id)
			}
		} else {
			tools.DB.MustExec("update gdzc set zcbm = ?, place = ?, zclx = ?, pinpai = ?, xinghao = ?,pzmx = ?, sn = ?,buytime = ?,buyqd = ?,cfdd = ?, zcyz = ?, syzt = ?, whzt = ?, zsqk = ?,bfqx = ?,llbfrq = ?,bfsjc = ?,lubf = ?,sfbf = ?,bftime = ?,bz =? where id = ?", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, "---", "---", "---", "---", sfbf, bftime, bz, id)
			fmt.Println("gzsj存在，报废期限不存在")
		}
	}
	ctx.Redirect(zxsZ)
	ctx.View("gdzc_update.html")
}

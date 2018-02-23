package main

import (
	"fmt"
	"math"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

// Gdzc mysql表gdzc
type Gdzc struct {
	IDgdzc  int    `db:"id"`
	Zcbm    string `db:"zcbm"`
	Place   string `db:"place"`
	Zclx    string `db:"zclx"`
	Pinpai  string `db:"pinpai"`
	Xinghao string `db:"xinghao"`
	Pzmx    string `db:"pzmx"`
	Sn      string `db:"sn"`
	Buytime string `db:"buytime"`
	Buyqd   string `db:"buyqd"`
	Cfdd    string `db:"cfdd"`
	Zcyz    string `db:"zcyz"`
	Syzt    string `db:"syzt"`
	Whzt    string `db:"whzt"`
	Zsqk    string `db:"zsqk"`
	Bfqx    string `db:"bfqx"`
	Llbfrq  string `db:"llbfrq"`
	Bfsjc   string `db:"bfsjc"`
	Lubf    string `db:"lubf"`
	Sfbf    string `db:"sfbf"`
	Bftime  string `db:"bftime"`
	BZ      string `db:"bz"`
}

// User mysql表user
type User struct {
	IDUser int    `db:"id"`
	Name   string `db:"name"`
	Passwd string `db:"passwd"`
}

// BuyInfo mysql表buyinfo
type BuyInfo struct {
	IDbuyinfo int    `db:"id"`
	Zcmc      string `db:"zcmc"`
	Ggxh      string `db:"ggxh"`
	Jldw      string `db:"jldw"`
	Gzqd      string `db:"gzqd"`
	Buytime   string `db:"buytime"`
	Buynum    int    `db:"buynum"`
	Danjia    int    `db:"danjia"`
	Gzcb      int    `db:"gzcb"`
}

// Equipment mysql表equipment
type Equipment struct {
	IDequipment int    `db:"id"`
	Zclx        string `db:"zclx"`
	Pinpai      string `db:"pinpai"`
	Xinghao     string `db:"xinghao"`
	Pzmx        string `db:"pzmx"`
	Qudao       string `db:"qudao"`
	Syzt        string `db:"syzt"`
}

// Feigdzc mysql表feigdzc
type Feigdzc struct {
	IDfeigdzc int    `db:"id"`
	Zcmc      string `db:"zcmc"`
	Pinpai    string `db:"pinpai"`
	Ggxh      string `db:"ggxh"`
	jldw      string `db:"jldw"`
	buynum    int    `db:"buynum"`
	usenum    int    `db:"usenum"`
	yuliang   int    `db:"yuliang"`
	place     string `db:"place"`
	bz        string `db:"bz"`
}

// People mysql表people
type People struct {
	IDPeople  int    `db:"id"`
	Name      string `db:"name"`
	Bumen     string `db:"bumen"`
	Telephone string `db:"telephone"`
}

// Makeinfo mysql表makeinfo
type Makeinfo struct {
	IDMakeinfo int    `db:"id"`
	Zcbm       string `db:"zcbm"`
	Zclx       string `db:"zclx"`
	Pinpai     string `db:"pinpai"`
	Xinghao    string `db:"xinghao"`
	Pzmx       string `db:"pzmx"`
	Sn         string `db:"sn"`
	Maketime   string `db:"maketime"`
	Makemoney  string `db:"makemoney"`
	Makedw     string `db:"makedw"`
	Bz         string `db:"bz"`
}

// Gdinorout mysql表gdinorout
type Gdinorout struct {
	IDGdinorout int    `db:"id"`
	Zcbm        string `db:"zcbm"`
	Place       string `db:"place"`
	Leixing     string `db:"leixing"`
	Pinpai      string `db:"pinpai"`
	Xinghao     string `db:"xinghao"`
	Pzmx        string `db:"pzmx"`
	Sn          string `db:"sn"`
	Lybm        string `db:"lybm"`
	Ygbh        string `db:"ygbh"`
	User        string `db:"user"`
	Usetime     string `db:"usetime"`
	Backtime    string `db:"backtime"`
	Bz          string `db:"bz"`
}

// Fgdinorout mysql表fgdinorout
type Fgdinorout struct {
	IDFgdinorout int    `db:"id"`
	Gettime      string `db:"gettime"`
	Wpmc         string `db:"wpmc"`
	Num          int    `db:"num"`
	Danwei       string `db:"danwei"`
	Lysy         string `db:"lysy"`
	Lybm         string `db:"lybm"`
	User         string `db:"user"`
}

func main() {

	db, _ := sqlx.Connect("mysql", "root:1q2w3e@tcp(127.0.0.1:3306)/txwzglxt?charset=utf8")
	app := iris.New()
	app.RegisterView(iris.HTML("./templates", ".html"))
	sess := sessions.New(sessions.Config{
		Cookie:                      "mysessionid",
		Expires:                     time.Hour * 2,
		DisableSubdomainPersistence: false,
	})
	app.StaticWeb("/static", "./static")
	app.Use(func(ctx iris.Context) {
		ph := ctx.Path()
		if ph == "/login" {
			ctx.Next()
		} else {
			s := sess.Start(ctx).GetString("name")
			if s == "" {
				ctx.Redirect("/static/login.html", iris.StatusTemporaryRedirect)
			} else {
				ctx.Next()
			}

		}
	})
	app.Get("/gdzc_list", func(ctx iris.Context) {
		// s := sess.Start(ctx).GetString("name")
		// if s != "" {
		var page float64
		db.Get(&page, "SELECT count(*) FROM gdzc ")
		xsys := 15.0
		zxsys := math.Ceil(page / xsys)
		zxsZ := int(zxsys)
		var aaa []int
		for i := 0; i < zxsZ; i++ {
			aaa = append(aaa, i+1)

		}
		ctx.ViewData("sname", sess.Start(ctx).GetString("name"))
		ctx.ViewData("sqlnum", aaa)
		//ctx.FormValue() must don't be null
		qunayeen := ctx.URLParams()["qunaye"]

		if qunayeen == "" {
			qunayeen = "1"
		}
		//string to int
		qunayeZ, _ := strconv.Atoi(qunayeen)
		qunayeS := qunayeZ*15 - 15

		sqlList := []Gdzc{}
		questionlist := []Gdzc{}
		placelist := []Gdzc{}
		questiontypelist := []Gdzc{}
		countlist := []Gdzc{}
		db.Select(&sqlList, "SELECT * FROM gdzc limit ?,?", qunayeS, 15)
		db.Select(&questionlist, "SELECT question_list FROM moban where question_list !='' group by question_list")
		db.Select(&placelist, "SELECT place_list FROM moban where place_list !='' group by place_list")
		db.Select(&questiontypelist, "SELECT questiontype_list FROM moban where questiontype_list !='' group by questiontype_list")
		db.Select(&countlist, "SELECT count_list FROM moban where count_list !='' group by count_list")

		ctx.ViewData("sqll", sqlList)
		ctx.ViewData("question_list", questionlist)
		ctx.ViewData("place_list", placelist)
		ctx.ViewData("questiontype_list", questiontypelist)
		ctx.ViewData("count_list", countlist)
		ctx.View("gdzc_list.html")
		// }

	})
	app.Get("/delete_gdzc", func(ctx iris.Context) {
		// s := sess.Start(ctx).GetString("name")
		// if s != "" {
		deleteidgdzc := ctx.FormValue("deleteid_gdzc")
		deleteidgdzcurl := ctx.FormValue("deleteid_gdzc_url")
		db.MustExec("delete from gdzc where id=?", deleteidgdzc)
		fmt.Println(deleteidgdzc, "删除成功")
		fmt.Printf(deleteidgdzcurl)

		ctx.Redirect(deleteidgdzcurl, iris.StatusTemporaryRedirect)

		// }

	})
	app.Post("/update_gdzc", func(ctx iris.Context) {

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

		// zcbm2 := ctx.FormValue("zcbm2")
		// place2 := ctx.FormValue("place2")
		// zclx2 := ctx.FormValue("zclx2")
		// pinpai2 := ctx.FormValue("pinpai2")
		// ggxh2 := ctx.FormValue("ggxh2")
		// pzmx2 := ctx.FormValue("pzmx2")
		// sn2 := ctx.FormValue("sn2")
		gzsj2 := ctx.FormValue("gzsj2")
		// gmqd2 := ctx.FormValue("gmqd2")
		// cfdd2 := ctx.FormValue("cfdd2")
		// zcyz2 := ctx.FormValue("zcyz2")
		// sfsy2 := ctx.FormValue("sfsy2")
		// wxzt2 := ctx.FormValue("wxzt2")
		// zsqk2 := ctx.FormValue("zsqk2")
		// bfqx2 := ctx.FormValue("bfqx2")
		// sfbf2 := ctx.FormValue("sfbf2")
		// bftime2 := ctx.FormValue("bftime2")
		// bz2 := ctx.FormValue("bz2")
		var page float64
		db.Get(&page, "SELECT count(*) FROM gdzc ")
		xsys := 15.0
		zxsys := math.Ceil(page / xsys)
		zxsS := strconv.FormatFloat(zxsys, 'f', -1, 64)
		zxsZ := "/gdzc_list/?qunaye=" + zxsS
		fmt.Println("post方法：", gzsj2, id, zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, bfqx, sfbf, bftime, bz)
		if gzsj == "" {
			buytimeT, _ := time.Parse("2006-01-02", gzsj2)
			if len(bfqx) == 4 {
				bfqxI := string(bfqx[0])
				fmt.Println("报废期限为：", bfqxI)
				bfqxT, _ := strconv.Atoi(bfqxI)
				llbfrq := buytimeT.AddDate(bfqxT, 0, 0).Format("2006-01-02")
				llbfrqT, _ := time.Parse("2006-01-02", llbfrq)
				bfsjcT := (llbfrqT.Unix() - time.Now().Unix()) / 60 / 60 / 24
				bfqxS := bfqxI + "年"
				bfsjcS := strconv.FormatInt(bfsjcT, 10) + "天"

				// db.MustExec("insert into gdzc(zcbm, place, zclx, pinpai, xinghao, pzmx, sn, buytime, buyqd, cfdd, zcyz, syzt, whzt, zsqk, bfqx,llbfrq,bfsjc,lubf,sfbf,bftime,bz) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, bfqxS, llbfrq, bfsjcS, llbfrq, "--", "--", bz)

				fmt.Println("购买日期", gzsj2)
				fmt.Println("理论到期时间", llbfrq)
				fmt.Println("报废期限", bfqxS)
				fmt.Println("报废时间差", bfsjcS)
			} else if len(bfqx) == 5 {
				bfqxI := string(bfqx[0]) + string(bfqx[1])
				fmt.Println("报废期限为：", bfqxI)
				bfqxT, _ := strconv.Atoi(bfqxI)
				llbfrq := buytimeT.AddDate(bfqxT, 0, 0).Format("2006-01-02")
				llbfrqT, _ := time.Parse("2006-01-02", llbfrq)
				bfsjcT := (llbfrqT.Unix() - time.Now().Unix()) / 60 / 60 / 24
				bfqxS := bfqxI + "年"
				bfsjcS := strconv.FormatInt(bfsjcT, 10) + "天"

				// db.MustExec("insert into gdzc(zcbm, place, zclx, pinpai, xinghao, pzmx, sn, buytime, buyqd, cfdd, zcyz, syzt, whzt, zsqk, bfqx,llbfrq,bfsjc,lubf,sfbf,bftime,bz) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, bfqxS, llbfrq, bfsjcS, llbfrq, "--", "--", bz)

				fmt.Println("购买日期", gzsj2)
				fmt.Println("理论到期时间", llbfrq)
				fmt.Println("报废期限", bfqxS)
				fmt.Println("报废时间差", bfsjcS)
			} else {
				bfqxI := bfqx
				fmt.Println("报废期限为：", bfqxI)
				bfqxT, _ := strconv.Atoi(bfqxI)
				llbfrq := buytimeT.AddDate(bfqxT, 0, 0).Format("2006-01-02")
				llbfrqT, _ := time.Parse("2006-01-02", llbfrq)
				bfsjcT := (llbfrqT.Unix() - time.Now().Unix()) / 60 / 60 / 24
				bfqxS := bfqxI + "年"
				bfsjcS := strconv.FormatInt(bfsjcT, 10) + "天"

				// db.MustExec("insert into gdzc(zcbm, place, zclx, pinpai, xinghao, pzmx, sn, buytime, buyqd, cfdd, zcyz, syzt, whzt, zsqk, bfqx,llbfrq,bfsjc,lubf,sfbf,bftime,bz) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, bfqxS, llbfrq, bfsjcS, llbfrq, "--", "--", bz)

				fmt.Println("购买日期", gzsj2)
				fmt.Println("理论到期时间", llbfrq)
				fmt.Println("报废期限", bfqxS)
				fmt.Println("报废时间差", bfsjcS)
			}

			// db.MustExec("insert into gdzc(zcbm, place, zclx, pinpai, xinghao, pzmx, sn, buytime, buyqd, cfdd, zcyz, syzt, whzt, zsqk, bfqx,llbfrq,bfsjc,lubf,sfbf,bftime,bz) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, bfqxS, llbfrq, bfsjcS, llbfrq, "--", "--", bz)

			// db.MustExec("update gdzc set zcbm = ?, place = ?, zclx = ?, pinpai = ?, xinghao = ?, pzmx = ?, sn = ?, buytime = ?, buyqd = ?, cfdd = ?, zcyz = ?, syzt = ?, whzt = ?, zsqk = ?, bfqx = ?,sfbf = ?,bftime = ?,bz =? where id = ?", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj2, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, bfqx, sfbf, bftime, bz, id)
			// } else {
			// 	if bfqx == "" {
			// 		db.MustExec("insert into gdzc(zcbm, place, zclx, pinpai, xinghao, pzmx, sn, buytime, buyqd, cfdd, zcyz, syzt, whzt, zsqk, bfqx,llbfrq,bfsjc,lubf,sfbf,bftime,bz) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, "--", "--", "--", "--", "--", "--", bz)
			// 	} else {
			// 		buytimeT, _ := time.Parse("2006-01-02", gzsj)
			// 		bfqxT, _ := strconv.Atoi(bfqx)
			// 		llbfrq := buytimeT.AddDate(bfqxT, 0, 0).Format("2006-01-02")
			// 		llbfrqT, _ := time.Parse("2006-01-02", llbfrq)
			// 		bfsjcT := (llbfrqT.Unix() - time.Now().Unix()) / 60 / 60 / 24
			// 		bfqxS := bfqx + "年"
			// 		bfsjcS := strconv.FormatInt(bfsjcT, 10) + "天"

			// 		db.MustExec("insert into gdzc(zcbm, place, zclx, pinpai, xinghao, pzmx, sn, buytime, buyqd, cfdd, zcyz, syzt, whzt, zsqk, bfqx,llbfrq,bfsjc,lubf,sfbf,bftime,bz) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, bfqxS, llbfrq, bfsjcS, llbfrq, "--", "--", bz)

			// 		fmt.Println("购买日期", gzsj)
			// 		fmt.Println("理论到期时间", llbfrq)
			// 		fmt.Println("报废时间差", bfsjcT, "天")

		}
		ctx.Redirect(zxsZ, iris.StatusTemporaryRedirect)
		ctx.View("gdzc_update.html")
		// var page float64
		// db.Get(&page, "SELECT count(*) FROM gdzc ")
		// xsys := 15.0
		// zxsys := math.Ceil(page / xsys)
		// zxsS := strconv.FormatFloat(zxsys, 'f', -1, 64)
		// zxsZ := "/gdzc_list/?qunaye=" + zxsS
		// fmt.Println("去哪页", zxsZ)

		// ctx.Redirect(zxsZ, iris.StatusTemporaryRedirect)
	})
	app.Get("/update_gdzc", func(ctx iris.Context) {
		updategdzcid := ctx.FormValue("updateid_gdzc")
		updategdzcsql := []Gdzc{}
		db.Select(&updategdzcsql, "select * from gdzc where id = ?", updategdzcid)
		ctx.ViewData("updategdzcsql", updategdzcsql)
		ctx.View("gdzc_update.html")
	})
	app.Get("/question_moban_update", func(ctx iris.Context) {
		// s := sess.Start(ctx).GetString("name")
		// if s != "" {
		Updateid := ctx.FormValue("updateid")
		updatesql := []Gdzc{}
		db.Select(&updatesql, "select * from moban where id = ?", Updateid)
		ctx.ViewData("question_moban_update", updatesql)
		ctx.ViewData("sname", sess.Start(ctx).GetString("name"))
		wtnrg := ctx.FormValue("wtnrg")
		wtddg := ctx.FormValue("wtddg")
		wtysg := ctx.FormValue("wtysg")
		fscsg := ctx.FormValue("fscsg")

		if wtnrg == "" {
			fmt.Println("空", Updateid)
		} else {
			qmid := ctx.FormValue("qmid")
			fmt.Println("g:", qmid, wtnrg, wtddg, wtysg, fscsg)
			db.MustExec("update moban set question_list=?,place_list=?,questiontype_list=?,count_list=? where id =?", wtnrg, wtddg, wtysg, fscsg, qmid)
			ctx.Redirect("/question_moban/", iris.StatusTemporaryRedirect)
		}

		ctx.View("question_moban_update.html")

		// }

	})

	app.Post("/insert_gdzc", func(ctx iris.Context) {
		// s := sess.Start(ctx).GetString("name")
		// if s != "" {
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
			db.MustExec("insert into gdzc(zcbm, place, zclx, pinpai, xinghao, pzmx, sn, buytime, buyqd, cfdd, zcyz, syzt, whzt, zsqk, bfqx,llbfrq,bfsjc,lubf,sfbf,bftime,bz) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, "--", gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, "--", "--", "--", "--", "--", "--", bz)
		} else {
			if bfqx == "" {
				db.MustExec("insert into gdzc(zcbm, place, zclx, pinpai, xinghao, pzmx, sn, buytime, buyqd, cfdd, zcyz, syzt, whzt, zsqk, bfqx,llbfrq,bfsjc,lubf,sfbf,bftime,bz) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, "--", "--", "--", "--", "--", "--", bz)
			} else {
				buytimeT, _ := time.Parse("2006-01-02", gzsj)
				bfqxT, _ := strconv.Atoi(bfqx)
				llbfrq := buytimeT.AddDate(bfqxT, 0, 0).Format("2006-01-02")
				llbfrqT, _ := time.Parse("2006-01-02", llbfrq)
				bfsjcT := (llbfrqT.Unix() - time.Now().Unix()) / 60 / 60 / 24
				bfqxS := bfqx + "年"
				bfsjcS := strconv.FormatInt(bfsjcT, 10) + "天"

				db.MustExec("insert into gdzc(zcbm, place, zclx, pinpai, xinghao, pzmx, sn, buytime, buyqd, cfdd, zcyz, syzt, whzt, zsqk, bfqx,llbfrq,bfsjc,lubf,sfbf,bftime,bz) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", zcbm, place, zclx, pinpai, ggxh, pzmx, sn, gzsj, gmqd, cfdd, zcyz, sfsy, wxzt, zsqk, bfqxS, llbfrq, bfsjcS, llbfrq, "--", "--", bz)

				fmt.Println("购买日期", gzsj)
				fmt.Println("理论到期时间", llbfrq)
				fmt.Println("报废时间差", bfsjcT, "天")

			}
		}
		var page float64
		db.Get(&page, "SELECT count(*) FROM gdzc ")
		xsys := 15.0
		zxsys := math.Ceil(page / xsys)
		zxsS := strconv.FormatFloat(zxsys, 'f', -1, 64)
		zxsZ := "/gdzc_list/?qunaye=" + zxsS
		fmt.Println("去哪页", zxsZ)

		ctx.Redirect(zxsZ, iris.StatusTemporaryRedirect)

	})
	app.Post("/insert_moban", func(ctx iris.Context) {
		// s := sess.Start(ctx).GetString("name")
		// if s != "" {
		qtlbmb := ctx.FormValue("qtlbmb")
		qtddmb := ctx.FormValue("qtddmb")
		wtlx := ctx.FormValue("wtlx")
		fscs := ctx.FormValue("fscs")
		// fmt.Println(qtlbmb, qtddmb, wtlx, fscs)
		ctx.ViewData("sname", sess.Start(ctx).GetString("name"))
		db.MustExec("insert into moban(question_list,place_list,questiontype_list,count_list) values(?,?,?,?)", qtlbmb, qtddmb, wtlx, fscs)
		ctx.Redirect("/question_moban/", iris.StatusTemporaryRedirect)
		// }

	})
	app.Get("/question_moban", func(ctx iris.Context) {
		// s := sess.Start(ctx).GetString("name")
		// if s != "" {
		var pagem float64
		db.Get(&pagem, "SELECT count(*) FROM moban ")
		xsysm := 10.0
		zxsysm := math.Ceil(pagem / xsysm)
		zxsZm := int(zxsysm)
		var aaam []int
		for j := 0; j < zxsZm; j++ {
			aaam = append(aaam, j+1)

		}
		ctx.ViewData("snamem", sess.Start(ctx).GetString("name"))
		ctx.ViewData("sqlnumm", aaam)
		//ctx.FormValue() must don't be null
		qunayeenm := ctx.URLParams()["qunayem"]

		if qunayeenm == "" {
			qunayeenm = "1"
		}
		//string to int
		qunayeZm, _ := strconv.Atoi(qunayeenm)
		qunayeSm := qunayeZm*10 - 10

		questionxianshi := []Gdzc{}
		db.Select(&questionxianshi, "select * from moban limit ?,?", qunayeSm, 10)
		// fmt.Println("moban:", questionxianshi)
		ctx.ViewData("question_xianshi", questionxianshi)
		ctx.ViewData("sname", sess.Start(ctx).GetString("name"))
		ctx.View("question_moban.html")
		// }

	})

	app.Post("/login", func(ctx iris.Context) {
		user := ctx.FormValue("lguser")
		passwd := ctx.FormValue("lgpasswd")
		fmt.Println("login:", user, passwd)
		var loginc int
		db.Get(&loginc, "select count(*) from user where name = ? and passwd = ?", user, passwd)
		fmt.Println("检测用户是否存在", loginc)
		if loginc == 1 {
			s := sess.Start(ctx)
			s.Set("name", user)
			ctx.Redirect("/gdzc_list/?qunaye=1", iris.StatusTemporaryRedirect)
		} else {
			ctx.HTML("用户不存在")
		}

	})
	app.Get("/search", func(ctx iris.Context) {
		// s := sess.Start(ctx).GetString("name")
		// if s != "" {
		search := []Gdzc{}

		searchnr := ctx.URLParams()["search_nr"]
		starttimes := ctx.URLParams()["starttimes"]
		endtimes := ctx.URLParams()["endtimes"]
		fmt.Println("shijians:", starttimes, endtimes)
		if starttimes == "" || endtimes == "" {
			db.Select(&search, "select * from txlyd where question like  '%"+searchnr+"%'")
		} else {
			if searchnr == "" {
				db.Select(&search, "select * from txlyd where starttime >= ? and endtime <= ?", starttimes, endtimes)
			} else {
				db.Select(&search, "select * from txlyd where starttime >= ? and endtime <= ? and question like  '%"+searchnr+"%'", starttimes, endtimes)
			}
		}

		ctx.ViewData("search_list", search)
		ctx.ViewData("sname", sess.Start(ctx).GetString("name"))
		ctx.View("search.html")
		// }

	})
	app.Get("/logout", func(ctx iris.Context) {
		s := sess.Start(ctx).GetString("name")
		if s != "" {
			sess.Start(ctx).Delete("name")
		}
		ctx.Redirect("/static/login.html", iris.StatusTemporaryRedirect)
	})

	app.Get("/", func(ctx iris.Context) {
		// s := sess.Start(ctx).GetString("name")
		// if s != "" {
		// ctx.Redirect("question_list", iris.StatusTemporaryRedirect)
		// } else {
		ctx.Redirect("/static/login.html", iris.StatusTemporaryRedirect)
		// }
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

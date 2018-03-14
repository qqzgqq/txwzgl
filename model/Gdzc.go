package model

// Gdzc mysql_固定资产
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

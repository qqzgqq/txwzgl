package model

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

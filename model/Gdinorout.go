package model

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

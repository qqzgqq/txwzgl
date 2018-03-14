package model

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

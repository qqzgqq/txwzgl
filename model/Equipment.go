package model

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

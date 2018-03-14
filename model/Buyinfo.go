package model

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

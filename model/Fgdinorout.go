package model

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

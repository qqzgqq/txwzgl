package model

// User mysql表user
type User struct {
	IDUser int    `db:"id"`
	Name   string `db:"name"`
	Passwd string `db:"passwd"`
}

package model

// People mysql表people
type People struct {
	IDPeople  int    `db:"id"`
	Name      string `db:"name"`
	Bumen     string `db:"bumen"`
	Telephone string `db:"telephone"`
}

package model

// People mysql表people
type People struct {
	IDPeople  int    `db:"id"`
	Name      string `db:"name"`
	Position  string `db:"position"`
	Bumen     string `db:"bumen"`
	Telephone string `db:"telephone"`
}

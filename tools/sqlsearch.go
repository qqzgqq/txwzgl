package tools

import (
	"math"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetSqlxsys(db *sqlx.DB) int {
	var page float64
	db.Get(&page, "SELECT count(*) FROM gdzc ")
	xsys := 15.0
	zxsys := math.Ceil(page / xsys)
	zxsZ := int(zxsys)
	return zxsZ
}

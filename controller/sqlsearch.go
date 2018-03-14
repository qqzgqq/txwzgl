package controller

import (
	"math"
	"txwzgl/tools"
)

// GetSqlxsys 获取固定资产显示信息页数
func GetSqlxsys() int {
	var page float64
	tools.DB.Get(&page, "SELECT count(*) FROM gdzc ")
	xsys := 15.0
	zxsys := math.Ceil(page / xsys)
	zxsZ := int(zxsys)
	return zxsZ
}

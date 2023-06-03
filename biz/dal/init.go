package dal

import (
	"github.com/lyleshaw/open-plugin/biz/dal/mysql"
	"github.com/lyleshaw/open-plugin/biz/model/query"
)

func Init() {
	mysql.Init()
	query.SetDefault(mysql.DB)
}

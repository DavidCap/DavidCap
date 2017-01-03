package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	_DATA_BASE_TYPE  = "mysql"
	_DATA_BASE_LOGIN = "root:123456@/DavidCap?charset=utf8"
)

func init() {
	orm.RegisterModel(new(Blog))
	orm.RegisterDriver(_DATA_BASE_TYPE, orm.DRMySQL)
	orm.RegisterDataBase("default", _DATA_BASE_TYPE, _DATA_BASE_LOGIN)
	orm.RunSyncdb("default", false, true)
}

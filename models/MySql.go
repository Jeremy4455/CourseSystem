package models

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func RegisterMySql() {
	// 注册数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	mysqluser, _ := beego.AppConfig.String("mysql::userName")
	mysqlpass, _ := beego.AppConfig.String("mysql::password")
	mysqladdr, _ := beego.AppConfig.String("mysql::addr")
	mysqldb, _ := beego.AppConfig.String("mysql::databaseName")

	// 设置数据库连接
	orm.RegisterDataBase("default", "mysql", mysqluser+":"+mysqlpass+"@tcp("+mysqladdr+")/"+mysqldb+"?charset=utf8")
	orm.RegisterModel(new(User), new(Teacher), new(Student), new(Course), new(Admin), new(Class), new(ClassStudent))
}

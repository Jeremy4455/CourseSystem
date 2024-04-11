package models

import (
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func RegisterDB() {
	// 注册数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	mysqluser, _ := beego.AppConfig.String("mysql::userName")
	mysqlpass, _ := beego.AppConfig.String("mysql::password")
	mysqlurls, _ := beego.AppConfig.String("mysql::urls")
	mysqldb, _ := beego.AppConfig.String("mysql::databaseName")

	// 设置数据库连接
	orm.RegisterDataBase("default", "mysql", mysqluser+":"+mysqlpass+"@tcp("+mysqlurls+")/"+mysqldb+"?charset=utf8")
}
func init() {
	orm.RegisterModel(new(Teacher), new(Student), new(Course), new(Admin), new(User), new(Class))
	RegisterDB()
}

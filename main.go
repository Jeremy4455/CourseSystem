package main

import (
	_ "CourseSystem/models"
	_ "CourseSystem/routers"

	"github.com/beego/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RunSyncdb("default", false, true)
	beego.Run()
}

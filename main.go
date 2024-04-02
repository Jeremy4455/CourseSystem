package main

import (
	_ "CourseSystem/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

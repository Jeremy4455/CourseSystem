package routers

import (
	"CourseSystem/controllers"
	"CourseSystem/controllers/AdminControllers"
	"CourseSystem/controllers/StudentControllers"
	"CourseSystem/controllers/TeacherControllers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	// 在 Beego 中注册过滤器
	var FilterUser = func(ctx *context.Context) {
		if ctx.Request.RequestURI != "/login" {
			ctx.Redirect(302, "/login")
		}
	}

	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.Router("/", &controllers.BaseController{})
	beego.Router("/stu/:stuid", &StudentControllers.StudentController{})
	beego.Router("/tea/:teaid", &TeacherControllers.TeacherController{})
	beego.Router("/admin", &AdminControllers.AdminController{})
	beego.Router("/login", &controllers.LoginController{})
}

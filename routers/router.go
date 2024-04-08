package routers

import (
	"CourseSystem/controllers"
	"CourseSystem/controllers/AdminControllers"
	"CourseSystem/controllers/StudentControllers"
	"CourseSystem/controllers/TeacherControllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.BaseController{})
	beego.Router("/stu/:stuid", &StudentControllers.StudentController{})
	beego.Router("/tea/:teaid", &TeacherControllers.TeacherController{})
	beego.Router("/admin", &AdminControllers.AdminController{})
}

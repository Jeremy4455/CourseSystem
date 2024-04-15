package routers

import (
	"CourseSystem/controllers"
	"CourseSystem/controllers/AdminControllers"
	AdminCourseControllers "CourseSystem/controllers/AdminControllers/AdminCourseControllers"
	AdminStudentControllers "CourseSystem/controllers/AdminControllers/AdminStudentControllers"
	AdminTeacherControllers "CourseSystem/controllers/AdminControllers/AdminTeacherControllers"
	"CourseSystem/controllers/StudentControllers"
	"CourseSystem/controllers/TeacherControllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 在 Beego 中注册过滤器
	//var FilterUser = func(ctx *context.Context) {
	//	if ctx.Request.RequestURI != "/login" {
	//		ctx.Redirect(302, "/login")
	//	}
	//}
	//
	//beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.Router("/", &controllers.BaseController{})
	beego.Router("/login", &controllers.LoginController{})

	// 管理员路由，默认管理为查询操作
	beego.Router("/admin", &AdminControllers.AdminController{})

	beego.Router("/admin/course", &AdminCourseControllers.AdminCourseControllerSearch{})
	beego.Router(`/admin/course/search`, &AdminCourseControllers.AdminCourseControllerSearch{})
	beego.Router("/admin/course/create", &AdminCourseControllers.AdminCourseControllerAdd{})
	beego.Router("/admin/course/update", &AdminCourseControllers.AdminCourseControllerSet{})
	beego.Router("/admin/course/delete", &AdminCourseControllers.AdminCourseControllerDelete{})

	beego.Router("/admin/teacher", &AdminTeacherControllers.AdminTeacherControllerSearch{})
	beego.Router("/admin/student", &AdminStudentControllers.AdminStudentControllerSearch{})

	beego.Router("/stu/:stuid", &StudentControllers.StudentController{})
	beego.Router("/tea/:teaid", &TeacherControllers.TeacherController{})
}

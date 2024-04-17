package routers

import (
	"CourseSystem/controllers"
	"CourseSystem/controllers/AdminControllers"
	AdminCourseControllers "CourseSystem/controllers/AdminControllers/AdminCourseControllers"
	AdminStudentControllers "CourseSystem/controllers/AdminControllers/AdminStudentControllers"
	AdminTeacherControllers "CourseSystem/controllers/AdminControllers/AdminTeacherControllers"
	"CourseSystem/controllers/StudentControllers"
	"CourseSystem/controllers/TeacherControllers"

	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func FilterUser(ctx *context.Context) {
	userID := ctx.Input.Session("userID")
	if userID == nil && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
		return
	}
}
func init() {
	web.InsertFilter("/*", web.BeforeRouter, FilterUser)

	beego.Router("/", &controllers.BaseController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})

	// 管理员路由，默认管理为查询操作
	beego.Router("/admin", &AdminControllers.AdminController{})

	beego.Router("/admin/course", &AdminCourseControllers.AdminCourseControllerRetrieve{})
	beego.Router(`/admin/course/retrieve`, &AdminCourseControllers.AdminCourseControllerRetrieve{})
	beego.Router("/admin/course/create", &AdminCourseControllers.AdminCourseControllerCreate{})
	beego.Router("/admin/course/update", &AdminCourseControllers.AdminCourseControllerUpdate{})
	beego.Router("/admin/course/delete", &AdminCourseControllers.AdminCourseControllerDelete{})

	beego.Router("/admin/teacher", &AdminTeacherControllers.AdminTeacherControllerRetrieve{})
	beego.Router("/admin/teacher/retrieve", &AdminTeacherControllers.AdminTeacherControllerRetrieve{})
	beego.Router("/admin/teacher/delete", &AdminTeacherControllers.AdminTeacherControllerDelete{})
	beego.Router("/admin/teacher/create", &AdminTeacherControllers.AdminTeacherControllerCreate{})
	beego.Router("/admin/teacher/update", &AdminTeacherControllers.AdminTeacherControllerUpdate{})

	beego.Router("/admin/student", &AdminStudentControllers.AdminStudentControllerRetrieve{})
	beego.Router("/admin/student/retrieve", &AdminStudentControllers.AdminStudentControllerRetrieve{})
	beego.Router("/admin/student/create", &AdminStudentControllers.AdminStudentControllerCreate{})
	beego.Router("/admin/student/update", &AdminStudentControllers.AdminStudentControllerUpdate{})
	beego.Router("/admin/student/delete", &AdminStudentControllers.AdminStudentControllerDelete{})

	beego.Router("/stu/:stuid", &StudentControllers.StudentController{})
	beego.Router("/tea/:teaid", &TeacherControllers.TeacherController{})
}

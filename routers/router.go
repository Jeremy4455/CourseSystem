package routers

import (
	"CourseSystem/controllers"
	"CourseSystem/controllers/AdminControllers"
	"CourseSystem/controllers/AdminControllers/AdminClassControllers"
	AdminCourseControllers "CourseSystem/controllers/AdminControllers/AdminCourseControllers"
	AdminStudentControllers "CourseSystem/controllers/AdminControllers/AdminStudentControllers"
	AdminTeacherControllers "CourseSystem/controllers/AdminControllers/AdminTeacherControllers"
	"CourseSystem/controllers/AdminControllers/AdminTransactionControllers"
	"CourseSystem/controllers/StudentControllers"
	"CourseSystem/controllers/StudentControllers/StudentClassControllers"
	"CourseSystem/controllers/TeacherControllers"
	"strings"

	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func FilterUser(ctx *context.Context) {
	userId := ctx.Input.Session("userId")
	if userId == nil && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
		return
	}
	if userId == nil || ctx.Request.RequestURI == "/logout" {
		return
	}

	role := userId.(string)[0]

	if role == '0' && !strings.HasPrefix(ctx.Request.RequestURI, "/admin") {
		ctx.Redirect(302, "/admin")
		return
	}
	if role == '1' && !strings.HasPrefix(ctx.Request.RequestURI, "/student") {
		ctx.Redirect(302, "/student")
		return
	}
	if role == '2' && !strings.HasPrefix(ctx.Request.RequestURI, "/teacher") {
		ctx.Redirect(302, "/teacher")
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

	beego.Router("/admin/class", &AdminClassControllers.AdminClassControllerRetrieve{})
	beego.Router("/admin/class/retrieve", &AdminClassControllers.AdminClassControllerRetrieve{})
	beego.Router("/admin/class/create", &AdminClassControllers.AdminClassControllerCreate{})
	beego.Router("/admin/class/update", &AdminClassControllers.AdminClassControllerUpdate{})
	beego.Router("/admin/class/delete", &AdminClassControllers.AdminClassControllerDelete{})

	beego.Router("/admin/trans", &AdminTransactionControllers.AdminTransactionControllerPick{})
	beego.Router("/admin/trans/drop", &AdminTransactionControllers.AdminTransactionControllerDrop{})
	beego.Router("/admin/trans/pick", &AdminTransactionControllers.AdminTransactionControllerPick{})
	beego.Router("/admin/trans/update", &AdminTransactionControllers.AdminTransactionControllerUpdate{})
	beego.Router("/admin/trans/upgrade", &AdminTransactionControllers.AdminTransactionControllerUpgrade{})
	beego.Router("/admin/trans/commit", &AdminTransactionControllers.AdminTransactionControllerCommit{})

	beego.Router("/student", &StudentControllers.StudentController{})
	beego.Router("/student/index", &StudentControllers.StudentIndexController{})
	beego.Router("/student/pick", &StudentClassControllers.StudentClassControllerPick{})
	beego.Router("/student/drop", &StudentClassControllers.StudentClassControllerDrop{})
	beego.Router("/student/grade", &StudentClassControllers.StudentClassControllerRetrieveGrade{})

	beego.Router("/teacher", &TeacherControllers.TeacherController{})
}

package AdminTeacherControllers

import "CourseSystem/controllers"

type AdminTeacherControllerAdd struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminTeacherControllerAdd) Get() {
	c.viewpath = "AdminViews/AdminStudentViews/addstudent.tpl"
	c.TplName = c.viewpath
}

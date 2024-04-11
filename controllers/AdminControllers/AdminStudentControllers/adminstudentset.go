package AdminStudentControllers

import "CourseSystem/controllers"

type AdminStudentControllerSet struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminStudentControllerSet) Get() {
	c.viewpath = "AdminViews/AdminStudentViews/addstudent.tpl"
	c.TplName = c.viewpath
}

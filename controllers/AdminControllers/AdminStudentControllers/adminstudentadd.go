package AdminStudentControllers

import "CourseSystem/controllers"

type AdminStudentControllerAdd struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminStudentControllerAdd) Get() {
	c.viewpath = "AdminViews/AdminStudentViews/addstudent.tpl"
	c.TplName = c.viewpath
}

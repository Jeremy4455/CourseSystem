package AdminStudentControllers

import "CourseSystem/controllers"

type AdminStudentControllerSearch struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminStudentControllerSearch) Get() {
	c.viewpath = "AdminViews/AdminStudentViews/addstudent.tpl"
	c.TplName = c.viewpath
}

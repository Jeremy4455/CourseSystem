package AdminStudentControllers

import "CourseSystem/controllers"

type AdminStudentControllerDelete struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminStudentControllerDelete) Get() {
	c.viewpath = "AdminViews/AdminStudentViews/addstudent.tpl"
	c.TplName = c.viewpath
}

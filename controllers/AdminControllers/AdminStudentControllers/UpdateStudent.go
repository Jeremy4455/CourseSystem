package AdminStudentControllers

import "CourseSystem/controllers"

type AdminStudentControllerUpdate struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminStudentControllerUpdate) Get() {
	c.viewpath = "AdminViews/AdminStudentViews/CreateStudent.tpl"
	c.TplName = c.viewpath
}

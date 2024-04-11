package AdminTeacherControllers

import "CourseSystem/controllers"

type AdminTeacherControllerSearch struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminTeacherControllerSearch) Get() {
	c.viewpath = "AdminViews/AdminTeacherViews/addstudent.tpl"
	c.TplName = c.viewpath
}

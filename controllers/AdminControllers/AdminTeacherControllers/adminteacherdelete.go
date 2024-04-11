package AdminTeacherControllers

import "CourseSystem/controllers"

type AdminTeacherControllerDelete struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminTeacherControllerDelete) Get() {
	c.viewpath = "AdminViews/AdminTeacherViews/addstudent.tpl"
	c.TplName = c.viewpath
}

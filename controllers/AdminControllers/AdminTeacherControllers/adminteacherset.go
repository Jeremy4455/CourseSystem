package AdminTeacherControllers

import "CourseSystem/controllers"

type AdminTeacherControllerSet struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminTeacherControllerSet) Get() {
	c.viewpath = "AdminViews/AdminTeacherViews/addstudent.tpl"
	c.TplName = c.viewpath
}

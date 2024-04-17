package AdminTeacherControllers

import "CourseSystem/controllers"

type AdminTeacherControllerUpdate struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminTeacherControllerUpdate) Get() {
	c.viewpath = "AdminViews/AdminTeacherViews/UpdateTeacher.tpl"
	c.TplName = c.viewpath
}

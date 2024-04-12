package AdminStudentControllers

import (
	"CourseSystem/controllers"
)

type AdminCourseControllerSet struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminCourseControllerSet) Get() {
	c.viewpath = "AdminViews/AdminStudentViews/addstudent.tpl"
	c.TplName = c.viewpath
}

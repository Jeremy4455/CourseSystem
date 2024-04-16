package AdminStudentControllers

import (
	"CourseSystem/controllers"
)

type AdminCourseControllerSet struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminCourseControllerSet) Get() {
	c.TplName = "AdminViews/AdminCourseViews/setcourse.tpl"
}

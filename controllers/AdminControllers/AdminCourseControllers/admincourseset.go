package AdminCourseControllers

import (
	"CourseSystem/controllers"
)

type AdminCourseControllerSet struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminCourseControllerSet) Get() {
	c.viewpath = "AdminViews/AdminCourseViews/setcourse.tpl"
	c.TplName = c.viewpath
}

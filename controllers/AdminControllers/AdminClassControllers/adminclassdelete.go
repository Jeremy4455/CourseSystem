package AdminCourseControllers

import (
	"CourseSystem/controllers"
)

type AdminClassControllerDelete struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminClassControllerDelete) Get() {
	c.viewpath = "AdminViews/AdminCourseViews/deletecourse.tpl"
	c.TplName = c.viewpath
}

func (c *AdminClassControllerDelete) Post() {
	c.TplName = c.viewpath
}

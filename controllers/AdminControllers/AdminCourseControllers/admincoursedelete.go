package AdminStudentControllers

import (
	"CourseSystem/controllers"
)

type AdminCourseControllerDelete struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminCourseControllerDelete) Get() {
	c.viewpath = "AdminViews/AdminCourseViews/deletecourse.tpl"
	c.TplName = c.viewpath
}

func (c *AdminCourseControllerDelete) Post() {
	c.TplName = c.viewpath

}

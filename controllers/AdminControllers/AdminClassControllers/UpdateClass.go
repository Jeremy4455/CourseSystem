package AdminCourseControllers

import (
	"CourseSystem/controllers"
)

type AdminClassControllerUpdate struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminClassControllerUpdate) Get() {
	c.viewpath = "AdminViews/AdminCourseViews/UpdateCourse.tpl"
	c.TplName = c.viewpath
}
func (c *AdminClassControllerUpdate) Post() {

}

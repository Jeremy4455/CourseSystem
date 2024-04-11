package AdminCourseControllers

import (
	"CourseSystem/controllers"
)

type AdminClassControllerSet struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminClassControllerSet) Get() {
	c.viewpath = "AdminViews/AdminCourseViews/setcourse.tpl"
	c.TplName = c.viewpath
}
func (c *AdminClassControllerSet) Post() {

}

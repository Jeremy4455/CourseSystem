package AdminStudentControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminCourseControllerAdd struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminCourseControllerAdd) Get() {
	c.viewpath = "AdminViews/AdminCourseViews/addcourse.tpl"
	c.TplName = c.viewpath
}

func (c *AdminCourseControllerAdd) Post() {
	c.TplName = c.viewpath
	courseCode := c.GetString("CourseCode")
	name := c.GetString("Name")
	college := c.GetString("College")
	credit := c.GetString("Credit")

	err := models.AddCourse(courseCode, name, college, credit)
	if err != nil {
		return
	}
}

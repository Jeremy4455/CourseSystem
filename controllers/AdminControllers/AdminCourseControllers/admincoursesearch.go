package AdminStudentControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminCourseControllerSearch struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminCourseControllerSearch) Get() {
	c.viewpath = "AdminViews/AdminStudentViews/addstudent.tpl"
	c.TplName = c.viewpath
}

func (c *AdminCourseControllerSearch) Post() {
	c.TplName = c.viewpath
	courseCode := c.GetString("CourseCode")
	name := c.GetString("Name")
	course, _ := models.GetCourse(courseCode, name)
	if course == nil {
		return
	}
	c.Data["Course"] = course
}

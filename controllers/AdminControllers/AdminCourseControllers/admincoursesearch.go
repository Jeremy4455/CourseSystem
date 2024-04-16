package AdminStudentControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminCourseControllerSearch struct {
	controllers.BaseController
}

func (c *AdminCourseControllerSearch) Get() {
	c.TplName = "AdminViews/AdminCourseViews/searchcourse.tpl"

}

func (c *AdminCourseControllerSearch) Post() {
	c.TplName = "AdminViews/AdminCourseViews/searchcourse.tpl"
	courseCode := c.GetString("CourseCode")
	name := c.GetString("Name")
	courses, _ := models.GetCourse(courseCode, name)
	if courses == nil {
		return
	}
	c.Data["Courses"] = courses
}

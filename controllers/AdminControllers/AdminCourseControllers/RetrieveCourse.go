package AdminCourseControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminCourseControllerRetrieve struct {
	controllers.BaseController
}

func (c *AdminCourseControllerRetrieve) Get() {
	c.TplName = "AdminViews/AdminCourseViews/RetrieveCourse.tpl"

}

func (c *AdminCourseControllerRetrieve) Post() {
	c.TplName = "AdminViews/AdminCourseViews/RetrieveCourse.tpl"
	courseCode := c.GetString("CourseCode")
	name := c.GetString("Name")
	courses, _ := models.GetCourses(courseCode, name)
	if courses == nil {
		return
	}
	c.Data["Courses"] = courses
}

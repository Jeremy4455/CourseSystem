package AdminCourseControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminCourseControllerDelete struct {
	controllers.BaseController
}

func (c *AdminCourseControllerDelete) Get() {
	c.TplName = "AdminViews/AdminCourseViews/DeleteCourse.tpl"
	c.searchCourse()
}

func (c *AdminCourseControllerDelete) searchCourse() {
	courseCode := c.GetString("CourseCode")
	name := c.GetString("Name")

	courses, err := models.GetCourses(courseCode, name)
	if err != nil {
		c.Err(err)
		return
	}
	c.Data["Courses"] = courses
	c.Sucess()
}

func (c *AdminCourseControllerDelete) Post() {
	c.TplName = "AdminViews/AdminCourseViews/DeleteCourse.tpl"
	courseCode := c.GetString("CourseCode")

	err := models.DeleteCourse(courseCode)
	if err != nil {
		c.Err(err)
		return
	}
	c.Sucess()
}

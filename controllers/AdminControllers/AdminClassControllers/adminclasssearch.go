package AdminCourseControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminClassControllerSearch struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminClassControllerSearch) Get() {
	c.viewpath = "AdminViews/AdminCourseViews/searchcourse.tpl"
	c.TplName = c.viewpath
}
func (c *AdminClassControllerSearch) Post() {
	c.TplName = c.viewpath

	courseCode := c.GetString("courseCode")
	courseName := c.GetString("courseName")
	courseTeacherId := c.GetString("courseTeacherId")
	courseSemester := c.GetString("courseSemester")
	courseTime := c.GetString("courseTime")
	classroom := c.GetString("classroom")

	classes, _ := models.SearchClasses(courseCode, courseName, courseTeacherId, courseSemester, courseTime, classroom)
	c.Data["Classes"] = classes
}

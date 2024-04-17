package AdminCourseControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminClassControllerRetrieve struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminClassControllerRetrieve) Get() {
	c.viewpath = "AdminViews/AdminCourseViews/RetrieveCourse.tpl"
	c.TplName = c.viewpath
}
func (c *AdminClassControllerRetrieve) Post() {
	c.TplName = c.viewpath

	courseCode := c.GetString("courseCode")
	courseName := c.GetString("courseName")
	courseTeacherId := c.GetString("courseTeacherId")
	courseSemester := c.GetString("courseSemester")
	courseTime := c.GetString("courseTime")
	classroom := c.GetString("classroom")

	classes, _ := models.GetClasses(courseCode, courseName, courseTeacherId, courseSemester, courseTime, classroom)
	c.Data["Classes"] = classes
}

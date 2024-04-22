package AdminClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminClassControllerRetrieve struct {
	controllers.BaseController
}

func (c *AdminClassControllerRetrieve) Get() {
	c.TplName = "AdminViews/AdminClassViews/RetrieveClass.tpl"
}
func (c *AdminClassControllerRetrieve) Post() {
	c.TplName = "AdminViews/AdminClassViews/RetrieveClass.tpl"
	courseCode := c.GetString("courseCode")
	courseName := c.GetString("courseName")
	courseTeacherId := c.GetString("courseTeacherId")
	courseTeacherName := c.GetString("CourseTeacherName")
	courseSemester := c.GetString("courseSemester")
	courseTime := c.GetString("courseTime")
	classroom := c.GetString("classroom")

	classes, _ := models.GetClasses(courseCode, courseName, courseTeacherId, courseTeacherName, courseSemester, courseTime, classroom)
	c.Data["Classes"] = classes
}

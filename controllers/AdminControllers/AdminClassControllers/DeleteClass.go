package AdminClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminClassControllerDelete struct {
	controllers.BaseController
}

func (c *AdminClassControllerDelete) Get() {
	c.SearchClasses()
	c.TplName = "AdminViews/AdminClassViews/DeleteClass.tpl"
}

func (c *AdminClassControllerDelete) SearchClasses() {
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

func (c *AdminClassControllerDelete) Post() {
	c.TplName = "AdminViews/AdminClassViews/DeleteClass.tpl"

	courseCode := c.GetString("CourseCode")
	courseTeacherId := c.GetString("CourseTeacherId")
	courseSemester := c.GetString("CourseSemester")
	err := models.DeleteClass(courseCode, courseTeacherId, courseSemester)
	if err == false {
		return
	}
}

package AdminCourseControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminClassControllerUpdate struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminClassControllerUpdate) Get() {
	c.TplName = "AdminViews/AdminCourseViews/UpdateCourse.tpl"
}
func (c *AdminClassControllerUpdate) Post() {
	c.TplName = "AdminViews/AdminCourseViews/UpdateCourse.tpl"

	courseCode := c.GetString("courseCode")
	courseName := c.GetString("courseName")
	courseTeacherId := c.GetString("courseTeacherId")
	courseTeacherName := c.GetString("CourseTeacherName")
	courseSemester := c.GetString("courseSemester")
	classTime := c.GetString("courseTime")
	capacity := c.GetString("Capacity")
	classroom := c.GetString("classroom")

	class, _ := models.GetClasses(courseCode, courseName, courseTeacherId, courseTeacherName, courseSemester, classTime, classroom)
	if len(class) != 1 {
		return
	}
	err := models.ReviseClass(class[0], courseTeacherId, classTime, capacity, classroom)
	if err == false {
		return
	}
}

package AdminCourseControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminClassControllerCreate struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminClassControllerCreate) Get() {
	c.viewpath = "AdminViews/AdminCourseViews/CreateCourse.tpl"
	c.TplName = c.viewpath
}

func (c *AdminClassControllerCreate) Post() {
	c.TplName = c.viewpath

	courseCode := c.GetString("CourseCode")
	courseName := c.GetString("CourseName")
	courseTeacherId := c.GetString("CourseTeacherId")
	courseSemester := c.GetString("CourseSemester")
	classTime := c.GetString("ClassTime")
	capacity := c.GetString("Capacity")
	classroom := c.GetString("Classroom")

	err := models.AddClass(courseCode, courseName, courseTeacherId, courseSemester, classTime, capacity, classroom)
	if err == false {
		return
	}
}

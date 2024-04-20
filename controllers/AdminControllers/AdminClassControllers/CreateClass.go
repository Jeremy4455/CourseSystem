package AdminCourseControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminClassControllerCreate struct {
	controllers.BaseController
}

func (c *AdminClassControllerCreate) Get() {
	c.TplName = "AdminViews/AdminCourseViews/CreateCourse.tpl"
}

func (c *AdminClassControllerCreate) Post() {
	c.TplName = "AdminViews/AdminCourseViews/CreateCourse.tpl"

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

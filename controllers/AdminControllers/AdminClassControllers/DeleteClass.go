package AdminClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminClassControllerDelete struct {
	controllers.BaseController
}

func (c *AdminClassControllerDelete) Get() {
	c.TplName = "AdminViews/AdminCourseViews/DeleteCourse.tpl"
}

func (c *AdminClassControllerDelete) Post() {
	c.TplName = "AdminViews/AdminCourseViews/DeleteCourse.tpl"

	courseCode := c.GetString("CourseCode")
	courseTeacherId := c.GetString("CourseTeacherId")
	courseSemester := c.GetString("CourseSemester")
	err := models.DeleteClass(courseCode, courseTeacherId, courseSemester)
	if err == false {
		return
	}
}

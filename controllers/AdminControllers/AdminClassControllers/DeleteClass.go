package AdminClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminClassControllerDelete struct {
	controllers.BaseController
}

func (c *AdminClassControllerDelete) Get() {
	c.TplName = "AdminViews/AdminClassViews/DeleteClass.tpl"
}

func (c *AdminClassControllerDelete) Post() {
	c.TplName = "AdminViews/AdminClassViews/DeleteClass.tpl"
	courseCode := c.GetString("CourseCode")
	courseTeacherId := c.GetString("CourseTeacherId")
	courseSemester := c.GetString("CourseSemester")

	err := models.DeleteClass(courseCode, courseTeacherId, courseSemester)
	if err != nil {
		return
	}
}

package AdminTeacherControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminTeacherControllerCreate struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminTeacherControllerCreate) Get() {
	c.TplName = "AdminViews/AdminTeacherViews/CreateTeacher.tpl"
}

func (c *AdminTeacherControllerCreate) Post() {
	c.TplName = "AdminViews/AdminTeacherViews/CreateTeacher.tpl"
	teacherId := c.GetString("TeacherId")
	name := c.GetString("Name")
	mobile := c.GetString("Mobile")
	email := c.GetString("Email")

	err := models.AddTeacher(teacherId, name, mobile, email)
	if err != nil {
		return
	}
}

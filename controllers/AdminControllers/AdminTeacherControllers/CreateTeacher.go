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
	c.viewpath = "AdminViews/AdminTeacherViews/CreateTeacher.tpl"
	c.TplName = c.viewpath
}
func (c *AdminTeacherControllerCreate) Post() {
	c.TplName = c.viewpath
	teacherId := c.GetString("TeacherId")
	name := c.GetString("Name")
	mobile := c.GetString("Mobile")
	email := c.GetString("Email")

	err := models.AddTeacher(teacherId, name, mobile, email)
	if err != nil {
		return
	}
}

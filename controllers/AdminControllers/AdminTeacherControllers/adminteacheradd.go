package AdminTeacherControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminTeacherControllerAdd struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminTeacherControllerAdd) Get() {
	c.viewpath = "AdminViews/AdminTeacherViews/addteacher.tpl"
	c.TplName = c.viewpath
}
func (c *AdminTeacherControllerAdd) Post() {
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

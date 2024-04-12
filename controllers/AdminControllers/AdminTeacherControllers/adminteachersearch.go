package AdminTeacherControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminTeacherControllerSearch struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminTeacherControllerSearch) Get() {
	c.viewpath = "AdminViews/AdminTeacherViews/addstudent.tpl"
	c.TplName = c.viewpath
}
func (c *AdminTeacherControllerSearch) Post() {
	c.TplName = c.viewpath

	teacherId := c.GetString("TeacherId")
	teacher, _ := models.GetTeacher(teacherId)
	if teacher != nil {
		return
	}
	c.Data["Teacher"] = teacher
}

package AdminTeacherControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminTeacherControllerUpdate struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminTeacherControllerUpdate) Get() {
	c.TplName = "AdminViews/AdminTeacherViews/UpdateTeacher.tpl"
}

func (c *AdminTeacherControllerUpdate) Post() {
	c.TplName = c.viewpath

	teacherId := c.GetString("TeacherId")
	teacherName := c.GetString("TeacherName")
	mobile := c.GetString("Mobile")
	email := c.GetString("Email")
	teacher, _ := models.GetTeacher(teacherId, teacherName, mobile, email)
	if teacher != nil {
		return
	}
	c.Data["Teacher"] = teacher
}

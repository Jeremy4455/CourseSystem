package AdminTeacherControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminTeacherControllerDelete struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminTeacherControllerDelete) Get() {
	c.TplName = "AdminViews/AdminTeacherViews/DeleteTeacher.tpl"
	teacherId := c.GetString("TeacherId")
	teacherName := c.GetString("Name")
	mobile := c.GetString("Mobile")
	email := c.GetString("Email")
	teachers, _ := models.GetTeachers(teacherId, teacherName, mobile, email)
	c.Data["Teachers"] = teachers
}

func (c *AdminTeacherControllerDelete) Post() {
	c.TplName = "AdminViews/AdminTeacherViews/DeleteTeacher.tpl"
	teacherId := c.GetString("TeacherId")
	err := models.DeleteTeacher(teacherId)
	if err != nil {
		return
	}
}

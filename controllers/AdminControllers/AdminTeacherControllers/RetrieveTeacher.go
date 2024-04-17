package AdminTeacherControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminTeacherControllerRetrieve struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminTeacherControllerRetrieve) Get() {
	c.TplName = "AdminViews/AdminTeacherViews/RetrieveTeacher.tpl"
}
func (c *AdminTeacherControllerRetrieve) Post() {
	c.TplName = "AdminViews/AdminTeacherViews/RetrieveTeacher.tpl"

	teacherId := c.GetString("TeacherId")
	teacherName := c.GetString("TeacherName")
	mobile := c.GetString("Mobile")
	email := c.GetString("Email")
	teacher, _ := models.GetTeacher(teacherId, "", "", "")
	if len(teacher) != 1 {
		return
	}
	err := models.ReviseTeacher(teacher[0], teacherName, mobile, email)
	if err == false {
		return
	}
}

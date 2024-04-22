package AdminTeacherControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminTeacherControllerRetrieve struct {
	controllers.BaseController
}

func (c *AdminTeacherControllerRetrieve) Get() {
	c.TplName = "AdminViews/AdminTeacherViews/RetrieveTeacher.tpl"
}
func (c *AdminTeacherControllerRetrieve) Post() {
	c.TplName = "AdminViews/AdminTeacherViews/RetrieveTeacher.tpl"
	teacherId := c.GetString("TeacherId")
	teacherName := c.GetString("Name")
	mobile := c.GetString("Mobile")
	email := c.GetString("Email")

	teachers, _ := models.GetTeachers(teacherId, teacherName, mobile, email)
	c.Data["Teachers"] = teachers
}

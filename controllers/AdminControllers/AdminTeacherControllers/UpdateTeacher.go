package AdminTeacherControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminTeacherControllerUpdate struct {
	controllers.BaseController
}

func (c *AdminTeacherControllerUpdate) Get() {
	c.TplName = "AdminViews/AdminTeacherViews/UpdateTeacher.tpl"
	teacherId := c.GetString("TeacherId")
	c.searchTeacher(teacherId)
}

func (c *AdminTeacherControllerUpdate) searchTeacher(teacherId string) {
	if teacherId == "" {
		teachers, _ := models.GetAllTeachers()
		c.Data["Teachers"] = teachers
	} else {
		teachers, err := models.GetTeachers(teacherId, "", "", "")
		if err != nil {
			c.Err(err)
			return
		}
		c.Data["Teachers"] = teachers
	}
}

func (c *AdminTeacherControllerUpdate) Post() {
	c.TplName = "AdminViews/AdminTeacherViews/UpdateTeacher.tpl"
	teacherId := c.GetString("TeacherId")
	name := c.GetString("Name")
	mobile := c.GetString("Mobile")
	email := c.GetString("Email")

	err := models.ReviseTeacher(teacherId, name, mobile, email)
	if err != nil {
		c.Err(err)
		return
	}
	c.Sucess()
	c.searchTeacher("")
}

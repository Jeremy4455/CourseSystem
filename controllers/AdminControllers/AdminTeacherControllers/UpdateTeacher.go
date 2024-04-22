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
	c.searchTeacher()
}

func (c *AdminTeacherControllerUpdate) searchTeacher() {
	teacherId := c.GetString("TeacherId")
	if teacherId == "" {
		teachers, _ := models.GetAllTeachers()
		c.Data["Teachers"] = teachers
	} else {
		teachers, _ := models.GetTeachers(teacherId, "", "", "")
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
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"message": "教师更新成功"}
	}
}

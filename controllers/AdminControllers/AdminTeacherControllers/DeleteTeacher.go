package AdminTeacherControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminTeacherControllerDelete struct {
	controllers.BaseController
}

func (c *AdminTeacherControllerDelete) Get() {
	c.TplName = "AdminViews/AdminTeacherViews/DeleteTeacher.tpl"
	c.searchTeacher()
}

func (c *AdminTeacherControllerDelete) searchTeacher() {
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
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"message": "教师删除成功"}
	}
}

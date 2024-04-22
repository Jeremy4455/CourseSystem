package AdminTeacherControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminTeacherControllerCreate struct {
	controllers.BaseController
}

func (c *AdminTeacherControllerCreate) Get() {
	c.TplName = "AdminViews/AdminTeacherViews/CreateTeacher.tpl"
}

func (c *AdminTeacherControllerCreate) Post() {
	c.TplName = "AdminViews/AdminTeacherViews/CreateTeacher.tpl"
	teacherId := c.GetString("TeacherId")
	name := c.GetString("Name")
	mobile := c.GetString("Mobile")
	email := c.GetString("Email")

	err := models.CreateTeacher(teacherId, name, mobile, email)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"message": "教师添加成功"}
	}
}

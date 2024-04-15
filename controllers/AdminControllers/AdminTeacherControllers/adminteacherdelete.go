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
	c.viewpath = "AdminViews/AdminTeacherViews/deleteteacher.tpl"
	c.TplName = c.viewpath
}

func (c *AdminTeacherControllerDelete) Post() {
	c.TplName = c.viewpath

	teacherId := c.GetString("TeacherId")
	err := models.DeleteTeacher(teacherId)
	if err != nil {
		return
	}
}

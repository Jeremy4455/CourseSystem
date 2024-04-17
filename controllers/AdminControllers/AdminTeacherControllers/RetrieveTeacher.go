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
	c.viewpath = "AdminViews/AdminTeacherViews/RetrieveTeacher.tpl"
	c.TplName = c.viewpath
}
func (c *AdminTeacherControllerRetrieve) Post() {
	c.TplName = c.viewpath

	teacherId := c.GetString("TeacherId")
	teacher, _ := models.GetTeacher(teacherId)
	if teacher != nil {
		return
	}
	c.Data["Teacher"] = teacher
}

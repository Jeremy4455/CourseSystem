package AdminStudentControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminStudentControllerSearch struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminStudentControllerSearch) Get() {
	c.viewpath = "AdminViews/AdminStudentViews/searchstudent.tpl"
	c.TplName = c.viewpath
}

func (c *AdminStudentControllerSearch) Post() {
	c.TplName = c.viewpath

	studentId := c.GetString("StudentId")
	student, _ := models.GetStudent(studentId)
	if student != nil {
		return
	}
	c.Data["Student"] = student
}

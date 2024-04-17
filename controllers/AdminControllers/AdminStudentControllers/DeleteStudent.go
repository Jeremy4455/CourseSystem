package AdminStudentControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminStudentControllerDelete struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminStudentControllerDelete) Get() {
	c.viewpath = "AdminViews/AdminStudentViews/CreateStudent.tpl"
	c.TplName = c.viewpath
}

func (c *AdminStudentControllerDelete) Post() {
	c.TplName = c.viewpath
	studentId := c.GetString("StudentId")

	err := models.DeleteStudent(studentId)
	if err != nil {
		return
	}
}

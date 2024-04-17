package AdminStudentControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminStudentControllerRetrieve struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminStudentControllerRetrieve) Get() {
	c.viewpath = "AdminViews/AdminStudentViews/RetrieveStudent.tpl"
	c.TplName = c.viewpath
}

func (c *AdminStudentControllerRetrieve) Post() {
	c.TplName = c.viewpath

	studentId := c.GetString("StudentId")
	student, _ := models.GetStudent(studentId)
	if student != nil {
		return
	}
	c.Data["Student"] = student
}

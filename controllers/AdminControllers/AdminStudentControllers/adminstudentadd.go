package AdminStudentControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminStudentControllerAdd struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminStudentControllerAdd) Get() {
	c.viewpath = "AdminViews/AdminStudentViews/addstudent.tpl"
	c.TplName = c.viewpath
}

func (c *AdminStudentControllerAdd) Post() {
	c.TplName = c.viewpath
	studentId := c.GetString("StudentId")
	name := c.GetString("Name")
	class := c.GetString("Class")
	grade := c.GetString("Grade")

	err := models.AddStudent(studentId, name, class, grade)
	if err != nil {
		return
	}
}

package AdminStudentControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminStudentControllerCreate struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminStudentControllerCreate) Get() {
	c.viewpath = "AdminViews/AdminStudentViews/CreateStudent.tpl"
	c.TplName = c.viewpath
}

func (c *AdminStudentControllerCreate) Post() {
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

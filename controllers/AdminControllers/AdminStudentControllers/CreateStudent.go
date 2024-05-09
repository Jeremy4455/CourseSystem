package AdminStudentControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminStudentControllerCreate struct {
	controllers.BaseController
}

func (c *AdminStudentControllerCreate) Get() {
	c.TplName = "AdminViews/AdminStudentViews/CreateStudent.tpl"
}

func (c *AdminStudentControllerCreate) Post() {
	c.TplName = "AdminViews/AdminStudentViews/CreateStudent.tpl"
	studentId := c.GetString("StudentId")
	name := c.GetString("Name")
	class := c.GetString("Class")
	grade := c.GetString("Grade")

	err := models.CreateStudent(studentId, name, class, grade)
	if err != nil {
		c.Err(err)
	} else {
		c.Sucess()
	}
}

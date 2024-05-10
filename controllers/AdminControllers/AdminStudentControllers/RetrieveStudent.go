package AdminStudentControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminStudentControllerRetrieve struct {
	controllers.BaseController
}

func (c *AdminStudentControllerRetrieve) Get() {
	c.TplName = "AdminViews/AdminStudentViews/RetrieveStudent.tpl"
}

func (c *AdminStudentControllerRetrieve) Post() {
	c.TplName = "AdminViews/AdminStudentViews/RetrieveStudent.tpl"
	studentId := c.GetString("StudentId")
	name := c.GetString("Name")
	class := c.GetString("Class")

	students, err := models.GetStudents(studentId, name, class)
	if err != nil {
		c.Err(err)
		return
	}
	c.Data["Students"] = students
}

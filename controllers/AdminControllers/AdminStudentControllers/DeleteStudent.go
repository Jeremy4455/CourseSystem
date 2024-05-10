package AdminStudentControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminStudentControllerDelete struct {
	controllers.BaseController
}

func (c *AdminStudentControllerDelete) Get() {
	c.TplName = "AdminViews/AdminStudentViews/DeleteStudent.tpl"
	c.searchStudent()
}
func (c *AdminStudentControllerDelete) searchStudent() {
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

func (c *AdminStudentControllerDelete) Post() {
	c.TplName = "AdminViews/AdminStudentViews/DeleteStudent.tpl"
	studentId := c.GetString("StudentId")

	err := models.DeleteStudent(studentId)
	if err != nil {
		c.Err(err)
	} else {
		c.Sucess()
	}
}

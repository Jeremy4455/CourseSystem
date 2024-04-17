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
	c.SearchStudent()
	c.TplName = "AdminViews/AdminStudentViews/DeleteStudent.tpl"
}
func (c *AdminStudentControllerDelete) SearchStudent() {
	studentId := c.GetString("StudentId")
	name := c.GetString("Name")
	class := c.GetString("Class")
	students, _ := models.GetStudents(studentId, name, class)
	if students == nil {
		return
	}
	c.Data["Students"] = students
}

func (c *AdminStudentControllerDelete) Post() {
	c.TplName = "AdminViews/AdminStudentViews/DeleteStudent.tpl"
	studentId := c.GetString("StudentId")

	err := models.DeleteStudent(studentId)
	if err != nil {
		return
	}
}

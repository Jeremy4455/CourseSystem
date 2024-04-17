package AdminStudentControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminStudentControllerUpdate struct {
	controllers.BaseController
}

func (c *AdminStudentControllerUpdate) Get() {
	c.TplName = "AdminViews/AdminStudentViews/UpdateStudent.tpl"
	studentId := c.GetString("StudentId")
	if studentId == "" {
		students, _ := models.GetAllStudents()
		c.Data["Students"] = students
	} else {
		students, _ := models.GetStudents(studentId, "", "")
		c.Data["Students"] = students
	}
}

func (c *AdminStudentControllerUpdate) Post() {
	c.TplName = "AdminViews/AdminStudentViews/UpdateStudent.tpl"

	studentId := c.GetString("StudentId")
	name := c.GetString("Name")
	class := c.GetString("Class")
	s, err := models.GetStudent(studentId)
	if err != nil {
		return
	}
	models.ReviseStudent(s, name, class)
}

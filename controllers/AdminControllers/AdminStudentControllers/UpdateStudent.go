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
		students, err := models.GetStudents(studentId, "", "")
		if err != nil {
			c.Err(err)
			return
		}
		c.Data["Students"] = students
	}
}

func (c *AdminStudentControllerUpdate) Post() {
	c.TplName = "AdminViews/AdminStudentViews/UpdateStudent.tpl"

	studentId := c.GetString("StudentId")
	name := c.GetString("Name")
	class := c.GetString("Class")

	err := models.ReviseStudent(studentId, name, class)
	if err != nil {
		c.Err(err)
	} else {
		c.Sucess()
		students, _ := models.GetAllStudents()
		c.Data["Students"] = students
	}
}

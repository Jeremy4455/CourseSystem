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
	students, _ := models.GetStudents(studentId, "", "")
	if students == nil {
		return
	}
	c.Data["Students"] = students
}

func (c *AdminStudentControllerUpdate) Post() {
	c.TplName = "AdminViews/AdminStudentViews/UpdateStudent.tpl"
}

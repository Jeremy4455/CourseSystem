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
	c.TplName = "AdminViews/AdminStudentViews/RetrieveStudent.tpl"
}

func (c *AdminStudentControllerRetrieve) Post() {
	c.TplName = "AdminViews/AdminStudentViews/RetrieveStudent.tpl"
	// 获取请求参数
	studentId := c.GetString("StudentId")
	name := c.GetString("Name")
	class := c.GetString("Class")
	students, _ := models.GetStudents(studentId, name, class)
	c.Data["Students"] = students
}

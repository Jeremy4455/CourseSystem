package AdminStudentControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminStudentControllerDelete struct {
	controllers.BaseController
}

func (c *AdminStudentControllerDelete) Get() {
	c.searchStudent()
	c.TplName = "AdminViews/AdminStudentViews/DeleteStudent.tpl"
}
func (c *AdminStudentControllerDelete) searchStudent() {
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
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"message": "学生删除成功"}
	}
}

package AdminTransactionControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminTransactionControllerDrop struct {
	controllers.BaseController
}

func (c *AdminTransactionControllerDrop) Get() {
	c.TplName = ""
	c.searchClassStudent()
}
func (c *AdminTransactionControllerDrop) searchClassStudent() {
	studentId := c.GetString("StudentId")
	courseCode := c.GetString("CourseCode")
	courseName := c.GetString("CourseName")
	teacherId := c.GetString("TeacherId")
	teacherName := c.GetString("CourseTeacherName")
	semester := c.GetString("Semester")

	student, err := models.GetStudent(studentId)
	if err != nil {
		return
	}
	class, err := models.GetClasses(courseCode, courseName, teacherId, teacherName, semester, "", "")
	if err != nil {
		return
	}
	if len(class) != 1 {
		return
	}
	cs, err := models.GetClassStudent(student, semester, class[0])
	if err != nil {
		return
	}

	x := make(map[string]interface{})
	x["CourseCode"] = class[0].Course.CourseCode
	x["CourseName"] = class[0].Course.Name
	x["Credit"] = class[0].Course.Credit
	x["TeacherId"] = class[0].Teacher.TeacherId
	x["TeacherName"] = class[0].Teacher.Name
	x["StudentId"] = cs[0].Student.StudentId
	x["StudentName"] = cs[0].Student.Name
	c.Data["ClassStudent"] = x
}

func (c *AdminTransactionControllerDrop) Post() {
	c.TplName = ""
	studentId := c.GetString("StudentId")
	courseCode := c.GetString("CourseCode")
	courseName := c.GetString("CourseName")
	teacherId := c.GetString("TeacherId")
	teacherName := c.GetString("CourseTeacherName")
	semester := c.GetString("Semester")

	student, err := models.GetStudent(studentId)
	if err != nil {
		return
	}
	class, err := models.GetClasses(courseCode, courseName, teacherId, teacherName, semester, "", "")
	if err != nil {
		return
	}
	models.DropClass(student, class[0], models.ADMIN_PICK_DROP_CLASS)
}

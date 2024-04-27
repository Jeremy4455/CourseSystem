package AdminTransactionControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminTransactionControllerPick struct {
	controllers.BaseController
}

func (c *AdminTransactionControllerPick) Get() {
	c.TplName = "AdminViews/AdminTransViews/PickClass.tpl"
	c.Data["Semesters"] = models.Semesters
	c.searchClass()
}
func (c *AdminTransactionControllerPick) searchClass() {
	courseCode := c.GetString("CourseCode")
	courseName := c.GetString("CourseName")
	teacherId := c.GetString("TeacherId")
	teacherName := c.GetString("CourseTeacherName")
	semester := c.GetString("Semester")

	class, err := models.GetClasses(courseCode, courseName, teacherId, teacherName, semester, "", "")
	if err != nil {
		return
	}
	x := make(map[string]interface{})
	x["CourseCode"] = class[0].Course.CourseCode
	x["CourseName"] = class[0].Course.Name
	x["Credit"] = class[0].Course.Credit
	x["TeacherId"] = class[0].Teacher.TeacherId
	x["TeacherName"] = class[0].Teacher.Name

	c.Data["ClassStudent"] = x
}

func (c *AdminTransactionControllerPick) Post() {
	c.TplName = "AdminViews/AdminTransViews/PickClass.tpl"
	c.Data["Semesters"] = models.Semesters
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
	cnt, err := models.GetPickedCount(class[0])
	if err != nil {
		return
	}
	if cnt >= class[0].Capacity {
		return
	}
	err = models.PickClass(student, class[0], models.ADMIN_PICK_DROP_CLASS)
	if err != nil {
		return
	}
}

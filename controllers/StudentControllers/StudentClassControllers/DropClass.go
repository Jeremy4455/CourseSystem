package StudentClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type StudentClassControllerDrop struct {
	controllers.BaseController
}

func (c *StudentClassControllerDrop) Get() {
	c.TplName = "StudentViews/StudentClassViews/"
	c.showClass()
}
func (c *StudentClassControllerDrop) showClass() {
	studentId := c.GetSession("userId").(string)[1:]

	student, err := models.GetStudent(studentId)
	if err != nil {
		return
	}
	var classes []*models.Class
	for _, class := range student.Classes {
		classes = append(classes, class.Class)
	}
	c.Data["Classes"] = classes
}
func (c *StudentClassControllerDrop) Post() {
	c.TplName = "StudentViews/StudentClassViews/"
	courseCode := c.GetString("courseCode")
	courseName := c.GetString("courseName")
	courseTeacherId := c.GetString("courseTeacherId")
	courseTeacherName := c.GetString("CourseTeacherName")
	courseSemester := c.GetString("courseSemester")
	courseTime := c.GetString("courseTime")
	classroom := c.GetString("classroom")

	class, _ := models.GetClasses(courseCode, courseName, courseTeacherId, courseTeacherName, courseSemester, courseTime, classroom)
	studentId := c.GetSession("userId").(string)[1:]
	student, err := models.GetStudent(studentId)
	if err != nil {
		return
	}
	models.DropClass(student, class[0])
}

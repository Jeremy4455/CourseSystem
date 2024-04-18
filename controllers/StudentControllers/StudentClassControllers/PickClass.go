package StudentClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type StudentClassControllerPick struct {
	controllers.BaseController
}

func (c *StudentClassControllerPick) Get() {
	c.TplName = "StudentViews/StudentClassViews/"
	c.searchClass()
}
func (c *StudentClassControllerPick) searchClass() {
	courseCode := c.GetString("courseCode")
	courseName := c.GetString("courseName")
	courseTeacherId := c.GetString("courseTeacherId")
	courseTeacherName := c.GetString("CourseTeacherName")
	courseSemester := c.GetString("courseSemester")
	courseTime := c.GetString("courseTime")
	classroom := c.GetString("classroom")

	classes, _ := models.GetClasses(courseCode, courseName, courseTeacherId, courseTeacherName, courseSemester, courseTime, classroom)

	c.Data["Classes"] = classes
}

func (c *StudentClassControllerPick) Post() {
	c.TplName = "StudentViews/StudentClassViews/"

	courseCode := c.GetString("courseCode")

	class, _ := models.GetClasses(courseCode, "", "", "", "", "", "")
	studentId := c.GetSession("userId").(string)[1:]

	student, err := models.GetStudent(studentId)
	if err != nil {
		return
	}

	models.StudentTimeConflict(student, class[0])
	models.PickClass(student, class[0])
}

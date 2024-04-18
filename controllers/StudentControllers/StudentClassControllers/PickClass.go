package StudentClassControllers

import (
	"CourseSystem/controllers"
)

type StudentClassControllerPick struct {
	controllers.BaseController
}

func (c *StudentClassControllerPick) Get() {
	c.TplName = "StudentViews/StudentClassViews/"
}

// func (c *StudentClassControllerPick) Post() {
// 	c.TplName = "StudentViews/StudentClassViews/"

// 	courseCode := c.GetString("courseCode")
// 	courseName := c.GetString("courseName")
// 	courseTeacherId := c.GetString("courseTeacherId")
// 	courseTeacherName := c.GetString("CourseTeacherName")
// 	courseSemester := c.GetString("courseSemester")
// 	courseTime := c.GetString("courseTime")
// 	classroom := c.GetString("classroom")

// 	classes, _ := models.GetClasses(courseCode, courseName, courseTeacherId, courseTeacherName, courseSemester, courseTime, classroom)
// 	models.PickClass()
// }

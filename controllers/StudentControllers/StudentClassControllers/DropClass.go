package StudentClassControllers

import (
	"CourseSystem/controllers"
)

type StudentClassControllerDrop struct {
	controllers.BaseController
}

func (c *StudentClassControllerDrop) Get() {
	c.TplName = "StudentViews/StudentClassViews/"
}

// func (c *StudentClassControllerDrop) Post() {
// 	c.TplName = "StudentViews/StudentClassViews/"

// 	courseCode := c.GetString("courseCode")
// 	courseName := c.GetString("courseName")
// 	courseTeacherId := c.GetString("courseTeacherId")
// 	courseTeacherName := c.GetString("CourseTeacherName")
// 	courseSemester := c.GetString("courseSemester")
// 	courseTime := c.GetString("courseTime")
// 	classroom := c.GetString("classroom")

// 	classes, _ := models.GetClasses(courseCode, courseName, courseTeacherId, courseTeacherName, courseSemester, courseTime, classroom)
// }

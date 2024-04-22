package AdminClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
	"fmt"
)

type AdminClassControllerRetrieve struct {
	controllers.BaseController
}

func (c *AdminClassControllerRetrieve) Get() {
	c.TplName = "AdminViews/AdminClassViews/RetrieveClass.tpl"
	semesters := []string{"23春季", "23夏季", "23秋季", "23冬季", "24春季", "24夏季"}
	c.Data["Semesters"] = semesters
}
func (c *AdminClassControllerRetrieve) Post() {
	c.TplName = "AdminViews/AdminClassViews/RetrieveClass.tpl"
	courseCode := c.GetString("courseCode")
	courseName := c.GetString("courseName")
	courseTeacherId := c.GetString("courseTeacherId")
	courseTeacherName := c.GetString("CourseTeacherName")
	courseSemester := c.GetString("courseSemester")
	courseTime := c.GetString("courseTime")
	classroom := c.GetString("classroom")
	semesters := []string{"23春季", "23夏季", "23秋季", "23冬季", "24春季", "24夏季"}
	fmt.Println(c.Input())
	c.Data["Semesters"] = semesters
	classes, _ := models.GetClasses(courseCode, courseName, courseTeacherId, courseTeacherName, courseSemester, courseTime, classroom)

	c.Data["Classes"] = classes
}

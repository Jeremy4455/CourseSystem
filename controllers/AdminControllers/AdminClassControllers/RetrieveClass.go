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
	c.Data["Semesters"] = models.Semesters
}
func (c *AdminClassControllerRetrieve) Post() {
	c.TplName = "AdminViews/AdminClassViews/RetrieveClass.tpl"
	courseCode := c.GetString("courseCode")
	courseName := c.GetString("courseName")
	courseTeacherId := c.GetString("courseTeacherId")
	courseTeacherName := c.GetString("courseTeacherName")
	courseSemester := c.GetString("courseSemester")
	courseTime := c.GetString("courseTime")
	classroom := c.GetString("classroom")
	c.Data["Semesters"] = models.Semesters

	classes, err := models.GetClasses(courseCode, courseName, courseTeacherId, courseTeacherName, courseSemester, courseTime, classroom)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.Data["Classes"] = classes
}

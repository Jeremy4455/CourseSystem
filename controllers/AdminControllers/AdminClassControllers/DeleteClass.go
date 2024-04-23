package AdminClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
	"fmt"
)

type AdminClassControllerDelete struct {
	controllers.BaseController
}

func (c *AdminClassControllerDelete) Get() {
	c.SearchClasses()
	c.TplName = "AdminViews/AdminClassViews/DeleteClass.tpl"
}

func (c *AdminClassControllerDelete) SearchClasses() {
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

func (c *AdminClassControllerDelete) Post() {
	c.TplName = "AdminViews/AdminClassViews/DeleteClass.tpl"
	courseCode := c.GetString("courseCode")
	courseTeacherId := c.GetString("courseTeacherId")
	courseSemester := c.GetString("courseSemester")
	c.Data["Semesters"] = models.Semesters

	err := models.DeleteClass(courseCode, courseTeacherId, courseSemester)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

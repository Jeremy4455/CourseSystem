package AdminClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
	"fmt"
)

type AdminClassControllerUpdate struct {
	controllers.BaseController
}

func (c *AdminClassControllerUpdate) Get() {
	c.TplName = "AdminViews/AdminClassViews/UpdateClass.tpl"
	courseCode := c.GetString("courseCode")
	TeacherId := c.GetString("TeacherId")
	courseSemester := c.GetString("courseSemester")
	semesters := models.Semesters
	c.Data["Semesters"] = semesters
	if courseCode == "" || TeacherId == "" || courseSemester == "" {
		classes, err := models.GetAllClasses()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		c.Data["Classes"] = classes
	} else {
		classes, err := models.GetClasses(courseCode, "", TeacherId, "", courseSemester, "", "")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		c.Data["Classes"] = classes
	}
}

func (c *AdminClassControllerUpdate) Post() {
	c.TplName = "AdminViews/AdminClassViews/UpdateClass.tpl"
	courseCode := c.GetString("CourseCode")
	courseName := c.GetString("CourseName")
	courseTeacherId := c.GetString("CourseTeacherId")
	courseTeacherName := c.GetString("CourseTeacherName")
	courseSemester := c.GetString("CourseSemester")
	classTime := c.GetString("CourseTime")
	capacity := c.GetString("Capacity")
	classroom := c.GetString("Classroom")
	c.Data["Semesters"] = models.Semesters
	class, _ := models.GetClasses(courseCode, courseName, courseTeacherId, courseTeacherName, courseSemester, classTime, classroom)
	if len(class) != 1 {
		return
	}
	err := models.ReviseClass(class[0], courseTeacherId, classTime, capacity, classroom)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

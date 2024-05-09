package AdminClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
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
		classes, _ := models.GetAllClasses()
		c.Data["Classes"] = classes
	} else {
		classes, err := models.GetClasses(courseCode, "", TeacherId, "", courseSemester, "", "", models.PER_CLASS)
		if err != nil {
			c.Err(err)
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
	class, _ := models.GetClasses(courseCode, courseName, courseTeacherId, courseTeacherName, courseSemester, "", "", models.PER_CLASS)
	if len(class) != 1 {
		return
	}
	err := models.ReviseClass(class[0], courseTeacherId, classTime, capacity, classroom)
	if err != nil {
		c.Err(err)
	} else {
		c.Sucess()
		classes, _ := models.GetAllClasses()
		c.Data["Classes"] = classes
	}
}

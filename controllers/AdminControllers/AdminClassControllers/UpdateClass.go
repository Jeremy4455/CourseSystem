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
	c.searchClass()
}
func (c *AdminClassControllerUpdate) searchClass() {
	courseCode := c.GetString("CourseCode")
	teacherId := c.GetString("TeacherId")
	semester := c.GetString("Semester")
	if courseCode == "" && teacherId == "" && semester == "" {
		courses, err := models.GetAllClasses()
		if err != nil {
			return
		}
		c.Data["Classes"] = courses
	} else {
		courses, err := models.GetClasses(courseCode, "", teacherId, "", semester, "", "")
		if err != nil {
			return
		}
		c.Data["Classes"] = courses
	}
}

func (c *AdminClassControllerUpdate) Post() {
	c.TplName = "AdminViews/AdminClassViews/UpdateClass.tpl"
	courseCode := c.GetString("courseCode")
	courseName := c.GetString("courseName")
	courseTeacherId := c.GetString("courseTeacherId")
	courseTeacherName := c.GetString("CourseTeacherName")
	courseSemester := c.GetString("courseSemester")
	classTime := c.GetString("courseTime")
	capacity := c.GetString("Capacity")
	classroom := c.GetString("classroom")

	class, _ := models.GetClasses(courseCode, courseName, courseTeacherId, courseTeacherName, courseSemester, classTime, classroom)
	if len(class) != 1 {
		return
	}
	err := models.ReviseClass(class[0], courseTeacherId, classTime, capacity, classroom)
	if err != nil {
		return
	}
}

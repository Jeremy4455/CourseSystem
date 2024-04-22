package AdminClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminClassControllerCreate struct {
	controllers.BaseController
}

func (c *AdminClassControllerCreate) Get() {
	c.TplName = "AdminViews/AdminClassViews/CreateClass.tpl"
	c.searchCourse()
}

func (c *AdminClassControllerCreate) searchCourse() {
	courseCode := c.GetString("courseCode")
	courseName := c.GetString("courseName")
	if courseCode == "" && courseName == "" {
		courses, err := models.GetAllCourses()
		if err != nil {
			return
		}
		c.Data["Courses"] = courses
	} else {
		courses, err := models.GetCourses(courseCode, courseName)
		if err != nil {
			return
		}
		c.Data["Courses"] = courses
	}
}

func (c *AdminClassControllerCreate) Post() {
	c.TplName = "AdminViews/AdminClassViews/CreateClass.tpl"
	courseCode := c.GetString("CourseCode")
	courseName := c.GetString("CourseName")
	courseTeacherId := c.GetString("CourseTeacherId")
	courseSemester := c.GetString("CourseSemester")
	classTime := c.GetString("ClassTime")
	capacity := c.GetString("Capacity")
	classroom := c.GetString("Classroom")

	err := models.CreateClass(courseCode, courseName, courseTeacherId, courseSemester, classTime, capacity, classroom)
	if err != nil {
		return
	}
}

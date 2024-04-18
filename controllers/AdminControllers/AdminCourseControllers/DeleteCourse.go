package AdminStudentControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
	"fmt"
)

type AdminCourseControllerDelete struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminCourseControllerDelete) Get() {
	c.searchCourse()
	c.TplName = "AdminViews/AdminCourseViews/DeleteCourse.tpl"
}

func (c *AdminCourseControllerDelete) searchCourse() {
	courseCode := c.GetString("CourseCode")
	name := c.GetString("Name")
	courses, _ := models.GetCourses(courseCode, name)
	if courses == nil {
		return
	}
	c.Data["Courses"] = courses
}

func (c *AdminCourseControllerDelete) Post() {
	c.TplName = "AdminViews/AdminCourseViews/DeleteCourse.tpl"
	courseCode := c.GetString("CourseCode")
	fmt.Println(courseCode)
	err := models.DeleteCourse(courseCode)
	if err != nil {
		return
	}
}

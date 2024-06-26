package AdminCourseControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminCourseControllerUpdate struct {
	controllers.BaseController
}

func (c *AdminCourseControllerUpdate) Get() {
	c.TplName = "AdminViews/AdminCourseViews/UpdateCourse.tpl"
	courseCode := c.GetString("CourseCode")
	c.searchCourse(courseCode)
}

func (c *AdminCourseControllerUpdate) searchCourse(courseCode string) {
	if courseCode == "" {
		courses, _ := models.GetAllCourses()
		c.Data["Courses"] = courses
	} else {
		courses, err := models.GetCourses(courseCode, "")
		if err != nil {
			c.Err(err)
			return
		}
		c.Data["Courses"] = courses
	}
}

func (c *AdminCourseControllerUpdate) Post() {
	c.TplName = "AdminViews/AdminCourseViews/UpdateCourse.tpl"

	courseCode := c.GetString("CourseCode")
	name := c.GetString("Name")
	college := c.GetString("College")
	credit := c.GetString("Credit")
	proportion := c.GetString("Proportion")
	if proportion == "" {
		proportion = "0.3"
	}

	err := models.ReviseCourse(courseCode, name, college, credit, proportion)
	if err != nil {
		c.Err(err)
	} else {
		c.Sucess()
		c.searchCourse("")
	}
}

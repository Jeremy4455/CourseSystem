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
	c.searchCourse()
}

func (c *AdminCourseControllerUpdate) searchCourse() {
	courseCode := c.GetString("CourseCode")
	if courseCode == "" {
		courses, _ := models.GetAllCourses()
		c.Data["Courses"] = courses
	} else {
		courses, _ := models.GetCourses(courseCode, "")
		c.Data["Courses"] = courses
	}
}

func (c *AdminCourseControllerUpdate) Post() {
	c.TplName = "AdminViews/AdminCourseViews/UpdateCourse.tpl"
	courseCode := c.GetString("CourseCode")
	name := c.GetString("Name")
	college := c.GetString("College")
	credit := c.GetString("Credit")

	err := models.ReviseCourse(courseCode, name, college, credit)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"message": "课程更新成功"}
	}
}

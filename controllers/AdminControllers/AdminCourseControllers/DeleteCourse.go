package AdminCourseControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminCourseControllerDelete struct {
	controllers.BaseController
}

func (c *AdminCourseControllerDelete) Get() {
	c.TplName = "AdminViews/AdminCourseViews/DeleteCourse.tpl"
	c.searchCourse()
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

	err := models.DeleteCourse(courseCode)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"message": "课程删除成功"}
	}
}

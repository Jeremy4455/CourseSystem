package AdminCourseControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminCourseControllerCreate struct {
	controllers.BaseController
}

func (c *AdminCourseControllerCreate) Get() {
	c.TplName = "AdminViews/AdminCourseViews/CreateCourse.tpl"
}

func (c *AdminCourseControllerCreate) Post() {
	c.TplName = "AdminViews/AdminCourseViews/CreateCourse.tpl"
	courseCode := c.GetString("CourseCode")
	name := c.GetString("Name")
	college := c.GetString("College")
	creditStr := c.GetString("Credit")

	err := models.CreateCourse(courseCode, name, college, creditStr)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"message": "课程添加成功"}
	}
}

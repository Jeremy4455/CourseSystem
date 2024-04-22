package AdminCourseControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminCourseControllerCreate struct {
	controllers.BaseController
	viewpath string
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
	// 添加课程到数据库
	err := models.CreateCourse(courseCode, name, college, creditStr)
	if err != nil {
		return
	}
}

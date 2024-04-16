package AdminStudentControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
	"fmt"
)

type AdminCourseControllerAdd struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminCourseControllerAdd) Get() {
	c.TplName = "AdminViews/AdminCourseViews/addcourse.tpl"
}

func (c *AdminCourseControllerAdd) Post() {
	c.TplName = "AdminViews/AdminCourseViews/addcourse.tpl"
	courseCode := c.GetString("CourseCode")
	name := c.GetString("Name")
	college := c.GetString("College")
	creditStr := c.GetString("Credit")

	// 添加课程到数据库
	err := models.AddCourse(courseCode, name, college, creditStr)
	fmt.Println(err)
	if err != nil {
		return
	}
}

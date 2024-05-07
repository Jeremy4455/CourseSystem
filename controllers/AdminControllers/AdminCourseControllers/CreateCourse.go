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
	credit := c.GetString("Credit")
	proportion := c.GetString("Proportion")
	if proportion == "" {
		proportion = "0.3"
	}

	err := models.CreateCourse(courseCode, name, college, credit, proportion)
	if err != nil {
		c.Err(err)
	} else {
		c.Data["json"] = map[string]interface{}{"message": "课程添加成功"}
	}
}

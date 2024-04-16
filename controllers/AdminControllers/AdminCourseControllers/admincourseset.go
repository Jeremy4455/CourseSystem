package AdminStudentControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type AdminCourseControllerSet struct {
	controllers.BaseController
}

func (c *AdminCourseControllerSet) Get() {
	c.TplName = "AdminViews/AdminCourseViews/setcourse.tpl"
	courseCode := c.GetString("CourseCode")
	courses, _ := models.GetCourses(courseCode, "")
	if courses == nil {
		return
	}
	c.Data["Courses"] = courses
}

func (c *AdminCourseControllerSet) Post() {
	c.TplName = "AdminViews/AdminCourseViews/setcourse.tpl"
	name := c.GetString("Name")
	courseCode := c.GetString("CourseCode")
	college := c.GetString("College")
	creditStr := c.GetString("Credit")

	// 查询要更新的课程
	course := models.Course{CourseCode: courseCode}
	err := orm.NewOrm().Read(&course, "CourseCode")
	if err != nil {
		return
	}

	// 判断表单字段是否为空，如果为空则保持原来的数据
	if name != "" {
		course.Name = name
	}
	if college != "" {
		course.College = college
	}
	if creditStr != "" {
		credit, _ := strconv.Atoi(creditStr)
		course.Credit = credit
	}

	_, err = orm.NewOrm().Update(&course)

	if err != nil {
		// 处理更新错误
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	} else {
		// 更新成功
		c.Data["json"] = map[string]interface{}{"message": "课程更新成功"}
	}
}

package AdminCourseControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
	"fmt"

	"github.com/beego/beego/orm"
)

type AdminCourseControllerSearch struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminCourseControllerSearch) Get() {
	c.viewpath = "AdminViews/AdminCourseViews/searchcourse.tpl"
	c.TplName = c.viewpath
}
func (c *AdminCourseControllerSearch) Post() {
	c.TplName = c.viewpath

	courseNo := c.GetString("courseNo")
	courseName := c.GetString("courseName")
	courseTeacher := c.GetString("courseTeacher")
	courseTime := c.GetString("courseTime")
	credit := c.GetString("credit")
	classroom := c.GetString("classroom")

	o := orm.NewOrm()
	q := o.QueryTable("course")

	if courseNo != "" {
		q = q.Filter("CourseCode", courseNo)
	}
	if courseName != "" {
		q = q.Filter("Name", courseName)
	}
	if courseTeacher != "" {
		q = q.Filter("Teacher", courseTeacher)
	}
	if courseTime != "" {
		q = q.Filter("ClassTime", courseTime)
	}
	if credit != "" {
		q = q.Filter("Credit", credit)
	}
	if classroom != "" {
		q = q.Filter("Location", classroom)
	}
	var courses []models.Course
	_, err := q.All(&courses)

	if err != nil {
		// 处理错误
		fmt.Println("Error querying courses from database:", err)
	}
	c.Data["Courses"] = courses
}

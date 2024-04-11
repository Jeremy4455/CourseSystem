package AdminCourseControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
	"fmt"

	"github.com/beego/beego/orm"
)

type AdminClassControllerDelete struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminClassControllerDelete) Get() {
	c.viewpath = "AdminViews/AdminCourseViews/deletecourse.tpl"
	c.TplName = c.viewpath
}

func (c *AdminClassControllerDelete) Post() {
	c.TplName = c.viewpath

	courseCode := c.GetString("courseCode")
	courseName := c.GetString("courseName")
	courseTeacherId := c.GetString("courseTeacherId")
	courseSemester := c.GetString("courseSemester")
	courseTime := c.GetString("courseTime")
	classroom := c.GetString("classroom")

	if courseCode == "" || courseTeacherId == "" || courseSemester == "" || courseTime == "" || classroom == "" {
		return
	}

	o := orm.NewOrm()
	q := o.QueryTable("class").Filter("Semester", courseSemester)

	if courseCode != "" {
		course, err := models.GetCourse(courseCode, courseName)
		if err != nil {
			return
		}
		q = q.Filter("Course", course)
	}

	if courseTeacherId != "" {
		teacher, err := models.GetTeacher(courseTeacherId)
		if err != nil {
			return
		}
		q = q.Filter("Teacher", teacher)
	}

	if courseTime != "" {
		q = q.Filter("ClassTime__contains", courseTime)
	}

	if classroom != "" {
		q = q.Filter("Location__contains", classroom)
	}

	var classes []models.Course
	_, err := q.All(&classes)

	if err != nil {
		// 处理错误
		fmt.Println("Error querying courses from database:", err)
	}
	c.Data["Classes"] = classes

}

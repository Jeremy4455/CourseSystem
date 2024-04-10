package AdminCourseControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
	"fmt"
	"strconv"

	"github.com/beego/beego/orm"
)

type AdminCourseControllerAdd struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminCourseControllerAdd) Get() {
	c.viewpath = "AdminViews/AdminCourseViews/addcourse.tpl"
	c.TplName = c.viewpath
}

func (c *AdminCourseControllerAdd) Post() {
	c.TplName = c.viewpath

	courseNo := c.GetString("courseNo")
	courseName := c.GetString("courseName")
	courseTeacher := c.GetString("courseTeacher")
	courseTime := c.GetString("courseTime")
	credit := c.GetString("credit")
	classroom := c.GetString("classroom")

	if courseNo != "" || courseName != "" || courseTeacher != "" || courseTime != "" || credit != "" || classroom != "" {
		return
	}

	o := orm.NewOrm()
	var courses []models.Course
	_, err := o.QueryTable("course").
		Filter("CourseCode", courseNo).
		Filter("Name", courseName).
		Filter("Teacher", courseTeacher).
		Filter("ClassTime", courseTime).
		Filter("Credit", credit).
		Filter("Location", classroom).
		All(&courses)
	if err != nil {
		// 处理错误
		fmt.Println("Error querying courses from database:", err)
		return
	}
	if len(courses) != 0 {
		return
	}

	// code below remain to be done
	var teacher []models.Teacher

	creditint, err := strconv.Atoi(credit)
	newcourse := models.Course{Name: courseName, CourseCode: courseNo, Teacher: &teacher[0], ClassTime: courseTime, Credit: creditint, Location: classroom}
	id, err := o.Insert(&newcourse)
	if err == nil {
		fmt.Printf("插入成功，新记录的ID为: %d\n", id)
	} else {
		fmt.Println("插入失败:", err)
	}
}

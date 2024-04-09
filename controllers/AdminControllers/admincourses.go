package AdminControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type AdminCourseController struct {
	controllers.BaseController
}

func (this *AdminCourseController) CourseList() {
	this.TplName = "AdminViews/course.tpl"
	courseNo := this.GetString("courseNo")
	courseName := this.GetString("courseName")
	courseTeacher := this.GetString("courseTeacher")
	courseTime := this.GetString("courseTime")
	credit := this.GetString("credit")
	classroom := this.GetString("classroom")

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
	this.Data["Courses"] = courses
	this.TplName = "AdminViews/course.tpl"
}

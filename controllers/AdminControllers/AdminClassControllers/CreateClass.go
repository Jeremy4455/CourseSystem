package AdminCourseControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type AdminClassControllerCreate struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminClassControllerCreate) Get() {
	c.viewpath = "AdminViews/AdminCourseViews/CreateCourse.tpl"
	c.TplName = c.viewpath
}

func (c *AdminClassControllerCreate) Post() {
	c.TplName = c.viewpath

	courseCode := c.GetString("courseCode")
	courseName := c.GetString("courseName")
	courseTeacherId := c.GetString("courseTeacherId")
	courseSemester := c.GetString("courseSemester")
	courseTime := c.GetString("courseTime")
	classroom := c.GetString("classroom")
	capacity := c.GetString("capacity")

	if courseCode == "" || courseTeacherId == "" || courseSemester == "" || courseTime == "" || classroom == "" || capacity == "" {
		return
	}

	course, _ := models.GetCourses(courseCode, courseName)
	if course == nil || len(course) > 1 {
		return
	}

	teacher, _ := models.GetTeacher(courseTeacherId)
	if teacher == nil {
		return
	}

	if models.ExistClass(course[0], teacher, courseSemester) == true {
		return
	}

	o := orm.NewOrm()
	var lastRecord models.Class

	err := o.QueryTable("class").OrderBy("-id").Limit(1).One(&lastRecord)
	if err != nil {
		return
	}

	cap, _ := strconv.Atoi(capacity)
	class := &models.Class{
		Id:        lastRecord.Id + 1,
		Course:    course[0],
		Teacher:   teacher,
		Semester:  courseSemester,
		ClassTime: courseTime,
		Capacity:  cap,
		Location:  classroom,
	}
	o.Insert(class)
}

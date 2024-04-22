package AdminClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminClassControllerCreate struct {
	controllers.BaseController
}

func (c *AdminClassControllerCreate) Get() {
	c.TplName = "AdminViews/AdminClassViews/CreateClass.tpl"
	classes := models.GetAllClasses()
	semesters := []string{"23春季", "23夏季", "23秋季", "23冬季", "24春季", "24夏季"}
	// 将班级和学期列表传递到模板中
	c.Data["Classes"] = classes
	c.Data["Semesters"] = semesters
}
func (c *AdminClassControllerCreate) Post() {
	c.TplName = "AdminViews/AdminClassViews/CreateClass.tpl"
	courseCode := c.GetString("CourseCode")
	courseName := c.GetString("CourseName")
	courseTeacherId := c.GetString("CourseTeacherId")
	courseSemester := c.GetString("CourseSemester")
	classTime := c.GetString("ClassTime")
	capacity := c.GetString("Capacity")
	classroom := c.GetString("Classroom")

	err := models.CreateClass(courseCode, courseName, courseTeacherId, courseSemester, classTime, capacity, classroom)
	if err != nil {
		return
	}
}

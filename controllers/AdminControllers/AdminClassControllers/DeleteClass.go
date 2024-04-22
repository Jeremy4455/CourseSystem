package AdminClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminClassControllerDelete struct {
	controllers.BaseController
}

func (c *AdminClassControllerDelete) Get() {
	c.SearchClasses()
	c.TplName = "AdminViews/AdminClassViews/DeleteClass.tpl"
}

func (c *AdminClassControllerDelete) SearchClasses() {
	courseCode := c.GetString("courseCode")
	courseName := c.GetString("courseName")
	courseTeacherId := c.GetString("courseTeacherId")
	courseTeacherName := c.GetString("courseTeacherName")
	courseSemester := c.GetString("courseSemester")
	courseTime := c.GetString("courseTime")
	classroom := c.GetString("classroom")
	semesters := []string{"23春季", "23夏季", "23秋季", "23冬季", "24春季", "24夏季"}
	c.Data["Semesters"] = semesters
	classes, _ := models.GetClasses(courseCode, courseName, courseTeacherId, courseTeacherName, courseSemester, courseTime, classroom)
	c.Data["Classes"] = classes
}

func (c *AdminClassControllerDelete) Post() {
	c.TplName = "AdminViews/AdminClassViews/DeleteClass.tpl"
	courseCode := c.GetString("courseCode")
	courseTeacherId := c.GetString("courseTeacherId")
	courseSemester := c.GetString("courseSemester")
	semesters := []string{"23春季", "23夏季", "23秋季", "23冬季", "24春季", "24夏季"}
	c.Data["Semesters"] = semesters
	err := models.DeleteClass(courseCode, courseTeacherId, courseSemester)
	if err != nil {
		return
	}
}

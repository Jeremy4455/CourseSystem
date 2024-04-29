package StudentClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type StudentClassControllerCourseTable struct {
	controllers.BaseController
}

func (c *StudentClassControllerCourseTable) Get() {
	c.TplName = "StudentViews/StudentClassViews/CourseTable.tpl"
	c.searchClass()

	userInfo := c.GetUserInfo()
	c.Data["UserInfo"] = userInfo
}

func (c *StudentClassControllerCourseTable) searchClass() {
	studentId := c.GetSession("userId").(string)[1:]
	semester := c.GetSession("semester").(string)
	s, err := models.GetStudent(studentId)
	if err != nil {
		return
	}
	classes, err := models.GetClassStudent(s, semester, nil)
	if err != nil {
		return
	}
	var class []*models.Class
	for _, v := range classes {
		class = append(class, v.Class)
	}
	c.Data["Classes"] = class
}

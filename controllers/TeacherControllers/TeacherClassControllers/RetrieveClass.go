package TeacherClassController

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type TeacherClassControllerRetrieve struct {
	controllers.BaseController
}

func (c *TeacherClassControllerRetrieve) Get() {
	c.TplName = ""
	c.searchClass()

	userInfo := c.GetUserInfo()
	c.Data["UserInfo"] = userInfo
}

func (c *TeacherClassControllerRetrieve) searchClass() {
	teacherId := c.GetSession("userId").(string)[1:]
	semester := c.GetSession("semester").(string)

	classes, err := models.GetClasses("", "", teacherId, "", semester, "", "")
	if err != nil {
		return
	}
	c.Data["Classes"] = classes
}

// 转SetGrade
func (c *TeacherClassControllerRetrieve) Post() {
	c.TplName = ""
	teacherId := c.GetSession("userId").(string)[1:]
	semester := c.GetSession("semester").(string)
	courseCode := c.GetString("CourseCode")

	_, err := models.GetClasses(courseCode, "", teacherId, "", semester, "", "")
	if err != nil {
		return
	}
}

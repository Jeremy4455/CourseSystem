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

	userInfo := c.GetUserInfo()
	c.Data["UserInfo"] = userInfo
}

func (c *TeacherClassControllerRetrieve) Post() {
	c.TplName = ""
	teacherId := c.GetSession("userId").(string)[1:]
	semester := c.GetString("Semester")

	classes, err := models.GetClasses("", "", teacherId, "", semester, "", "")
	if err != nil {
		return
	}
	c.Data["Classes"] = classes
}

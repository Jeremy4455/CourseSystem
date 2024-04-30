package TeacherClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type TeacherClassControllerRetrieve struct {
	controllers.BaseController
}

func (c *TeacherClassControllerRetrieve) Get() {
	c.TplName = "TeacherViews/TeacherClassViews/RetrieveClass.tpl"
	c.searchClass()

	userInfo := c.GetUserInfo()
	c.Data["UserInfo"] = userInfo
}

func (c *TeacherClassControllerRetrieve) searchClass() {
	teacherId := c.GetSession("userId").(string)[1:]
	semester := c.GetSession("semester").(string)

	classes, err := models.GetClasses("", "", teacherId, "", semester, "", "", models.TEACHER_UPDATE_GRADE)
	if err != nil {
		return
	}
	c.Data["Classes"] = classes
}

// 转SetGrade
func (c *TeacherClassControllerRetrieve) Post() {
	c.TplName = "TeacherViews/TeacherClassViews/RetrieveClass.tpl"
	teacherId := c.GetSession("userId").(string)[1:]
	semester := c.GetSession("semester").(string)
	courseCode := c.GetString("CourseCode")

	_, err := models.GetClasses(courseCode, "", teacherId, "", semester, "", "", models.TEACHER_UPDATE_GRADE)
	if err != nil {
		return
	}
}

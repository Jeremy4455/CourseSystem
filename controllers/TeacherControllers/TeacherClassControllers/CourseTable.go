package TeacherClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type TeacherClassControllerCourseTable struct {
	controllers.BaseController
}

func (c *TeacherClassControllerCourseTable) Get() {
	c.TplName = "TeacherViews/TeacherClassViews/CourseTable.tpl"
	c.searchClass()

	userInfo := c.GetUserInfo()
	c.Data["UserInfo"] = userInfo
}

func (c *TeacherClassControllerCourseTable) searchClass() {
	teacherId := c.GetSession("userId").(string)[1:]
	semester := c.GetSession("semester").(string)

	classes, err := models.GetClasses("", "", teacherId, "", semester, "", "", models.TEACHER_UPDATE_GRADE)
	if err != nil {
		return
	}
	c.Data["Classes"] = classes
}

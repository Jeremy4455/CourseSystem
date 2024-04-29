package TeacherControllers

import (
	"CourseSystem/controllers"
)

type TeacherIndexController struct {
	controllers.BaseController
}

func (c *TeacherIndexController) Get() {
	c.TplName = "TeacherViews/Index.tpl"
	userInfo := c.GetUserInfo()
	c.Data["UserInfo"] = userInfo
}

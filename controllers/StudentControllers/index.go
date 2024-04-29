package StudentControllers

import (
	"CourseSystem/controllers"
)

type StudentIndexController struct {
	controllers.BaseController
}

func (c *StudentIndexController) Get() {
	c.TplName = "StudentViews/Index.tpl"
	userInfo := c.GetUserInfo()
	c.Data["UserInfo"] = userInfo
}

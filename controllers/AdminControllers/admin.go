package AdminControllers

import "CourseSystem/controllers"

type AdminController struct {
	controllers.BaseController
	viewpath string
}

func (c *AdminController) Get() {
	c.viewpath = "AdminViews/direct.tpl"
	c.TplName = c.viewpath
}

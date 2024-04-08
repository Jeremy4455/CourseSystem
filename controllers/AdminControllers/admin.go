package AdminControllers

import "CourseSystem/controllers"

type AdminController struct {
	controllers.BaseController
}

func (c *AdminController) Get() {
	c.TplName = "index.tpl"
}

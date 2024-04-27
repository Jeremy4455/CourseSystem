package StudentControllers

import (
	"CourseSystem/controllers"
)

type StudentIndexController struct {
	controllers.BaseController
}

func (C *StudentIndexController) Get() {
	C.TplName = "StudentViews/Index.tpl"
}

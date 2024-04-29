package StudentClassControllers

import (
	"CourseSystem/controllers"
)

type StudentClassControllerSelectedClass struct {
	controllers.BaseController
}

func (c *StudentClassControllerSelectedClass) Get() {
	c.TplName = "StudentViews/StudentClassViews/SelectedClass.tpl"
}

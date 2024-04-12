package StudentClassControllers

import (
	"CourseSystem/controllers"
)

type StudentClassControllerPick struct {
	controllers.BaseController
	viewpath string
}

func (c *StudentClassControllerPick) Get() {
	c.viewpath = "StudentViews/StudentClassViews/"
}
func (c *StudentClassControllerPick) Post() {

}

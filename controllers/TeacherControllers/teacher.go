package TeacherControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type TeacherController struct {
	controllers.BaseController
}

func (c *TeacherController) Get() {
	c.TplName = "TeacherViews/SemesterSelection.tpl"
	c.Data["Semesters"] = models.Semesters
}

func (c *TeacherController) Post() {
	semester := c.GetString("semester")
	c.SetSession("semester", semester)
	c.Redirect("/teacher/index", 302)
}
func (c *TeacherController) GetUserInfo() map[string]interface{} {
	name := c.GetSession("name").(string)
	userId := c.GetSession("userId").(string)[1:]
	semester := c.GetSession("semester").(string)

	userInfo := map[string]interface{}{
		"name":     name,
		"userId":   userId,
		"semester": semester,
	}
	return userInfo
}

package StudentControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type StudentController struct {
	controllers.BaseController
}

func (StudentController *StudentController) Get() {
	StudentController.TplName = "StudentViews/SemesterSelection.tpl"
	StudentController.Data["Semesters"] = models.Semesters
}

func (c *StudentController) Post() {
	semester := c.GetString("semester")
	c.SetSession("semester", semester)
	c.Redirect("/student/index", 302)
}
func (c *StudentController) GetUserInfo() map[string]interface{} {
	name := c.GetSession("name").(string)
	studentId := c.GetSession("userId").(string)[1:]
	semester := c.GetSession("semester").(string)

	userInfo := map[string]interface{}{
		"name":      name,
		"studentID": studentId,
		"semester":  semester,
	}
	return userInfo
}

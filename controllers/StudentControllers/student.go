package StudentControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type StudentController struct {
	controllers.BaseController
	Semester string
}

func (StudentController *StudentController) Get() {
	StudentController.TplName = "StudentViews/SemesterSelection.tpl"
	StudentController.Data["Semesters"] = models.Semesters
}

func (StudentController *StudentController) Post() {
	semester := StudentController.GetString("semester")
	StudentController.Semester = semester //这里要保存下选择的学期
	StudentController.Redirect("/student/index", 302)
}

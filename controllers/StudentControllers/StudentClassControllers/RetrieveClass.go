package StudentClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type StudentClassControllerRetrieve struct {
	controllers.BaseController
}

func (c *StudentClassControllerRetrieve) Get() {
	c.TplName = ""

	userInfo := c.GetUserInfo()
	c.Data["UserInfo"] = userInfo
}
func (c *StudentClassControllerRetrieve) Post() {
	c.TplName = ""
	studentId := c.GetSession("userId").(string)[1:]
	semester := c.GetString("Semester")

	student, err := models.GetStudent(studentId)
	if err != nil {
		return
	}
	cs, err := models.GetClassStudent(student, semester, nil)
	if err != nil {
		return
	}
	var classes []*models.Class
	for _, class := range cs {
		classes = append(classes, class.Class)
	}
	c.Data["Classes"] = classes
}

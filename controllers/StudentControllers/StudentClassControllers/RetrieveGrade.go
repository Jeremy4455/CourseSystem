package StudentClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type StudentClassControllerRetrieveGrade struct {
	controllers.BaseController
}

func (c *StudentClassControllerRetrieveGrade) Get() {
	c.TplName = "StudentViews/StudentClassViews/RetrieveGrade.tpl"

	userInfo := c.GetUserInfo()
	c.Data["UserInfo"] = userInfo

	studentId := c.GetSession("userId").(string)[1:]
	semester := c.GetSession("semester").(string)
	s, err := models.GetStudent(studentId)
	if err != nil {
		return
	}
	cs, err := models.GetClassStudent(s, semester, nil)
	if err != nil {
		return
	}

	classes := []map[string]interface{}{}
	for _, class := range cs {
		t := make(map[string]interface{})
		t["CourseCode"] = class.Class.Course.CourseCode
		t["CourseName"] = class.Class.Course.Name
		t["Teacher"] = class.Class.Teacher.Name
		if class.Class.Level <= models.TEACHER_UPDATE_GRADE {
			t["Performance"] = "暂未发布"
			t["Score"] = "暂未发布"
			t["Grade"] = "暂未发布"
			t["Level"] = "暂未发布"
		} else {
			t["Performance"] = class.Performance
			t["Score"] = class.Score
			proportion := class.Class.Course.Proportion
			t["Grade"] = proportion*class.Performance + (1-proportion)*class.Score
			t["Level"] = models.CalculateGrade(float64(proportion*class.Performance + (1-proportion)*class.Score))
		}
		classes = append(classes, t)
	}
	c.Data["Classes"] = classes
}

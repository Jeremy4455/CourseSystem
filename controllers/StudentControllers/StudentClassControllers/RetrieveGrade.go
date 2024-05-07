package StudentClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
	"fmt"
)

type StudentClassControllerRetrieveGrade struct {
	controllers.BaseController
}

func (c *StudentClassControllerRetrieveGrade) Get() {
	c.TplName = "StudentViews/StudentClassViews/RetrieveGrade.tpl"

	userInfo := c.GetUserInfo()
	c.Data["UserInfo"] = userInfo
	studentIdInterface, ok := userInfo["studentID"].(string)
	if !ok {
		return
	}
	studentId := studentIdInterface

	semesterInterface, ok := userInfo["semester"].(string)
	if !ok {
		return
	}
	semester := semesterInterface

	student, err := models.GetStudent(studentId)
	if err != nil {
		return
	}
	cs, err := models.GetClassStudent(student, semester, nil)
	if err != nil {
		return
	}
	var classes []map[string]interface{}
	for _, class := range cs {
		t := make(map[string]interface{})
		t["CourseCode"] = class.Class.Course.CourseCode
		t["CourseName"] = class.Class.Course.Name
		t["Teacher"] = class.Class.Teacher.Name
		t["Performance"] = class.Performance
		t["Score"] = class.Score
		proportion := class.Class.Course.Proportion
		t["Grade"] = proportion*class.Performance + (1-proportion)*class.Score
	}
	c.Data["Classes"] = classes
	fmt.Println(c.Data["Classes"])
}

package TeacherClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type TeacherClassControllerSet struct {
	controllers.BaseController
}

func (c *TeacherClassControllerSet) Get() {
	c.TplName = "TeacherViews/TeacherClassViews/SetGrade.tpl"
	c.searchCS()

	userInfo := c.GetUserInfo()
	c.Data["UserInfo"] = userInfo
}

func (c *TeacherClassControllerSet) searchCS() {
	teacherId := c.GetSession("userId").(string)[1:]
	semester := c.GetSession("semester").(string)
	courseCode := c.GetString("CourseCode")
	classes, err := models.GetClasses(courseCode, "", teacherId, "", semester, "", "", models.TEACHER_UPDATE_GRADE)
	c.Data["Class"] = classes[0]
	if err != nil {
		return
	}
	cs, err := models.GetClassStudent(nil, "", classes[0])
	if err != nil {
		return
	}

	var info []map[string]interface{}
	for _, t := range cs {
		tinfo := make(map[string]interface{})
		tinfo["StudentId"] = t.Student.StudentId
		tinfo["Name"] = t.Student.Name
		tinfo["Class"] = t.Student.Class
		tinfo["Performance"] = t.Performance
		tinfo["Score"] = t.Score
		info = append(info, tinfo)
	}

	c.Data["Info"] = info
}

func (c *TeacherClassControllerSet) Post() {
	c.TplName = "TeacherViews/TeacherClassViews/SetGrade.tpl"
	teacherId := c.GetSession("userId").(string)[1:]
	semester := c.GetSession("semester").(string)
	courseCode := c.GetString("CourseCode")
	studentId := c.GetString("StudentId")
	performance := c.GetString("Performance")
	score := c.GetString("Score")

	classes, err := models.GetClasses(courseCode, "", teacherId, "", semester, "", "", models.TEACHER_UPDATE_GRADE)
	if err != nil {
		return
	}
	student, err := models.GetStudent(studentId)
	if err != nil {
		return
	}

	models.UpdateClass(student, classes[0], performance, score, models.TEACHER_UPDATE_GRADE)
}

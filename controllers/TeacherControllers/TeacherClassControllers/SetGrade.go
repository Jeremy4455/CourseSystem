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
	userInfo := c.GetUserInfo()
	c.Data["UserInfo"] = userInfo
	teacherId := c.GetSession("userId").(string)[1:]
	semester := c.GetSession("semester").(string)
	courseCode := c.GetString("CourseCode")
	studentIds := c.GetStrings("StudentId")
	performances := c.GetStrings("Performance")
	scores := c.GetStrings("Score")

	classes, err := models.GetClasses(courseCode, "", teacherId, "", semester, "", "", models.TEACHER_UPDATE_GRADE)
	if err != nil {
		return
	}
	for i := range studentIds {
		student, err := models.GetStudent(studentIds[i])
		if err != nil {
			return
		}

		err = models.UpdateClass(student, classes[0], performances[i], scores[i], models.TEACHER_UPDATE_GRADE)
		if err != nil {
			c.Err(err)
			return
		}
	}

	err = models.UpgradeLevel(classes[0])
	if err != nil {
		c.Err(err)
		return
	}
	c.Sucess()
}

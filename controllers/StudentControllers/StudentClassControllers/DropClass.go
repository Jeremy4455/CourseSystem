package StudentClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type StudentClassControllerDrop struct {
	controllers.BaseController
}

func (c *StudentClassControllerDrop) Get() {
	c.TplName = "StudentViews/StudentClassViews/DropClass.tpl"
	c.showClass()

	userInfo := c.GetUserInfo()
	c.Data["UserInfo"] = userInfo
}
func (c *StudentClassControllerDrop) showClass() {
	studentId := c.GetSession("userId").(string)[1:]
	semester := c.GetSession("semester").(string)
	student, err := models.GetStudent(studentId)
	if err != nil {
		return
	}
	cs, err := models.GetClassStudent(student, semester, nil)
	if err != nil {
		return
	}
	var classes []*models.Class
	for _, v := range cs {
		if v.Class.Level <= models.STUDENT_PICK_DROP_CLASS {
			classes = append(classes, v.Class)
		}
	}

	c.Data["Classes"] = classes
}
func (c *StudentClassControllerDrop) Post() {
	c.TplName = "StudentViews/StudentClassViews/DropClass.tpl"
	c.showClass()

	userInfo := c.GetUserInfo()
	c.Data["UserInfo"] = userInfo
	courseCode := c.GetString("courseCode")
	courseName := c.GetString("courseName")
	courseTeacherId := c.GetString("courseTeacherId")
	courseTeacherName := c.GetString("CourseTeacherName")
	courseSemester := c.GetSession("semester").(string)
	courseTime := c.GetString("courseTime")
	classroom := c.GetString("classroom")

	class, _ := models.GetClasses(courseCode, courseName, courseTeacherId, courseTeacherName, courseSemester, courseTime, classroom, models.STUDENT_PICK_DROP_CLASS)
	studentId := c.GetSession("userId").(string)[1:]
	student, err := models.GetStudent(studentId)
	if err != nil {
		return
	}
	models.DropClass(student, class[0], models.STUDENT_PICK_DROP_CLASS)
}

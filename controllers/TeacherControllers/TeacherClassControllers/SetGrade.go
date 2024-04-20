package TeacherClassController

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type TeacherClassControllerSet struct {
	controllers.BaseController
}

func (c *TeacherClassControllerSet) Get() {
	c.TplName = ""
	c.SearchClass()
}

func (c *TeacherClassControllerSet) SearchClass() {
	semester := c.GetString("Semester")

	teacherId := c.GetSession("userId").(string)[1:]
	classes, err := models.GetClasses("", "", teacherId, "", semester, "", "")
	if err != nil {
		return
	}
	c.Data["Classes"] = classes
}

func (c *TeacherClassControllerSet) Post() {
	courseCode := c.GetString("CourseCode")
	semester := c.GetString("Semester")

	teacherId := c.GetSession("userId").(string)[1:]
	class, err := models.GetClasses(courseCode, "", teacherId, "", semester, "", "")
	if err != nil {
		return
	}
	class_student, err := models.GetClassStudent(class[0])
	if err != nil {
		return
	}

	var information []map[string]interface{}
	for _, cs := range class_student {
		var x map[string]interface{}
		x["StudentId"] = cs.Student.StudentId
		x["StudentName"] = cs.Student.Name
		x["StudentClass"] = cs.Student.Class
		x["Performance"] = cs.Performance
		x["Score"] = cs.Score

		information = append(information, x)
	}

	c.Data["Students"] = information

	studentId := c.GetString("StudentId")
	performance := c.GetString("Performance")
	score := c.GetString("Score")

	student, err := models.GetStudent(studentId)
	if err != nil {
		return
	}

	models.UpdateClass(student, class[0], performance, score)
}

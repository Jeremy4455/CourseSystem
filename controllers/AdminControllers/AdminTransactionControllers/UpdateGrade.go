package AdminTransactionControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminTransactionControllerUpdateGrade struct {
	controllers.BaseController
}

var (
	teacherId string
	semester  string
)

func (c *AdminTransactionControllerUpdateGrade) Get() {
	c.TplName = "AdminViews/AdminTransViews/UpdateGrade.tpl"
	c.Data["Semesters"] = models.Semesters
	c.searchCS()
}
func (c *AdminTransactionControllerUpdateGrade) searchCS() {
	studentId := c.GetString("StudentId")
	courseCode := c.GetString("CourseCode")
	teacherId = c.GetString("TeacherId")
	semester = c.GetString("Semester")

	classes, err := models.GetClasses(courseCode, "", teacherId, "", semester, "", "", models.ADMIN_UPDATE_GRADE)
	if err != nil {
		return
	}
	student, err := models.GetStudent(studentId)
	if err != nil {
		return
	}
	cs, err := models.GetClassStudent(student, semester, classes[0])

	info := map[string]interface{}{}
	info["StudentId"] = student.StudentId
	info["StudentName"] = student.Name
	info["CourseCode"] = courseCode
	info["CourseName"] = classes[0].Course.Name
	info["Performance"] = cs[0].Performance
	info["Score"] = cs[0].Score
	c.Data["ClassStudent"] = info
}

func (c *AdminTransactionControllerUpdateGrade) Post() {
	c.TplName = "AdminViews/AdminTransViews/UpdateGrade.tpl"
	score := c.GetString("Score")
	studentId := c.GetString("StudentId")
	courseCode := c.GetString("CourseCode")
	c.Data["Semesters"] = models.Semesters
	classes, err := models.GetClasses(courseCode, "", teacherId, "", semester, "", "", models.ADMIN_UPDATE_GRADE)
	if err != nil {
		return
	}
	student, err := models.GetStudent(studentId)
	if err != nil {
		return
	}
	err = models.UpdateClass(student, classes[0], "", score, models.ADMIN_UPDATE_GRADE)
	if err != nil {
		return
	}

}

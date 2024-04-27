package StudentClassControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type StudentClassControllerPick struct {
	controllers.BaseController
}

func (c *StudentClassControllerPick) Get() {
	c.TplName = "StudentViews/StudentClassViews/PickClass.tpl"
	c.searchClass()

	userInfo := c.GetUserInfo()
	c.Data["UserInfo"] = userInfo
}
func (c *StudentClassControllerPick) searchClass() {
	courseCode := c.GetString("courseCode")
	courseName := c.GetString("courseName")
	courseTeacherId := c.GetString("courseTeacherId")
	courseTeacherName := c.GetString("CourseTeacherName")
	courseSemester := c.GetSession("semester").(string)
	courseTime := c.GetString("courseTime")
	classroom := c.GetString("classroom")

	var result []map[string]interface{}
	classes, _ := models.GetClasses(courseCode, courseName, courseTeacherId, courseTeacherName, courseSemester, courseTime, classroom)
	for _, class := range classes {
		cnt, err := models.GetPickedCount(class)
		if err != nil {
			return
		}
		t := make(map[string]interface{})
		t["CourseCode"] = class.Course.CourseCode
		t["CourseName"] = class.Course.Name
		t["TeacherId"] = class.Teacher.TeacherId
		t["Teacher"] = class.Teacher.Name
		t["Time"] = class.ClassTime
		t["Classroom"] = class.Location
		t["PickedCount"] = cnt
		t["Capacity"] = class.Capacity
		result = append(result, t)
	}
	c.Data["Classes"] = result
}

func (c *StudentClassControllerPick) Post() {
	c.TplName = "StudentViews/StudentClassViews/PickClass.tpl"
	courseCode := c.GetString("courseCode")
	courseName := c.GetString("courseName")
	courseTeacherId := c.GetString("courseTeacherId")
	courseTeacherName := c.GetString("CourseTeacherName")
	courseSemester := c.GetSession("semester").(string)
	courseTime := c.GetString("courseTime")
	classroom := c.GetString("classroom")

	class, _ := models.GetClasses(courseCode, courseName, courseTeacherId, courseTeacherName, courseSemester, courseTime, classroom)
	studentId := c.GetSession("userId").(string)[1:]
	student, err := models.GetStudent(studentId)
	if err != nil {
		return
	}
	models.PickClass(student, class[0], models.STUDENT_PICK_DROP_CLASS)
}

package AdminTeacherControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
	"github.com/astaxie/beego/orm"
)

type AdminTeacherControllerUpdate struct {
	controllers.BaseController
}

func (c *AdminTeacherControllerUpdate) Get() {
	c.TplName = "AdminViews/AdminTeacherViews/UpdateTeacher.tpl"
	teacherId := c.GetString("TeacherId")
	if teacherId == "" {
		teachers, _ := models.GetAllTeachers()
		c.Data["Teachers"] = teachers
	} else {
		teachers, _ := models.GetTeachers(teacherId, "", "", "")
		c.Data["Teachers"] = teachers
	}
}

func (c *AdminTeacherControllerUpdate) Post() {
	c.TplName = "AdminViews/AdminTeacherViews/UpdateTeacher.tpl"

	teacherId := c.GetString("TeacherId")
	name := c.GetString("Name")
	mobile := c.GetString("Mobile")
	email := c.GetString("Email")

	// 查询要更新的课程
	teacher := models.Teacher{TeacherId: teacherId}
	err := orm.NewOrm().Read(&teacher, "TeacherId")
	if err != nil {
		return
	}

	// 判断表单字段是否为空，如果为空则保持原来的数据
	if name != "" {
		teacher.Name = name
	}
	if mobile != "" {
		teacher.Mobile = mobile
	}
	if email != "" {
		teacher.Email = email
	}

	_, err = orm.NewOrm().Update(&teacher)
}

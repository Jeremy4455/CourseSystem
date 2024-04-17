package models

import "github.com/astaxie/beego/orm"

func CheckTime(busytime, classtime string) bool {

	return true
}

func StudentTimeConflict(s *Student, c *Class) bool {
	if s == nil || c == nil {
		return false
	}
	classes := s.Classes
	for _, class := range classes {
		if CheckTime(class.Class.ClassTime, c.ClassTime) == false {
			return false
		}
	}
	return true
}

func TeacherTimeConflict(t *Teacher, c *Class) bool {
	if t == nil || c == nil {
		return false
	}
	classes := t.Classes
	for _, class := range classes {
		if CheckTime(class.ClassTime, c.ClassTime) == false {
			return false
		}
	}
	return true
}

func Syncronize() {
	o := orm.NewOrm()

	var students []*Student
	_, err := o.QueryTable("student").All(&students)
	if err != nil {
		return
	}

	for _, student := range students {
		err := o.Read(&User{Id: student.StudentId})
		if err == nil {
			continue
		}
		AddUser(student.StudentId, student.Name, "123456", "student")
	}

	var teachers []*Teacher
	_, err = o.QueryTable("teacher").All(&teachers)
	if err == nil {
		return
	}

	for _, teacher := range teachers {
		err := o.Read(&User{Id: teacher.TeacherId})
		if err != nil {
			continue
		}
		AddUser(teacher.TeacherId, teacher.Name, "teacher123", "teacher")
	}
}

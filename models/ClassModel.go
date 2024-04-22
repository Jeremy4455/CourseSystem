package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type Class struct {
	Id        int
	Course    *Course  `orm:"rel(fk)"` // Course作为外键
	Teacher   *Teacher `orm:"rel(fk)"` // Teacher作为外键
	Semester  string
	ClassTime string
	Capacity  int
	Location  string
}

func (c *Class) TableIndex() [][]string {
	return [][]string{
		{"Course", "Teacher", "Semester"},
	}
}
func (c *Class) TableUnique() [][]string {
	return [][]string{
		{"Course", "Teacher", "Semester"},
	}
}

func ExistClass(c *Course, t *Teacher, s string) bool {
	o := orm.NewOrm()
	q := o.QueryTable("class")
	exist := q.Filter("Course", c).Filter("Teacher", t).Filter("Semester", s).Exist()
	return exist
}

func GetAllClasses() []*Class {
	o := orm.NewOrm()
	q := o.QueryTable("class")
	var classes []*Class
	_, err := q.All(&classes)
	if err != nil {
		return nil
	}
	return classes
}

func GetClasses(courseCode, courseName, courseTeacherId, courseTeacherName, courseSemester, classTime, classroom string) ([]*Class, error) {
	if courseSemester == "" {
		return nil, nil
	}

	o := orm.NewOrm()
	q := o.QueryTable("class").Filter("Semester", courseSemester)

	if courseCode != "" || courseName != "" {
		courses, err := GetCourses(courseCode, courseName)
		if err != nil {
			return nil, nil
		}
		q = q.Filter("Course", courses)
	}

	if courseTeacherId != "" || courseTeacherName != "" {
		teacher, err := GetTeacher(courseTeacherId, courseTeacherName, "", "")
		if err != nil {
			return nil, nil
		}
		q = q.Filter("Teacher", teacher)
	}

	if classTime != "" {
		q = q.Filter("ClassTime__contains", classTime)
	}

	if classroom != "" {
		q = q.Filter("Location__contains", classroom)
	}

	var classes []*Class
	_, err := q.All(classes)

	if err != nil {
		// 处理错误
		fmt.Println("Error querying courses from database:", err)
	}
	return classes, nil
}

func AddClass(courseCode, courseName, courseTeacherId, courseSemester, classTime, capacity, classroom string) bool {
	if courseCode == "" || courseTeacherId == "" || courseSemester == "" || classTime == "" || classroom == "" || capacity == "" {
		return false
	}

	course, _ := GetCourses(courseCode, courseName)
	if len(course) != 1 {
		return false
	}

	teacher, _ := GetTeacher(courseTeacherId, "", "", "")
	if teacher == nil {
		return false
	}

	if ExistClass(course[0], teacher[0], courseSemester) == true {
		return false
	}

	cap, err := strconv.Atoi(capacity)
	if err != nil {
		return false
	}
	class := &Class{
		Course:    course[0],
		Teacher:   teacher[0],
		Semester:  courseSemester,
		ClassTime: classTime,
		Capacity:  cap,
		Location:  classroom,
	}

	o := orm.NewOrm()

	var existedClass []*Class
	_, err = o.QueryTable("Class").Filter("Location", classroom).All(&existedClass)
	if err != nil {
		return false
	}
	for _, c := range existedClass {
		if ClassConflict(class, c) == false {
			return false
		}
	}
	if TeacherTimeConflict(teacher[0], class) == false {
		return false
	}

	o.Insert(class)
	return true
}

func DeleteClass(courseCode, courseTeacherId, courseSemester string) bool {
	if courseCode == "" || courseTeacherId == "" || courseSemester == "" {
		return false
	}
	class, _ := GetClasses(courseCode, "", courseTeacherId, "", courseSemester, "", "")
	if class == nil {
		return false
	}
	o := orm.NewOrm()
	_, err := o.Delete(class[0])
	if err != nil {
		return false
	}
	return true
}

func ReviseClass(c *Class, courseTeacherId, classTime, capacity, classroom string) bool {
	if c == nil {
		return false
	}
	if courseTeacherId != "" {
		teacher, err := GetTeacher(courseTeacherId, "", "", "")
		if err != nil {
			return false
		}
		c.Teacher = teacher[0]
	}
	if classTime != "" {
		c.ClassTime = classTime
	}
	if capacity != "" {
		cap, err := strconv.Atoi(capacity)
		if err != nil {
			return false
		}
		c.Capacity = cap
	}
	if classroom != "" {
		c.Location = classroom
	}

	_, err := orm.NewOrm().Update(c)
	if err != nil {
		return false
	}
	return true
}

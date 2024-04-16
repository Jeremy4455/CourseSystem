package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type Class struct {
	Id        int      `orm:"pk"`
	Course    *Course  `orm:"rel(fk);index"` // Course作为外键
	Teacher   *Teacher `orm:"rel(fk);index"` // Teacher作为外键
	Semester  string   `orm:"index"`
	ClassTime string
	Capacity  int
	Location  string
}

// 多字段唯一键
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

func GetClasses(courseCode, courseName, courseTeacherId, courseSemester, courseTime, classroom string) ([]*Class, error) {
	if courseSemester == "" {
		return nil, nil
	}

	o := orm.NewOrm()
	q := o.QueryTable("class").Filter("Semester", courseSemester)

	if courseCode != "" || courseName != "" {
		courses, err := GetCourse(courseCode, courseName)
		if err != nil {
			return nil, nil
		}
		q = q.Filter("Course", courses)
	}

	if courseTeacherId != "" {
		teacher, err := GetTeacher(courseTeacherId)
		if err != nil {
			return nil, nil
		}
		q = q.Filter("Teacher", teacher)
	}

	if courseTime != "" {
		q = q.Filter("ClassTime__contains", courseTime)
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

func AddClass(courseCode, courseName, courseTeacherId, courseSemester, courseTime, capacity, classroom string) bool {
	if courseCode == "" || courseTeacherId == "" || courseSemester == "" || courseTime == "" || classroom == "" || capacity == "" {
		return false
	}

	course, _ := GetCourse(courseCode, courseName)
	if course == nil || len(course) > 1 {
		return false
	}

	teacher, _ := GetTeacher(courseTeacherId)
	if teacher == nil {
		return false
	}

	if ExistClass(course[0], teacher, courseSemester) == true {
		return false
	}

	o := orm.NewOrm()
	var lastRecord Class

	err := o.QueryTable("class").OrderBy("-id").Limit(1).One(&lastRecord)
	if err != nil {
		return false
	}

	cap, _ := strconv.Atoi(capacity)
	class := &Class{
		Id:        lastRecord.Id + 1,
		Course:    course[0],
		Teacher:   teacher,
		Semester:  courseSemester,
		ClassTime: courseTime,
		Capacity:  cap,
		Location:  classroom,
	}
	o.Insert(class)

	return true
}

func DeleteClass(courseCode, courseTeacherId, courseSemester string) bool {
	if courseCode == "" || courseTeacherId == "" || courseSemester == "" {
		return false
	}
	class, _ := GetClasses(courseCode, "", courseTeacherId, courseSemester, "", "")
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

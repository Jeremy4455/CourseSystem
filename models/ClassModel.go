package models

import (
	"errors"
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
		teacher, err := GetTeachers(courseTeacherId, courseTeacherName, "", "")
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

func CreateClass(courseCode, courseName, courseTeacherId, courseSemester, classTime, capacity, classroom string) error {
	if courseCode == "" && courseName == "" || courseTeacherId == "" || courseSemester == "" || classTime == "" || classroom == "" || capacity == "" {
		return errors.New("参数不能为空")
	}
	course, err := GetCourses(courseCode, courseName)
	fmt.Println(len(course))
	if len(course) != 1 {
		return errors.New("选定课程不唯一")
	}

	teacher, _ := GetTeachers(courseTeacherId, "", "", "")
	if len(teacher) == 0 {
		return errors.New("不存在该教师")
	}
	if ExistClass(course[0], teacher[0], courseSemester) == true {
		return errors.New("该教师本学期已授该课")
	}
	cap, err := strconv.Atoi(capacity)
	if err != nil {
		return err
	}

	idv, err := ClassId()
	if err != nil {
		return err
	}
	class := &Class{
		Id:        idv + 1,
		Course:    course[0],
		Teacher:   teacher[0],
		Semester:  courseSemester,
		ClassTime: classTime,
		Capacity:  cap,
		Location:  classroom,
	}
	o := orm.NewOrm()
	_, err = o.Insert(class)
	var existedClass []*Class
	_, err = o.QueryTable("Class").Filter("Location", classroom).All(&existedClass)
	if err != nil {
		return err
	}
	for _, c := range existedClass {
		if ClassConflict(class, c) == false {
			return errors.New("课程存在时间冲突")
		}
	}
	if TeacherTimeConflict(teacher[0], class) == false {
		return errors.New("该教师存在时间冲突")
	}
	if err != nil {
		fmt.Println(123)
		return err
	}
	return nil
}

func DeleteClass(courseCode, courseTeacherId, courseSemester string) error {
	if courseCode == "" || courseTeacherId == "" || courseSemester == "" {
		return errors.New("参数不能为空")
	}
	class, _ := GetClasses(courseCode, "", courseTeacherId, "", courseSemester, "", "")
	if len(class) == 0 {
		return errors.New("不存在该课程")
	}

	o := orm.NewOrm()
	_, err := o.Delete(class[0])
	if err != nil {
		return err
	}
	return nil
}

func ReviseClass(c *Class, courseTeacherId, classTime, capacity, classroom string) error {
	if courseTeacherId != "" {
		teacher, err := GetTeachers(courseTeacherId, "", "", "")
		if err != nil {
			return err
		}
		c.Teacher = teacher[0]
	}
	if classTime != "" {
		c.ClassTime = classTime
	}
	if capacity != "" {
		cap, err := strconv.Atoi(capacity)
		if err != nil {
			return err
		}
		c.Capacity = cap
	}
	if classroom != "" {
		c.Location = classroom
	}

	_, err := orm.NewOrm().Update(c)
	if err != nil {
		return err
	}
	return nil
}

func GetAllClasses() ([]*Class, error) {
	var classes []*Class
	_, err := orm.NewOrm().QueryTable("class").All(&classes)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

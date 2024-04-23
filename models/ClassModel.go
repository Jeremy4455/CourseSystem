package models

import (
	"errors"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
)

type Class struct {
	Id        int64    `orm:"pk;auto"`
	Course    *Course  `orm:"rel(fk);on_delete(cascade)"` // Course作为外键
	Teacher   *Teacher `orm:"rel(fk);on_delete(cascade)"` // Teacher作为外键
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

func GetClasses(courseCode, courseName, courseTeacherId, courseTeacherName, courseSemester, classTime, classroom string) ([]*Class, error) {
	if courseSemester == "" {
		return nil, errors.New("必须选择一个学期")
	}

	o := orm.NewOrm()
	q := o.QueryTable("class").RelatedSel().Filter("Semester", courseSemester)

	if courseCode != "" || courseName != "" {
		courses, err := GetCourses(courseCode, courseName)
		if err != nil {
			return nil, err
		}
		q = q.Filter("Course__in", courses)
	}

	if courseTeacherId != "" || courseTeacherName != "" {
		teachers, err := GetTeachers(courseTeacherId, courseTeacherName, "", "")
		if err != nil {
			return nil, err
		}
		q = q.Filter("Teacher__in", teachers)
	}

	if classTime != "" {
		q = q.Filter("ClassTime__contains", classTime)
	}

	if classroom != "" {
		q = q.Filter("Location__contains", classroom)
	}

	var classes []*Class
	_, err := q.All(&classes)

	if err != nil {
		return nil, err
	}
	return classes, nil
}

func CreateClass(courseCode, courseName, courseTeacherId, courseSemester, classTime, capacity, classroom string) error {
	if courseCode == "" || courseTeacherId == "" || courseSemester == "" || classTime == "" || classroom == "" || capacity == "" {
		return errors.New("参数不能为空")
	}

	course, err := GetCourses(courseCode, courseName)
	if err != nil {
		return err
	}
	if len(course) != 1 {
		return errors.New("选定课程不唯一")
	}

	teacher, err := GetTeachers(courseTeacherId, "", "", "")
	if err != nil {
		return err
	}
	if len(teacher) == 0 {
		return errors.New("不存在该教师")
	}

	if ExistClass(course[0], teacher[0], courseSemester) {
		return errors.New("该教师本学期已授该课")
	}

	cap, err := strconv.Atoi(capacity)
	if err != nil {
		return err
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
	_, err = o.QueryTable("Class").Filter("Semester", courseSemester).Filter("Location", classroom).All(&existedClass)
	if err != nil {
		return err
	}
	for _, c := range existedClass {
		if !ClassConflict(class, c) {
			return errors.New("课程存在时间冲突")
		}
	}
	if !TeacherTimeConflict(teacher[0], class) {
		return errors.New("该教师存在时间冲突")
	}

	_, err = o.Insert(class)
	if err != nil {
		return err
	}

	// teacher[0].Classes = append(teacher[0].Classes, class)
	// _, err = o.Update(teacher[0])
	// if err != nil {
	// 	return err
	// }
	// course[0].Classes = append(course[0].Classes, class)
	// _, err = o.Update(course[0])
	// if err != nil {
	// 	return err
	// }
	return nil
}

func DeleteClass(courseCode, courseTeacherId, courseSemester string) error {
	if courseCode == "" || courseTeacherId == "" || courseSemester == "" {
		return errors.New("参数不能为空")
	}
	class, err := GetClasses(courseCode, "", courseTeacherId, "", courseSemester, "", "")
	if err != nil {
		return err
	}
	if len(class) == 0 {
		return errors.New("不存在该课程")
	}

	o := orm.NewOrm()
	_, err = o.Delete(class[0])
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
	_, err := orm.NewOrm().QueryTable("class").RelatedSel().All(&classes)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

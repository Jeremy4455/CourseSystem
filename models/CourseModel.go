package models

import (
	"errors"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
)

// Course 课程模型
type Course struct {
	CourseCode string `orm:"unique;size(20);pk"`
	Name       string `orm:"size(100)"`
	College    string `orm:"size(100)"`
	Credit     int    // 学分
	Proportion float64
	Classes    []*Class `orm:"reverse(many)"`
}

func GetCourses(courseCode, name string) ([]*Course, error) {
	if courseCode == "" && name == "" {
		return nil, errors.New("参数不能为空")
	}

	o := orm.NewOrm()
	q := o.QueryTable("course")
	var courses []*Course

	if courseCode != "" {
		q = q.Filter("CourseCode", courseCode)
	} else {
		q = q.Filter("Name__contains", name)
	}
	_, err := q.All(&courses)

	return courses, err
}

func CreateCourse(courseCode, name, college, credit, proportion string) error {
	course, err := GetCourses(courseCode, "")
	if len(course) != 0 {
		return err
	}
	c, err := strconv.Atoi(credit)
	if err != nil {
		return err
	}
	p, err := strconv.ParseFloat(proportion, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	newcourse := &Course{
		CourseCode: courseCode,
		Name:       name,
		College:    college,
		Credit:     c,
		Proportion: p,
	}
	_, err = o.Insert(newcourse)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCourse(courseCode string) error {
	o := orm.NewOrm()

	course := &Course{CourseCode: courseCode}
	err := o.Read(course)
	if err != nil {
		return err
	}
	o.Delete(course)
	return nil
}

func ReviseCourse(courseCode, name, college, credit, proportion string) error {
	courses, err := GetCourses(courseCode, "")
	if err != nil {
		return err
	}

	course := courses[0]
	if name != "" {
		course.Name = name
	}
	if college != "" {
		course.College = college
	}
	if credit != "" {
		c, err := strconv.Atoi(credit)
		if err != nil {
			return err
		}
		course.Credit = c
	}
	if proportion != "" {
		p, err := strconv.ParseFloat(proportion, 64)
		if err != nil {
			return err
		}
		course.Proportion = p
	}

	_, err = orm.NewOrm().Update(course)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCourses() ([]*Course, error) {
	var courses []*Course
	_, err := orm.NewOrm().QueryTable("course").All(&courses)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

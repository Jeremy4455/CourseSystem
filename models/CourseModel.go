package models

import (
	"strconv"

	"github.com/astaxie/beego/orm"
)

// Course 课程模型
type Course struct {
	CourseCode string   `orm:"unique;size(20);pk"`
	Name       string   `orm:"size(100)"`
	College    string   `orm:"size(100)"`
	Credit     int      // 学分
	Classes    []*Class `orm:"reverse(many)"`
}

func GetCourse(courseCode, name string) ([]*Course, error) {
	if courseCode == "" && name == "" {
		return nil, nil
	}

	o := orm.NewOrm()
	q := o.QueryTable("course")
	var courses []*Course

	if courseCode != "" {
		_, err := q.Filter("CourseCode", courseCode).All(&courses)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := q.Filter("Name__contains", name).All(&courses)
		if err != nil {
			return nil, err
		}
	}

	return courses, nil
}

func AddCourse(courseCode, name, college, credit string) error {
	course, err := GetCourse(courseCode, "")
	if course != nil {
		return err
	}
	c, err := strconv.Atoi(credit)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	newcourse := &Course{CourseCode: courseCode, Name: name, College: college, Credit: c}
	o.Insert(newcourse)
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

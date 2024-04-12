package models

import (
	"strconv"

	"github.com/beego/beego/v2/client/orm"
)

// Student 学生模型
type Student struct {
	StudentId string   `orm:"unique;size(20);pk"`
	Name      string   `orm:"size(100)"`
	Class     string   `orm:"size(50)"`
	Grade     float64  `orm:"size(20)"`
	Classes   []*Class `orm:"rel(m2m)"`
}

func GetStudent(studentId string) (*Student, error) {
	o := orm.NewOrm()

	student := &Student{StudentId: studentId}

	err := o.Read(student)
	if err != nil {
		// 处理错误
		return nil, err
	}

	return student, nil
}

func AddStudent(studentId, name, class, grade string) error {
	student, err := GetStudent(studentId)
	if student != nil {
		return err
	}

	g, err := strconv.ParseFloat(grade, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	student = &Student{StudentId: studentId, Name: name, Class: class, Grade: g}
	_, err = o.Insert(student)
	if err != nil {
		return err
	}

	AddUser(studentId, name, "123456", "student")
	return nil
}

func DeleteStudent(studentId string) error {
	o := orm.NewOrm()

	student := &Student{StudentId: studentId}
	err := o.Read(student)
	if err != nil {
		return err
	}
	_, err = o.Delete(student)
	if err != nil {
		return err
	}

	DeleteUser(studentId)
	return nil
}

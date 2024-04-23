package models

import (
	"errors"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type ClassStudent struct {
	Id          int64
	Class       *Class   `orm:"rel(fk);on_delete(cascade)"`
	Student     *Student `orm:"rel(fk)"`
	Performance float64
	Score       float64
}

func (c *ClassStudent) TableIndex() [][]string {
	return [][]string{
		[]string{"Class", "Student"},
	}
}
func (c *ClassStudent) TableUnique() [][]string {
	return [][]string{
		[]string{"Class", "Student"},
	}
}
func (c *ClassStudent) TableName() string {
	return "class_student"
}

func GetClassStudent(s *Student, courseSemester string, c *Class) ([]*ClassStudent, error) {
	var class_student []*ClassStudent
	o := orm.NewOrm()
	if s == nil {
		_, err := o.QueryTable("class_student").Filter("class_id", c.Id).All(&class_student)
		if err != nil {
			return nil, err
		}
	}
	if c == nil {
		var temp []*ClassStudent
		_, err := o.QueryTable("class_student").Filter("student_id", s.StudentId).All(&temp)
		if err != nil {
			return nil, err
		}
		for _, t := range temp {
			if t.Class.Semester != courseSemester {
				continue
			}
			class_student = append(class_student, t)
		}
	}
	return class_student, nil
}

func PickClass(s *Student, c *Class) error {
	if s == nil || c == nil {
		return errors.New("不存在该课程")
	}
	o := orm.NewOrm()
	var classstuent []*ClassStudent
	_, err := o.QueryTable("class_student").All(&classstuent)
	if err != nil {
		return err
	}
	if len(classstuent) >= c.Capacity {
		return errors.New("该课程人数已满")
	}
	// if !StudentTimeConflict(s, c) {
	// 	return errors.New("该课程与已选课程冲突")
	// }

	id, err := GetId("ClassStudent")
	if err != nil {
		return err
	}
	newclassstudent := &ClassStudent{
		Id:          id + 1,
		Class:       c,
		Student:     s,
		Performance: 0,
		Score:       0,
	}
	o.Insert(newclassstudent)
	return nil
}

func DropClass(s *Student, c *Class) error {
	if s == nil || c == nil {
		return errors.New("不存在该课程")
	}
	o := orm.NewOrm()

	classstudent := &ClassStudent{Class: c, Student: s}
	err := o.Read(classstudent)
	if err != nil {
		return err
	}
	o.Delete(classstudent)
	return nil
}

func UpdateClass(s *Student, c *Class, performance, score string) error {
	if c == nil || s == nil || performance == "" && score == "" {
		return errors.New("")
	}
	o := orm.NewOrm()
	classstudent := &ClassStudent{Class: c, Student: s}
	o.Read(classstudent)

	if performance != "" {
		perform, err := strconv.ParseFloat(performance, 64)
		if err != nil {
			return err
		}
		classstudent.Performance = perform
	}

	if score != "" {
		sco, err := strconv.ParseFloat(score, 64)
		if err != nil {
			return err
		}
		classstudent.Score = sco
	}

	return nil
}

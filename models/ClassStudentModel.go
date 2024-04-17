package models

import (
	"strconv"

	"github.com/astaxie/beego/orm"
)

type ClassStudent struct {
	Id          int
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

func PickClass(s *Student, c *Class) bool {
	if s == nil || c == nil {
		return false
	}
	o := orm.NewOrm()
	var classstuent []*ClassStudent
	_, err := o.QueryTable("class_student").All(&classstuent)
	if err != nil {
		return false
	}
	if len(classstuent) >= c.Capacity {
		return false
	}
	if StudentTimeConflict(s, c) == false {
		return false
	}

	newclassstudent := &ClassStudent{
		Class:       c,
		Student:     s,
		Performance: 0,
		Score:       0,
	}
	o.Insert(newclassstudent)
	return true
}

func DropClass(s *Student, c *Class) bool {
	if s == nil || c == nil {
		return false
	}
	o := orm.NewOrm()

	classstudent := &ClassStudent{Class: c, Student: s}
	err := o.Read(classstudent)
	if err != nil {
		return false
	}
	o.Delete(classstudent)
	return true
}

func UpdateClass(s *Student, c *Class, performance, score string) bool {
	if c == nil || s == nil || performance == "" && score == "" {
		return false
	}
	o := orm.NewOrm()
	classstudent := &ClassStudent{Class: c, Student: s}
	o.Read(classstudent)

	if performance != "" {
		perform, err := strconv.ParseFloat(performance, 64)
		if err != nil {
			return false
		}
		classstudent.Performance = perform
	}

	if score != "" {
		sco, err := strconv.ParseFloat(score, 64)
		if err != nil {
			return false
		}
		classstudent.Score = sco
	}

	return true
}

package models

import (
	"errors"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
)

type ClassStudent struct {
	Id          int64    `orm:"pk;auto"`
	Class       *Class   `orm:"rel(fk);on_delete(cascade)"`
	Student     *Student `orm:"rel(fk);on_delete(do_nothing)"`
	Performance float64
	Score       float64
	Fscore      float64
}

func (c *ClassStudent) TableIndex() [][]string {
	return [][]string{
		{"Class", "Student"},
	}
}
func (c *ClassStudent) TableUnique() [][]string {
	return [][]string{
		{"Class", "Student"},
	}
}
func (c *ClassStudent) TableName() string {
	return "class_student"
}

func GetClassStudent(s *Student, courseSemester string, c *Class) ([]*ClassStudent, error) {
	var class_student []*ClassStudent
	o := orm.NewOrm()
	if s != nil && c != nil {
		_, err := o.QueryTable("ClassStudent").RelatedSel().Filter("Class", c).Filter("Student", s).All(&class_student)
		if err != nil {
			return nil, err
		}
		if len(class_student) == 0 {
			return nil, errors.New("没有该记录")
		}
		return class_student, nil
	}

	if s == nil {
		_, err := o.QueryTable("ClassStudent").RelatedSel().Filter("Class", c).All(&class_student)
		if err != nil {
			return nil, err
		}
		if len(class_student) == 0 {
			return nil, errors.New("没有该记录")
		}
		return class_student, nil
	}
	if c == nil {
		var temp []*ClassStudent
		_, err := o.QueryTable("ClassStudent").RelatedSel().Filter("Student", s).All(&temp)
		if err != nil {
			return nil, err
		}
		for _, t := range temp {
			if t.Class.Semester != courseSemester {
				continue
			}
			class_student = append(class_student, t)
		}
		if len(class_student) == 0 {
			return nil, errors.New("没有该记录")
		}
		return class_student, nil
	}
	return nil, errors.New("参数错误")
}

func PickClass(s *Student, c *Class, level int) error {
	if s == nil || c == nil {
		return errors.New("不存在该课程")
	}
	if c.Level > level {
		return errors.New("权限不足")
	}
	if level == STUDENT_PICK_DROP_CLASS && c.Count >= c.Capacity {
		return errors.New("该课程人数已满")
	}

	o := orm.NewOrm()
	newclassstudent := &ClassStudent{
		Class:       c,
		Student:     s,
		Performance: 0,
		Score:       0,
	}
	_, err := o.Insert(newclassstudent)
	if err != nil {
		return err
	}
	return nil
}

func DropClass(s *Student, c *Class, level int) error {
	if s == nil || c == nil {
		return errors.New("不存在该课程")
	}
	o := orm.NewOrm()

	classstudent := &ClassStudent{Class: c, Student: s}
	err := o.Read(classstudent, "Class", "Student")
	if err != nil {
		return err
	}
	if level < c.Level {
		return errors.New("权限不足")
	}
	_, err = o.Delete(classstudent)
	if err != nil {
		return err
	}
	return nil
}

func UpdateClass(s *Student, c *Class, performance, score string, level int) error {
	if c == nil || s == nil || performance == "" && score == "" {
		return errors.New("输入有误")
	}

	o := orm.NewOrm()
	classstudent := &ClassStudent{Class: c, Student: s}
	if err := o.Read(classstudent, "Class", "Student"); err != nil {
		return err
	}
	if level < c.Level {
		return errors.New("权限不足")
	}

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
	classstudent.Fscore = c.Course.Proportion*classstudent.Performance + (1-c.Course.Proportion)*classstudent.Score
	if _, err := o.Update(classstudent); err != nil {
		return err
	}
	return nil
}

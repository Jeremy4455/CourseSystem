package models

import (
	"errors"

	"github.com/beego/beego/v2/client/orm"
)

// Teacher 老师模型
type Teacher struct {
	TeacherId string   `orm:"unique;size(20);pk"`
	Name      string   `orm:"size(100)"`
	Mobile    string   `orm:"size(100)"`
	Email     string   `orm:"size(100)"`
	Classes   []*Class `orm:"reverse(many);on_delete(cascade)"`
}

func GetTeachers(teacherId, name, mobile, email string) ([]*Teacher, error) {
	o := orm.NewOrm()
	q := o.QueryTable("teacher")

	var teacher []*Teacher
	if teacherId != "" {
		_, err := q.Filter("TeacherId", teacherId).All(&teacher)
		if err != nil {
			return nil, err
		}
	}
	if name != "" {
		_, err := q.Filter("Name__contains", name).All(&teacher)
		if err != nil {
			return nil, err
		}
	}
	if mobile != "" {
		_, err := q.Filter("Mobile", mobile).All(&teacher)
		if err != nil {
			return nil, err
		}
	}
	if email != "" {
		_, err := q.Filter("Email", email).All(&teacher)
		if err != nil {
			return nil, err
		}
	}

	if len(teacher) == 0 {
		return nil, errors.New("没有该教师")
	}

	return teacher, nil
}

func CreateTeacher(teacherId, name, mobile, email string) error {
	teachers, err := GetTeachers(teacherId, "", "", "")
	if len(teachers) != 0 {
		return err
	}

	o := orm.NewOrm()
	teacher := &Teacher{
		TeacherId: teacherId,
		Name:      name,
		Mobile:    mobile,
		Email:     email,
	}
	_, err = o.Insert(teacher)
	if err != nil {
		return err
	}

	AddUser(teacherId, name, "teacher123", "teacher")
	return nil
}

func DeleteTeacher(teacherId string) error {
	o := orm.NewOrm()

	teacher := &Teacher{TeacherId: teacherId}
	err := o.Read(teacher)
	if err != nil {
		return err
	}
	_, err = o.Delete(teacher)
	if err != nil {
		return err
	}

	DeleteUser(teacherId)
	return nil
}

func ReviseTeacher(teacherId, name, mobile, email string) error {
	teachers, err := GetTeachers(teacherId, "", "", "")
	if err != nil {
		return err
	}

	t := teachers[0]
	if name != "" {
		t.Name = name
	}
	if mobile != "" {
		t.Mobile = mobile
	}
	if email != "" {
		t.Email = email
	}
	_, err = orm.NewOrm().Update(t)
	if err != nil {
		return err
	}
	return nil
}

func GetAllTeachers() ([]*Teacher, error) {
	var teachers []*Teacher
	_, err := orm.NewOrm().QueryTable("teacher").All(&teachers)
	if err != nil {
		return nil, err
	}
	return teachers, nil
}
